package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/prompt"
	"github.com/mickaelvieira/saxifrage/template"
)

func runRemove(a *App) error {
	files, err := parser.ParseFiles()
	if err != nil {
		return err
	}

	file := files.GetUserConfig()
	if file == nil {
		return config.ErrMissingUserConfig
	}

	sections, err := file.BuildSections()
	if err != nil {
		return err
	}

	s := promptui.Select{
		Label:        prompt.MsgPromptChooseSection,
		Items:        sections.GetMatchingValues(),
		HideSelected: true,
	}

	_, r, err := s.Run()
	if err != nil {
		return err
	}

	section := sections.FindSectionByMatchingValue(r)
	if section == nil {
		return fmt.Errorf("Cannot find section matching %s", r)
	}

	lines := file.FindSectionLines(section.Matching)
	keys, err := config.GetKeyFiles(section)
	if err != nil {
		return err
	}

	d := struct {
		KeyFiles []string
		Lines    config.Lines
	}{
		KeyFiles: keys,
		Lines:    lines,
	}

	if len(keys) > 0 {
		if err := template.Output("files", d); err != nil {
			return err
		}

		confirm, err := prompt.Confirm(prompt.MsgConfirmDeleteFiles)
		if err != nil {
			return err
		}

		if confirm {
			for _, keyFile := range keys {
				if err := os.Remove(keyFile); err != nil {
					return err
				}
			}
		}
	}

	if len(lines) > 0 {
		if err := template.Output("lines", d); err != nil {
			return err
		}

		r, err := prompt.Prompt(prompt.MsgConfirmDeleteLines, "ignore")
		if err != nil {
			return err
		}

		if r == "d" {
			file.RemoveLineNumbers(lines.GetNumbers())
			if err := config.WriteToFile(file.Bytes()); err != nil {
				return err
			}
		}
		if r == "c" {
			file.CommentLineNumbers(lines.GetNumbers())
			if err := config.WriteToFile(file.Bytes()); err != nil {
				return err
			}
		}
	}

	return nil
}