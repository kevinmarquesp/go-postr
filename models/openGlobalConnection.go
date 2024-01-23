package models

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func OpenGlobalConnection(host, port, user, password, dbname string) {
	var err error

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err = sql.Open("postgres", conn)
	if err != nil {
		log.Println("Could not open the postgres connection")
		log.Panic(err)
	}
}
