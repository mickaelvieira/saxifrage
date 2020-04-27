package cmd

import (
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/mickaelvieira/saxifrage/keys"
	"github.com/mickaelvieira/saxifrage/prompt"
	"github.com/mickaelvieira/saxifrage/template"
)

var (
	msgConfirmOverride = "The key already exists. Do you want to override it"
	msgConfirmContinue = "Do you want to continue"
	msgConfirmKeyType  = "Select the type of key you want to generate"
	msgConfirmDir      = "Enter the directory"
	msgConfirmFilename = "Enter the file name"
)

type options struct {
	KeyType    keys.Type
	PrivateKey string
	PublicKey  string
	Directory  string
}

func askForKeyType(o *options) error {
	s := promptui.Select{
		Label:        msgConfirmKeyType,
		Items:        keys.Types,
		HideSelected: true,
	}

	_, r, err := s.Run()
	if err != nil {
		return err
	}
	t := keys.GetKeyType(r)
	if t == keys.INVALID {
		return keys.ErrInvalidKeyType
	}

	o.KeyType = t

	return nil
}

func askForDirectory(o *options) error {
	o.Directory = keys.GetDir("")

	r, err := prompt.Prompt(msgConfirmDir, o.Directory)
	if err != nil {
		return err
	}
	if r != "" {
		o.Directory = keys.GetDir(r)
	}

	return nil
}

func askForFilename(o *options) error {
	s, _ := keys.GetFilenamesFromType(o.KeyType)
	r, err := prompt.Prompt(msgConfirmFilename, s)
	if err != nil {
		return err
	}

	fn1, fn2 := keys.GetFilenamesFromType(o.KeyType)
	if r != "" {
		fn1, fn2 = keys.GetFilenamesFromString(r)
	}

	o.PrivateKey = filepath.Join(o.Directory, fn1)
	o.PublicKey = filepath.Join(o.Directory, fn2)

	return nil
}

func runGen(a *App) error {
	o := &options{}
	if err := askForKeyType(o); err != nil {
		return err
	}
	if err := askForDirectory(o); err != nil {
		return err
	}
	if err := askForFilename(o); err != nil {
		return err
	}
	if err := template.Render("summary", o); err != nil {
		return err
	}
	c, err := prompt.Confirm(msgConfirmContinue)
	if err != nil {
		return err
	}

	if c {
		// make sure we don't override an exiting key
		if _, err := os.Stat(o.PrivateKey); err == nil {
			o, err := prompt.Confirm(msgConfirmOverride)
			if err != nil {
				return err
			}
			if !o {
				return keys.ErrKeyOverrideNotAllowed
			}
		}

		prompt.Msg("Generating SSH keys...")

		g := keys.GetGenerator(o.KeyType)

		privateKey, err := g.GenPrivateKey()
		if err != nil {
			return err
		}
		publicKey, err := g.GenPublicKey()
		if err != nil {
			return err
		}
		if err := keys.MakeDir(o.Directory); err != nil {
			return err
		}
		if err := keys.WriteToFile(privateKey, o.PrivateKey); err != nil {
			return err
		}
		if err := keys.WriteToFile(publicKey, o.PublicKey); err != nil {
			return err
		}

		prompt.Msg("The SSH keys were generated successfully!")
	}

	return nil
}
