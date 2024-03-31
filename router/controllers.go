package router

import (
	"fmt"
	"go-postr/db"
	"go-postr/templates"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
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
