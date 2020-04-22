package formatter

import (
	"fmt"
	"sync"

	"github.com/mickaelvieira/saxifrage/config"
)

// List formats the list of sections
func List(f *config.File) {
	var wg sync.WaitGroup
	wg.Add(1)

	var s string
	go func() {
		for _, c := range f.Sections {
			s += fmt.Sprintf("%s %s\n", string(c.Type), c.Matching)
			for k, v := range c.Configs {
				s += fmt.Sprintf("    %s %s\n", k, v)
			}
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("===========================================")
	fmt.Printf("%s\n", f.Path)
	fmt.Println("===========================================")
	fmt.Printf("%s", s)
}
