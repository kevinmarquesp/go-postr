package models

type ValidateUsernameResponse struct {
	IsEmpty         bool
	BootstrapStatus string
	InfoMessage     string
}
