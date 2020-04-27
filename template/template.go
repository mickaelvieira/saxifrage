package template

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/manifoldco/promptui"
)

// GetMaxLen returns the length of the longest string
func GetMaxLen(i []string) int {
	var max int
	for _, s := range i {
		l := len(s)
		if l > max {
			max = l
		}
	}
	return max
}

// Divider helper tp create a dividing line
func Divider() string {
	return fmt.Sprintf("=====================================================")
}

// Line draws a line of the given length
func Line(l int) string {
	return strings.Repeat("-", l)
}

var (
	helpTemplate = `
 NAME:
  {{ .AppName }} {{ .AppVersion }} - {{ .AppUsage }}

 USAGE:
  {{ .AppName }} [command]

 COMMANDS:
{{ range $name, $usage := .Commands}}
  {{ $name }}    {{ $usage }}{{ end }}
`
	dumpTemplate = `{{ range .Files }}
{{ divider | bold }}
 {{ "File" | bold }} {{ .Path | bold | green  }}
{{ divider | bold }}

{{ . }}
{{ end }}`
	summaryTemplate = `
 {{ "You are about to create the following SSH key" | bold }}
 {{ "Type:    " | bold }} {{ .KeyType | bold | green }}
 {{ "Private: " | bold }} {{ .PrivateKey | bold | green }}
 {{ "Public:  " | bold }} {{ .PublicKey | bold | green }}

`
	listTemplate = `{{ range .Files }}
{{ divider | bold }}
 {{ "File" | bold }} {{ .Path | bold | green  }}
{{ divider | bold }}{{ $l := len .Sections }}
{{ if eq $l 0 }}
 No sections have been defined in this file
{{ else }}{{ range .Sections }}
 {{ .Type | bold }}{{ .Separator }}{{ .Matching | green | bold }}
{{ range .Options }}
     {{ .Name | bold }}{{ .Separator }}{{ .Value | green | bold }}{{ end }}
{{ end }}{{ end }}{{ end }}`
	readInputTemplate  = `{{ "✔ " | green }}{{ .Text | bold }}{{ if .Default }} {{ printf "%s%s%s" "(" .Default ")" | faint }}{{ end }}{{ ": " | bold }}`
	askConfirmTemplate = `{{ "? " | bold | blue }}{{ .Text | bold }} (y/N): `
)

var templates = map[string]string{
	"dump":        dumpTemplate,
	"help":        helpTemplate,
	"summary":     summaryTemplate,
	"list":        listTemplate,
	"read-input":  readInputTemplate,
	"ask-confirm": askConfirmTemplate,
}

// Render renders the template corresponding to the name with the provided data
func Render(n string, d interface{}) error {
	s, ok := templates[n]
	if !ok {
		return fmt.Errorf("Template '%s' does not exist", n)
	}

	fn := promptui.FuncMap
	fn["divider"] = Divider
	fn["line"] = Line

	t, err := template.New(n).Funcs(fn).Parse(s)
	if err != nil {
		return err
	}

	if err := t.Execute(os.Stdout, d); err != nil {
		return err
	}
	return nil
}
