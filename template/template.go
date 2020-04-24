package template

import (
	"fmt"
	"os"
	"text/template"
)

// Divider helper tp create a dividing line
func Divider() string {
	return fmt.Sprintf("=====================================================")
}

// fn defines template helpers
var fn = template.FuncMap{
	"black":     Styler(FGBlack),
	"red":       Styler(FGRed),
	"green":     Styler(FGGreen),
	"yellow":    Styler(FGYellow),
	"blue":      Styler(FGBlue),
	"magenta":   Styler(FGMagenta),
	"cyan":      Styler(FGCyan),
	"white":     Styler(FGWhite),
	"bgBlack":   Styler(BGBlack),
	"bgRed":     Styler(BGRed),
	"bgGreen":   Styler(BGGreen),
	"bgYellow":  Styler(BGYellow),
	"bgBlue":    Styler(BGBlue),
	"bgMagenta": Styler(BGMagenta),
	"bgCyan":    Styler(BGCyan),
	"bgWhite":   Styler(BGWhite),
	"bold":      Styler(FGBold),
	"faint":     Styler(FGFaint),
	"italic":    Styler(FGItalic),
	"underline": Styler(FGUnderline),
	"divider":   Divider,
}

var (
	helpTemplate = `
 NAME:
  {{ .AppName }} - {{ .AppUsage }}

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
 {{ "Type:" | bold }} {{ .Type | bold | green }}
 {{ "Private:" | bold }} {{ .PrivateKey | bold | green }}
 {{ "Public:" | bold }} {{ .PublicKey | bold | green }}
`
	listTemplate = `{{ range .Files }}
{{ divider | bold }}
 {{ "File" | bold }} {{ .Path | bold | green  }}
{{ divider | bold }}{{ $l := len .Sections }}
{{ if eq $l 0 }}
 No sections have been defined in this file
{{ else }}{{ range .Sections }}
 {{ .Type | bold }} {{ .Matching | green | bold }}
{{ range $key, $value := .Configs }}
     {{ $key | bold }} {{ $value | green | bold }}{{ end }}
{{ end }}{{ end }}{{ end }}`
)

var templates = map[string]string{
	"dump":    dumpTemplate,
	"help":    helpTemplate,
	"summary": summaryTemplate,
	"list":    listTemplate,
}

// Render renders the template corresponding to the name with the provided data
func Render(n string, d interface{}) error {
	s, ok := templates[n]
	if !ok {
		return fmt.Errorf("Template '%s' does not exist", n)
	}

	t, err := template.New(n).Funcs(fn).Parse(s)
	if err != nil {
		return err
	}

	if err := t.Execute(os.Stdout, d); err != nil {
		return err
	}
	return nil
}
