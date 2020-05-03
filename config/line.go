package config

import "github.com/mickaelvieira/saxifrage/lexer"

// Lines are the lines in the configuration file
type Lines []*Line

// GetNumbers returns the lines' numbers
func (l Lines) GetNumbers() []int {
	n := make([]int, len(l))
	for i, line := range l {
		n[i] = line.Number
	}
	return n
}

// Line is a line in the configuration file
type Line struct {
	Number int
	tokens []*lexer.Token
}

// Comment comments the line out
func (l *Line) Comment() {
	c := &lexer.Token{Type: lexer.Comment, Value: "# "}
	l.tokens = append([]*lexer.Token{c}, l.tokens...)
}

// String returns the line content as a string
func (l *Line) String() (s string) {
	for _, t := range l.tokens {
		s += t.String()
	}
	return s
}

// Bytes returns the line content as a slice of bytes
func (l *Line) Bytes() (b []byte) {
	for _, t := range l.tokens {
		b = append(b, t.Bytes()...)
	}
	return b
}

// Add adds a token to the list
func (l *Line) Add(t *lexer.Token) {
	l.tokens = append(l.tokens, t)
}

// IsComment is the line a comment?
func (l *Line) IsComment() bool {
	if len(l.tokens) == 0 {
		return false
	}
	for _, t := range l.tokens {
		if !t.IsWhitespace() && !t.IsComment() && !t.IsEOL() {
			return false
		}
	}
	return true
}

// IsEmpty is the line empty?
func (l *Line) IsEmpty() bool {
	if len(l.tokens) == 0 {
		return true
	}
	for _, t := range l.tokens {
		if !t.IsEOL() && !t.IsWhitespace() {
			return false
		}
	}
	return true
}

// IsSection is the line the beginning of a section?
func (l *Line) IsSection() bool {
	if len(l.tokens) == 0 {
		return false
	}

	for _, t := range l.tokens {
		if t.IsSection() {
			return true
		}
	}

	return false
}

// IsSectionMatching @TODO to refactor
func (l *Line) IsSectionMatching(s string) bool {
	if len(l.tokens) == 0 {
		return false
	}

	var found bool

	for _, t := range l.tokens {
		if t.IsSection() {
			found = true
		}
		if found && t.IsValue() && t.Value == s {
			return true
		}
	}

	return false
}
