package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/mickaelvieira/saxifrage/template"
)

type command struct {
	Name   string
	Usage  string
	Action func(*App) error
}

// App a cli application
type App struct {
	Name     string
	Usage    string
	Commands []*command
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
func New() *App {
	a := &App{
		Usage: "Manage your ssh_config",
	}

	a.Commands = make([]*command, 4)
	a.Commands[0] = generate()
	a.Commands[1] = list()
	a.Commands[2] = dump()
	a.Commands[3] = help()

	return a
}
