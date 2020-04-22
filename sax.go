package main

import (
	"log"
	"os"

	"github.com/mickaelvieira/saxifrage/cmd"
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
		log.Fatal(err)
	}
}
