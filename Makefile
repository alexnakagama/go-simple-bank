.PHONY: all build test run migrate createdb dropdb

DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

all: build

build:
	go build -v ./...

test:
	DB_SOURCE=$(DB_SOURCE) go test -v ./...

run:
	DB_SOURCE=$(DB_SOURCE) go run main.go

migrate:
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose up

createdb:
	createdb -U root simple_bank

dropdb:
	dropdb -U root simple_bank
