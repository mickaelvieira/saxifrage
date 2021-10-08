package cmd

import (
	"fmt"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
)

type file struct {
	Path     string
	Sections config.Sections
}

func runList(a *App) error {
	files, err := parser.ParseFiles()
	if err != nil {
		return err
	}

	data := make([]*file, 0)

	for _, f := range files {
		s, err := f.BuildSections()
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, &file{Path: f.Path, Sections: s})
	}

	d := struct {
		Files []*file
	}{
		Files: data,
	}

	if err := a.Templates.Output("list", d); err != nil {
		return err
	}

	return nil
}
