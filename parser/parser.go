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
	tokens chan *lexer.Token // shared channel, I need to find a better name
	lexer  *lexer.Lexer      // our lexer to perform the lexical analysis
}

// Parse --
func (p *Parser) Parse() (config.Sections, []*lexer.Token, error) {
	var tokens []*lexer.Token
	var sections config.Sections
	var keyword string
	var section config.SectionType

	go p.lexer.Lex()

	for t := range p.tokens {
		tokens = append(tokens, t)
	}

	for _, t := range tokens {
		if !t.IsComment() && !t.IsEOF() && !t.IsEOL() && !t.IsWhitespace() && !t.IsSeparator() {
			switch {
			case t.IsSection():
				section = config.HostType
				if t.IsMatchSection() {
					section = config.MatchType
				}
			case t.IsKeyword():
				keyword = t.Value
			case t.IsValue() && section != "":
				sections = append(sections, config.NewSection(section, t.Value))
				section = ""
			case t.IsValue() && keyword != "":
				idx := len(sections) - 1
				sections[idx].Configs[keyword] = t.Value
				keyword = ""
			case keyword != "" && !t.IsValue():
				return sections, tokens, fmt.Errorf(MsgMissingKeywordValue, keyword)
			case section != "" && !t.IsValue():
				return sections, tokens, fmt.Errorf(MsgMissingSectionValue, section)
			case t.IsError():
				return sections, tokens, fmt.Errorf(MsgLexerError, t.Value)
			case t.IsIllegal():
				return sections, tokens, fmt.Errorf(MsgIllegalToken, t.Value)
			default:
				return sections, tokens, fmt.Errorf(MsgUnexpectedToken)
			}
		}
	}

	return sections, tokens, nil
}

// New creates a new parser
func New(i string) *Parser {
	c := make(chan *lexer.Token)
	return &Parser{
		tokens: c,
		lexer:  lexer.New(i, c),
	}
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
	p := New(string(c))

	s, t, err := p.Parse()
	if err != nil {
		return nil, err
	}

	return &config.File{Path: path, Sections: s, Tokens: t}, nil
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
