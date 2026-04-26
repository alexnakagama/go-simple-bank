package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open("postgres", os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	testQueries = New(testDB)
	defer testDB.Close()
	os.Exit(m.Run())
}
