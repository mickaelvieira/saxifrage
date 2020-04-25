package parser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/lexer"
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
				p.lastError = fmt.Errorf(MsgLexerError, t.Value)
				return
			}
			if t.IsIllegal() {
				p.lastError = fmt.Errorf(MsgIllegalToken, t.Value)
				return
			}
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
			k := tokens[0]
			s := tokens[1]
			v := tokens[2]

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

// ParseFile parses the file content and returns configuration file structure
func ParseFile(path string) (*config.File, error) {
	c, err := loadContent(path)
	if err != nil {
		return nil, err
	}
	p := &Parser{lexer: lexer.New(string(c))}

	if err := p.Parse(); err != nil {
		return nil, err
	}

	return &config.File{Path: path, Sections: p.Sections, Tokens: p.Tokens}, nil
}

// ParseFiles parses configuration files
func ParseFiles() ([]*config.File, error) {
	f := make([]*config.File, 2)

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
