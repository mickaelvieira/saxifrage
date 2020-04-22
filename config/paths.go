package config

import (
	"os"
	"path/filepath"
)

// GetGlobalConfigPath returns the global ssh configuration path
func GetGlobalConfigPath() string {
	return filepath.Join("/", "etc", "ssh", "ssh_config")
}

// GetUserConfigPath returns the user ssh configuration path
func GetUserConfigPath() string {
	return filepath.Join(os.Getenv("HOME"), ".ssh", "config")
}
