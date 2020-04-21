package lexer

import (
	"errors"
	"unicode/utf8"

	"github.com/mickaelvieira/saxifrage/token"
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
type Lexer struct {
	input       string
	position    int // current position
	width       int // last rune's width
	expectValue bool
	tokens      chan *token.Token
}

// New create a new Lexer
func New(i string, c chan *token.Token) *Lexer {
	return &Lexer{
		input:  i,
		tokens: c,
	}
}

// next returns the next rune
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

// get the character without moving
func (l *Lexer) peek() rune {
	r := l.next()
	l.rewind()
	return r
}

// rewind move back to the previous position
func (l *Lexer) rewind() {
	l.position -= l.width
}

func (l *Lexer) lexWhitespaces() string {
	for isWhitespace(l.next()) {
	}
	l.rewind()
	return " "
}

// move to the end of line
func (l *Lexer) lexComments() string {
	var s string
	c := l.next()
	for !isEOL(c) {
		s += string(c)
		c = l.next()
	}
	l.rewind()
	return s
}

func (l *Lexer) lexWord() string {
	var s string
	c := l.next()
	for {
		if isWhitespace(c) || isEOL(c) || isEOF(c) {
			break
		}

		s += string(c)
		c = l.next()
	}
	l.rewind()
	l.expectValue = true

	return s
}

func (l *Lexer) lexValue() (string, error) {
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
		return "", ErrUnclosedQuote
	}

	l.expectValue = false

	return s, nil
}

func (l *Lexer) lexChar() string {
	c := l.next()
	return string(c)
}

// scan returns each token at each invocation
func (l *Lexer) scan(r rune) *token.Token {
	if isWhitespace(r) {
		return &token.Token{Type: token.Whitespace, Value: l.lexWhitespaces()}
	}
	if isHash(r) {
		return &token.Token{Type: token.Comment, Value: l.lexComments()}
	}
	if isEOL(r) {
		return &token.Token{Type: token.EOL, Value: l.lexChar()}
	}
	if isEOF(r) {
		return &token.Token{Type: token.EOF, Value: l.lexChar()}
	}

	if l.expectValue {
		v, err := l.lexValue()
		if err != nil {
			return &token.Token{Type: token.Err, Value: err.Error()}
		}
		return &token.Token{Type: token.Value, Value: v}
	}

	w := l.lexWord()
	t := token.Section
	if isKeyword(w) {
		t = token.Keyword
	}

	return &token.Token{Type: t, Value: w}
}

// Lex sends token over the tokens channel
func (l *Lexer) Lex() {
	for t := l.scan(l.peek()); !t.IsEOF(); {
		l.tokens <- t
		t = l.scan(l.peek())
	}
	close(l.tokens)
}
