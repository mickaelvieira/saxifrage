package formatter

import (
	"fmt"

	"github.com/mickaelvieira/saxifrage/config"
)

// List formats the list of sections
func List(f *config.File) {
	var s string
	for _, c := range f.Sections {
		s += fmt.Sprintf("%s %s\n", string(c.Type), c.Matching)
		for k, v := range c.Configs {
			s += fmt.Sprintf("    %s %s\n", k, v)
		}
	}

	fmt.Println("===========================================")
	fmt.Printf("%s\n", f.Path)
	fmt.Println("===========================================")
	fmt.Printf("%s", s)
}
