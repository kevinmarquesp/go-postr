package models_test

import (
	"fmt"
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

	// Define users details

	users := []struct {
		fullname   string
		username   string
		password   string
		expectFail bool
	}{
		{
			fullname:   "John Doe",
			username:   "johndoe",
			password:   "Password1234$",
			expectFail: false,
		},
		{
			fullname:   "Fulano de Tal",
			username:   "fulano",
			password:   "Password1234$",
			expectFail: false,
		},
		{
			fullname:   "Invalid Username",
			username:   "^^ Invalid >w<",
			password:   "Password1234$",
			expectFail: true,
		},
		{
			fullname:   "",
			username:   "invalid_fullname",
			password:   "Password1234$",
			expectFail: true,
		},
		{
			fullname:   "Invalid Password",
			username:   "invalid_password",
			password:   "d6873fde9e01402d8c403b876350d911",
			expectFail: true,
		},
		{
			fullname:   "John Doe",
			username:   "johndoe",
			password:   "Password1234$",
			expectFail: true, // user already exists
		},
	}

	for _, user := range users {
		testDescription := fmt.Sprintf("fullname:%s; username:%s; password:%s; expectFail:%t",
			user.fullname, user.username, user.password, user.expectFail)

		t.Run(testDescription, func(t *testing.T) {
			// Try to register the user

			publicID, userSessionToken, err := db.RegisterNewUser(user.fullname, user.username, user.password)

			if user.expectFail {
				assert.NotNil(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, publicID)
			assert.NotEmpty(t, userSessionToken)

			// Query the database to verify if the user was inserted

			dbField := struct {
				publicID       string
				fullname       string
				username       string
				password       string
				sessionToken   string
				sessionExpires time.Time
			}{}

			const SELECT_QUERY = "SELECT public_id, fullname, username, password, session_token," +
				"session_expires FROM users WHERE username = ?"

			err = db.Conn.QueryRow(SELECT_QUERY, user.username).Scan(&dbField.publicID,
				&dbField.fullname, &dbField.username, &dbField.password, &dbField.sessionToken,
				&dbField.sessionExpires)
			assert.NoError(t, err)

			// Verify the user details

			assert.Equal(t, publicID, dbField.publicID)
			assert.Equal(t, user.fullname, dbField.fullname)
			assert.Equal(t, user.username, dbField.username)
			assert.Equal(t, userSessionToken, dbField.sessionToken)

			// Verify the passwords

			err = bcrypt.CompareHashAndPassword([]byte(dbField.password),
				[]byte(user.username+user.password))
			assert.NoError(t, err)

			// Verify the session expiration time is within the expected range

			expectedExpiration := time.Now().Add(models.SESSION_MAX_DURATION)
			assert.WithinDuration(t, expectedExpiration, dbField.sessionExpires, time.Minute)
		})
	}
}

func mockSqliteWithJohnDoe(t *testing.T) (models.Sqlite, string, string, string, string, string) {
	const (
		TARGET_DIR = "../../tmp/"
		MOCK_FILE  = "../../db/sqlite3/mock_setup.sql"
	)

	// Create a temporary database for mocking tests

	db, dbFile, err := models.MockSqlite(TARGET_DIR, MOCK_FILE)
	assert.NoError(t, err)

	// Insert a user into the database

	fullname := "John Doe"
	username := "johndoe"
	password := "password123"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(username+password), bcrypt.DefaultCost)
	assert.NoError(t, err)

	publicID, err := utils.GenerateTokenID()
	assert.NoError(t, err)

	initSessionToken, initSessionExpires, err := utils.GenerateNewSessionToken(time.Second)

	const INSERT_QUERY = "INSERT INTO users (public_id, fullname, username, password," +
		"session_token, session_expires) VALUES (?, ?, ?, ?, ?, ?)"

	_, err = db.Conn.Exec(INSERT_QUERY, publicID, fullname, username, hashedPassword,
		initSessionToken, initSessionExpires)
	assert.NoError(t, err)

	return db, dbFile, fullname, username, password, initSessionToken
}

func TestSqliteAuthorizeUserWithSessionToken(t *testing.T) {
	db, dbFile, _, username, _, sessionToken := mockSqliteWithJohnDoe(t)

	defer os.Remove(dbFile)

	// Authorize user with the session token

	newSessionToken, err := db.AuthorizeUserWithSessionToken(sessionToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, newSessionToken)

	// Query the database to verify the session_token and session_expires were updated

	var (
		dbSessionToken   string
		dbSessionExpires time.Time
	)

	const SELECT_QUERY = "SELECT session_token, session_expires FROM users WHERE username = ?"

	err = db.Conn.QueryRow(SELECT_QUERY, username).Scan(&dbSessionToken, &dbSessionExpires)
	assert.NoError(t, err)

	// Verify the session token

	assert.Equal(t, newSessionToken, dbSessionToken)

	// Verify the session expiration time is within the expected range

	expectedExpiration := time.Now().Add(models.SESSION_MAX_DURATION)

	assert.WithinDuration(t, expectedExpiration, dbSessionExpires, time.Minute)
}

func TestSqliteAuthorizeUserWithSessionTokenFail(t *testing.T) {
	db, dbFile, _, _, _, sessionToken := mockSqliteWithJohnDoe(t)

	defer os.Remove(dbFile)

	// Fail with invalid session token

	newSessiontoken, err := db.AuthorizeUserWithSessionToken("blah-blah-blah-blah-blah")
	assert.NotNil(t, err)
	assert.Empty(t, newSessiontoken)

	// Fail with an expired, but valid, session token

	_, err = db.Conn.Exec("UPDATE users SET session_expires = ?"+
		"WHERE session_token IS ?", time.Now().Add(-1*time.Hour), sessionToken)
	assert.NoError(t, err)

	sessionToken, err = db.AuthorizeUserWithSessionToken(sessionToken)
	assert.NotNil(t, err)
	assert.Empty(t, sessionToken)
}

func TestSqliteAuthorizeUserWithCredentials(t *testing.T) {
	db, dbFile, _, username, password, _ := mockSqliteWithJohnDoe(t)

	defer os.Remove(dbFile)

	// Authorize user with credentials

	sessionToken, err := db.AuthorizeUserWithCredentials(username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, sessionToken)

	// Query the database to verify the session_token and session_expires were updated

	var (
		dbSessionToken   string
		dbSessionExpires time.Time
	)

	const SELECT_QUERY = "SELECT session_token, session_expires FROM users WHERE username = ?"

	err = db.Conn.QueryRow(SELECT_QUERY, username).Scan(&dbSessionToken, &dbSessionExpires)
	assert.NoError(t, err)

	// Verify the session token

	assert.Equal(t, sessionToken, dbSessionToken)

	// Verify the session expiration time is within the expected range

	expectedExpiration := time.Now().Add(models.SESSION_MAX_DURATION)

	assert.WithinDuration(t, expectedExpiration, dbSessionExpires, time.Minute)
}

func TestSqliteAuthorizeUserWithCredentialsFail(t *testing.T) {
	db, dbFile, _, username, password, _ := mockSqliteWithJohnDoe(t)

	defer os.Remove(dbFile)

	// Fail with incorrect username

	sessionToken, err := db.AuthorizeUserWithCredentials("NonExisting", password)
	assert.NotNil(t, err)
	assert.Empty(t, sessionToken)

	// Fail with incorrect password

	sessionToken, err = db.AuthorizeUserWithCredentials(username, "6607cc3df0ec4abfb2e57f8334ca30e3")
	assert.NotNil(t, err)
	assert.Empty(t, sessionToken)
}
