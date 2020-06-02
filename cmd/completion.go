package cmd

import (
	"strings"
)

func runCompletion(a *App) error {
	d := struct {
		App      *App
		Commands string
	}{
		App:      a,
		Commands: strings.Join(a.Commands.getNames(), " "),
	}

	if err := a.Templates.Output("completion", d); err != nil {
		return err
	}

	return nil
}
