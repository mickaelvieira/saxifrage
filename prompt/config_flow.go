package prompt

import "github.com/mickaelvieira/saxifrage/config"

func askForKeyHost(o *config.Generated) error {
	r, err := Prompt(MsgPromptKeyHost, "")
	if err != nil {
		return err
	}
	o.Host = r
	return nil
}

func askForPort(o *config.Generated) error {
	o.Port = "22"
	r, err := Prompt(MsgPromptKeyPort, o.Port)
	if err != nil {
		return err
	}
	if r != "" {
		o.Port = r
	}
	return nil
}

func askForUser(o *config.Generated) error {
	r, err := Prompt(MsgPromptKeyUser, "")
	if err != nil {
		return err
	}
	o.User = r
	return nil
}

// ConfigFlow prompts user to get the configuration options
func ConfigFlow(p string) (*config.Generated, error) {
	o := &config.Generated{IdentityFile: p}

	if err := askForKeyHost(o); err != nil {
		return nil, err
	}
	if err := askForPort(o); err != nil {
		return nil, err
	}
	if err := askForUser(o); err != nil {
		return nil, err
	}

	return o, nil
}
