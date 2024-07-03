package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/kevinmarquesp/go-postr/internal/utils"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type Sqlite struct {
	GenericDatabaseProvider

	Conn *sql.DB
}

func (s *Sqlite) Connect(url string) error {
	conn, err := sql.Open("sqlite3", url)
	if err != nil {
		return err
	}

	s.Conn = conn

	return nil
}

func (s *Sqlite) RegisterNewUser(fullname, username, password string) (string, string, error) {
	if fullname == "" || username == "" || password == "" {
		return "", "", errors.New(EMPTY_ARGUMENTS_ERROR)
	}

	if err := utils.ValidateUsernameString(username); err != nil {
		return "", "", err
	}

	if err := utils.ValidatePasswordString(password); err != nil {
		return "", "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(username+password), BCRYPT_COST)
	if err != nil {
		return "", "", err
	}

	publicID, err := utils.GenerateTokenID()
	if err != nil {
		return "", "", err
	}

	sessionToken, expirationDate, err := utils.GenerateNewSessionToken(SESSION_MAX_DURATION)
	if err != nil {
		return "", "", err
	}

	statement, err := s.Conn.Prepare(`INSERT INTO users (public_id, fullname, username, password,
        session_token, session_expires) VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return "", "", err
	}

	_, err = statement.Exec(publicID, fullname, username, hashedPassword, sessionToken, expirationDate)
	if err != nil {
		return "", "", err
	}

	return publicID, sessionToken, nil
}

func (s *Sqlite) AuthorizeUserWithSessionToken(sessionToken string) (string, error) {
	if sessionToken == "" {
		return "", errors.New(EMPTY_ARGUMENTS_ERROR)
	}

	newSessionToken, newExpirationDate, err := utils.GenerateNewSessionToken(SESSION_MAX_DURATION)
	if err != nil {
		return "", err
	}

	statement, err := s.Conn.Prepare(`UPDATE users SET session_token = ?, session_expires = ?
        WHERE session_token IS ? AND session_expires > ?`)
	if err != nil {
		return "", err
	}

	rows, err := statement.Exec(newSessionToken, newExpirationDate, sessionToken, time.Now())
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

func (s *Sqlite) AuthorizeUserWithCredentials(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New(EMPTY_ARGUMENTS_ERROR)
	}

	if err := s.comparePassword(username, password); err != nil {
		return "", err
	}

	newSessionToken, newExpirationDate, err := utils.GenerateNewSessionToken(SESSION_MAX_DURATION)
	if err != nil {
		return "", err
	}

	statement, err := s.Conn.Prepare(`UPDATE users SET session_token = ?, session_expires = ?
        WHERE username IS ?`)
	if err != nil {
		return "", err
	}

	rows, err := statement.Exec(newSessionToken, newExpirationDate, username)
	if err != nil {
		return "", err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected < 1 {
		return "", errors.New(INVALID_AUTH_CREDENTIALS_ERROR)
	}

	return newSessionToken, nil
}

func (s *Sqlite) comparePassword(username, password string) error {
	const SELECT_QUERY = "SELECT password FROM users WHERE username IS ?"

	var hashedPassword string

	if err := s.Conn.QueryRow(SELECT_QUERY, username).Scan(&hashedPassword); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(username+password)); err != nil {
		return err
	}

	return nil
}
