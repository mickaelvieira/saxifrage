package config

import (
	"errors"
	"io/ioutil"

	"github.com/mickaelvieira/saxifrage/lexer"
)

// Files errors
var (
	ErrMissingUserConfig = errors.New("Unable to find user configuration file")
)

// Files list of ssh_config files
type Files []*File

// GetUserConfig retrieves the user configuration file
func (f Files) GetUserConfig() *File {
	for _, f := range f {
		if f.IsUserConfig() {
			return f
		}
	}
	return nil
}

// File a SSH configuration file
type File struct {
	Path     string
	Sections Sections
	Tokens   []*lexer.Token
}

// String returns the file content as a string
func (f *File) String() (s string) {
	for _, t := range f.Tokens {
		s += t.String()
	}
	return s
}

// Bytes returns the file content as a slice of bytes
func (f *File) Bytes() (b []byte) {
	for _, t := range f.Tokens {
		b = append(b, t.ToBytes()...)
	}
	return b
}

// IsUserConfig identifies whether the file is the user local configuration
func (f *File) IsUserConfig() bool {
	return f.Path == GetUserConfigPath()
}

// WriteToFile writes the key to a file
func WriteToFile(b []byte) error {
	p := GetUserConfigPath()
	err := ioutil.WriteFile(p, b, 0600)
	if err != nil {
		return err
	}
	return nil
}
