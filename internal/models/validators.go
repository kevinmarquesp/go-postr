package models

import (
	"errors"
	"regexp"
)

func ValidateName(name string) error {
	return nil
}

func ValidateUsername(username string) error {
	if len(username) < 1 {
		return errors.New("Empty username not allowed.")

	} else if regexp.MustCompile(`\s`).Match([]byte(username)) {
		return errors.New("Username cannot have spaces included.")

	} else if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(username) {
		return errors.New("Username could not have special characters other than - and _.")
	}

	return nil
}

func ValidateEmail(email string) error {
	const regexMatchEmailString = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

	if len(email) < 1 {
		return errors.New("Empty email not allowed.")

	} else if !regexp.MustCompile(regexMatchEmailString).MatchString(email) {
		return errors.New("Empty email not allowed.")
	}

	return nil
}

func ValidatePassword(password string) error {
	length := len(password)

	if length == 0 {
		return errors.New("Empty password not allowed.")

	} else if length < 12 {
		return errors.New("Password is too short.")

	} else if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("The password should include uppercase letters.")

	} else if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("The password should include lowercase letters.")

	} else if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("The password should include numbers.")

	} else if !regexp.MustCompile(`[!@#$%^&*()_+{}[\]|:;<>,.?/\\]`).MatchString(password) {
		return errors.New("The password should include special characters.")
	}

	return nil
}
