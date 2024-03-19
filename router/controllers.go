package router

import (
	"fmt"
	"go-postr/db"
	"go-postr/templates"
	"net/http"

	"github.com/charmbracelet/log"
)

func RenderIndexController(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()

	templ.Render(w, "Index", nil)
}

func SearchUsernameController(w http.ResponseWriter, r *http.Request) {
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
