package cmd

import (
	"fmt"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/formatter"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/urfave/cli/v2"
)

func runList(ctx *cli.Context) error {

	t, err := askForKeyType()
	if err != nil {
		return err
	}

	fmt.Println(t)
	gc, err := parser.ParseFile(config.GetGlobalConfigPath())
	if err != nil {
		return err
	}

	uc, err := parser.ParseFile(config.GetUserConfigPath())
	if err != nil {
		return err
	}

	formatter.List(gc)
	formatter.List(uc)

	return nil
}

// List creates the list command
func List() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "List your ssh configuration",
		Action: runList,
	}
}
