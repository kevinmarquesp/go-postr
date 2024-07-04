package main

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"

	"github.com/charmbracelet/log"
	"github.com/kevinmarquesp/go-postr/internal/models"
)

func InsertDummyUsers(db models.GenericDatabaseProvider) {
	// Fetch the dummy users from the JSON Placeholder API.

	resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal("Couldn't not fetch the JSON Placeholder users API.", "error", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Expected status 200 (OK).", "received", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not read the response body.", "error", err)
	}

	var users []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Website  string `json:"website"`
		Address  struct {
			Street  string `json:"street"`
			Suite   string `json:"suite"`
			City    string `json:"city"`
			Zipcode string `json:"zipcode"`
			Geo     struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"geo"`
		} `json:"address"`
		Company struct {
			Name        string `json:"name"`
			CatchPhrase string `json:"catchPhrase"`
			Bs          string `json:"bs"`
		} `json:"company"`
	}

	if err = json.Unmarshal(body, &users); err != nil {
		log.Fatal("Couldn't convert the JSON to a Go's hash map.", "error", err)
	}

	// Loop over all the users and insert each of them in the database.

	const password = "Password!123"

	for _, user := range users {
		fullname := user.Name
		username := string(regexp.MustCompile(`[^a-zA-Z0-9_-]`).
			ReplaceAll([]byte(user.Username), []byte("_"))) // Make it valid.

		_, _, err := db.RegisterNewUser(models.RegisterForm{
			Fullname: fullname,
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Error("Could not register the "+username+" user.", "error", err)
			continue
		}

		log.Info("Inserted the user " + fullname + " (@" + username + ") with success!")
	}
}
