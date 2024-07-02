package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	BCRYPT_COST          = bcrypt.MinCost
	SESSION_MAX_DURATION = 10 * time.Second

	FULLNAME_REGITER_FIELD_NOT_SPECIFIED_ERROR = "a fullname should be provided to register the user"

	CANNOT_MATCH_TOKEN_TO_USERNAME_ERROR = "invalid token for username or session expired"
	INVALID_AUTH_CREDENTIALS_ERROR       = "invalid username and password credentials"
)

type GenericDatabaseProvider interface {
	Connect(url string) error
	RegisterNewUser(fullname, username, password string) (string, string, error)
	AuthorizeUserWithSessionToken(sessionToken string) (string, error)
	AuthorizeUserWithCredentials(username, password string) (string, error)
}
