package cmd

import "fmt"

func runVersion(a *App) error {
	fmt.Printf("%s %s\n", a.Name, a.Version)
	return nil
}
