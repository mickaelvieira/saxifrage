package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/parser"
	"github.com/mickaelvieira/saxifrage/prompt"
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
		return fmt.Errorf("cannot find section matching %s", r)
	}

	lines := file.FindSectionLines(section.Matching)
	keyPath := section.GetIdentityFile()
	keyFiles, err := config.GetKeyFiles(keyPath)
	if err != nil && err != config.ErrMissingIdentityFileValue {
		return err
	}

	d := struct {
		KeyFiles []string
		Lines    config.Lines
	}{
		KeyFiles: keyFiles,
		Lines:    lines,
	}

	if len(keyFiles) > 0 {
		if err := a.Templates.Output("files", d); err != nil {
			return err
		}

		confirm, err := a.Prompt.Confirm(prompt.MsgConfirmDeleteFiles)
		if err != nil {
			return err
		}

		if confirm {
			for _, keyFile := range keyFiles {
				if err := os.Remove(keyFile); err != nil {
					return err
				}
			}
			keyDir, err := config.GetKeyDir(keyPath)
			if err == nil {
				if err := os.Remove(keyDir); err != nil {
					return err
				}
			}
		}
	}

	if len(lines) > 0 {
		if err := a.Templates.Output("lines", d); err != nil {
			return err
		}

		r, err := a.Prompt.Prompt(prompt.MsgConfirmDeleteLines, "d")
		if err != nil {
			return err
		}

		if r == "c" {
			file.CommentLineNumbers(lines.GetNumbers())
			if err := config.WriteToFile(file.Bytes()); err != nil {
				return err
			}
		} else {
			file.RemoveLineNumbers(lines.GetNumbers())
			if err := config.WriteToFile(file.Bytes()); err != nil {
				return err
			}
		}
	}

	return nil
}
