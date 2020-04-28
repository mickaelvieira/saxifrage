package lexer

import (
	"fmt"
	"unicode/utf8"
)

const (
	msgIllegalCharacter    = "Illegal character '%s' at line %d, column %d"
	msgMissingClosingQuote = "Double quote is not closed at line %d, column %d"
)

// create a EOF constant for clarity
const eof = rune(0)
const eol = rune(10)

// Tokenizer describes a Lexer struct
type Tokenizer interface {
	Run() chan *Token
}

type tokenizer struct {
	input    string // input string on which the lexical analysis
	position int    // current position
	width    int    // last rune's width
	line     int    // current line
	column   int    // current position in the line
}

// New create a new Lexer
func New(i string) Tokenizer {
	return &tokenizer{input: i, line: 1, column: 1}
}

func (l *tokenizer) next() rune {
	if l.position >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.position:]) // returns the first rune

	l.position += w // set the next position in the string
	l.column += w   // set the next position in the current line
	l.width = w     // keep the last rune's width in order to be able to backup

	return r
}

func (l *tokenizer) peek() rune {
	r := l.next()
	l.rewind()
	return r
}

func (l *tokenizer) rewind() {
	l.position -= l.width
	l.column -= l.width
}

func (l *tokenizer) newLine() {
	l.line++
	l.column = 1
}

func (l *tokenizer) lexWhitespaces() *Token {
	var s string

	for c := l.next(); isWhitespace(c); c = l.next() {
		s += string(c)
	}
	l.rewind()

	return &Token{Type: Whitespace, Value: s}
}

func (l *tokenizer) lexSeparator() *Token {
	var s string

	for c := l.next(); isSeparator(c); c = l.next() {
		s += string(c)
	}
	l.rewind()

	return &Token{Type: Separator, Value: s}
}

func (l *tokenizer) lexComments() *Token {
	var s string

	for c := l.next(); !isEOL(c); c = l.next() {
		s += string(c)
	}
	l.rewind()

	return &Token{Type: Comment, Value: s}
}

func (l *tokenizer) lexWord() *Token {
	var col = l.column
	var line = l.line

	var s string

	for c := l.next(); !isSeparator(c) && !isEOL(c) && !isEOF(c); c = l.next() {
		s += string(c)
	}
	l.rewind()

	is := isSection(s)
	ik := isKeyword(s)

	if is || ik {
		t := Section
		if ik {
			t = Keyword
		}
		return &Token{Type: t, Value: s}
	}
	return &Token{
		Type:  Illegal,
		Value: fmt.Sprintf(msgIllegalCharacter, s, line, col),
	}
}

func (l *tokenizer) lexValue() *Token {
	var col = l.column
	var line = l.line

	var inQuote = false
	var s string

	c := l.next()

	for {
		if isDoubleQuote(c) {
			inQuote = !inQuote
		}
		if isEOL(c) || isEOF(c) {
			break
		}
		if isHash(c) && !inQuote {
			break
		}
		if isWhitespace(c) {
			n := l.peek()
			if isHash(n) && !inQuote {
				break
			}
			if isEOL(n) || isEOF(n) {
				break
			}
		}

		s += string(c)
		c = l.next()
	}
	l.rewind()

	if inQuote {
		return &Token{
			Type:  Err,
			Value: fmt.Sprintf(msgMissingClosingQuote, line, col),
		}
	}

	return &Token{Type: Value, Value: s}
}

func (l *tokenizer) lexEOL() *Token {
	c := l.next()
	l.newLine()
	return &Token{Type: EOL, Value: string(c)}
}

func (l *tokenizer) lexEOF() *Token {
	c := l.next()
	return &Token{Type: EOF, Value: string(c)}
}

func (l *tokenizer) lexIllegal() *Token {
	var col = l.column
	var line = l.line

	c := l.next()
	v := fmt.Sprintf(msgIllegalCharacter, string(c), line, col)

	if isEOL(c) {
		l.newLine()
	}
	return &Token{Type: Illegal, Value: v}
}

// Run performs the lexical analysis of the input text
// Tokens will be send over the channel returned by this method
// until it reaches the end of the string
func (l *tokenizer) Run() chan *Token {
	var es bool // are we expecting a separator?
	var ev bool // are we expecting a value?

	tokenize := func(r rune) *Token {
		if es && isSeparator(r) {
			es = false
			ev = true // next token must be a value
			return l.lexSeparator()
		}

		if isWhitespace(r) || isHash(r) || isEOL(r) || isEOF(r) {
			if ev || es {
				return l.lexIllegal()
			}
			if isWhitespace(r) {
				return l.lexWhitespaces()
			}
			if isHash(r) {
				return l.lexComments()
			}
			if isEOL(r) {
				return l.lexEOL()
			}
			if isEOF(r) {
				return l.lexEOF()
			}
		}

		if ev {
			ev = false
			return l.lexValue()
		}

		t := l.lexWord()

		if t.IsKeyword() || t.IsSection() {
			es = true // next token must be a separator
		}
		return t
	}

	out := make(chan *Token)

	go func() {
		defer close(out)
		for t := tokenize(l.peek()); !t.IsEOF(); t = tokenize(l.peek()) {
			out <- t
		}
	}()

	return out
}
