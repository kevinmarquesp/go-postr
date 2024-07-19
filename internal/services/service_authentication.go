package services

import "github.com/kevinmarquesp/go-postr/internal/repositories"

type GopostrAuthentication struct {
	AuthenticationService

	UserRepo repositories.UserRepository
}

func NewGopostrAuthentication(userRepo repositories.UserRepository) GopostrAuthentication {
	return GopostrAuthentication{
		UserRepo: userRepo,
	}
}

func (ga GopostrAuthentication) AuthenticateWithCredentials(name, username, email, password, confirmation string) (repositories.UserSchema, error) {
	empty := repositories.UserSchema{}

	return empty, nil
}
