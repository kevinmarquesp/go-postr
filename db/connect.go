package db

import (
	"database/sql"
	"fmt"
	"go-postr/utils"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const Dotenv = ".env"

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

// This function is meant to be used on the test files, it will create a dummy
// database with an unique UUID in the name (e.g.: test_799b1dc43f8d49d69341184b244a4013)
// and store its connection on the global `conn` variable in this package --
// like the `Connect()` function does.
func SetupDummyTestDatabaseConnection() (string, error) {
	const migrateSqlFile = "migrate.sql"
	testid := uuid.New().String()
	testdb := "test_" + strings.ReplaceAll(testid, "-", "")

	// go to the root of the project to load the environment file before anything

	os.Chdir("..")

	err := godotenv.Load(Dotenv)
	if err != nil {
		return "", err
	}

	creds, err := DefaultCredentials()
	if err != nil {
		return "", err
	}

	// connect to the admin database to create a new one

	creds.DatabaseName = "postgres"

	err = Connect(creds)
	if err != nil {
		return "", err
	}

	_, err = conn.Exec("CREATE DATABASE " + testdb)
	if err != nil {
		return "", err
	}

	// update the connection to use the database created on the previous step

	conn.Close()

	creds.DatabaseName = testdb

	err = Connect(creds)
	if err != nil {
		return "", err
	}

	migrate, err := os.ReadFile(migrateSqlFile)
	if err != nil {
		return "", err
	}

	_, err = conn.Exec(string(migrate))
	if err != nil {
		return "", err
	}

	return testdb, nil
}
