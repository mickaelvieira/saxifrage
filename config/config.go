package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func mustLoadContent(path string) string {
	b, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getGlobalConfigPath() string {
	return filepath.Join("etc", "ssh", "config")
}

func getUserConfigPath() string {
	return filepath.Join(os.Getenv("HOME"), ".ssh", "config")
}

// GetGlobalConfigContent returns the content of the global ssh_config
func GetGlobalConfigContent() string {
	return mustLoadContent(filepath.Join("etc", "ssh", "config"))
}

// GetUserConfigContent returns the content of the user ssh_config
func GetUserConfigContent() string {
	return mustLoadContent(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))
}
