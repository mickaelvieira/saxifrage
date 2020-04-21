package main

import (
	"fmt"
	"log"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
)

func main() {
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
}
