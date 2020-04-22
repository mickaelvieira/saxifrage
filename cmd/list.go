package cmd

import (
	"fmt"
	"log"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/urfave/cli/v2"
)

func runList(ctx *cli.Context) error {
	c := config.GetUserConfigContent()
	p := parser.New(string(c))

	sections, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range sections {
		fmt.Printf("%s %s\n", string(c.Type), c.Matching)
		for k, v := range c.Configs {
			fmt.Printf("    %s %s\n", k, v)
		}
	}

	return nil
}

// List creates the list command
func List() *cli.Command {
	return &cli.Command{
		Name:   "list",
		Usage:  "List your ssh configuration",
		Action: runList,
	}
}
