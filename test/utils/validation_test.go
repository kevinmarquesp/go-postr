package utils_test

import (
	"errors"
	"testing"

	"github.com/kevinmarquesp/go-postr/internal/utils"
)

func TestValidateUsernameString(t *testing.T) {
	tests := []struct {
		username string
		expected error
	}{
		{"", errors.New(utils.USERNAME_IS_EMPTY_ERROR)},
		{"user name", errors.New(utils.USERNAME_CONTAINS_SPACES_ERROR)},
		{"user@name", errors.New(utils.USERNAME_HAS_INVALID_CHARS_ERROR)},
		{"validUsername123", nil},
	}

	for _, test := range tests {
		t.Run(test.username, func(t *testing.T) {
			err := utils.ValidateUsernameString(test.username)
			if err != nil && err.Error() != test.expected.Error() {
				t.Errorf("Expected error: %v, got: %v", test.expected, err)
			}
		})
	}
}

func TestValidatePasswordString(t *testing.T) {
	tests := []struct {
		password string
		expected error
	}{
		{"", errors.New(utils.PASSWORD_IS_EMPTY_ERROR)},
		{"short", errors.New(utils.PASSWORD_TOO_SHORT_ERROR)},
		{"NoNumbersOrSpecials", errors.New(utils.PASSWORD_DONT_HAVE_ANY_NUMBERS_ERROR)},
		{"NoSpecials123", errors.New(utils.PASSWORD_DONT_HAVE_ANY_SPECIALS_ERROR)},
		{"ValidPassword123!", nil},
	}

	for _, test := range tests {
		t.Run(test.password, func(t *testing.T) {
			err := utils.ValidatePasswordString(test.password)
			if err != nil && err.Error() != test.expected.Error() {
				t.Errorf("Expected error: %v, got: %v", test.expected, err)
			}
		})
	}
}
