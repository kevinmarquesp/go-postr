package models

import (
	"database/sql"
	"time"

	"github.com/kevinmarquesp/go-postr/internal/utils"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	BCRYPT_COST     = bcrypt.MinCost
	SESSION_EXPIRES = 30 * time.Second
)

type Sqlite struct {
	conn *sql.DB
}

func (s *Sqlite) Connect(url string) error {
	conn, err := sql.Open("sqlite3", url)
	if err != nil {
		return err
	}

	s.conn = conn

	return nil
}

func (s *Sqlite) RegisterNewUser(username string, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(username+password), BCRYPT_COST)
	if err != nil {
		return "", err
	}

	token := utils.GenerateSessionToken(username)
	expirationDate := time.Now().Add(SESSION_EXPIRES)

	statement, err := s.conn.Prepare(`INSERT INTO users (username, password, session_token, session_expires)
        VALUES (?, ?, ?, ?)`)
	if err != nil {
		return "", err
	}

	_, err = statement.Exec(username, hashedPassword, token, expirationDate)
	if err != nil {
		return "", err
	}

	return token, nil
}
