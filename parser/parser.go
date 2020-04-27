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
	msgMissingKeywordValue = "A value was expected for keyword %s"
	msgMissingSectionValue = "A value was expected for section %s"
	msgLexerError          = "Lexing error: %s"
	msgIllegalToken        = "Illegal token: %s"
	msgUnexpectedToken     = "Unexpected token"
	msgNotEmptyGroup       = "Token is not empty but not full either"
)

// Parser --
type Parser struct {
	lastError error
	Sections  config.Sections
	Tokens    []*lexer.Token // shared channel, I need to find a better name
	lexer     *lexer.Lexer   // our lexer to perform the lexical analysis
}

func (p *Parser) store(in chan *lexer.Token) chan *lexer.Token {
	out := make(chan *lexer.Token)

	go func() {
		defer close(out)
		for t := range in {
			p.Tokens = append(p.Tokens, t)
			out <- t
		}
	}()

	return out
}

func (p *Parser) group(in chan *lexer.Token) chan []*lexer.Token {
	out := make(chan []*lexer.Token)

	go func() {
		defer close(out)

		b := &group{}
		for t := range in {
			if t.IsSection() ||
				t.IsKeyword() ||
				t.IsSeparator() ||
				t.IsValue() {
				if err := b.add(t); err != nil {
					p.lastError = err
					return
				}
				if b.isFull() {
					out <- b.tokens
					b.clear()
				}
			}
			if t.IsError() {
				p.lastError = fmt.Errorf(msgLexerError, t.Value)
				return
			}
			if t.IsIllegal() {
				p.lastError = fmt.Errorf(msgIllegalToken, t.Value)
				return
			}
		}

		if !b.isEmpty() {
			p.lastError = fmt.Errorf(msgNotEmptyGroup)
		}
	}()

	return out
}

func (p *Parser) build(in chan []*lexer.Token) chan *config.Section {
	out := make(chan *config.Section)

	go func() {
		defer close(out)

		var se *config.Section

		for tokens := range in {
			k := tokens[0] // keyword
			s := tokens[1] // separator
			v := tokens[2] // value

			if k.IsSection() {
				if se != nil {
					out <- se
				}
				t := config.HostType
				if k.IsMatchSection() {
					t = config.MatchType
				}
				se = config.NewSection(t, s.Value, v.Value)
			} else {
				se.Options = append(se.Options, config.NewOption(k.Value, s.Value, v.Value))
			}
		}
		if se != nil {
			out <- se
		}
	}()

	return out
}

// Parse --
func (p *Parser) Parse() error {
	t := p.store(p.lexer.Lex())
	g := p.group(t)
	s := p.build(g)

	for section := range s {
		p.Sections = append(p.Sections, section)
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
func ParseString(in string) (config.Sections, []*lexer.Token, error) {
	p := &Parser{lexer: lexer.New(in)}
	if err := p.Parse(); err != nil {
		return nil, nil, err
	}
	return p.Sections, p.Tokens, nil
}

// ParseFile parses the file content and returns configuration file structure
func ParseFile(path string) (*config.File, error) {
	c, err := loadContent(path)
	if err != nil {
		return nil, err
	}

	s, t, err := ParseString(string(c))
	if err != nil {
		return nil, err
	}

	return &config.File{Path: path, Sections: s, Tokens: t}, nil
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
