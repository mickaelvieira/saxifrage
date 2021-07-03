package template

import (
	"bytes"
	"embed"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
)

//go:embed templates/*.tmpl
var content embed.FS

// Divider helper tp create a dividing line
func Divider() string {
	return "====================================================="
}

// Line draws a line of the given length
func Line(l int) string {
	return strings.Repeat("-", l)
}

// Templates a template rendering
type Templates struct {
	templates *template.Template
}

// New creates a new Template struct
func New() *Templates {
	fn := promptui.FuncMap
	fn["divider"] = Divider
	fn["line"] = Line

	t := template.New("app-templates").Funcs(fn)
	t, err := t.ParseFS(content, "templates/*.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	return &Templates{
		templates: t,
	}
}

// Output renders the template corresponding to the name with the provided data
func (t *Templates) Output(n string, d interface{}) error {
	if err := t.templates.ExecuteTemplate(os.Stdout, n, d); err != nil {
		return err
	}
	return nil
}

// AsString renders the template corresponding to the name with the provided data
func (t *Templates) AsString(n string, d interface{}) (s string, err error) {
	var buf bytes.Buffer
	if err := t.templates.ExecuteTemplate(&buf, n, d); err != nil {
		return s, err
	}

	s = buf.String()

	return s, nil
}
