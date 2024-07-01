package models_test

import (
	"os"
	"testing"
	"time"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSqliteRegisterUser(t *testing.T) {
	const (
		TARGET_DIR = "../../tmp/"
		MOCK_FILE  = "../../db/sqlite3/mock_setup.sql"
	)

	// Create a temporary database for mocking tests

	db, dbFile, err := models.MockSqlite(TARGET_DIR, MOCK_FILE)
	assert.NoError(t, err)

	defer os.Remove(dbFile)

	// Define user details
	fullname := "John Doe"
	username := "johndoe"
	password := "password123"

	// Register new user
	publicID, sessionToken, err := db.RegisterNewUser(fullname, username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, publicID)
	assert.NotEmpty(t, sessionToken)

	// Query the database to verify the user was inserted
	var (
		dbPublicID       string
		dbFullname       string
		dbUsername       string
		dbPassword       string
		dbSessionToken   string
		dbSessionExpires time.Time
	)
	query := `SELECT public_id, fullname, username, password, session_token, session_expires FROM users WHERE username = ?`
	err = db.Conn.QueryRow(query, username).Scan(&dbPublicID, &dbFullname, &dbUsername, &dbPassword, &dbSessionToken, &dbSessionExpires)
	assert.NoError(t, err)

	// Verify the user details
	assert.Equal(t, publicID, dbPublicID)
	assert.Equal(t, fullname, dbFullname)
	assert.Equal(t, username, dbUsername)
	assert.Equal(t, sessionToken, dbSessionToken)

	// Verify the password
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(username+password))
	assert.NoError(t, err)

	// Verify the session expiration time is within the expected range
	expectedExpiration := time.Now().Add(models.SESSION_MAX_DURATION)
	assert.WithinDuration(t, expectedExpiration, dbSessionExpires, time.Minute)
}

func TestSqliteAuthorizeUserWithCredentials(t *testing.T) {
	const (
		TARGET_DIR = "../../tmp/"
		MOCK_FILE  = "../../db/sqlite3/mock_setup.sql"
	)

	// Create a temporary database for mocking tests

	db, dbFile, err := models.MockSqlite(TARGET_DIR, MOCK_FILE)
	assert.NoError(t, err)

	defer os.Remove(dbFile)

	// Insert a user into the database
	fullname := "John Doe"
	username := "johndoe"
	password := "password123"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(username+password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	publicID, err := utils.GenerateTokenID()
	assert.NoError(t, err)

	initSessionToken, initSessionExpires, err := utils.GenerateNewSessionToken(time.Hour)

	insertStmt := `INSERT INTO users (public_id, fullname, username, password, session_token,
        session_expires) VALUES (?, ?, ?, ?, ?, ?)`

	_, err = db.Conn.Exec(insertStmt, publicID, fullname, username, hashedPassword,
		initSessionToken, initSessionExpires)
	assert.NoError(t, err)

	// Authorize user with credentials
	sessionToken, err := db.AuthorizeUserWithCredentials(username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, sessionToken)

	// Query the database to verify the session_token and session_expires were updated
	var dbSessionToken string
	var dbSessionExpires time.Time
	query := `SELECT session_token, session_expires FROM users WHERE username = ?`
	err = db.Conn.QueryRow(query, username).Scan(&dbSessionToken, &dbSessionExpires)
	assert.NoError(t, err)

	// Verify the session token
	assert.Equal(t, sessionToken, dbSessionToken)

	// Verify the session expiration time is within the expected range
	expectedExpiration := time.Now().Add(models.SESSION_MAX_DURATION)
	assert.WithinDuration(t, expectedExpiration, dbSessionExpires, time.Minute)
}
