package html

import (
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	template *template.Template
}

func (t *Template) Execute(w http.ResponseWriter, r *http.Request, target string, data interface{}) {
	err := t.template.ExecuteTemplate(w, target, data)
	if err != nil {
		log.Println("could not execute", target, "template:", err)
		http.Error(w, err.Error(), http.StatusNoContent)
	}
}

func GetFiles(keys ...string) []string {
	filesMap := map[string]string{
		"Partials.Base": "templates/partials/base.html",
		"Home":          "templates/home.html",
	}

	var files []string

	for _, key := range keys {
		files = append(files, filesMap[key])
	}

	return files
}

func ParseFiles(files ...string) (*Template, error) {
	tmpl, err := template.New("").ParseFiles(files...)  //add custom functions here
	if err != nil {
		return nil, err
	}

	return &Template{
		template: tmpl,
	}, nil
}
