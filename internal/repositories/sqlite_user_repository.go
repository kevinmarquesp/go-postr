package repositories

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteUserRepository struct {
	UserRepository

	database *sql.DB
}

func NewSqliteUserRepository(path string) (*SqliteUserRepository, error) {
	connection, err := sql.Open("sqlite3", path)
	if err != nil {
		return &SqliteUserRepository{}, err
	}

	return &SqliteUserRepository{
		database: connection,
	}, nil
}

func (su SqliteUserRepository) CreateNewUser(id, name, username, email, password string) (UserSchema, error) {
	empty := UserSchema{}

	statement, err := su.database.Prepare(`
      INSERT
        INTO Users
          (id, name, username, email, password)
      VALUES
        (?1, ?2, ?3, ?4, ?5);
    `)
	if err != nil {
		return empty, err
	}
	defer statement.Close()

	_, err = statement.Exec(id, name, username, email, password)
	if err != nil {
		return empty, err
	}

	return UserSchema{
		Id:       id,
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
	}, nil
}

func (su SqliteUserRepository) FindUniqueByEmail(email string) (UserSchema, error) {
	return UserSchema{}, errors.New("Not implemented.")
}
