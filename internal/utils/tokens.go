package utils

import (
	"time"

	"github.com/google/uuid"
)

func GenerateTokenID() (string, error) {
	token, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return token.String(), nil
}

func GenerateNewSessionToken(sessionDuration time.Duration) (string, time.Time, error) {
	newSessionToken, err := GenerateTokenID()
	if err != nil {
		return "", time.Time{}, err
	}

	return newSessionToken, time.Now().Add(sessionDuration), nil
}
