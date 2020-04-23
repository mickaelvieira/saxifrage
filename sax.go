package main

import (
	"os"

	"github.com/mickaelvieira/saxifrage/cmd"
)

func main() {
	app := cmd.New()
	if err := app.Run(os.Args); err != nil {
		app.PrintError(err)
	}
}
