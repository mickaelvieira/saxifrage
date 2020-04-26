package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/mickaelvieira/saxifrage/keys"
	"github.com/mickaelvieira/saxifrage/template"
)

// App a cli application
type App struct {
	Name     string
	Version  string
	Usage    string
	Commands []*command
}

type command struct {
	Name   string
	Usage  string
	Action func(*App) error
}

// Run runs the application
func (a *App) Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Cannot run the application. Arguments are missing")
	}

	a.Name = filepath.Base(args[0])

	n := "help"
	if len(args) > 1 {
		n = args[1]
	}

	c := a.find(n)
	if c == nil {
		return fmt.Errorf("Command '%s' does not exist", n)
	}

	return c.Action(a)
}

// PrintError prints the error
func (a *App) PrintError(e error) {
	fn := template.Styler(template.FGBold, template.FGRed)
	fmt.Printf(" %s\n", fn(e))
}

func (a *App) find(name string) *command {
	for _, c := range a.Commands {
		if c.Name == name {
			return c
		}
	}
	return nil
}

// New creates a new application
func New(v string) *App {
	a := &App{
		Version: v,
		Usage:   "A CLI tool to manage your SSH keys",
	}

	a.Commands = make([]*command, 4)
	a.Commands[0] = &command{Name: "gen", Usage: fmt.Sprintf("Generate interactively a SSH key (%s)", keys.TypesToString()), Action: runGen}
	a.Commands[1] = &command{Name: "config", Usage: "Show your SSH configuration", Action: runConfig}
	a.Commands[2] = &command{Name: "dump", Usage: "Dump your SSH configuration", Action: runDump}
	a.Commands[3] = &command{Name: "help", Usage: "Show this help", Action: runHelp}

	return a
}
