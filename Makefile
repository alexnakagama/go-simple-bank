DB_URL=postgresql://root:secret@localhost:5432/root?sslmode=disable

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root root

dropdb:
	docker exec -it postgres15 dropdb root

migrateup:
	migrate -path internal/db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlc