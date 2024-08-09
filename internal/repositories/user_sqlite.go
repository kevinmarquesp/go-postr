package repositories

import (
	"context"
	"database/sql"

	"github.com/kevinmarquesp/go-postr/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/oklog/ulid/v2"
)

type UserSqliteRepository struct {
	UserRepository

	Database *sql.DB
}

func NewUserSqliteRepository(path string) (*UserSqliteRepository, error) {
	connection, err := sql.Open("sqlite3", path)
	if err != nil {
		return &UserSqliteRepository{}, err
	}

	return &UserSqliteRepository{
		Database: connection,
	}, nil
}

func (us UserSqliteRepository) RegisterWithCredentials(ctx context.Context, props CredentialsPropperties) (models.AccountSchema, error) {
	userStatement, err := us.Database.Prepare(`
      INSERT
        INTO Users
          (id, username, email, provider, role)
        VALUES
          (?1, ?2, ?3, ?4, ?5);`)
	if err != nil {
		return models.AccountSchema{}, err
	}
	defer userStatement.Close()

	credentialsStatement, err := us.Database.Prepare(`
      INSERT
        INTO Credentials
          (user_id, email, password)
        VALUES
          (?1, ?2, ?3);`)
	if err != nil {
		return models.AccountSchema{}, err
	}
	defer credentialsStatement.Close()

	id := ulid.Make().String()

	_, err = userStatement.ExecContext(ctx, id, props.Username, props.Email, models.CredentialsProvider, props.Role)
	if err != nil {
		return models.AccountSchema{}, err
	}

	_, err = credentialsStatement.ExecContext(ctx, id, props.Email, props.Password)
	if err != nil {
		return models.AccountSchema{}, err
	}

	return getAccountData(ctx, us.Database, id)
}

func (us UserSqliteRepository) VerifyAccount(ctx context.Context, userId string) (models.AccountSchema, error) {
	verifyStatement, err := us.Database.Prepare(`
      UPDATE Users
        SET
          verified_at = CURRENT_TIMESTAMP
        WHERE
          id IS ?1;`)
	if err != nil {
		return models.AccountSchema{}, err
	}
	defer verifyStatement.Close()

	_, err = verifyStatement.ExecContext(ctx, userId)
	if err != nil {
		return models.AccountSchema{}, err
	}

	return getAccountData(ctx, us.Database, userId)
}

func getAccountData(ctx context.Context, database *sql.DB, userId string) (models.AccountSchema, error) {
	rows, err := database.QueryContext(ctx, `
      SELECT
        id, username, email, role, provider, verified_at, created_at, updated_at
      FROM
        Users User
      WHERE
        User.id IS ?1;`, userId)
	if err != nil {
		return models.AccountSchema{}, err
	}

	var result models.AccountSchema

	for rows.Next() {
		if err = rows.Scan(
			&result.Id,
			&result.Username,
			&result.Email,
			&result.Role,
			&result.Provider,
			&result.VerifiedAt,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
			return models.AccountSchema{}, err
		}
	}

	return result, nil
}
