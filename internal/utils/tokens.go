package utils

import "github.com/google/uuid"

func GenerateTokenID() (string, error) {
	token, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return token.String(), nil
}
