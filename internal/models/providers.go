package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	BCRYPT_COST          = bcrypt.MinCost
	SESSION_MAX_DURATION = 10 * time.Second

	CANNOT_MATCH_TOKEN_TO_USERNAME_ERROR = "invalid token for username or session expired"
	INVALID_AUTH_CREDENTIALS_ERROR       = "invalid username and password credentials"
)

type GenericDatabaseProvider interface {
	Connect(url string) error
	RegisterNewUser(fullname, username, password string) (string, string, error)
	AuthorizeUserWithSessionToken(username, sessionToken string) (string, error)
	AuthorizeUserWithCredentials(username, password string) (string, error)
}
