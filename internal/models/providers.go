package models

type DatabaseProvider interface {
	Connect(url string) error
	RegisterNewUser(username string, password string) (string, error)
}
