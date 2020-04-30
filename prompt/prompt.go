package prompt

import (
	"bufio"
	"os"

	"github.com/mickaelvieira/saxifrage/template"
)

// Prompt messages
const (
	MsgConfirmOverride     = "The key already exists. Do you want to override it"
	MsgConfirmContinue     = "Do you want to continue"
	MsgConfirmAddition     = "Do you want to add this key to your config file"
	MsgPromptKeyType       = "Select the type of key you want to generate"
	MsgPromptKeyComplexity = "Select the key complexity"
	MsgPromptKeyDirectory  = "Enter the directory"
	MsgPromptKeyFilename   = "Enter the file name"
	MsgPromptKeyPassphrase = "Enter the passphrase"
	MsgPromptKeyHost       = "Enter the host to which you want to associate this key"
	MsgPromptKeyPort       = "Enter the port"
	MsgPromptKeyUser       = "Enter the user"
)

// Msg output a message in the console
func Msg(m string) error {
	o := struct {
		Text string
	}{
		Text: m,
	}

	if err := template.Output("message", o); err != nil {
		return err
	}

	return nil
}

// Confirm asks for user confirmation
func Confirm(t string) (bool, error) {
	o := struct {
		Text    string
		Default string
	}{
		Text: t,
	}

	var c bool

	if err := template.Output("ask-confirm", o); err != nil {
		return c, err
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := s.Text()
		if t == "y" {
			c = true
		}
		break
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return c, nil
}

// Prompt prompts user
func Prompt(t string, d string) (string, error) {
	o := struct {
		Text    string
		Default string
	}{
		Text:    t,
		Default: d,
	}

	var i string

	if err := template.Output("read-input", o); err != nil {
		return i, err
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		i = s.Text()
		break
	}

	if err := s.Err(); err != nil {
		return i, err
	}
	return i, nil
}
