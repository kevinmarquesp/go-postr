package models

// TODO: Update this methods to follow the SQLite3 provider todos comments...

type GenericDatabaseProvider interface {
	Connect(url string) error
	RegisterNewUser(fullname, username, password string) (string, string, error)
	AuthorizeUserWithSessionToken(username, sessionToken string) (string, error)
	AuthorizeUserWithCredentials(username, password string) (string, error)
}
