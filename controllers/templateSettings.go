package controllers

import "html/template"

var htmlDir string = "templates/*.html"
var Tmpl *template.Template = template.Must(template.ParseGlob(htmlDir))
