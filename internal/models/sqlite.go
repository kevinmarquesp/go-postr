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

func (s *Sqlite) RegisterNewUser(form RegisterForm) (RegisterResponse, error) {
	if err := utils.ValidateUsernameString(form.Username); err != nil {
		return RegisterResponse{}, err
	}

	if err := utils.ValidatePasswordString(form.Password); err != nil {
		return RegisterResponse{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Username+form.Password), BCRYPT_COST)
	if err != nil {
		return RegisterResponse{}, err
	}

	publicId, err := utils.GenerateTokenID()
	if err != nil {
		return RegisterResponse{}, err
	}

	sessionToken, expirationDate, err := utils.GenerateNewSessionToken(SESSION_MAX_DURATION)
	if err != nil {
		return RegisterResponse{}, err
	}

	statement, err := s.Conn.Prepare(`INSERT
        INTO users
            (public_id, fullname, username, password, session_token, session_expires)
        VALUES
            (?1, ?2, ?3, ?4, ?5, ?6)`)
	if err != nil {
		return RegisterResponse{}, err
	}

	_, err = statement.Exec(publicId, form.Fullname, form.Username,
		hashedPassword, sessionToken, expirationDate)
	if err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{
		Fullname:     form.Fullname,
		Username:     form.Username,
		PublicId:     publicId,
		SessionToken: sessionToken,
	}, nil
}

func (s *Sqlite) AuthorizeUserWithSessionToken(sessionToken string) (string, error) {
	newSessionToken, newExpirationDate, err := utils.GenerateNewSessionToken(SESSION_MAX_DURATION)
	if err != nil {
		return "", err
	}

	statement, err := s.Conn.Prepare(`UPDATE users
        SET
            session_token   = ?1,
            session_expires = ?2,
            updated_at      = CURRENT_TIMESTAMP
        WHERE
            session_token IS ?3
            AND session_expires > ?4`)
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
	if err := s.comparePassword(username, password); err != nil {
		return "", err
	}

	newSessionToken, newExpirationDate, err := utils.GenerateNewSessionToken(SESSION_MAX_DURATION)
	if err != nil {
		return "", err
	}

	statement, err := s.Conn.Prepare(`UPDATE users
        SET
            session_token   = ?1,
            session_expires = ?2,
            updated_at      = CURRENT_TIMESTAMP
        WHERE
            username IS ?3`)
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
	const SELECT_QUERY = `SELECT password
        FROM
            users
        WHERE
            username IS ?`

	var hashedPassword string

	if err := s.Conn.QueryRow(SELECT_QUERY, username).Scan(&hashedPassword); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(username+password)); err != nil {
		return err
	}

	return nil
}
