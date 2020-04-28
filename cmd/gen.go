package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/keys"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/prompt"
	"github.com/mickaelvieira/saxifrage/template"
)

const (
	msgConfirmOverride    = "The key already exists. Do you want to override it"
	msgConfirmContinue    = "Do you want to continue"
	msgConfirmAddition    = "Do you want to add this key to your config file"
	msgPromptKeyType      = "Select the type of key you want to generate"
	msgPromptKeyDirectory = "Enter the directory"
	msgPromptKeyFilename  = "Enter the file name"
	msgPromptKeyHost      = "Enter the host to which you want to associate this key"
)

type options struct {
	KeyType    keys.Type
	Host       string
	PrivateKey string
	PublicKey  string
	Directory  string
}

func askForKeyType(o *options) error {
	s := promptui.Select{
		Label:        msgPromptKeyType,
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

func askForKeyDirectory(o *options) error {
	o.Directory = keys.GetDir("")

	r, err := prompt.Prompt(msgPromptKeyDirectory, o.Directory)
	if err != nil {
		return err
	}
	if r != "" {
		o.Directory = keys.GetDir(r)
	}

	return nil
}

func askForKeyHost(o *options) error {
	r, err := prompt.Prompt(msgPromptKeyHost, "")
	if err != nil {
		return err
	}
	o.Host = r

	return nil
}

func askForKeyName(o *options) error {
	s, _ := keys.GetFilenamesFromType(o.KeyType)
	r, err := prompt.Prompt(msgPromptKeyFilename, s)
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
	if err := askForKeyDirectory(o); err != nil {
		return err
	}
	if err := askForKeyName(o); err != nil {
		return err
	}
	if err := template.Output("summary", o); err != nil {
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

		if err := prompt.Msg("Generating SSH keys..."); err != nil {
			return err
		}

		g := keys.GetGenerator(o.KeyType)

		privateKey, err := g.GenPrivateKey()
		if err != nil {
			return err
		}
		publicKey, err := g.GenPublicKey()
		if err != nil {
			return err
		}
		if e := keys.MakeDir(o.Directory); e != nil {
			return e
		}
		if e := keys.WriteToFile(privateKey, o.PrivateKey); e != nil {
			return e
		}
		if e := keys.WriteToFile(publicKey, o.PublicKey); e != nil {
			return e
		}
		if e := prompt.Msg("The SSH keys were generated successfully!"); e != nil {
			return e
		}

		c, err = prompt.Confirm(msgConfirmAddition)
		if err != nil {
			return err
		}

		if c {
			if err := askForKeyHost(o); err != nil {
				return err
			}

			files, err := parser.ParseFiles()
			if err != nil {
				return err
			}

			f := files.GetUserConfig()
			if f == nil {
				return config.ErrMissingUserConfig
			}

			b := struct {
				Host       string
				PrivateKey string
			}{
				Host:       o.Host,
				PrivateKey: strings.Replace(o.PrivateKey, os.Getenv("HOME"), "~", 1),
			}

			s, err := template.AsString("config", b)
			if err != nil {
				return err
			}

			_, tokens, err := parser.ParseString(s)
			if err != nil {
				return err
			}

			f.Tokens = append(f.Tokens, tokens...)

			if err := config.WriteToFile(f.Bytes()); err != nil {
				return err
			}
			if err := prompt.Msg("The SSH key was added to your config file"); err != nil {
				return nil
			}
		}
	}

	return nil
}
