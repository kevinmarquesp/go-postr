package router

import (
	"fmt"
	"go-postr/db"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func searchUsernameController(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	query := v.Get("query")

	if len(query) == 0 {
		fmt.Fprintf(w, "")  // insert an empty string in the results tag

		return
	}

	list, err := db.SearchByUsername(query)
	if err != nil {
		log.Println("ERROR: Could not search for user", query)
		log.Println("DETAIL:", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	fmt.Fprintf(w, list)
}

func getRecentArticlesController(w http.ResponseWriter, r *http.Request) {
	list, err := db.GetRecentArticles()
	if err != nil {
		log.Println("ERROR: Could not list recent articles")
		log.Println("DETAIL:", err)

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

func usernameValidationController(w http.ResponseWriter, r *http.Request) {
	usernameb, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("ERROR: Could not fetch the body from the username validation request")
		log.Println("DETAIL:", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	username := string(usernameb)

	wasTaken, err := db.WasUsernameAlreadyTaken(username)
	if err != nil {
		log.Println("ERROR: Database connection when trying to verify username name")
		log.Println("DETAIL:", err)

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
	err := r.ParseForm()
	if err != nil {
		log.Println("ERROR: Could not parse the new user credentials form")
		log.Println("DETAIL:", err)

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
		
		log.Println("ERROR: Could not insert the new user to the table")
		log.Println("DETAIL:", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	//end.

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
