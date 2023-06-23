#include <assert.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <netinet/in.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>
#define _OPEN_THREADS

#ifndef MYSERVERH

#define MYSERVERH

//CONSTANTS
#define PORTNUM 3000
#define MAXQUEUELEN 6
#define MAXLOBBIES 100000

enum infoMap {
	MAXGAMETYPE = 255,
	ISOPENBIT = 256,
	ISFINBIT = 512,
	PONEDC = 1024,
	PTWODC = 2048,
	ERRBITS = ~1023
};
//structs
typedef struct {
	unsigned int info;
	int playerOneSock;
	int playerTwoSock;
	int (*messageRules)(int, int);
	pthread_mutex_t lock;
} lobby;

//Variables
extern int serverSock;
extern lobby lobbies[MAXLOBBIES];

//functions
void setupInterruptHandlers();
void sigInterrupt();

//int getLobbyReadyStatus( lobby *);
//int isLobbyFull( lobby *);

#endif
