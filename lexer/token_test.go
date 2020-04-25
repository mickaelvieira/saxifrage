package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenIsSection(t *testing.T) {
	token := &Token{Type: Section}
	assert.True(t, token.IsSection(), "Test Case %v", token.Type)
}

func TestTokenIsKeyword(t *testing.T) {
	token := &Token{Type: Keyword}
	assert.True(t, token.IsKeyword(), "Test Case %v", token.Type)
}

func TestTokenIsWhitespace(t *testing.T) {
	token := &Token{Type: Whitespace}
	assert.True(t, token.IsWhitespace(), "Test Case %v", token.Type)
}

func TestTokenIsComment(t *testing.T) {
	token := &Token{Type: Comment}
	assert.True(t, token.IsComment(), "Test Case %v", token.Type)
}

func TestTokenIsError(t *testing.T) {
	token := &Token{Type: Err}
	assert.True(t, token.IsError(), "Test Case %v", token.Type)
}

func TestTokenIsValue(t *testing.T) {
	token := &Token{Type: Value}
	assert.True(t, token.IsValue(), "Test Case %v", token.Type)
}

func TestTokenIsEOL(t *testing.T) {
	token := &Token{Type: EOL}
	assert.True(t, token.IsEOL(), "Test Case %v", token.Type)
}

func TestTokenIsEOF(t *testing.T) {
	token := &Token{Type: EOF}
	assert.True(t, token.IsEOF(), "Test Case %v", token.Type)
}

func TestTokenIsSeparator(t *testing.T) {
	token := &Token{Type: Separator}
	assert.True(t, token.IsSeparator(), "Test Case %v", token.Type)
}

func TestTokenIsIllegal(t *testing.T) {
	token := &Token{Type: Illegal}
	assert.True(t, token.IsIllegal(), "Test Case %v", token.Type)
}

func TestTokenIsHostSection(t *testing.T) {
	token := &Token{Type: Section, Value: "Host"}
	assert.True(t, token.IsHostSection(), "Test Case %v", token.Type)
}

func TestTokenIsMatchSection(t *testing.T) {
	token := &Token{Type: Section, Value: "Match"}
	assert.True(t, token.IsMatchSection(), "Test Case %v", token.Type)
}

func TestStringConvertion(t *testing.T) {
	cases := []struct {
		input string
		t     Type
		want  string
	}{
		{string(rune(0)), EOF, "EOF"},
		{string(rune(10)), EOL, "\n"},
		{" ", Whitespace, " "},
	}

	for i, tc := range cases {
		token := Token{Value: tc.input, Type: tc.t}
		got := token.String()
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}
