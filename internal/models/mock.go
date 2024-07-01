package models

import (
	"os"

	"github.com/google/uuid"
)

func MockSqlite(target, setupFile string) (Sqlite, string, error) {
	dbFile := target + "/" + uuid.New().String()[:8] + ".sqlite3"
	db := Sqlite{}

	if err := db.Connect(dbFile); err != nil {
		return Sqlite{}, "", err
	}

	queryBytes, err := os.ReadFile(setupFile)
	if err != nil {
		return Sqlite{}, "", err
	}

	query := string(queryBytes)

	_, err = db.Conn.Exec(query)
	if err != nil {
		return Sqlite{}, "", err
	}

	return db, dbFile, nil
}
