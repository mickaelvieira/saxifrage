package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/mickaelvieira/saxifrage/keys"
)

// App a cli application
type App struct {
	Name       string
	Executable string
	Version    string
	Usage      string
	Commands   commands
}

type commands []*command

func (c *commands) getNames() []string {
	names := make([]string, len(*c))
	for i, cmd := range *c {
		names[i] = cmd.Name
	}
	return names
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

	a.Executable = filepath.Base(args[0])

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
	fn := promptui.Styler(promptui.FGBold, promptui.FGRed)
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
		Name:    "Saxifrage",
		Version: v,
		Usage:   "A CLI tool to manage your SSH keys",
	}

	a.Commands = make(commands, 8)
	a.Commands[0] = &command{Name: "gen", Usage: fmt.Sprintf("Generate interactively a SSH key (%s)", keys.TypesToString()), Action: runGenerate}
	a.Commands[1] = &command{Name: "ls", Usage: "List SSH configuration sections", Action: runList}
	a.Commands[2] = &command{Name: "dump", Usage: "Dump your SSH configuration", Action: runDump}
	a.Commands[3] = &command{Name: "rm", Usage: "Remove interactively a section and its related SSH keys", Action: runRemove}
	a.Commands[4] = &command{Name: "help", Usage: "Show this help", Action: runHelp}
	a.Commands[5] = &command{Name: "upgrade", Usage: fmt.Sprintf("Upgrade %s", a.Name), Action: runUpgrade}
	a.Commands[6] = &command{Name: "version", Usage: "Display the application version", Action: runVersion}
	a.Commands[7] = &command{Name: "completion", Usage: "Generate bash completion", Action: runCompletion}

	return a
}
