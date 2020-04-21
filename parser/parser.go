package parser

import (
	"fmt"

	"github.com/mickaelvieira/saxifrage/config"
	"github.com/mickaelvieira/saxifrage/lexer"
	"github.com/mickaelvieira/saxifrage/token"
)

// Parser --
type Parser struct {
	tokens chan *token.Token
	lexer  *lexer.Lexer
}

// Parse --
func (p *Parser) Parse() ([]*config.Section, error) {
	var sections []*config.Section
	var keyword string
	var section config.SectionType

	go p.lexer.Lex()

	for t := range p.tokens {
		if !t.IsComment() && !t.IsEOF() && !t.IsEOL() && !t.IsWhitespace() {
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
				return sections, fmt.Errorf("A value was expected for keyword %s", keyword)
			case section != "" && !t.IsValue():
				return sections, fmt.Errorf("A value was expected for section %s", section)
			case t.IsError():
				return sections, fmt.Errorf("Lexing error: %s", t.Value)
			default:
				return sections, fmt.Errorf("Unexpected token")
			}
		}
	}
	return sections, nil
}

// New creates a new parser
func New(i string) *Parser {
	c := make(chan *token.Token)
	return &Parser{
		tokens: c,
		lexer:  lexer.New(i, c),
	}
}
