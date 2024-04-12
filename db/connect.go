package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConnCredentials struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

var conn *sql.DB

// Expose the connection object to the other parts of the code, but this
// package should abstract the database interaction through specific functions.
func Connection() *sql.DB {
	return conn
}

func Connect(cred ConnCredentials) error {
	credentials := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cred.Host, cred.Port, cred.Username, cred.Password, cred.DatabaseName)

	var err error

	conn, err = sql.Open("postgres", credentials)
	if err != nil {
		return err
	}

	return nil
}
