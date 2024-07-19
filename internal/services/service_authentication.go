package services

import (
	"errors"
	"os"
	"strings"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type GopostrAuthenticationService struct {
	AuthenticationService

	UserRepo repositories.UserRepository
}

func NewGopostrAuthenticationService(userRepo repositories.UserRepository) GopostrAuthenticationService {
	return GopostrAuthenticationService{
		UserRepo: userRepo,
	}
}

func (ga GopostrAuthenticationService) AuthenticateWithCredentials(name, username, email, password, confirmation string) (repositories.UserSchema, error) {
	empty := repositories.UserSchema{}

	// Validate parameters request.

	name, err := models.ValidateName(name)
	if err != nil {
		return empty, err
	}

	username, err = models.ValidateUsername(username)
	if err != nil {
		return empty, err
	}

	email, err = models.ValidateEmail(email)
	if err != nil {
		return empty, err
	}

	password, err = models.ValidatePassword(password)
	if err != nil {
		return empty, err
	}

	if strings.Trim(confirmation, " ") != password {
		return empty, errors.New("Passwords did not match.")
	}

	// Hash the password before inserting.

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return empty, err
	}

	// Insert a new user to the repository & return the new user schema.

	user, err := ga.UserRepo.CreateNewUser(name, username, email, hashedPassword)
	if err != nil {
		return empty, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	cost := bcrypt.MinCost
	secret := os.Getenv("PASSWORD_SECRET")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+secret), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
