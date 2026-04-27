package main

import (
	"database/sql"
	"log"

	"github.com/alexnakagama/go-simple-bank/config"
	"github.com/alexnakagama/go-simple-bank/internal/api"
	db "github.com/alexnakagama/go-simple-bank/internal/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
