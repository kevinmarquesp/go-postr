package services_test

import (
	"testing"

	"github.com/kevinmarquesp/go-postr/internal/repositories"
	"github.com/kevinmarquesp/go-postr/internal/services"
	"github.com/stretchr/testify/assert"
)

type MockedUserRepository struct {
	repositories.UserRepository
}

func TestNewGopostrAuthentication(t *testing.T) {
	mockUserRepo := MockedUserRepository{}
	authenticationService := services.NewGopostrAuthentication(mockUserRepo)

	assert.IsType(t, services.GopostrAuthentication{}, authenticationService)
}
