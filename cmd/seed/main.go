package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/kevinmarquesp/go-postr/internal/models"
)

const (
	DOTENV   = ".env"
	PASSWORD = "password1234"
)

type UsersSchema struct {
	Username    string
	Password    string
	Description string
}

func main() {
	log.Info("Initializing the seed script.")

	if err := godotenv.Load(DOTENV); err != nil {
		log.Error("Could not load the "+DOTENV+" file.", "error", err)
	}

	users, err := JsonPlaceholderFetchUsers(PASSWORD)
	if err != nil {
		log.Fatal("Could not fetch users from teh JSON Placeholder API.", "error", err)
	}

	db_service := &models.Postgres{}

	if err := db_service.Connect(); err != nil {
		log.Fatal("Could not connect to the specifyed database service.", "error", err)
	}

	for _, user := range users {
		err := db_service.InsertUser(user.Username, user.Password, user.Description)
		if err != nil {
			log.Error("Could not insert user to the database", "error", err)
		} else {
			log.Info("Inserted user '" + user.Username + "' to the database with success!")
		}
	}
}

const FAKE_JSON_API = "https://jsonplaceholder.typicode.com/users"

type JsonPlaceholderUserAddress struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type JsonPlaceholderUserCompany struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type JsonPlaceholderUser struct {
	Name     string                     `json:"name"`
	Username string                     `json:"username"`
	Email    string                     `json:"email"`
	Address  JsonPlaceholderUserAddress `json:"address"`
	Phone    string                     `json:"phone"`
	Website  string                     `json:"website"`
	Company  JsonPlaceholderUserCompany `json:"company"`
}

func JsonPlaceholderFetchUsers(password string) ([]UsersSchema, error) {
	resp, err := http.Get(FAKE_JSON_API)
	if err != nil {
		return []UsersSchema{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []UsersSchema{}, fmt.Errorf("Invalid response status code: %d", resp.StatusCode)
	}

	resp_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []UsersSchema{}, err
	}

	var json_placeholder_users []JsonPlaceholderUser

	err = json.Unmarshal(resp_body, &json_placeholder_users)
	if err != nil {
		return []UsersSchema{}, err
	}

	var results []UsersSchema

	for _, user := range json_placeholder_users {
		username := user.Username

		description := fmt.Sprintf("Hello! My real name is %s, I work at the %s - \"%s\" - company as %s. You"+
			"can contact me by my email %s, or by my phone number %s. Checkout my portfolio in %s! My"+
			"personal office address: %s (%s), %s - %s", user.Name, user.Company.Name,
			user.Company.CatchPhrase, user.Company.Bs, user.Email, user.Phone, user.Website,
			user.Address.Street, user.Address.City, user.Address.Suite, user.Address.Zipcode)

		results = append(results, UsersSchema{
			Username:    username,
			Password:    password,
			Description: description,
		})
	}

	return results, nil
}
