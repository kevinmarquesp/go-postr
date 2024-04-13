package db

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func InsertNewUser(username string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(username+password), bcrypt.MinCost)
	bio := "Hello there, checkout my brand new profile! 🤓" // default user's bio text

	_, err = conn.Query(`INSERT INTO "user" (username, password, bio, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $4)`, username, hashedPassword, bio, time.Now())

	return err
}
