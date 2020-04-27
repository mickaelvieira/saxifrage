package prompt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/mickaelvieira/saxifrage/template"
)

// Msg output a message in the console
func Msg(m string) {
	f := promptui.Styler(promptui.FGBold, promptui.FGGreen)
	fmt.Printf(" %s\n", f(m))
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

	if err := template.Render("ask-confirm", o); err != nil {
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

	if err := template.Render("read-input", o); err != nil {
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
