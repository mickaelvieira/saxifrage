package cmd

import (
	"os"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/keys"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/prompt"
	"github.com/mickaelvieira/saxifrage/template"
)

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

		if err := prompt.Msg(prompt.MsgGeneratingKeys); err != nil {
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
		if e := prompt.Msg(prompt.MsgKeyGenerated); e != nil {
			return e
		}

		confirm, err := prompt.Confirm(prompt.MsgConfirmAddition)
		if err != nil {
			return err
		}

		if confirm {
			configOpts, err := prompt.ConfigFlow(config.ToRelativePath(keyOpts.PrivateKey))
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

			s, err := template.AsString("config", configOpts)
			if err != nil {
				return err
			}

			lines, err := parser.ParseString(s)
			if err != nil {
				return err
			}

			f.Lines = append(f.Lines, lines...)

			if err := config.WriteToFile(f.Bytes()); err != nil {
				return err
			}
			if err := prompt.Msg(prompt.MsgConfigGenerated); err != nil {
				return nil
			}
		}
	}

	return nil
}
