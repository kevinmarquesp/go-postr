package templates

import (
	"bytes"
	"html/template"
	"io"
)

const parseFilesLocation = "html/*.html"

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *TemplateRenderer) RenderString(name string, data interface{}) (string, error) {
	var buf bytes.Buffer

	err := t.templates.ExecuteTemplate(&buf, name, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(parseFilesLocation)),
	}
}

