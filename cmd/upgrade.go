package cmd

import (
	"fmt"
	"os"

	"github.com/mickaelvieira/saxifrage/upgrade"
)

func runUpgrade(a *App) error {
	filename, err := upgrade.GetPlatformArchiveName()
	if err != nil {
		return err
	}

	if e := a.Prompt.Msg("Checking for latest version"); e != nil {
		return e
	}

	latest, err := upgrade.GetLatestVersion()
	if err != nil {
		return err
	}

	if e := a.Prompt.Msg("Version has been found"); e != nil {
		return e
	}

	shouldUpdate, err := upgrade.CompareVersions(a.Version, latest)
	if err != nil {
		return err
	}

	if shouldUpdate {
		if e := a.Prompt.Msg(fmt.Sprintf("%s is upgrading to version %s", a.Name, latest)); e != nil {
			return e
		}
		binDir, err := upgrade.GetExecutableDir()
		if err != nil {
			return err
		}
		tempDir, err := os.MkdirTemp(binDir, "saxifrage-*")
		if err != nil {
			return err
		}
		if e := os.Chdir(tempDir); e != nil {
			return e
		}
		defer os.RemoveAll(tempDir)

		if e := upgrade.Download(a.Prompt, filename, latest); e != nil {
			return e
		}
		if e := upgrade.Unpack(filename); e != nil {
			return e
		}
		if e := upgrade.ReplaceBinary(binDir); e != nil {
			return e
		}
		if e := a.Prompt.Msg(fmt.Sprintf("%s was upgraded successfully", a.Name)); e != nil {
			return e
		}
	} else {
		if e := a.Prompt.Msg(fmt.Sprintf("%s is already up-to-date", a.Name)); e != nil {
			return e
		}
	}

	return nil
}
