package lexer

import (
	"fmt"
	"testing"
)

func TestTokenIsSection(t *testing.T) {
	token := &Token{Type: Section}
	if !token.IsSection() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestTokenIsKeyword(t *testing.T) {
	token := &Token{Type: Keyword}
	if !token.IsKeyword() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestTokenIsWhitespace(t *testing.T) {
	token := &Token{Type: Whitespace}
	if !token.IsWhitespace() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestTokenIsComment(t *testing.T) {
	token := &Token{Type: Comment}
	if !token.IsComment() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestTokenIsError(t *testing.T) {
	token := &Token{Type: Err}
	if !token.IsError() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestTokenIsValue(t *testing.T) {
	token := &Token{Type: Value}
	if !token.IsValue() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestTokenIsEOL(t *testing.T) {
	token := &Token{Type: EOL}
	if !token.IsEOL() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestTokenIsEOF(t *testing.T) {
	token := &Token{Type: EOF}
	if !token.IsEOF() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestTokenIsHostSection(t *testing.T) {
	token := &Token{Type: Section, Value: "Host"}
	if !token.IsHostSection() {
		t.Errorf("Failed for %v %v", token.Type, token.Value)
	}
}

func TestTokenIsMatchSection(t *testing.T) {
	token := &Token{Type: Section, Value: "Match"}
	if !token.IsMatchSection() {
		t.Errorf("Failed for %v %v", token.Type, token.Value)
	}
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			token := Token{Value: tt.input, Type: tt.t}
			got := token.String()
			if got != tt.want {
				t.Errorf("Want '%s', got '%s'", tt.want, got)
			}
		})
	}
}
