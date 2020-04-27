package prompt

import (
	"bufio"
	"os"

	"github.com/mickaelvieira/saxifrage/template"
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
