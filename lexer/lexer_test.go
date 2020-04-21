package lexer

import (
	"fmt"
	"testing"

	"github.com/mickaelvieira/saxifrage/token"
)

func TestNext(t *testing.T) {
	cases := []struct {
		want rune
	}{
		{'H'},
		{'e'},
		{'l'},
		{'l'},
		{'o'},
		{','},
		{' '},
		{'世'},
		{'界'},
		{eof},
	}

	l := Lexer{input: "Hello, 世界"}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := l.next()
			if got != tt.want {
				t.Errorf("Failed for %v", tt.want)
			}
		})
	}
}

func TestNextEmptyString(t *testing.T) {
	cases := []struct {
		want rune
	}{
		{eof},
	}

	l := Lexer{}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := l.next()
			if got != tt.want {
				t.Errorf("Failed for %v", tt.want)
			}
		})
	}
}

func TestRewind(t *testing.T) {
	l := Lexer{input: "世界"}

	got := l.next()
	if got != '世' {
		t.Errorf("Want %v got %v", '世', got)
	}

	l.rewind()

	got = l.next()
	if got != '世' {
		t.Errorf("Want %v got %v", '世', got)
	}
}

func TestPeek(t *testing.T) {
	l := Lexer{input: "世界"}

	got := l.next()
	if got != '世' {
		t.Errorf("Want %v got %v", '世', got)
	}

	got = l.peek()
	if got != '界' {
		t.Errorf("Want %v got %v", '界', got)
	}

	got = l.next()
	if got != '界' {
		t.Errorf("Want %v got %v", '界', got)
	}
}

func TestIgnoreWhitespaces(t *testing.T) {
	cases := []struct {
		input string
		want  rune
	}{
		{"   a", 'a'},
		{"	a", 'a'},
		{"  	a", 'a'},
		{"	  a", 'a'},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			l := Lexer{input: tt.input}
			ws := l.lexWhitespaces()
			got := l.next()
			if ws != " " {
				t.Errorf("Want whitespace, got %v", ws)
			}
			if got != tt.want {
				t.Errorf("Want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestComments(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{`# foo
`, "# foo"},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			l := Lexer{input: tt.input}
			got := l.lexComments()
			e := l.next()
			if got != tt.want {
				t.Errorf("Want %s, got %s", string(tt.want), string(got))
			}
			if e != eol {
				t.Errorf("Want eol, got %v", e)
			}
		})
	}
}

func TestWord(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{`foo`, "foo"},
		{"bar foo", "bar"},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			l := Lexer{input: tt.input}
			got := l.lexWord()
			if got != tt.want {
				t.Errorf("Want %s, got %s", string(tt.want), string(got))
			}
		})
	}
}

func TestValues(t *testing.T) {
	cases := []struct {
		input string
		want  string
		err   error
	}{
		{`foo
`, "foo", nil},
		{`foo # comment`, "foo", nil},
		{"\"foo # comment\"", "\"foo # comment\"", nil},
		{`bar foo
`, "bar foo", nil},
		{`"foo bar" foo, baz
`, "\"foo bar\" foo, baz", nil},
		{`"bar foo, baz`, "", ErrUnclosedQuote},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			l := Lexer{input: tt.input}
			got, err := l.lexValue()
			if got != tt.want {
				t.Errorf("Want '%s', got '%s'", string(tt.want), string(got))
			}
			if err != tt.err {
				t.Errorf("Want %s, got %s", tt.err, err)
			}
		})
	}
}

func TestScan(t *testing.T) {
	cases := []struct {
		want *token.Token
	}{
		{&token.Token{Type: token.EOL, Value: string(eol)}},
		{&token.Token{Type: token.EOL, Value: string(eol)}},
		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Section, Value: "Host"}},
		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Value, Value: "*"}},
		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Comment, Value: "# here is the first comment"}},
		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Keyword, Value: "VisualHostKey"}},
		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Value, Value: "foo"}},
		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Comment, Value: "# here is the second comment"}},
		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Comment, Value: "# This is the third comment"}},
		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Keyword, Value: "HostName"}},
		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Value, Value: "bar"}},
		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.Keyword, Value: "ServerAliveInterval"}},
		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.Value, Value: "foobar"}},
		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.EOL, Value: string(eol)}},

		{&token.Token{Type: token.Keyword, Value: "VerifyHostKeyDNS"}},
		{&token.Token{Type: token.Whitespace, Value: " "}},
		{&token.Token{Type: token.Value, Value: "baz"}},
		{&token.Token{Type: token.EOF, Value: string(eof)}},
	}

	tokens := make(chan *token.Token)
	l := Lexer{
		tokens: tokens,
		input: `

	Host * # here is the first comment
	VisualHostKey foo # here is the second comment
	# This is the third comment
	HostName bar

ServerAliveInterval

foobar

VerifyHostKeyDNS baz`}

	go l.Lex()

	var i int
	for got := range tokens {
		tt := cases[i]
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got.Type != tt.want.Type {
				t.Errorf("Want %v got %v", tt.want.Type, got.Type)
			}
			if got.Value != tt.want.Value {
				t.Errorf("Want %v got %v", tt.want.Value, got.Value)
			}
		})
		i++
	}
}
