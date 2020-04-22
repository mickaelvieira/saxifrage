package cmd

import (
	"fmt"
	"log"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/urfave/cli/v2"
)

func runDump(ctx *cli.Context) error {
	fmt.Println("added task: ", ctx.Args().First())

	c := config.GetUserConfigContent()
	p := parser.New(string(c))

	_, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}

	var s string
	for _, token := range p.Tokens {
		s += token.String()
	}

	fmt.Println(s)

	return nil
}

// Dump creates the dump command
func Dump() *cli.Command {
	return &cli.Command{
		Name:   "dump",
		Usage:  "Dump your ssh configuration",
		Action: runDump,
	}
}
