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

	// This function will return the public ID of the inserted user and its
	// session token ID (which has an expiration date defined by the
	// models.SESSION_MAX_DURATION constant)
	RegisterNewUser(form RegisterForm) (string, string, error)

	AuthorizeUserWithSessionToken(sessionToken string) (string, error)
	AuthorizeUserWithCredentials(username, password string) (string, error)
}

type RegisterForm struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}
