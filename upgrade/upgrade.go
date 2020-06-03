package upgrade

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/mickaelvieira/saxifrage/prompt"
)

const binFile = "sax"
const archiveURL = "https://github.com/mickaelvieira/saxifrage/releases/download/v%s/%s"
const versionFile = "https://raw.githubusercontent.com/mickaelvieira/saxifrage/master/.github/.version"

// GetExecutableDir returns the directory from which the binary is running
func GetExecutableDir() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	return filepath.Dir(path), nil
}

// CompareVersions compares the current and the latest versions
func CompareVersions(c, l string) (bool, error) {
	current, err := version.NewVersion(c)
	if err != nil {
		return false, err
	}
	latest, err := version.NewVersion(l)
	if err != nil {
		return false, err
	}
	return current.LessThan(latest), nil
}

// Download download the zip file from github and write it into a file
func Download(p *prompt.Prompt, filename, version string) error {
	url := fmt.Sprintf(archiveURL, version, filename)

	if e := p.Msg(fmt.Sprintf("downloading %s", url)); e != nil {
		return e
	}
	r, err := http.Get(url) // #nosec
	if err != nil {
		return err
	}
	defer r.Body.Close()

	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if e := ioutil.WriteFile(filepath.Clean(filename), c, 0600); e != nil {
		return e
	}

	return err
}

// Unpack unzips the file
func Unpack(file string) error {
	cmd := exec.Command("unzip", file) // #nosec
	out, err := cmd.Output()
	if err == nil {
		fmt.Printf("%s", out)
	}
	return err
}

// GetLatestVersion fetches the latest application's version from the github repository
func GetLatestVersion() (string, error) {
	r, err := http.Get(versionFile)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	s := strings.Trim(string(b), " \n")
	if s == "" {
		return s, fmt.Errorf("Cannot find version")
	}

	return s, nil
}

// GetPlatformArchiveName retrieves the zip filename for the current platform
func GetPlatformArchiveName() (string, error) {
	if runtime.GOARCH != "amd64" {
		return "", fmt.Errorf("Arch %s is not supported", runtime.GOARCH)
	}
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		return "", fmt.Errorf("Operating system %s is not supported", runtime.GOOS)
	}
	return fmt.Sprintf("saxifrage-%s-%s.zip", runtime.GOOS, runtime.GOARCH), nil
}

// ReplaceBinary replaces the existing binary file
func ReplaceBinary(binDir string) error {
	source := filepath.Clean(fmt.Sprintf("./%s", binFile))
	destination := filepath.Clean(fmt.Sprintf("%s/%s", binDir, binFile))

	if e := os.Rename(source, destination); e != nil {
		return e
	}

	return nil
}
