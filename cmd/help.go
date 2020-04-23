package cmd

import (
	"github.com/mickaelvieira/saxifrage/template"
)

func runHelp(a *App) error {
	if err := template.Render("help", a); err != nil {
		return err
	}
	return nil
}

// help creates the help command
func help() *command {
	return &command{
		Name:   "help",
		Usage:  "Show usage",
		Action: runHelp,
	}
}
