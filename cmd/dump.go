package cmd

import (
	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/template"
)

func runDump(a *App) error {
	files, err := parser.ParseFiles()
	if err != nil {
		return err
	}

	d := struct {
		Files []*config.File
	}{
		Files: files,
	}

	if err := template.Output("dump", d); err != nil {
		return err
	}

	return nil
}
