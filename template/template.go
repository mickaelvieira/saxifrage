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

const lineLen = 80

// Divider helper tp create a dividing line
func Divider() string {
	return strings.Repeat("─", lineLen)
}

func TopLine() string {
	return "┌" + strings.Repeat("─", lineLen) + "\n"
}

func MiddleLine() string {
	return "├" + strings.Repeat("─", lineLen) + "\n"
}

func BottomLine() string {
	return "└" + strings.Repeat("─", lineLen) + "\n"
}

func Border() string {
	return "│"
}

func NewLine() string {
	return "\n"
}

func IsLastOption(i1, t1, i2, t2 int) bool {
	return i1 == t1-1 && i2 == t2-1
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
	fn["border"] = Border
	fn["divider"] = Divider
	fn["topLine"] = TopLine
	fn["bottomLine"] = BottomLine
	fn["middleLine"] = MiddleLine
	fn["newline"] = NewLine
	fn["isLastOption"] = IsLastOption
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
