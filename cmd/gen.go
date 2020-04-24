package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mickaelvieira/saxifrage/keys"
	"github.com/mickaelvieira/saxifrage/template"
)

var (
	msgConfirmOverride = "The key already exists. Do you want to override it?"
	msgConfirmContinue = "Do you want to continue?"
	msgConfirmKeyType  = "Enter the type of key you want to generate %s (default: %s)?"
	msgConfirmDir      = "Enter the subdirectory (default: %s):"
	msgConfirmFilename = "Enter the file name (default: %s):"
)

func askForKeyType() (keys.Type, error) {
	f := template.Styler(template.FGBold, template.FGGreen)
	t := keys.GetDefaultType()
	s := keys.TypesToString()
	i := readInput(fmt.Sprintf(msgConfirmKeyType, s, f(string(t))))

	if i != "" {
		t = keys.GetKeyType(i)
		if t == keys.INVALID {
			return t, keys.ErrInvalidKeyType
		}
	}
	return t, nil
}

func msg(m string) {
	f := template.Styler(template.FGBold, template.FGGreen)
	fmt.Printf(" %s\n", f(m))
}

func askForDirectory() string {
	f := template.Styler(template.FGBold, template.FGGreen)
	s := keys.GetDir("")
	m := fmt.Sprintf(msgConfirmDir, f(s))
	return readInput(m)
}

func askForFilename(t keys.Type) string {
	f := template.Styler(template.FGBold, template.FGGreen)
	s, _ := keys.GetFilenamesFromType(t)
	m := fmt.Sprintf(msgConfirmFilename, f(s))
	return readInput(m)
}

func runGenerate(a *App) error {
	t, err := askForKeyType()
	if err != nil {
		return err
	}

	dn := askForDirectory()
	fn := askForFilename(t)
	dp := keys.GetDir(dn)

	fn1, fn2 := keys.GetFilenamesFromType(t)
	if fn != "" {
		fn1, fn2 = keys.GetFilenamesFromString(fn)
	}

	p1 := filepath.Join(dp, fn1)
	p2 := filepath.Join(dp, fn2)

	d := struct {
		Type       string
		PrivateKey string
		PublicKey  string
	}{
		Type:       string(t),
		PrivateKey: p1,
		PublicKey:  p2,
	}

	if err := template.Render("summary", d); err != nil {
		return err
	}

	c := askConfirm(msgConfirmContinue)

	if c {
		// make sure we don't override an exiting key
		if _, err := os.Stat(p1); err == nil {
			o := askConfirm(msgConfirmOverride)
			if !o {
				return keys.ErrKeyOverrideNotAllowed
			}
		}

		msg("Generating SSH keys...")

		g := keys.GetGenerator(t)

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

		msg("The SSH keys were generated successfully!")
	}

	return nil
}

// generate creates the dump command
func generate() *command {
	return &command{
		Name:   "gen",
		Usage:  "Generate interactively a SSH key",
		Action: runGenerate,
	}
}
