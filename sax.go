//go:generate pkger
package main

import (
	"os"

	"github.com/markbates/pkger"
	"github.com/mickaelvieira/saxifrage/cmd"
)

var version = "0.0.0" // version is set at build time

func main() {
	pkger.Include("/template/templates")
	app := cmd.New(version)
	if err := app.Run(os.Args); err != nil {
		app.PrintError(err)
	}
}
