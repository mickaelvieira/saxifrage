package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mickaelvieira/saxifrage/template"
)

func askConfirm(t string) (c bool) {
	fn := template.Styler(template.FGBold)
	fmt.Printf(fn(" %s (y/N) "), t)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := s.Text()
		if t == "y" {
			c = true
		}
		break
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return c
}

func readInput(t string) (i string) {
	fn := template.Styler(template.FGBold)
	fmt.Printf(fn(" %s "), t)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i = s.Text()
		break
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return i
}
