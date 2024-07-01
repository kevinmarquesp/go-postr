package models_test

import (
	"os"
	"testing"

	"github.com/kevinmarquesp/go-postr/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestMockSqlite(t *testing.T) {
	const (
		TARGET_DIR = "../../tmp/"
		MOCK_FILE  = "../../db/sqlite3/mock_setup.sql"
	)

	db, dbFile, err := models.MockSqlite(TARGET_DIR, MOCK_FILE)
	assert.NoError(t, err)

	defer os.Remove(dbFile)

	_, err = os.ReadFile(dbFile)
	assert.NoError(t, err)

	err = db.Conn.QueryRow("SELECT * FROM users").Err()
	assert.NoError(t, err)

	err = db.Conn.QueryRow("SELECT * FROM followers").Err()
	assert.NoError(t, err)
}
