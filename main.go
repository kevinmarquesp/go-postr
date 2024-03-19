package main

import (
	"html/template"
	"io"
	"net/http"
)

const parseFilesLocation = "html/*.html"

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(parseFilesLocation)),
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := NewTemplateRenderer()

		templ.Render(w, "Index", nil)
	})

	http.ListenAndServe(":8000", nil)
}
