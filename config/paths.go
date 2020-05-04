package config

import (
	"os"
	"path/filepath"
	"strings"
)

var base = filepath.Join(os.Getenv("HOME"), ".ssh")

// GetGlobalConfigPath returns the global ssh configuration path
func GetGlobalConfigPath() string {
	return filepath.Join("/", "etc", "ssh", "ssh_config")
}

// GetUserConfigPath returns the user ssh configuration path
func GetUserConfigPath() string {
	return filepath.Join(base, "config")
}

// IsBaseSSHDirectory returns true if the path is the base SSH directory
func IsBaseSSHDirectory(p string) bool {
	return p == base || p == ToRelativePath(base)
}

// ToRelativePath replaces path to home directory with tilde
func ToRelativePath(p string) string {
	return strings.Replace(p, os.Getenv("HOME"), "~", 1)
}

// ToAbsolutePath replaces tilde with path to home directory
func ToAbsolutePath(p string) string {
	return strings.Replace(p, "~", os.Getenv("HOME"), 1)
}
