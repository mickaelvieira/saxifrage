package config

import (
	"testing"

	"github.com/mickaelvieira/saxifrage/lexer"
	"github.com/stretchr/testify/assert"
)

func TestKeyValueFirstToken(t *testing.T) {
	b := &keyValue{}

	err := b.add(&lexer.Token{})
	assert.Equal(t, err, ErrExpectedKeyword)

	err = b.add(&lexer.Token{Type: lexer.Keyword})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(b.tokens))
}

func TestKeyValueFirstTokenWithSection(t *testing.T) {
	b := &keyValue{}

	err := b.add(&lexer.Token{})
	assert.Equal(t, err, ErrExpectedKeyword)

	err = b.add(&lexer.Token{Type: lexer.Section})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(b.tokens))
}

func TestKeyValueSecondToken(t *testing.T) {
	b := &keyValue{}

	err := b.add(&lexer.Token{Type: lexer.Keyword})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(b.tokens))

	err = b.add(&lexer.Token{})
	assert.Equal(t, err, ErrExpectedSeparated)

	err = b.add(&lexer.Token{Type: lexer.Separator})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(b.tokens))
}

func TestKeyValueThirdToken(t *testing.T) {
	b := &keyValue{}

	err := b.add(&lexer.Token{Type: lexer.Keyword})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(b.tokens))

	err = b.add(&lexer.Token{Type: lexer.Separator})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(b.tokens))

	err = b.add(&lexer.Token{})
	assert.Equal(t, err, ErrExpectedValue)

	err = b.add(&lexer.Token{Type: lexer.Value})
	assert.Nil(t, err)
	assert.Equal(t, 3, len(b.tokens))
}
