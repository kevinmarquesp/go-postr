package router

import (
	"fmt"
	"go-postr/db"
	"log"
	"net/http"
)

func searchUsernameController(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	query := v.Get("query")

	if len(query) == 0 {
		fmt.Fprintf(w, "") // insert an empty string in the results tag if there is no users

		return
	}

	list, err := db.SearchByUsername(query)
	if err != nil {
		fmt.Println(err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	fmt.Fprintf(w, list)
}

func getRecentArticlesController(w http.ResponseWriter, r *http.Request) {
	list, err := db.GetRecentArticles()
	if err != nil {
		log.Println(err)

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

// The username should be in the URL path. Will response with a 400 status code
// error if the user exists and with a 200 if the user doesn't exists (or, if
// you will), if the username is available.
func usernameValidationController(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	wasTaken, err := db.WasUsernameAlreadyTaken(username)
	if err != nil {
		log.Println(err)

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
		log.Println(err)

		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	err = db.InsertNewUser(username, password)
	if err != nil {
		log.Println(err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
