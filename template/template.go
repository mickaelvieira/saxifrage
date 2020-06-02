package template

import (
	"bytes"
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
	configTemplate = `
Host {{ .Host }}
{{ if .User }}    User {{ .User }}{{ end }}
    Port {{ .Port }}
    IdentityFile {{ .IdentityFile }}
`
	helpTemplate = `
 NAME:
  {{ .AppName }} {{ .AppVersion }} - {{ .AppUsage }}

 USAGE:
  {{ .AppExecutable }} [command]

 COMMANDS:
{{ range $name, $usage := .Commands}}
  {{ $name }}    {{ $usage }}{{ end }}
`
	dumpTemplate = `{{ range .Files }}
{{ divider | bold }}
 {{ "File" | bold }} {{ .Path | bold | green  }}
{{ divider | bold }}

{{ range .Lines}}{{ .Number | green | bold }}  {{ . }}{{ end }}
{{ end }}`
	summaryTemplate = `
 {{ "You are about to create the following SSH key" | bold }}
 {{ "Type:    " | bold }} {{ .KeyType | bold | green }}
 {{ "Private: " | bold }} {{ .PrivateKey | bold | green }}
 {{ "Public:  " | bold }} {{ .PublicKey | bold | green }}

`
	filesTemplate = `
Files to be deleted:

{{ range .KeyFiles}}{{ . }}
{{ end }}
`
	linesTemplate = `
Lines to be deleted:

{{ range .Lines}}{{ .Number | green | bold }} {{ . | bold }}{{ end }}
`
	listTemplate = `{{ range .Files }}
{{ divider | bold }}
 {{ "File" | bold }} {{ .Path | bold | green  }}
{{ divider | bold }}{{ $l := len .BuildSections }}
{{ if eq $l 0 }}
 No sections have been defined in this file
{{ else }}{{ range .BuildSections }}
 {{ .Type | bold }}{{ .Separator }}{{ .Matching | green | bold }}
{{ range .Options }}
     {{ .Name | bold }}{{ .Separator }}{{ .Value | green | bold }}{{ end }}
{{ end }}{{ end }}{{ end }}`
	readInputTemplate  = `{{ "✔ " | green }}{{ .Text | bold }}{{ if .Default }} {{ printf "%s%s%s" "(" .Default ")" | faint }}{{ end }}{{ ": " | bold }}`
	askConfirmTemplate = `{{ "? " | bold | blue }}{{ printf "%s %s" .Text "(y/N)?" | bold }} `
	messageTempate     = `{{ "=> " | bold | green }}{{ .Text | bold }}
`
	completionTemplate = `#!/bin/bash

_saxifrage()
{
  IFS=" " read -r -a COMPREPLY <<<"$(compgen -W "{{ .Commands}}" "${COMP_WORDS[1]}")"
}

complete -F _saxifrage {{ .AppExecutable }}
	`
)

var templates = map[string]string{
	"dump":        dumpTemplate,
	"help":        helpTemplate,
	"summary":     summaryTemplate,
	"list":        listTemplate,
	"read-input":  readInputTemplate,
	"ask-confirm": askConfirmTemplate,
	"message":     messageTempate,
	"config":      configTemplate,
	"files":       filesTemplate,
	"lines":       linesTemplate,
	"completion":  completionTemplate,
}

func getTemplate(n string) (*template.Template, error) {
	s, ok := templates[n]
	if !ok {
		return nil, fmt.Errorf("Template '%s' does not exist", n)
	}

	fn := promptui.FuncMap
	fn["divider"] = Divider
	fn["line"] = Line

	t, err := template.New(n).Funcs(fn).Parse(s)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Output renders the template corresponding to the name with the provided data
func Output(n string, d interface{}) error {
	t, err := getTemplate(n)
	if err != nil {
		return err
	}

	if err := t.Execute(os.Stdout, d); err != nil {
		return err
	}
	return nil
}

// AsString renders the template corresponding to the name with the provided data
func AsString(n string, d interface{}) (s string, err error) {
	t, err := getTemplate(n)
	if err != nil {
		return s, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, d); err != nil {
		return s, err
	}

	s = buf.String()

	return s, nil
}
