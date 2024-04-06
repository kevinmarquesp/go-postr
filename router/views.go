package router

import (
	"go-postr/templates"
	"net/http"
)

func renderIndexController(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()

	templ.Render(w, "Index", nil)
}

func renderSignupController(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()
	templ.Render(w, "Signup", nil)
}
