package utils

import (
	"time"

	"github.com/google/uuid"
)

// This function utilizes the github.com/google/uuid package to generate a
// version 7 UUID, which is then returned as a string.
//
// Example usage:
//
//	tokenID, err := utils.GenerateTokenID()
//	if err != nil {
//	    log.Fatalf("Failed to generate token ID: %v", err)
//	}
//
//	fmt.Println("Generated Token ID:", tokenID)
//
// The returned token ID will be a unique string that can be used for
// identifying sessions, users, etc.
func GenerateTokenID() (string, error) {
	token, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return token.String(), nil
}

// This function generates a new token ID using utils.GenerateTokenID and calculates
// the expiration time by adding the given session duration to the current time.
//
// Returns:
//   - A string representing the generated session token.
//   - A time.Time value representing the expiration time of the session.
//   - An error if there is an issue generating the session token.
//
// Example usage:
//
//	const SESSION_DURATION = 24 * time.Hour
//
//	sessionToken, expiryTime, err := utils.GenerateNewSessionToken(SESSION_DURATION)
//	if err != nil {
//	    log.Fatalf("Failed to generate session token: %v", err)
//	}
//
//	fmt.Printf("Generated Session Token: %s, Expires at: %s\n", sessionToken, expiryTime)
//
// The generated session token is a unique string, and the expiry time is calculated
// based on the current time plus the provided session duration.
func GenerateNewSessionToken(sessionDuration time.Duration) (string, time.Time, error) {
	newSessionToken, err := GenerateTokenID()
	if err != nil {
		return "", time.Time{}, err
	}

	return newSessionToken, time.Now().Add(sessionDuration), nil
}
