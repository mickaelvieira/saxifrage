package cmd

import (
	"strings"

	"github.com/mickaelvieira/saxifrage/template"
)

func runCompletion(a *App) error {
	d := struct {
		AppName       string
		AppUsage      string
		AppVersion    string
		AppExecutable string
		Commands      string
	}{
		AppName:       a.Name,
		AppUsage:      a.Usage,
		AppVersion:    a.Version,
		AppExecutable: a.Executable,
		Commands:      strings.Join(a.Commands.getNames(), " "),
	}

	if err := template.Output("completion", d); err != nil {
		return err
	}

	return nil
}
