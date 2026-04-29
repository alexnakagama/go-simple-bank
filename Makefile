.PHONY: all build test run migrate createdb dropdb mock migrate1 migratedown migratedown1

DB_SOURCE=postgresql://root:secret@localhost:5432/root?sslmode=disable

all: build

build:
	go build -v ./...

test:
	DB_SOURCE=$(DB_SOURCE) go test -v ./...

run:
	DB_SOURCE=$(DB_SOURCE) go run main.go

migrate:
	migrate -path internal/db/migrations -database "$(DB_SOURCE)" -verbose up

migrate1:
	migrate -path internal/db/migrations -database "$(DB_SOURCE)" -verbose up 1

migratedown:
	migrate -path internal/db/migrations -database "$(DB_SOURCE)" -verbose down

migratedown1:
	migrate -path internal/db/migrations -database "$(DB_SOURCE)" -verbose down 1

createdb:
	createdb -U root simple_bank

dropdb:
	dropdb -U root simple_bank

mock:
	mockgen -package mock_db -destination internal/db/mock/store.go github.com/alexnakagama/go-simple-bank/internal/db/sqlc Store
