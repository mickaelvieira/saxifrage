package parser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/lexer"
)

// Parser messages list
const (
	msgLexerError   = "lexing error: %s"
	msgIllegalToken = "illegal token: %s"
)

// parser --
type parser struct {
	// The last error triggered during parsing
	lastError error

	// The Tokenizer
	tokenizer lexer.Tokenizer

	// Lines contains the tokens grouped by line
	Lines config.Lines
}

func (p *parser) handleErrors(in chan *lexer.Token) chan *lexer.Token {
	out := make(chan *lexer.Token)

	go func() {
		defer close(out)
		for t := range in {
			if t.IsError() {
				p.lastError = fmt.Errorf(msgLexerError, t.Value)
				return
			} else if t.IsIllegal() {
				p.lastError = fmt.Errorf(msgIllegalToken, t.Value)
				return
			} else {
				out <- t
			}
		}
	}()

	return out
}

func (p *parser) parse() error {
	n := 1
	l := &config.Line{Number: n}
	for t := range p.handleErrors(p.tokenizer.Run()) {
		l.Add(t)
		if t.IsEOL() || t.IsEOF() {
			p.Lines = append(p.Lines, l)
			n++
			l = &config.Line{Number: n}
		}
	}
	return p.lastError
}

// loadContent loads the file's content or panic if there are any errors
func loadContent(path string) (s string, err error) {
	b, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return s, err
	}
	s = string(b)
	return s, nil
}

// ParseString parses the input string
func ParseString(in string) (config.Lines, error) {
	p := &parser{tokenizer: lexer.New(in)}
	if err := p.parse(); err != nil {
		return nil, err
	}
	return p.Lines, nil
}

// ParseFile parses the file content and returns configuration file structure
func ParseFile(path string) (*config.File, error) {
	c, err := loadContent(path)
	if err != nil {
		return nil, err
	}

	l, err := ParseString(string(c))
	if err != nil {
		return nil, err
	}

	return &config.File{Path: path, Lines: l}, nil
}

// ParseFiles parses configuration files
func ParseFiles() (config.Files, error) {
	f := make(config.Files, 2)

	gc, err := ParseFile(config.GetGlobalConfigPath())
	if err != nil {
		return f, err
	}

	uc, err := ParseFile(config.GetUserConfigPath())
	if err != nil {
		return f, err
	}

	f[0] = gc
	f[1] = uc

	return f, nil
}
