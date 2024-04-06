package router

import (
	"go-postr/templates"
	"net/http"
)

func renderIndexView(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()

	templ.Render(w, "Index", nil)
}

func renderSignupView(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplateRenderer()
	templ.Render(w, "Signup", nil)
}
