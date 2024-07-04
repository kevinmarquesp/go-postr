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

	t.Log("Creating a temporary database for mocking tests.")

	db, dbFile, err := models.MockSqlite(TARGET_DIR, MOCK_FILE)
	assert.NoError(t, err)

	defer os.Remove(dbFile)

	t.Log("Defining the users array with credentials details.")

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
		{
			fullname:   "",
			username:   "",
			password:   "Password1234$",
			expectFail: true,
		},
	}

	for _, user := range users {
		testDescription := fmt.Sprintf("fullname:%s; username:%s; password:%s; expectFail:%t",
			user.fullname, user.username, user.password, user.expectFail)

		t.Run(testDescription, func(t *testing.T) {
			t.Log("Try to register a the new user to the database.")

			publicID, userSessionToken, err := db.RegisterNewUser(models.RegisterForm{
				Fullname: user.fullname,
				Username: user.username,
				Password: user.password,
			})

			if user.expectFail {
				assert.NotNil(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, publicID)
			assert.NotEmpty(t, userSessionToken)

			t.Log("Query the database to verify if the user was inserted with success.")

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

			t.Log("Comparing the selected user details with the provided data.")

			assert.Equal(t, publicID, dbField.publicID)
			assert.Equal(t, user.fullname, dbField.fullname)
			assert.Equal(t, user.username, dbField.username)
			assert.Equal(t, userSessionToken, dbField.sessionToken)

			t.Log("Verifying the password hash.")

			err = bcrypt.CompareHashAndPassword([]byte(dbField.password),
				[]byte(user.username+user.password))
			assert.NoError(t, err)

			t.Log("Verifying the session expiration date is within the expected range.")

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

	t.Log("Creating a temporary database for mocking tests.")

	db, dbFile, err := models.MockSqlite(TARGET_DIR, MOCK_FILE)
	assert.NoError(t, err)

	t.Log("Insert a John Doe user to the database.")

	fullname := "John Doe"
	username := "johndoe"
	password := "Password123!"

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

	// Many tests may require all this information to run.

	return db, dbFile, fullname, username, password, initSessionToken
}

func TestSqliteAuthorizeUserWithSessionToken(t *testing.T) {
	db, dbFile, _, username, _, sessionToken := mockSqliteWithJohnDoe(t)

	defer os.Remove(dbFile)

	t.Log("Try to authorize with the session token string.")

	newSessionToken, err := db.AuthorizeUserWithSessionToken(sessionToken)
	assert.NoError(t, err)
	assert.NotEmpty(t, newSessionToken)

	t.Log("Query the database to verify if the session_token & session_expires fields were updated.")

	var (
		dbSessionToken   string
		dbSessionExpires time.Time
	)

	const SELECT_QUERY = "SELECT session_token, session_expires FROM users WHERE username = ?"

	err = db.Conn.QueryRow(SELECT_QUERY, username).Scan(&dbSessionToken, &dbSessionExpires)
	assert.NoError(t, err)

	t.Log("Verify if the session token was updated with success.")

	assert.Equal(t, newSessionToken, dbSessionToken)
	assert.NotEqual(t, sessionToken, dbSessionToken)

	t.Log("Verify if the session expiration date is within the expected range.")

	expectedExpiration := time.Now().Add(models.SESSION_MAX_DURATION)
	assert.WithinDuration(t, expectedExpiration, dbSessionExpires, time.Minute)
}

func TestSqliteAuthorizeUserWithSessionTokenFail(t *testing.T) {
	db, dbFile, _, _, _, sessionToken := mockSqliteWithJohnDoe(t)

	defer os.Remove(dbFile)

	t.Log("Should fail with a invalid session token string.")

	newSessiontoken, err := db.AuthorizeUserWithSessionToken("blah-blah-blah-blah-blah")
	assert.NotNil(t, err)
	assert.Empty(t, newSessiontoken)

	t.Log("Should fail with an expired, but still valid, session token string.")

	_, err = db.Conn.Exec("UPDATE users SET session_expires = ?"+
		"WHERE session_token IS ?", time.Now().Add(-1*time.Hour), sessionToken)
	assert.NoError(t, err)

	sessionToken, err = db.AuthorizeUserWithSessionToken(sessionToken)
	assert.NotNil(t, err)
	assert.Empty(t, sessionToken)
}

func TestSqliteAuthorizeUserWithCredentials(t *testing.T) {
	db, dbFile, _, username, password, sessionToken := mockSqliteWithJohnDoe(t)

	defer os.Remove(dbFile)

	t.Log("Try to authorize the user with the credentials.")

	newSessionToken, err := db.AuthorizeUserWithCredentials(username, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, newSessionToken)

	t.Log("Query the database to verify if the session_token & session_expires fields were updated.")

	var (
		dbSessionToken   string
		dbSessionExpires time.Time
	)

	const SELECT_QUERY = "SELECT session_token, session_expires FROM users WHERE username = ?"

	err = db.Conn.QueryRow(SELECT_QUERY, username).Scan(&dbSessionToken, &dbSessionExpires)
	assert.NoError(t, err)

	t.Log("Verify if the session token was updated with success.")

	assert.Equal(t, newSessionToken, dbSessionToken)
	assert.NotEqual(t, sessionToken, dbSessionToken)

	t.Log("Verify if the session expiration date is within the expected range.")

	expectedExpiration := time.Now().Add(models.SESSION_MAX_DURATION)
	assert.WithinDuration(t, expectedExpiration, dbSessionExpires, time.Minute)
}

func TestSqliteAuthorizeUserWithCredentialsFail(t *testing.T) {
	db, dbFile, _, username, password, _ := mockSqliteWithJohnDoe(t)

	defer os.Remove(dbFile)

	t.Log("Should fail with an incorrect username.")

	sessionToken, err := db.AuthorizeUserWithCredentials("NonExisting", password)
	assert.NotNil(t, err)
	assert.Empty(t, sessionToken)

	t.Log("Should fail with an incorrect password.")

	sessionToken, err = db.AuthorizeUserWithCredentials(username, "6607cc3df0ec4abfb2e57f8334ca30e3")
	assert.NotNil(t, err)
	assert.Empty(t, sessionToken)
}
