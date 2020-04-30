package cmd

import (
	"os"
	"strings"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/keys"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/prompt"
	"github.com/mickaelvieira/saxifrage/template"
)

func makeHomePathRelative(p string) string {
	return strings.Replace(p, os.Getenv("HOME"), "~", 1)
}

func runGen(a *App) error {
	keyOpts, err := prompt.KeyFlow()
	if err != nil {
		return err
	}

	if e := template.Output("summary", keyOpts); e != nil {
		return e
	}

	confirm, err := prompt.Confirm(prompt.MsgConfirmContinue)
	if err != nil {
		return err
	}

	if confirm {
		// make sure we don't override an exiting key
		if _, err := os.Stat(keyOpts.PrivateKey); err == nil {
			override, err := prompt.Confirm(prompt.MsgConfirmOverride)
			if err != nil {
				return err
			}
			if !override {
				return keys.ErrKeyOverrideNotAllowed
			}
		}

		if err := prompt.Msg("Generating SSH keys..."); err != nil {
			return err
		}

		generator := keys.GetGenerator(keyOpts)

		privateKey, err := generator.GenPrivateKey()
		if err != nil {
			return err
		}
		publicKey, err := generator.GenPublicKey()
		if err != nil {
			return err
		}
		if e := keys.Writekeys(publicKey, privateKey, keyOpts); e != nil {
			return e
		}
		if e := prompt.Msg("The SSH keys were generated successfully!"); e != nil {
			return e
		}

		confirm, err = prompt.Confirm(prompt.MsgConfirmAddition)
		if err != nil {
			return err
		}

		if confirm {
			genOpts, err := prompt.ConfigFlow(makeHomePathRelative(keyOpts.PrivateKey))
			if err != nil {
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

			s, err := template.AsString("config", genOpts)
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
