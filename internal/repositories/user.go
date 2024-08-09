package repositories

import (
	"context"

	"github.com/kevinmarquesp/go-postr/internal/models"
)

type UserRepository interface {
	RegisterWithCredentials(ctx context.Context, props CredentialsPropperties) (models.AccountSchema, error)
	VerifyAccount(ctx context.Context, userId string) (models.AccountSchema, error)
}

type CredentialsPropperties struct {
	Username string
	Email    string
	Password string
	Role     models.Role
}
