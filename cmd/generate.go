package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func runGenerate(ctx *cli.Context) error {
	fmt.Println("Generate SSH Key: ", ctx.Args().First())

	return nil
}

// Generate creates the dump command
func Generate() *cli.Command {
	return &cli.Command{
		Name:   "generate",
		Usage:  "Generate an SSH key",
		Action: runGenerate,
	}
}
