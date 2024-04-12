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

// Render the HTML output in the response writter. The `name` is the name of
// the template defined with `{{ define "TEMPLATE_NAME" }}` inside the `html`
// directory.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Will return the generated string basing on the template files The `name` is
// the name of the template defined with `{{ define "TEMPLATE_NAME" }}` inside
// the `html` directory.
func (t *TemplateRenderer) RenderString(name string, data interface{}) (string, error) {
	var buf bytes.Buffer

	err := t.templates.ExecuteTemplate(&buf, name, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Build a new TemplateRenderer object, it will use all the HTML files inside
// the `html/` directory by default - be aware of that.
func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(parseFilesLocation)),
	}
}
