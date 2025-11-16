APP_NAME=avito-tech-fall-2025
BINARY=server

.PHONY: build run docker-up docker-down fmt test

build:
	go build -o bin/$(BINARY) ./cmd/server

run:
	go run ./cmd/server

docker-up:
	cd deployments && docker-compose up --build

docker-down:
	cd deployments && docker-compose down

fmt:
	go fmt ./...

test:
	go test ./...