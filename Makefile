build:
	CGO_ENABLED=0 GOOS=linux go build -o audio
	./audio

server:
	./audio