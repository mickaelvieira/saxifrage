package cmd

import (
	"fmt"
	"strconv"

	"github.com/mickaelvieira/saxifrage/template"
)

func getCommands(commands []*command) map[string]string {
	var max int
	for _, c := range commands {
		l := len(c.Name)
		if l > max {
			max = l
		}
	}

	m := make(map[string]string)
	for _, c := range commands {
		f := "%-" + strconv.Itoa(max) + "s"
		n := fmt.Sprintf(f, c.Name)
		m[n] = c.Usage
	}

	return m
}

func runHelp(a *App) error {
	d := struct {
		AppName    string
		AppUsage   string
		AppVersion string
		Commands   map[string]string
	}{
		AppName:    a.Name,
		AppUsage:   a.Usage,
		AppVersion: a.Version,
		Commands:   getCommands(a.Commands),
	}

	if err := template.Render("help", d); err != nil {
		return err
	}
	return nil
}
