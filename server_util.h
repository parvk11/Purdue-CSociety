#ifndef MYSERVERUTILH

#define MYSERVERUTILH

typedef enum {
	INVALIDPASS = 0,
	VALIDPASS = 1,
	BADREQ = 2
} requestStatus;


void printHello();
//functions
void initConnections();
void initLobbies();
void * startLobby(void *);

//Lobby Getters
unsigned int getLobbyGameType( lobby * );
int isLobbyOpen( lobby * );
int isLobbyOpen_THR( lobby * );
int isLobbyFinished( lobby * );
unsigned int getLobbyErrorStatus( lobby * );

//Lobby setters
void setLobbyGameType( lobby * , int );
void setLobbyClosed( lobby * );
void setLobbyOpen( lobby * );
void setLobbyFinished( lobby * );
void setLobbyPOneDC( lobby * );
void setLobbyPTwoDC( lobby * );

void setLobbyMessagingRules( lobby * );
void resetLobby( lobby * );
void closeLobby( lobby * );

#endif
