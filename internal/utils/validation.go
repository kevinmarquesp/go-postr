package utils

import (
	"errors"
	"regexp"
)

const (
	USERNAME_IS_EMPTY_ERROR               = "unexpected empty username string"
	USERNAME_CONTAINS_SPACES_ERROR        = "username string has spaces"
	USERNAME_HAS_INVALID_CHARS_ERROR      = "username has invalid characters"
	PASSWORD_TOO_SHORT_ERROR              = "password length is too short"
	PASSWORD_IS_EMPTY_ERROR               = "unexpected empty password string"
	PASSWORD_DONT_HAVE_ANY_UPERS_ERROR    = "password should include at least one uppercase character"
	PASSWORD_DONT_HAVE_ANY_LOWERS_ERROR   = "password should include at least one lowercase character"
	PASSWORD_DONT_HAVE_ANY_NUMBERS_ERROR  = "password should include at least one number"
	PASSWORD_DONT_HAVE_ANY_SPECIALS_ERROR = "password should include at least one special character"
)

func ValidateUsernameString(username string) error {
	if len(username) < 1 {
		return errors.New(USERNAME_IS_EMPTY_ERROR)
	}

	if regexp.MustCompile(`\s`).Match([]byte(username)) {
		return errors.New(USERNAME_CONTAINS_SPACES_ERROR)
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(username) {
		return errors.New(USERNAME_HAS_INVALID_CHARS_ERROR)
	}

	return nil
}

func ValidatePasswordString(password string) error {
	lenght := len(password)

	if lenght == 0 {
		return errors.New(PASSWORD_IS_EMPTY_ERROR)
	}

	if lenght < 12 {
		return errors.New(PASSWORD_TOO_SHORT_ERROR)
	}

	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New(PASSWORD_DONT_HAVE_ANY_UPERS_ERROR)
	}

	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New(PASSWORD_DONT_HAVE_ANY_LOWERS_ERROR)
	}

	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New(PASSWORD_DONT_HAVE_ANY_NUMBERS_ERROR)
	}

	if !regexp.MustCompile(`[!@#$%^&*()_+{}[\]|:;<>,.?/\\]`).MatchString(password) {
		return errors.New(PASSWORD_DONT_HAVE_ANY_SPECIALS_ERROR)
	}

	return nil
}
