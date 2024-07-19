package repositories

import "time"

type UserSchema struct {
	Name      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	CreateNewUser(name, username, email, password string) (UserSchema, error) // register
	FindUniqueByEmail(email string) (UserSchema, error)                       // enter
}
