include .env

up:
	docker compose up -d

down:
	docekr compose down

build:
	go build -o ${BINARY} ./cmd

start:
	./${BINARY}

restart: build start
