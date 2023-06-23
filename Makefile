
CFLAGS = -Wall -Werror -g -pthread
SOURCES = server.c server_util.c server_rules.c
OBJECTS = $(SOURCES:.c=.o)
HFILES = $(SOURCES:.c=.h)
EXEC = run

all: $(EXEC)
	@make clean
	@make $(EXEC)

$(EXEC): $(OBJECTS)
	gcc $(CFLAGS) -o $@ $^

$(OBJECTS): $(SOURCES)
	gcc $(CFLAGS) -c $^

clean:
	rm run *.o

.PHONY: clean
