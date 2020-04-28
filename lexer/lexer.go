package lexer

import (
	"errors"
	"unicode/utf8"
)

// Lexing errors
var (
	ErrUnclosedQuote = errors.New("Double quote is not closed")
	ErrExpectedValue = errors.New("A value was expected")
)

// create a EOF constant for clarity
const eof = rune(0)
const eol = rune(10)

// Lexer - lexical analysis
// @TODO adds the current line number, to help debug parsing error?
type Lexer struct {
	input    string // input string on which the lexical analysis
	position int    // current position
	width    int    // last rune's width
}

// New create a new Lexer
func New(i string) *Lexer {
	return &Lexer{input: i}
}

func (l *Lexer) next() rune {
	if l.position >= len(l.input) {
		l.width = 0
		return eof
	}

	r, w := utf8.DecodeRuneInString(l.input[l.position:]) // returns the first rune

	l.position += w // set the next position
	l.width = w     // keep the last rune's width in order to be able to backup

	return r
}

func (l *Lexer) peek() rune {
	r := l.next()
	l.rewind()
	return r
}

func (l *Lexer) rewind() {
	l.position -= l.width
}

func (l *Lexer) lexWhitespaces() *Token {
	var s string

	for c := l.next(); isWhitespace(c); c = l.next() {
		s += string(c)
	}
	l.rewind()

	return &Token{Type: Whitespace, Value: s}
}

func (l *Lexer) lexSeparator() *Token {
	var s string

	for c := l.next(); isSeparator(c); c = l.next() {
		s += string(c)
	}
	l.rewind()

	return &Token{Type: Separator, Value: s}
}

func (l *Lexer) lexComments() *Token {
	var s string

	for c := l.next(); !isEOL(c); c = l.next() {
		s += string(c)
	}
	l.rewind()

	return &Token{Type: Comment, Value: s}
}

func (l *Lexer) lexWord() *Token {
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

	return &Token{Type: Illegal, Value: s}
}

func (l *Lexer) lexValue() *Token {
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
		return &Token{Type: Err, Value: ErrUnclosedQuote.Error()}
	}

	return &Token{Type: Value, Value: s}
}

func (l *Lexer) lexEOL() *Token {
	c := l.next()
	return &Token{Type: EOL, Value: string(c)}
}

func (l *Lexer) lexEOF() *Token {
	c := l.next()
	return &Token{Type: EOF, Value: string(c)}
}

// Lex performs the lexical analysis of the input text
// Tokens will be send over the channel returned by this method
// until it reaches the end of the string
func (l *Lexer) Lex() chan *Token {
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
				return &Token{Type: Illegal, Value: string(l.next())}
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
