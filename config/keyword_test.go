package config

import (
	"testing"

	"github.com/mickaelvieira/saxifrage/lexer"
	"github.com/stretchr/testify/assert"
)

func TestKeywordFirstToken(t *testing.T) {
	b := &keyword{}

	err := b.add(&lexer.Token{})
	assert.Equal(t, err, ErrExpectedKeyword)

	err = b.add(&lexer.Token{Type: lexer.Keyword})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(b.tokens))
}

func TestKeywordFirstTokenWithSection(t *testing.T) {
	b := &keyword{}

	err := b.add(&lexer.Token{})
	assert.Equal(t, err, ErrExpectedKeyword)

	err = b.add(&lexer.Token{Type: lexer.Section})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(b.tokens))
}

func TestKeywordSecondToken(t *testing.T) {
	b := &keyword{}

	err := b.add(&lexer.Token{Type: lexer.Keyword})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(b.tokens))

	err = b.add(&lexer.Token{})
	assert.Equal(t, err, ErrExpectedSeparated)

	err = b.add(&lexer.Token{Type: lexer.Separator})
	assert.Nil(t, err)
	assert.Equal(t, 2, len(b.tokens))
}

func TestKeywordThirdToken(t *testing.T) {
	b := &keyword{}

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
