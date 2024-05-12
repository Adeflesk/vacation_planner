package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Adeflesk/vacation_planner/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDRIVER, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	fmt.Println(conn)
}
