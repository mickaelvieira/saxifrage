package cmd

import (
	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/template"
)

func runDump(a *App) error {
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

	if err := template.Render("dump", d); err != nil {
		return err
	}

	return nil
}

// dump creates the dump command
func dump() *command {
	return &command{
		Name:   "dump",
		Usage:  "Dump your SSH configuration",
		Action: runDump,
	}
}
