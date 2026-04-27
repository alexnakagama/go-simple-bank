package main

import (
	"database/sql"
	"log"
	"os"
	"github.com/alexnakagama/go-simple-bank/internal/api"
	db "github.com/alexnakagama/go-simple-bank/internal/db/sqlc"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load .env file:", err)
	}
	dbSource := os.Getenv("DB_SOURCE")

	sqlDB, err := sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(sqlDB)
	server := api.NewServer(store)

	if err := server.Start(":8080"); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
