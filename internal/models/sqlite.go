package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
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

	token := GenerateSessionToken(username)
	expirationDate := time.Now().Add(SESSION_EXPIRES)

	_, err = s.conn.Query(`INSERT INTO users (username, password, session_token, session_expires)
        VALUES ($1, $2, $3, $4)`, username, hashedPassword, token, expirationDate)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateSessionToken(username string) string {
	shaAlgorithm := sha256.New()

	shaAlgorithm.Write([]byte(username))

	hashedUsername := hex.EncodeToString(shaAlgorithm.Sum(nil))
	tokenId := uuid.New().String()

	return hashedUsername[:12] + "." + tokenId
}
