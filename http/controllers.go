package http

import (
	"go-postr/html"
	"net/http"
)

func homePageController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed: expected get", http.StatusMethodNotAllowed)
		return
	}

	files := html.GetFiles("Partials.Base", "Home")

	tmpl, err := html.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct{PartialBaseParams html.PartialsBaseParams}{
		PartialBaseParams: html.PartialsBaseParams{
			DisplayHeader: true,
		},
	}

	tmpl.Execute(w, r, "Partials.Base", data)
}

func signupPageController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed: expected get", http.StatusMethodNotAllowed)
		return
	}

	files := html.GetFiles("Partials.Base", "Signup")

	tmpl, err := html.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := struct{PartialBaseParams html.PartialsBaseParams}{
		PartialBaseParams: html.PartialsBaseParams{
			DisplayHeader: true,
		},
	}

	tmpl.Execute(w, r, "Partials.Base", data)
}
