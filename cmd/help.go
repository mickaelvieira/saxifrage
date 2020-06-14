package cmd

import (
	"fmt"
	"strconv"
)

func getMaxLen(i []string) int {
	var max int
	for _, s := range i {
		l := len(s)
		if l > max {
			max = l
		}
	}
	return max
}

func getCommandList(a commands) map[string]string {
	max := getMaxLen(a.getNames())

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
		App      *App
		Commands map[string]string
	}{
		App:      a,
		Commands: getCommandList(a.Commands),
	}

	if err := a.Templates.Output("help", d); err != nil {
		return err
	}
	return nil
}
