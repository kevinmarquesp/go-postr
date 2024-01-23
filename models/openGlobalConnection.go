package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

var Db *sql.DB

func OpenGlobalConnection(cre Connection) {
	var err error

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cre.Host, cre.Port, cre.User, cre.Password, cre.Dbname)

	Db, err = sql.Open("postgres", conn)
	if err != nil {
		log.Println("Could not open the postgres connection")
		log.Panic(err)
	}
}
