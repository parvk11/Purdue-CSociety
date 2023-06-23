#include "server.h"
#include "server_util.h"
#include "server_rules.h"

inline unsigned int getLobbyGameType( lobby * lobby ) {
	return lobby->info & MAXGAMETYPE;
}

inline int isLobbyOpen( lobby * lobby ) {
	return ( lobby->info & ISOPENBIT) != 0;
}


int isLobbyOpen_THR( lobby * lobby) {
	pthread_mutex_lock(&lobby->lock);
	int isOpen = isLobbyOpen( lobby );
	pthread_mutex_unlock(&lobby->lock);
	return isOpen;
}

inline int isLobbyFinished( lobby * lobby ) {
	return ( lobby->info & ISFINBIT) != 0;
}

inline unsigned int getLobbyErrorStatus( lobby * lobby ) {
	return lobby->info & ERRBITS;
}

//Lobby Setters
inline void setLobbyGameType( lobby * lobby, int gameType ) {
	lobby->info = lobby->info | gameType;
}

inline void setLobbyClosed( lobby * lobby ) {
	lobby->info = lobby->info & ~ISOPENBIT;
}

inline void setLobbyOpen( lobby * lobby ) {
	lobby->info = lobby->info | ISOPENBIT;
}

inline void setLobbyFinished( lobby * lobby ) {
	lobby->info = lobby->info | ISFINBIT;
}

inline void setLobbyPOneDC( lobby * lobby ) {
	lobby->info = lobby->info | PONEDC;
}

inline void setLobbyPTwoDC( lobby * lobby ) {
	lobby->info = lobby->info | PTWODC;
}

void setLobbyMessagingRules( lobby * lobby ) {
	switch ( getLobbyGameType( lobby ) ) {
		case 5:
			lobby->messageRules = randomChoice;
			break;
		default:
			lobby->messageRules = relay;
	}
}

void resetLobby( lobby * lobby ) {
	lobby->info = ISOPENBIT;
	lobby->playerOneSock = 0;
	lobby->playerTwoSock = 0;
	lobby->messageRules = NULL;
}

void closeLobby( lobby * lobby ) {
	//close player communication
	shutdown( lobby->playerOneSock, SHUT_RDWR);
	shutdown( lobby->playerTwoSock, SHUT_RDWR);
	close( lobby->playerOneSock );
	close( lobby->playerTwoSock );
	//open the lobby back up for a new game
	pthread_mutex_lock( &lobby->lock);
	resetLobby( lobby );
	pthread_mutex_unlock( &lobby->lock);
}

//initializes the lobbies array to all be open with default values
void initLobbies() {
	for( int i = 0; i < MAXLOBBIES; i++) {
		resetLobby( lobbies + i);
		pthread_mutex_init(&(lobbies + i)->lock, NULL);
	}
}


