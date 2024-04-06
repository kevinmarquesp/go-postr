package router

import (
	"go-postr/templates"
	"net/http"
)

func renderIndexView(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()

	templ.Render(w, "Index", nil)
}

func redirectToHomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func renderSignupView(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()
	templ.Render(w, "Signup", nil)
}
