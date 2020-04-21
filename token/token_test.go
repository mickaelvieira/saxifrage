package token

import (
	"testing"
)

func TestIsSection(t *testing.T) {
	token := &Token{Type: Section}
	if !token.IsSection() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestIsKeyword(t *testing.T) {
	token := &Token{Type: Keyword}
	if !token.IsKeyword() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestIsWhitespace(t *testing.T) {
	token := &Token{Type: Whitespace}
	if !token.IsWhitespace() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestIsComment(t *testing.T) {
	token := &Token{Type: Comment}
	if !token.IsComment() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestIsError(t *testing.T) {
	token := &Token{Type: Err}
	if !token.IsError() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestIsValue(t *testing.T) {
	token := &Token{Type: Value}
	if !token.IsValue() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestIsEOL(t *testing.T) {
	token := &Token{Type: EOL}
	if !token.IsEOL() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestIsEOF(t *testing.T) {
	token := &Token{Type: EOF}
	if !token.IsEOF() {
		t.Errorf("Failed for %v", token.Type)
	}
}

func TestIsHostSection(t *testing.T) {
	token := &Token{Type: Section, Value: "Host"}
	if !token.IsHostSection() {
		t.Errorf("Failed for %v %v", token.Type, token.Value)
	}
}

func TestIsMatchSection(t *testing.T) {
	token := &Token{Type: Section, Value: "Match"}
	if !token.IsMatchSection() {
		t.Errorf("Failed for %v %v", token.Type, token.Value)
	}
}
