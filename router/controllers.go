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
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if len(list) == 0 {
		fmt.Fprintf(w, "Any posts created yet")
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

	log.Info(username)

	w.WriteHeader(http.StatusOK)
	// w.WriteHeader(http.StatusBadRequest)
}
