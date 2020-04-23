package cmd

import (
	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/template"
)

func runList(a *App) error {
	files, err := parser.ParseFiles()
	if err != nil {
		return err
	}

	d := struct {
		Files []*config.File
	}{
		Files: files,
	}

	if err := template.Render("list", d); err != nil {
		return err
	}

	return nil
}

// list creates the list command
func list() *command {
	return &command{
		Name:   "show",
		Usage:  "Show your SSH configuration",
		Action: runList,
	}
}
