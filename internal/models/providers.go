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

// GenericDatabaseProvider defines the interface for a generic database provider.
//
// This interface outlines the required methods for connecting to a database,
// registering new users, and authorizing users using either session tokens or
// username and password credentials. Implementations of this interface should
// provide concrete logic for each method based on the specific database being used.
type GenericDatabaseProvider interface {
	// Establishes a connection to the database using the provided URL.
	Connect(url string) error

	// Registers a new user in the database with the provided details.
	//
	// Returns:
	//   - A string representing the public ID of the new user.
	//   - A string representing the session token for the new user.
	//   - An error if the registration fails, or nil if successful.
	RegisterNewUser(fullname, username, password string) (string, string, error)

	// Authorizes a user based on the provided session token, extending the session duration.
	AuthorizeUserWithSessionToken(sessionToken string) (string, error)

	// Authorizes a user based on the provided username and password, issuing a new session token.
	AuthorizeUserWithCredentials(username, password string) (string, error)
}
