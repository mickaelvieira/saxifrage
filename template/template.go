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

// funcMap defines template helpers
var funcMap = template.FuncMap{
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

// Templates templates renderer
type Templates struct {
	templates *template.Template
}

// Render renders corresponding to the name provided with data
func (r *Templates) Render(name string, data interface{}) error {
	return r.templates.ExecuteTemplate(os.Stdout, fmt.Sprintf("%s.tmpl", name), data)
}

// NewRenderer creates a new renderer
func NewRenderer() *Templates {
	return &Templates{
		templates: template.Must(template.New("html-tmpl").Funcs(funcMap).ParseGlob("./template/tmpl/*.tmpl")),
	}
}
