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

type keyValue struct {
	tokens []*lexer.Token
}

func (b *keyValue) add(t *lexer.Token) (err error) {
	switch len(b.tokens) {
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
		b.tokens = append(b.tokens, t)
	}
	return err
}

func (b *keyValue) isEmpty() bool {
	return len(b.tokens) == 0
}

func (b *keyValue) isComplete() bool {
	return len(b.tokens) == 3
}
