package parser

import (
	"fmt"

	"github.com/mickaelvieira/saxifrage/lexer"
	"github.com/mickaelvieira/saxifrage/token"
)

// SSHConfig --
type SSHConfig struct {
	Host   string
	Config map[string]string
}

// NewConfig --
func NewConfig(hosts string) *SSHConfig {
	return &SSHConfig{
		Host:   hosts,
		Config: make(map[string]string),
	}
}

// Parser --
type Parser struct {
	tokens chan *token.Token
	lexer  *lexer.Lexer
}

// Parse --
func (p *Parser) Parse() ([]*SSHConfig, error) {
	var configs []*SSHConfig
	var lastKey string
	var expectValue bool
	var expectMatch bool

	go p.lexer.Lex()

	for t := range p.tokens {
		if !t.IsComment() && !t.IsEOF() && !t.IsEOL() && !t.IsWhitespace() {
			if t.IsSection() {
				expectMatch = true
			} else if t.IsKeyword() {
				lastKey = t.Value
				expectValue = true
			} else if t.IsValue() && expectMatch {
				configs = append(configs, NewConfig(t.Value))
				expectMatch = false
			} else if t.IsValue() && expectValue {
				if lastKey == "" {
					return configs, fmt.Errorf("lastKey is not defined")
				}
				configs[len(configs)-1].Config[lastKey] = t.Value
				expectValue = false
			} else if (expectMatch || expectValue) && !t.IsValue() {
				return configs, fmt.Errorf("A value was expected")
			} else {
				return configs, fmt.Errorf("Unexpected token")
			}
		}
	}
	return configs, nil
}

// New creates a new parser
func New(i string) *Parser {
	c := make(chan *token.Token)
	return &Parser{
		tokens: c,
		lexer:  lexer.New(i, c),
	}
}
