package utils

import (
	"errors"
	"os"
)

// Wrapper for the `os.Getenv()` function, but it returns an error if the
// variable value is an empty string (it happens when the variable doesn't
// exists).
func ProtectedGetenv(varname string) (string, error) {
	value := os.Getenv(varname)
	if value == "" {
		return "", errors.New("Environment variable " + varname + " is empty or doesn't exists")
	}

	return value, nil
}
