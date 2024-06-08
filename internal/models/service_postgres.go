package models

import (
	"database/sql"
	"fmt"

	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// PostgreSQL service (used in production).
type Postgres struct {
	conn        *sql.DB
	bcrypt_cost int
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

	// Bcrypt encription settings.
	pg.bcrypt_cost = bcrypt.MinCost

	return pg.conn.Ping()
}

func (pg *Postgres) InsertUser(username, password, description string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(username+password), pg.bcrypt_cost)
	if err != nil {
		return err
	}

	_, err = pg.conn.Query("INSERT INTO users (username, password, description) VALUES ($1, $2, $3)",
		username, hashed_password, description)

	return err
}

func (pg *Postgres) RecentlyCreatedUsers(size int) ([]UserBasicInfo, error) {
	var users_result []UserBasicInfo

	rows, err := pg.conn.Query("SELECT username, description FROM users ORDER BY created_at DESC LIMIT $1", size)
	if err != nil {
		return []UserBasicInfo{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserBasicInfo

		if err = rows.Scan(&user.Username, &user.Description); err != nil {
			return []UserBasicInfo{}, err
		}

		users_result = append(users_result, user)

	}

	return users_result, nil
}
