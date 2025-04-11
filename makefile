.PHONY: build migrate up down test

build:
	go build -o bin/book-service cmd/book-service/main.go

migrate:
	migrate -path migrations -database "postgres://bookuser:bookpass@localhost:5432/bookdb?sslmode=disable" up

up:
	docker-compose -f ../docker-compose.yml up --build

down:
	docker-compose -f ../docker-compose.yml down

test:
	go test ./...