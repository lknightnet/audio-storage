build:
	CGO_ENABLED=0 GOOS=linux go build -o audio

server:
	./audio

all:
	make build
	make server