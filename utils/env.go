package utils

import (
	"errors"
	"os"
)

func RequireEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("Required environment variable " + key + " doesn't exist or is empty.")
	}

	return value, nil
}
