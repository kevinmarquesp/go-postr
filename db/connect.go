package db

import (
	"database/sql"
	"fmt"
	"go-postr/utils"

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

func DefaultCredentials() (ConnCredentials, error) {
	pgHost, err := utils.RequireEnv("POSTGRES_HOST")
	pgPort, err := utils.RequireEnv("POSTGRES_PORT")
	pgUsernmae, err := utils.RequireEnv("POSTGRES_USER")
	pgPassword, err := utils.RequireEnv("POSTGRES_PASSWORD")
	pgDatabase, err := utils.RequireEnv("POSTGRES_DB")

	if err != nil {
		return ConnCredentials{}, err
	}

	creds := ConnCredentials{
		Host:         pgHost,
		Port:         pgPort,
		Username:     pgUsernmae,
		Password:     pgPassword,
		DatabaseName: pgDatabase,
	}

	return creds, nil
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
