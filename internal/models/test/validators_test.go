package models_test

import (
	"testing"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{"Empty username", "", true},
		{"Username with spaces", "user name", true},
		{"Username with special characters", "user@name", true},
		{"Valid username", "username", false},
		{"Valid username with dash", "user-name", false},
		{"Valid username with underscore", "user_name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := models.ValidateUsername(tt.username)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"Empty email", "", true},
		{"Invalid email format no @", "userexample.com", true},
		{"Invalid email format no domain", "user@", true},
		{"Invalid email format multiple @", "user@@example.com", true},
		{"Valid email", "user@example.com", false},
		{"Valid email with subdomain", "user@mail.example.com", false},
		{"Valid email with +", "user+mail@example.com", false},
		{"Valid email with . in local part", "user.name@example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := models.ValidateEmail(tt.email)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"Empty password", "", true},
		{"Too short password", "Short1!", true},
		{"No uppercase letter", "lowercase1!", true},
		{"No lowercase letter", "UPPERCASE1!", true},
		{"No number", "NoNumbers!", true},
		{"No special character", "NoSpecial1", true},
		{"Valid password", "ValidPassword1!", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := models.ValidatePassword(tt.password)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
		})
	}
}
