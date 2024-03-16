package config

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// WriteToFile writes the key to a file
func WriteToFile(b []byte) error {
	p := GetUserConfigPath()
	err := os.WriteFile(p, b, 0600)
	if err != nil {
		return err
	}
	return nil
}

// IsDirEmpty is the directory empty?
func IsDirEmpty(path string) (bool, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return false, err
	}

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}

	err = f.Close()
	if err != nil {
		return false, err
	}

	return false, err
}

// GetKeyFiles returns the private and public keys paths
// as well as their parent directory when the key is stored in a subdirectory
func GetKeyFiles(s string) ([]string, error) {
	files := make([]string, 0)
	if s == "" {
		return files, ErrMissingIdentityFileValue
	}

	privateKey := ToAbsolutePath(s)
	publickey := privateKey + ".pub"

	if _, err := os.Stat(privateKey); err == nil {
		files = append(files, privateKey)
	}
	if _, err := os.Stat(privateKey); err == nil {
		files = append(files, publickey)
	}

	return files, nil
}

// GetKeyDir returns the keys' directory path
func GetKeyDir(s string) (dp string, err error) {
	if s == "" {
		return dp, ErrMissingIdentityFileValue
	}

	fp := ToAbsolutePath(strings.Trim(s, "\""))
	dp = filepath.Dir(fp)

	if IsBaseSSHDirectory(dp) {
		return dp, ErrIsSSHBasedDirection
	}

	empty, err := IsDirEmpty(dp)
	if err != nil || !empty {
		return dp, ErrDirectoryIsNotEmpty
	}

	return dp, nil
}
