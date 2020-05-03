package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// WriteToFile writes the key to a file
func WriteToFile(b []byte) error {
	p := GetUserConfigPath()
	err := ioutil.WriteFile(p, b, 0600)
	if err != nil {
		return err
	}
	return nil
}

// GetKeyFiles returns the private and public keys paths
// as well as their parent directory when the key is stored in a subdirectory
func GetKeyFiles(s *Section) ([]string, error) {
	files := make([]string, 0)
	option := s.Options.Find("IdentityFile")

	if option == nil {
		return files, nil
	}

	// @TODO there is a bug here
	// The value might be between quotes

	privateKey := ToAbsolutePath(option.Value)
	publickey := privateKey + ".pub"
	directory := filepath.Dir(privateKey)

	if _, err := os.Stat(privateKey); err == nil {
		files = append(files, privateKey)
	}
	if _, err := os.Stat(privateKey); err == nil {
		files = append(files, publickey)
	}

	if !IsBaseSSHDirectory(directory) {
		keyFiles, err := ioutil.ReadDir(directory)
		if err != nil {
			if os.IsNotExist(err) {
				return files, nil
			}
			return files, err
		}
		n := len(keyFiles) / 2
		if n < 2 {
			files = append(files, directory)
		}
	}

	return files, nil
}
