package config

import (
	"errors"

	"github.com/mickaelvieira/saxifrage/lexer"
)

// Parsing errors
var (
	ErrExpectedKeyword   = errors.New("A keyword was expected")
	ErrExpectedSeparated = errors.New("A separator was expected")
	ErrExpectedValue     = errors.New("A value was expected")
)

// keyword is composed of a 3 tokens, a name, a separator, and a value
type keyword struct {
	tokens []*lexer.Token
}

func (k *keyword) name() *lexer.Token {
	return k.tokens[0]
}

func (k *keyword) separator() *lexer.Token {
	return k.tokens[1]
}

func (k *keyword) value() *lexer.Token {
	return k.tokens[2]
}

func (k *keyword) add(t *lexer.Token) (err error) {
	switch len(k.tokens) {
	case 0:
		if !t.IsKeyword() && !t.IsSection() {
			err = ErrExpectedKeyword
		}
	case 1:
		if !t.IsSeparator() {
			err = ErrExpectedSeparated
		}
	case 2:
		if !t.IsValue() {
			err = ErrExpectedValue
		}
	}

	if err == nil {
		k.tokens = append(k.tokens, t)
	}
	return err
}

func (k *keyword) isEmpty() bool {
	return len(k.tokens) == 0
}

func (k *keyword) isComplete() bool {
	return len(k.tokens) == 3
}
