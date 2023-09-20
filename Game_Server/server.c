#include "server.h"
#include "server_util.h"

int serverSock;
lobby lobbies[MAXLOBBIES];


int main(int argc, char** argv) {
	
	int port = PORTNUM;

	

	// Set the IP addr and port for the server
	struct sockaddr_in serverIPAddr;
	memset( &serverIPAddr, 0, sizeof(serverIPAddr) );
	serverIPAddr.sin_family = AF_INET;
	serverIPAddr.sin_addr.s_addr = INADDR_ANY;
	serverIPAddr.sin_port = htons( (u_short) port);

	//Allocate a socket to represent the server
	serverSock = socket( PF_INET, SOCK_STREAM, 0);
	assert( serverSock >= 0);

	//interrupt handlers
	setupInterruptHandlers();

	//establishing socket options
	int optval = 1;
	assert( setsockopt( serverSock, SOL_SOCKET, SO_REUSEADDR,
							(char *) &optval, sizeof( optval ) ) == 0);
	
	//Binding the socket
	assert( bind( serverSock, (struct sockaddr *) &serverIPAddr,
								sizeof( serverIPAddr) ) == 0);
	
	//display the hostname:port of the server
	char hostname[1024];
	hostname[1023] = '\0';
	gethostname(hostname, 1023); 
	//converting to a IPv4 address
	struct hostent* hostEntry = gethostbyname(hostname);
	char * IPBuff = inet_ntoa( *( ( struct in_addr*) hostEntry->h_addr_list[0]));

	printf("IP Address -> %s:%d\n", IPBuff, port);
	printf("URL -> %s:%d\n", hostname, port);

	//set server to begin listening
	assert( listen( serverSock, MAXQUEUELEN) == 0 );
	initConnections();
	//close down resources
	close( serverSock);

  return 0;
}



void setupInterruptHandlers() {
	struct sigaction signalAction;
	signalAction.sa_handler = sigInterrupt;
	sigemptyset( &signalAction.sa_mask);
	signalAction.sa_flags = SA_RESTART;

	//redirect interrupts to sigInterrupt funtion
	assert( sigaction( SIGINT, &signalAction, NULL) == 0 );
	assert( sigaction( SIGPIPE, &signalAction, NULL) == 0);
	assert( sigaction( SIGCHLD, &signalAction, NULL) == 0);
	assert( sigaction( SIGINT, &signalAction, NULL) == 0);
}

void sigInterrupt( int sigNum ) {
	switch ( sigNum) {
		case SIGINT:
			close( serverSock);
			while( waitpid(-1, NULL, WNOHANG) > 0);
			printf("Exiting from ^C\n");
			exit( 1);
			break;
		case SIGCHLD:
			printf("Received a SIGCHLD\n");
		case SIGPIPE:
			while (waitpid(-1, NULL, WNOHANG) > 0);
			printf("Received a SIGPIPE\n");
			break;
		default:
			printf("Received unknown error: %d\n", sigNum);
	}
}
