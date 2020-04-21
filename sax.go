package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mickaelvieira/saxifrage/parser"
)

func main() {

	path := filepath.Join(os.Getenv("HOME"), ".ssh", "config")
	content, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		log.Fatal(err)
	}

	p := parser.New(string(content))

	all, err := p.Parse()
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range all {
		fmt.Printf("Host %s\n", c.Host)
		for k, v := range c.Config {
			fmt.Printf("    %s %s\n", k, v)
		}
	}
}
