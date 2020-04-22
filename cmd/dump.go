package cmd

import (
	"fmt"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/urfave/cli/v2"
)

func runDump(ctx *cli.Context) error {
	gc, err := parser.ParseFile(config.GetGlobalConfigPath())
	if err != nil {
		return err
	}

	fmt.Printf("\nFile: %s\n\n", gc.Path)
	fmt.Printf("%s", gc)

	uc, err := parser.ParseFile(config.GetUserConfigPath())
	if err != nil {
		return err
	}

	fmt.Printf("\nFile: %s\n\n", uc.Path)
	fmt.Printf("%s", uc)

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
