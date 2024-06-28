package models

type GenericDatabaseProvider interface {
	Connect(url string) error
	RegisterNewUser(username string, password string) (string, error)
	AuthorizeUserWithSessionToken(username string, sessionToken string) (string, error)
}
