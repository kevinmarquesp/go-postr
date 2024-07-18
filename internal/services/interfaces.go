package services

import "github.com/kevinmarquesp/go-postr/internal/repositories"

type AuthenticationService interface {
	AuthenticateWithCredentials(name, username, email, password, confirmation string) (repositories.UserSchema, error)
}

type AuthorizationService interface {
	AuthorizeCredentials(email, password string) (repositories.UserSchema, error)
}
