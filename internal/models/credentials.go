package models

import "github.com/kevinmarquesp/go-postr/internal/utils"

// The `Get*Credentials()` functions should act as a parser that reads the
// environment variables and constructs a `struct{}` based on that, so the code
// can easely use this values without doing a bunch of `os.Getenv()` calls.

type PostgresCredentials struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func GetPostgresCredentials() (PostgresCredentials, error) {
	postgres_host, err := utils.ProtectedGetenv("POSTGRES_HOST")
	if err != nil {
		return PostgresCredentials{}, err
	}

	postgres_port, err := utils.ProtectedGetenv("POSTGRES_PORT")
	if err != nil {
		return PostgresCredentials{}, err
	}

	postgres_user, err := utils.ProtectedGetenv("POSTGRES_USER")
	if err != nil {
		return PostgresCredentials{}, err
	}

	postgres_password, err := utils.ProtectedGetenv("POSTGRES_PASSWORD")
	if err != nil {
		return PostgresCredentials{}, err
	}

	postgres_db, err := utils.ProtectedGetenv("POSTGRES_DB")
	if err != nil {
		return PostgresCredentials{}, err
	}

	postgres_sslmode, err := utils.ProtectedGetenv("POSTGRES_SSLMODE")
	if err != nil {
		return PostgresCredentials{}, err
	}

	return PostgresCredentials{
		Host:     postgres_host,
		Port:     postgres_port,
		Username: postgres_user,
		Password: postgres_password,
		Database: postgres_db,
		SSLMode:  postgres_sslmode,
	}, nil
}
