package repositories

import "time"

type UserSchema struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserRepository interface {
	CreateNewUser(id, name, username, email, password string) (UserSchema, error) // register
	FindUniqueByEmail(email string) (UserSchema, error)                           // enter
}
