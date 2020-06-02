package prompt

import "github.com/mickaelvieira/saxifrage/config"

func askForKeyHost(p *Prompt, o *config.Generated) error {
	r, err := p.Prompt(MsgPromptKeyHost, "")
	if err != nil {
		return err
	}
	o.Host = r
	return nil
}

func askForPort(p *Prompt, o *config.Generated) error {
	o.Port = "22"
	r, err := p.Prompt(MsgPromptKeyPort, o.Port)
	if err != nil {
		return err
	}
	if r != "" {
		o.Port = r
	}
	return nil
}

func askForUser(p *Prompt, o *config.Generated) error {
	r, err := p.Prompt(MsgPromptKeyUser, "")
	if err != nil {
		return err
	}
	o.User = r
	return nil
}

// ConfigFlow prompts user to get the configuration options
func ConfigFlow(p *Prompt, k string) (*config.Generated, error) {
	o := &config.Generated{IdentityFile: k}

	if err := askForKeyHost(p, o); err != nil {
		return nil, err
	}
	if err := askForPort(p, o); err != nil {
		return nil, err
	}
	if err := askForUser(p, o); err != nil {
		return nil, err
	}

	return o, nil
}
