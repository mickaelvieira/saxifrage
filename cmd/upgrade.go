package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mickaelvieira/saxifrage/prompt"
	"github.com/mickaelvieira/saxifrage/upgrade"
)

func runUpgrade(a *App) error {
	filename, err := upgrade.GetPlatformArchiveName()
	if err != nil {
		return err
	}

	if e := prompt.Msg("Checking for latest version"); e != nil {
		return e
	}

	latest, err := upgrade.GetLatestVersion()
	if err != nil {
		return err
	}

	if e := prompt.Msg("Version has been found"); e != nil {
		return e
	}

	shouldUpdate, err := upgrade.CompareVersions(a.Version, latest)
	if err != nil {
		return err
	}

	if !shouldUpdate {
		if e := prompt.Msg(fmt.Sprintf("%s is upgrading to version %s", a.Name, latest)); e != nil {
			return e
		}

		tempDir, err := ioutil.TempDir(os.TempDir(), "saxifrage-*")
		if err != nil {
			return err
		}
		if e := os.Chdir(tempDir); e != nil {
			return e
		}
		defer os.RemoveAll(tempDir)

		if e := upgrade.Download(filename, latest); e != nil {
			return e
		}

		out, err := upgrade.Unpack(filename)
		fmt.Printf("%s", out)
		if err != nil {
			return err
		}

		if e := upgrade.ReplaceBinary(); e != nil {
			return e
		}
	} else {
		if e := prompt.Msg(fmt.Sprintf("%s is already up-to-date", a.Name)); e != nil {
			return e
		}
	}

	return nil
}
