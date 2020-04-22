package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/mickaelvieira/saxifrage/keys"
	"github.com/mickaelvieira/saxifrage/keys/genrsa"
	"github.com/urfave/cli/v2"
)

func askForKeyType() (keys.Type, error) {
	t := keys.GetDefaultType()
	s := keys.TypesToString()
	m := fmt.Sprintf("Enter the type of key you want to generate %s (default: %s)?", s, string(t))
	c := readInput(m)

	if c != "" {
		t = keys.GetKeyType(c)
		if t == keys.INVALID {
			return t, keys.ErrInvalidKeyType
		}
	}
	if t != keys.RSA {
		return t, keys.ErrNotImplementedKeyType
	}
	return t, nil
}

func askForDirectory() string {
	s := keys.GetDir("")
	m := fmt.Sprintf("Enter the subdirectory (default: %s): ", s)
	return readInput(m)
}

func askForFilename(t keys.Type) string {
	s, _ := keys.GetDefaultFilenames(t)
	m := fmt.Sprintf("Enter the file name (default: %s) :", s)
	return readInput(m)
}

func runGenerate(ctx *cli.Context) error {
	fmt.Println("Generating SSH keys...")

	t, err := askForKeyType()
	if err != nil {
		return err
	}

	dn := askForDirectory()
	fn := askForFilename(t)
	dp := keys.GetDir(dn)

	fn1, fn2 := keys.GetDefaultFilenames(t)
	if fn != "" {
		fn1, fn2 = keys.GetUserFilenames(fn)
	}

	p1 := filepath.Join(dp, fn1)
	p2 := filepath.Join(dp, fn2)

	fmt.Printf("You are about to create the following key\nType: %s\nPrivate: %s\nPublic: %s\n", string(t), p1, p2)
	c := askConfirm("Do you want to continue?")

	if c {
		g := genrsa.New(4096)

		privateKey, err := g.GenPrivateKey()
		if err != nil {
			return err
		}
		publicKey, err := g.GenPublicKey()
		if err != nil {
			return err
		}
		if err := keys.MakeDir(dp); err != nil {
			return err
		}
		if err := keys.WriteToFile(privateKey, p1); err != nil {
			return err
		}
		if err := keys.WriteToFile(publicKey, p2); err != nil {
			return err
		}
	}

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
