package models

type DatabaseProvider interface {
	Connect(url string) error
}
