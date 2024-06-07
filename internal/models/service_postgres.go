package models

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
)

// PostgreSQL service (used in production).
type Postgres struct {
	conn *sql.DB
}

func (pg *Postgres) Connect() error {
	log.Info("Connecting to the Postgres database service, credentials given by the environment variables.")

	credentials, err := GetPostgresCredentials()
	if err != nil {
		return err
	}

	credentials_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		credentials.Host, credentials.Port, credentials.Username, credentials.Password, credentials.Database, credentials.SSLMode)

	pg.conn, err = sql.Open("postgres", credentials_string)
	if err != nil {
		return err
	}

	// Connection pooling settings.
	pg.conn.SetMaxOpenConns(25)
	pg.conn.SetMaxIdleConns(25)
	pg.conn.SetConnMaxLifetime(0)

	return pg.conn.Ping()
}
