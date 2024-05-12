package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDb, err = sql.Open(config.DBDRIVER, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	testQueries = New(testDb)
	os.Exit(m.Run())
}
