package config

import (
	"github.com/mickaelvieira/saxifrage/lexer"
)

// File a SSH configuration file
type File struct {
	Path     string
	Sections Sections
	Tokens   []*lexer.Token
}

func (f *File) String() (s string) {
	for _, t := range f.Tokens {
		s += t.String()
	}
	return s
}
