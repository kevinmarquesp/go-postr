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
	RegisterNewUser(form RegisterForm) (RegisterResponse, error)
	AuthorizeUserWithSessionToken(sessionToken string) (SessionToken, error)
	AuthorizeUserWithCredentials(username, password string) (SessionToken, error)
}

type RegisterForm struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	PublicId     string `json:"publicId"`
	SessionToken string `json:"sessionToken"`
}

type SessionToken struct {
	SessionToken string `json:"sessionToken"`
}

// Data types that may be useful for the Auth related routes.

type Auth struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	SessionToken string `json:"sessionToken"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
