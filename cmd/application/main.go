package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/kevinmarquesp/go-postr/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	userRepo, err := repositories.NewUserSqliteRepository("tmp/application.db")
	fatalOnError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	password, err := hashPassword("Password123!")
	fatalOnError(err)

	props := repositories.CredentialsPropperties{
		Role:     models.StandardRole,
		Username: "fulano",
		Email:    "me@email.com",
		Password: password,
	}

	registerResult, err := userRepo.RegisterWithCredentials(ctx, props)
	fatalOnError(err)

	verifyResult, err := userRepo.VerifyAccount(ctx, registerResult.Id)
	fatalOnError(err)

	data, err := json.MarshalIndent(verifyResult, "", "  ")
	fatalOnError(err)

	fmt.Println(string(data))
}

func fatalOnError(err error) {
	if err != nil {
		fmt.Printf("\nFATAL!\n\t%s\n\n", err)
		os.Exit(1)
	}
}

func hashPassword(password string) (string, error) {
	cost := bcrypt.MinCost
	secret := os.Getenv("PASSWORD_SECRET")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+secret), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
