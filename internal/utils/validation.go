package utils

import (
	"errors"
	"regexp"
)

const (
	USERNAME_IS_EMPTY_ERROR          = "unexpected empty username string"
	USERNAME_CONTAINS_SPACES_ERROR   = "username string has spaces"
	USERNAME_HAS_INVALID_CHARS_ERROR = "username has invalid characters"

	PASSWORD_TOO_SHORT_ERROR              = "password length is too short"
	PASSWORD_IS_EMPTY_ERROR               = "unexpected empty password string"
	PASSWORD_DONT_HAVE_ANY_UPERS_ERROR    = "password should include at least one uppercase character"
	PASSWORD_DONT_HAVE_ANY_LOWERS_ERROR   = "password should include at least one lowercase character"
	PASSWORD_DONT_HAVE_ANY_NUMBERS_ERROR  = "password should include at least one number"
	PASSWORD_DONT_HAVE_ANY_SPECIALS_ERROR = "password should include at least one special character"
)

// This function checks if the username is not empty, does not contain spaces,
// and consists only of valid characters (letters, numbers, underscores, and hyphens).
//
// Example usage:
//
//	if err := utils.ValidateUsernameString("valid_username"); err != nil {
//	    fmt.Println("Username validation error:", err)
//
//	} else {
//	    fmt.Println("Username is valid")
//	}
//
// The function returns specific error messages defined in constants ended with
// the specific "_ERROR" sufix for different validation failures.
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

// This function checks if the password is not empty, has a minimum length of 12 characters,
// and includes at least one uppercase letter, one lowercase letter, one number, and one special character.
//
// Example usage:
//
//	if err := utils.ValidatePasswordString("Valid@12345Password"); err != nil {
//	    fmt.Println("Password validation error:", err)
//
//	} else {
//	    fmt.Println("Password is valid")
//	}
//
// The function returns specific error messages defined in constants for different validation failures.
func ValidatePasswordString(password string) error {
	length := len(password)

	if length == 0 {
		return errors.New(PASSWORD_IS_EMPTY_ERROR)
	}

	if length < 12 {
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
