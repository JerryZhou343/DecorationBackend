.PHONY: all clean idl uidl

OUTPUT=decoration-backend-server

all: clean
	go build -o ./bin/${OUTPUT} main.go

clean:
	rm -f bin/${OUTPUT}

