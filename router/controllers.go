package router

import (
	"fmt"
	"go-postr/db"
	"go-postr/templates"
	"io"
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"golang.org/x/crypto/bcrypt"
)

func renderIndexController(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()

	templ.Render(w, "Index", nil)
}

func searchUsernameController(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	query := v.Get("query")

	if len(query) == 0 {
		fmt.Fprintf(w, "")  // insert an empty string in the results tag

		return
	}

	list, err := db.SearchByUsername(query)
	if err != nil {
		log.Error("Couldn't search for user " + query, "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	fmt.Fprintf(w, list)
}

func getRecentArticlesController(w http.ResponseWriter, r *http.Request) {
	list, err := db.GetRecentArticles()
	if err != nil {
		log.Error("Couldn't list recent articles", "error", err)
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)

		return
	}

	if len(list) == 0 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusNoContent)

		return
	}

	fmt.Fprintf(w, list)
}

func renderSignupController(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()
	templ.Render(w, "Signup", nil)
}

func usernameValidationController(w http.ResponseWriter, r *http.Request) {
	usernameb, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Couldn't fetch the body from the username validation request", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	username := string(usernameb)

	wasTaken, err := db.WasUsernameAlreadyTaken(username)
	if err != nil {
		log.Error("Database connection when trying to verify username name", "error",  err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if wasTaken {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func createNewUserController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Error("Create new user controller, the expected method was 'POST' and not '" + r.Method + "'")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	err := r.ParseForm()
	if err != nil {
		log.Error("Couldn't parse the new user credentials form", "error", err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	//todo: abstract this section into a function in db/insert.go

	conn := db.Connection()  //note: global in the db package

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(username + password), bcrypt.MinCost)
	bio := "Hello there, checkout my brand new profile! 🤓"

	_, err = conn.Query(`INSERT INTO "user" (username, password, bio, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $4)`, username, hashedPassword, bio, time.Now())
	if err != nil {
		//note: this block should be replace by a simple return "", err  on the db package
		//note: this block is what it should be in this package (router)
		
		log.Error("Couldn't insert the new user to the table", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	//end.

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
