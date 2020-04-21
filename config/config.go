package config

import (
	"bytes"
	"io"
	"io/ioutil"
	"path/filepath"
)

// LoadContent loads ssh_config file's content
func LoadContent(path string) (io.Reader, error) {
	b, err := ioutil.ReadFile(filepath.Clean(path))

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}
