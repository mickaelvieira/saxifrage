package cmd

import "fmt"

func runVersion(a *App) error {
	fmt.Print(a.Version)
	return nil
}
