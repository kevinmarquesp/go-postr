package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/kevinmarquesp/go-postr/internal/utils"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	BCRYPT_COST                          = bcrypt.MinCost
	SESSION_EXPIRES                      = 10 * time.Second
	CANNOT_MATCH_TOKEN_TO_USERNAME_ERROR = "invalid token for username or session expired"
)

type Sqlite struct {
	GenericDatabaseProvider

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

func (s *Sqlite) AuthorizeUserWithSessionToken(username string, sessionToken string) (string, error) {
	newSessionToken := utils.GenerateSessionToken(username)
	newExpirationDate := time.Now().Add(SESSION_EXPIRES)

	statement, err := s.conn.Prepare(`UPDATE users SET session_token = ?, session_expires = ?
        WHERE username IS ? AND session_token IS ? AND session_expires > ?`)
	if err != nil {
		return "", err
	}

	rows, err := statement.Exec(newSessionToken, newExpirationDate, username, sessionToken, time.Now())
	if err != nil {
		return "", err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected < 1 {
		return "", errors.New(CANNOT_MATCH_TOKEN_TO_USERNAME_ERROR)
	}

	return newSessionToken, nil
}
