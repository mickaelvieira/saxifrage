package main

import (
	"fmt"
	"os"

	"github.com/mickaelvieira/saxifrage/cmd"
	"github.com/mickaelvieira/saxifrage/template"
	"github.com/urfave/cli/v2"
)

func main() {
	commands := make([]*cli.Command, 3)
	commands[0] = cmd.List()
	commands[1] = cmd.Generate()
	commands[2] = cmd.Dump()

	app := &cli.App{
		Usage:    "Manage your ssh_config",
		Commands: commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		fn := template.Styler(template.FGBold, template.FGRed)
		fmt.Printf("\n %s\n", fn(err))
	}
}