//Creates the server logics for accepting new requests
//spawns a new thread that represents a player
void initConnections() {
	//memset( lobbies, 0, sizeof(lobbies) ); //I don't think this is neccessary
	initLobbies();
	while ( 1) {
		printf("Listening for connections:  %d\n", serverSock);
		struct sockaddr_in clientIPAddr;
		int alen = sizeof( clientIPAddr);
		int playerSock = accept( serverSock, (struct sockaddr *) &clientIPAddr,
											(socklen_t *) &alen);
		assert( playerSock >= 0);
		printf("made a connection\n");

		//get and check the game type
		char initConBuffer[14];
		initConBuffer[13] = '\0';
		int mesLen = read( playerSock, initConBuffer, 13);
		if ( mesLen <= 0)  {
			close( playerSock);
			continue;
		}
		printf("%s\n", initConBuffer);

		if ( strncmp(initConBuffer, "CONNECT:", 8) != 0 || 
				 strncmp(initConBuffer + mesLen - 2, "\r\r", 2) != 0 ) {
			//Error handling message here
			printf("Bad Format\n");
			write( playerSock, "REJECTED: Bad Format", 20);
			close( playerSock);
			continue;
		}

		int gameNum = atoi( initConBuffer + 8 );
		printf("%d\n", gameNum);

		if ( gameNum <= 0 || gameNum > MAXGAMETYPE ) {
			//Error handling message here
			printf("Bad game type\n");
			write( playerSock, "REJECTED: Bad Game Type", 23);
			close( playerSock);
			continue;
		}
		//add a player to an empty lobby
		int skipThread = 0;
		int i = 0;
		for( i = 0; i < MAXLOBBIES; i++) {
			lobby * lobbyToJoin = lobbies + i;
			pthread_mutex_lock(&lobbyToJoin->lock);
			//TODO: break this into two if statements and remove the nesting
			if ( isLobbyOpen(lobbyToJoin) && 
					(getLobbyGameType(lobbyToJoin) == 0 || 
					getLobbyGameType(lobbyToJoin) == gameNum) ) {
				if ( getLobbyGameType( lobbyToJoin) == 0 ) { //first person to join
					setLobbyGameType( lobbyToJoin, gameNum );
					setLobbyMessagingRules( lobbyToJoin );
					lobbyToJoin->playerOneSock = playerSock;
					mesLen = write( playerSock, "ACCEPTED: Joined a lobby", 24 );
					if ( mesLen <= 0) {
						close( playerSock);
						resetLobby( lobbyToJoin);
						skipThread = 1;
					}
				} else {
					lobbyToJoin->playerTwoSock = playerSock;
					setLobbyClosed( lobbyToJoin );
					skipThread = 1;
					write( playerSock, "ACCEPTED: Joined a lobby", 24);
					//no error check as the game loop should already do the check
				}
				pthread_mutex_unlock(&lobbyToJoin->lock);
				break;
			}
			pthread_mutex_unlock(&lobbyToJoin->lock);
		}

		if ( i == MAXLOBBIES ) {
			printf("all lobbies full\n");
			write( playerSock, "REJECTED: Lobbies Full", 22 );
			close( playerSock);
		}

		//create a new thread
		if ( !skipThread ) {
			pthread_t thread;
			pthread_attr_t attr;
			pthread_attr_init(&attr);
			//not using pthread_join() in this server
			pthread_attr_setdetachstate( &attr, PTHREAD_CREATE_DETACHED);
			pthread_create( &thread, &attr,
					startLobby, (void *)(lobbies + i) ); //free this later... maybe?
		}
	}
	return;
}

void * startLobby( void * returnThis ) {
	
	lobby * lobby = returnThis;
	while ( isLobbyOpen( lobby) ) {
		//write( lobby->playerOneSock, "WAITING:", 8);
		sched_yield();
	}
	//TODO: checking if connections are still alive
	char startBuff[10];
	startBuff[9] = '\0';
	write( lobby->playerOneSock, "VALIDATE:", 9);
	write( lobby->playerTwoSock, "VALIDATE:", 9);
	read( lobby->playerOneSock, startBuff, 9);
	read( lobby->playerTwoSock, startBuff, 9);
	lobby->messageRules(lobby->playerOneSock, lobby->playerTwoSock);

	//lobby is finished
	// by error or by end of game
	closeLobby( lobby);
	
	return returnThis;
}

//TODO:
// create function that waits for two players to join a lobby thread
// begin communication with game rules

void * joinThreadToLobby( void * playerSock) {
	void * returnVal = NULL;
	//int player = * (int *) playerSock;
	return returnVal;
}

void * testing( void * playerSock) {
	void * returnVal = NULL;
	int player = *(int *) playerSock;

	int charsRead = 0;
	//char buffer[2024];
	char * buffer = malloc( 2048 );
	memset( buffer, 0, 2048);
	//printf("here\n");
	//charsRead = read( player, buffer, 5);
	//printf("num read: %d\n", charsRead);
	//printf("%s\n", buffer);
	while ( ( charsRead = read( player, buffer, 2023) ) > 0 ) {
		printf("%s\n", buffer);
		buffer[charsRead] = '\0';
		write( player, buffer, 2023);
	}

	close( player);
	free( playerSock);
	free( buffer);
	pthread_exit(NULL);

	return returnVal;
}


