package controllers

import (
	"html/template"
	"log"
)

var Tmpl *template.Template

func InitializeHtmlTemplates() {
	log.Println("Setting up template files...")

	pagesDir := "templates/*.html"
	componentsDir := "templates/components/*.html"

	Tmpl = template.Must(template.ParseGlob(pagesDir))
	Tmpl = template.Must(Tmpl.ParseGlob(componentsDir))
}
