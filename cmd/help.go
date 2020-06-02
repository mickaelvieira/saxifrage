package cmd

import (
	"fmt"
	"strconv"

	"github.com/mickaelvieira/saxifrage/template"
)

func getCommandList(a commands) map[string]string {
	max := template.GetMaxLen(a.getNames())

	m := make(map[string]string)
	for _, c := range a {
		f := "%-" + strconv.Itoa(max) + "s"
		n := fmt.Sprintf(f, c.Name)
		m[n] = c.Usage
	}

	return m
}

func runHelp(a *App) error {
	d := struct {
		AppName       string
		AppUsage      string
		AppVersion    string
		AppExecutable string
		Commands      map[string]string
	}{
		AppName:       a.Name,
		AppUsage:      a.Usage,
		AppVersion:    a.Version,
		AppExecutable: a.Executable,
		Commands:      getCommandList(a.Commands),
	}

	if err := template.Output("help", d); err != nil {
		return err
	}
	return nil
}
