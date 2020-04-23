package cmd

import (
	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/template"
	"github.com/urfave/cli/v2"
)

func runDump(ctx *cli.Context) error {
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

	t := template.NewRenderer()
	err = t.Render("dump", d)
	if err != nil {
		return err
	}

	return nil
}

// Dump creates the dump command
func Dump() *cli.Command {
	return &cli.Command{
		Name:   "dump",
		Usage:  "Dump your ssh configuration",
		Action: runDump,
	}
}
