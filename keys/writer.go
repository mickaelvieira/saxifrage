package keys

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// GetDefaultFilenames returns a pair of string based on the type of key
// The first value is the name of the private key
// The second value is the name of the public key
func GetDefaultFilenames(t Type) (pk string, pub string) {
	pk = fmt.Sprintf("id_%s", string(t))
	pub = fmt.Sprintf("%s.pub", pk)
	return pk, pub
}

// GetUserFilenames returns a pair of string based on user's input
// The first value is the name of the private key
// The second value is the name of the public key
func GetUserFilenames(t string) (pk string, pub string) {
	pk = fmt.Sprintf("%s", string(t))
	pub = fmt.Sprintf("%s.pub", pk)
	return pk, pub
}

// GetDir get the destination directory
func GetDir(path string) string {
	return filepath.Join(os.Getenv("HOME"), ".ssh", path)
}

// WriteToFile writes the key to a file
func WriteToFile(b []byte, path string) error {
	err := ioutil.WriteFile(filepath.Clean(path), b, 0600)
	if err != nil {
		return err
	}
	return nil
}

// MakeDir makes destination directory if it does not exist
func MakeDir(path string) error {
	i, err := os.Lstat(path)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		if err := os.Mkdir(path, 0700); err != nil {
			return err
		}
	} else if !i.IsDir() {
		return fmt.Errorf("File '%s' exists but it is not a directory", path)
	}
	return nil
}
