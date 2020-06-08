package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/markbates/pkger"
)

// Divider helper tp create a dividing line
func Divider() string {
	return fmt.Sprintf("=====================================================")
}

// Line draws a line of the given length
func Line(l int) string {
	return strings.Repeat("-", l)
}

func compile(dir string) (*template.Template, error) {
	fn := promptui.FuncMap
	fn["divider"] = Divider
	fn["line"] = Line

	t := template.New("app-templates").Funcs(fn)

	err := pkger.Walk(dir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		f, err := pkger.Open(path)
		if err != nil {
			return err
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}

		_, err = t.Parse(string(b))
		if err != nil {
			return err
		}

		return nil
	})

	return t, err
}

// Templates a template rendering
type Templates struct {
	templates *template.Template
}

// New creates a new Template struct
func New() *Templates {
	t, err := compile("/template/templates/")
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
