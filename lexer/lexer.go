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
	input    string      // input string on which the lexical analysis
	position int         // current position
	width    int         // last rune's width
	tokens   chan *Token // channel to send tokens over
}

// New create a new Lexer
func New(i string, c chan *Token) *Lexer {
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

func (l *Lexer) lexSeparator() (s string) {
	for c := l.next(); isSeparator(c); c = l.next() {
		s += string(c)
	}
	l.rewind()
	return s
}

// move to the end of line
func (l *Lexer) lexComments() (s string) {
	for c := l.next(); !isEOL(c); c = l.next() {
		s += string(c)
	}
	l.rewind()
	return s
}

func (l *Lexer) lexWord() (s string) {
	for c := l.next(); !isSeparator(c) && !isEOL(c) && !isEOF(c); c = l.next() {
		s += string(c)
	}
	l.rewind()
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

	return s, nil
}

func (l *Lexer) lexChar() string {
	c := l.next()
	return string(c)
}

// Lex sends token over the tokens channel
func (l *Lexer) Lex() {
	var es bool // are we expecting a separator?
	var ev bool // are we expecting a value?

	tokenize := func(r rune) *Token {
		if es && isSeparator(r) {
			es = false
			ev = true // next token must be a value
			return &Token{Type: Separator, Value: l.lexSeparator()}
		}

		if ev {
			v, err := l.lexValue()
			ev = false

			if err != nil {
				return &Token{Type: Err, Value: err.Error()}
			}
			return &Token{Type: Value, Value: v}
		}

		if isWhitespace(r) || isHash(r) || isEOL(r) || isEOF(r) {
			if ev || es {
				return &Token{Type: Illegal, Value: l.lexChar()}
			}
			if isWhitespace(r) {
				return &Token{Type: Whitespace, Value: l.lexWhitespaces()}
			}
			if isHash(r) {
				return &Token{Type: Comment, Value: l.lexComments()}
			}
			if isEOL(r) {
				return &Token{Type: EOL, Value: l.lexChar()}
			}
			if isEOF(r) {
				return &Token{Type: EOF, Value: l.lexChar()}
			}
		}

		w := l.lexWord()

		is := isSection(w)
		ik := isKeyword(w)

		if is || ik {
			t := Section
			if ik {
				t = Keyword
			}
			es = true // next token must be a separator
			return &Token{Type: t, Value: w}
		}
		return &Token{Type: Illegal, Value: w}
	}

	for t := tokenize(l.peek()); !t.IsEOF(); t = tokenize(l.peek()) {
		l.tokens <- t
	}
	close(l.tokens)
}
