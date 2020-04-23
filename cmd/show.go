package cmd

import (
	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/template"
)

func runList(a *App) error {
	gc, err := parser.ParseFile(config.GetGlobalConfigPath())
	if err != nil {
		return err
	}

	uc, err := parser.ParseFile(config.GetUserConfigPath())
	if err != nil {
		return err
	}

	f := []*config.File{gc, uc}
	d := struct {
		Files []*config.File
	}{
		Files: f,
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
