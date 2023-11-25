build:
	CGO_ENABLED=0 GOOS=linux go build -o audio

server:
	docker compose up --build

all:
	make build
	make server