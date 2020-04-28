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
	msgNullSection         = "Cannot add options before declaring a section"
)

// Parser --
type Parser struct {
	// The last error triggered during parsing
	lastError error

	// The Tokenizer
	tokenizer lexer.Tokenizer

	// Sections gathers the sections
	// built with the tokens received
	// from the tokenizer
	Sections config.Sections

	// Tokens gathers all the tokens
	// by the tokenizer
	Tokens []*lexer.Token
}

func (p *Parser) collect(in chan *lexer.Token) chan *lexer.Token {
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

func (p *Parser) groupKeysValues(in chan *lexer.Token) chan []*lexer.Token {
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

func (p *Parser) buildSections(in chan []*lexer.Token) chan *config.Section {
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
				if se == nil {
					p.lastError = fmt.Errorf(msgNullSection)
					return
				}
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
	t := p.collect(p.tokenizer.Run())

	for section := range p.buildSections(p.groupKeysValues(t)) {
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
	p := &Parser{tokenizer: lexer.New(in)}
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
