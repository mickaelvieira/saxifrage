package prompt

import (
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/mickaelvieira/saxifrage/keys"
)

func askForKeyType(o *keys.Options) error {
	s := promptui.Select{
		Label:        MsgPromptKeyType,
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

func askForKeySize(o *keys.Options) error {
	keySizes := keys.GetKeySize(o.KeyType)
	values := keySizes.GetValues()

	if len(values) == 0 {
		return nil
	}

	s := promptui.Select{
		Label:        MsgPromptKeyComplexity,
		Items:        values,
		HideSelected: true,
	}

	_, r, err := s.Run()
	if err != nil {
		return err
	}

	v := keySizes.GetValue(r)
	if v == nil {
		return keys.ErrInvalidKeySize
	}
	o.KeySize = v

	return nil
}

func askForKeyDirectory(p *Prompt, o *keys.Options) error {
	o.Directory = keys.GetDir("")

	r, err := p.Prompt(MsgPromptKeyDirectory, o.Directory)
	if err != nil {
		return err
	}
	if r != "" {
		o.Directory = keys.GetDir(r)
	}

	return nil
}

func askForKeyPassPhrase(p *Prompt, o *keys.Options) error {
	r, err := p.Prompt(MsgPromptKeyPassphrase, "")
	if err != nil {
		return err
	}
	o.PassPhrase = r

	return nil
}

func askForKeyName(p *Prompt, o *keys.Options) error {
	s, _ := keys.GetFilenamesFromType(o.KeyType)
	r, err := p.Prompt(MsgPromptKeyFilename, s)
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

// KeyFlow prompts user to get the keys options
func KeyFlow(p *Prompt) (*keys.Options, error) {
	o := &keys.Options{}

	if err := askForKeyType(o); err != nil {
		return nil, err
	}
	if err := askForKeyDirectory(p, o); err != nil {
		return nil, err
	}
	if err := askForKeyName(p, o); err != nil {
		return nil, err
	}
	if err := askForKeyPassPhrase(p, o); err != nil {
		return nil, err
	}
	if err := askForKeySize(o); err != nil {
		return nil, err
	}

	return o, nil
}
