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
func Unpack(file string) ([]byte, error) {
	cmd := exec.Command("unzip", file) // #nosec
	return cmd.Output()
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
func ReplaceBinary() error {
	dir, err := os.Executable()
	if err != nil {
		return err
	}
	dir, err = filepath.EvalSymlinks(dir)
	if err != nil {
		return err
	}
	dir = filepath.Dir(dir)

	source := filepath.Clean(fmt.Sprintf("./%s", binFile))
	destination := filepath.Clean(fmt.Sprintf("%s/%s", dir, binFile))

	// Since it is not possible to rename a file across different partitions
	// this line will fail if the binary is not on the same partition than the "tmp" directory/
	// A solution to this issue could be to copy the file to the partition first and then rename it.
	if e := os.Rename(source, destination); e != nil {
		return e
	}

	return nil
}
