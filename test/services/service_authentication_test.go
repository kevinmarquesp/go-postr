package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/kevinmarquesp/go-postr/internal/repositories"
	"github.com/kevinmarquesp/go-postr/internal/services"
	"github.com/stretchr/testify/assert"
)

type MockedUserRepository struct {
	repositories.UserRepository

	MockFailCreateNewUser bool
}

func (mur MockedUserRepository) CreateNewUser(id, name, username, email, password string) (repositories.UserSchema, error) {
	empty := repositories.UserSchema{}

	if mur.MockFailCreateNewUser {
		return empty, errors.New("System under test: Forced fail.")
	}

	return repositories.UserSchema{
		Id:        id,
		Name:      name,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// The test units starts here!

func TestNewGopostrAuthentication(t *testing.T) {
	mockUserRepo := MockedUserRepository{}
	authenticationService := services.NewGopostrAuthenticationService(mockUserRepo)

	assert.IsType(t, &services.GopostrAuthenticationService{}, authenticationService)
}

func TestAuthenticateWithCredentials(t *testing.T) {
	t.Run("Should fail with invalid credentials.", func(t *testing.T) {
		mockUserRepo := MockedUserRepository{}
		authenticationService := services.NewGopostrAuthenticationService(mockUserRepo)
		_, err := authenticationService.AuthenticateWithCredentials("Fulano", "user", "email.com", "password", "password")
		assert.Error(t, err)
	})

	t.Run("Should fail with wrong passwords.", func(t *testing.T) {
		mockUserRepo := MockedUserRepository{}
		authenticationService := services.NewGopostrAuthenticationService(mockUserRepo)
		_, err := authenticationService.AuthenticateWithCredentials("Fulano de Tal", "fulano", "me@email.com", "Password123!", "Password123$")
		assert.Error(t, err)
	})

	t.Run("Should create a user with a hashed password with bcrypt.", func(t *testing.T) {
		mockUserRepo := MockedUserRepository{}
		authenticationService := services.NewGopostrAuthenticationService(mockUserRepo)
		createdUser, err := authenticationService.AuthenticateWithCredentials("Fulano de Tal", "fulano", "me@email.com", "Password123!", "Password123!")
		assert.NoError(t, err)

		assert.Len(t, createdUser.Password, 60)
	})

	t.Run("Should generate a valid ULID with 26 characters for the ID field.", func(t *testing.T) {
		mockUserRepo := MockedUserRepository{}
		authenticationService := services.NewGopostrAuthenticationService(mockUserRepo)
		createdUser, err := authenticationService.AuthenticateWithCredentials("Fulano de Tal", "fulano", "me@email.com", "Password123!", "Password123!")
		assert.NoError(t, err)

		assert.Len(t, createdUser.Id, 26)
	})
}
