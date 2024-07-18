package repositories

import "time"

type UserSchema struct {
	Name      string
	Username  string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

type UserRepository interface {
	CreateNewUser(name, username, email, password string) (UserSchema, error) // register
	FindUniqueByEmail(UserSchema, error) (UserSchema, error)                  // enter
}
