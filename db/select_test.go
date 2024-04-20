package db

import (
	"go-postr/utils"
	"testing"
)

func TestSearchUsername(t *testing.T) {
	testdb, err := SetupDummyTestDatabaseConnection()
	utils.FailTestNowIfErrorIsNotNil(t, err)

	t.Log("Working with the " + testdb + " database.")

	t.Run("Should return an empty array", func(t *testing.T) {
		result, err := SearchUsername("test")
		utils.FailTestNowIfErrorIsNotNil(t, err)

		if len(result) != 0 {
			t.Errorf("Expected empty array, got %v", result)
		}
	})

	_, err = conn.Exec(`INSERT INTO "user" (username, password) VALUES
		('user1', 'dummy_pass'),
		('user2', 'dummy_pass'),
		('user3', 'dummy_pass')`)
	utils.FailTestNowIfErrorIsNotNil(t, err)

	t.Run("Should find each username one by one", func(t *testing.T) {
		for _, username := range []string{"user1", "user2", "user3"} {
			result, err := SearchUsername(username)
			utils.FailTestNowIfErrorIsNotNil(t, err)

			if len(result) != 1 || result[0] != username {
				t.Errorf("Expected [%s], got %v", username, result)
			}
		}
	})

	t.Run("Should return an empty array for non existing username", func(t *testing.T) {
		result, err := SearchUsername("nonexistent")
		utils.FailTestNowIfErrorIsNotNil(t, err)

		if len(result) != 0 {
			t.Errorf("Expected empty array, got %v", result)
		}
	})
}
