package main

import (
	"os"

	"github.com/mickaelvieira/saxifrage/cmd"
)

var version = "0.0.0" // version is set at build time

func main() {
	app := cmd.New(version)
	if err := app.Run(os.Args); err != nil {
		app.PrintError(err)
	}
}
