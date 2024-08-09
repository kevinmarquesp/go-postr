package models

import "time"

type Provider string

const (
	CredentialsProvider Provider = "Credentials"
)

type Role string

const (
	StandardRole Role = "Standard"
	BannedRole   Role = "Banned"
)

type AccountSchema struct {
	Id          string     `json:"id"`
	Username    string     `json:"username"`
	Bio         string     `json:"bio"`
	DisplayName string     `json:"displayName"`
	Email       string     `json:"email"`
	Provider    Provider   `json:"provider"`
	Role        Role       `json:"role"`
	VerifiedAt  *time.Time `json:"verifyedAt"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}
