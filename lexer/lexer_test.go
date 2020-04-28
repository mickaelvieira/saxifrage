package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	for i, tc := range cases {
		got := l.next()
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestNextEmptyString(t *testing.T) {
	cases := []struct {
		want rune
	}{
		{eof},
	}

	l := Lexer{}

	for i, tc := range cases {
		got := l.next()
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestRewind(t *testing.T) {
	l := Lexer{input: "世界"}

	got := l.next()
	assert.Equal(t, '世', got, "Runes don't match")

	l.rewind()

	got = l.next()
	assert.Equal(t, '世', got, "Runes don't match")
}

func TestPeek(t *testing.T) {

	l := Lexer{input: "世界"}

	got := l.next()
	assert.Equal(t, '世', got, "Runes don't match")

	got = l.peek()
	assert.Equal(t, '界', got, "Runes don't match")

	got = l.next()
	assert.Equal(t, '界', got, "Runes don't match")
}

func TestWhitespaces(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"   a", "   "},
		{"	a", "	"},
		{"  	a", "  	"},
		{"	  a", "	  "},
	}

	for i, tc := range cases {
		l := Lexer{input: tc.input}
		got := l.lexWhitespaces()
		assert.Equal(t, Whitespace, got.Type, "Test Case %d %v", i, tc)
		assert.Equal(t, tc.want, got.Value, "Test Case %d %v", i, tc)
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

	for i, tc := range cases {
		l := Lexer{input: tc.input}
		got := l.lexComments()
		assert.Equal(t, Comment, got.Type, "Test Case %d %v", i, tc)
		assert.Equal(t, tc.want, got.Value, "Test Case %d %v", i, tc)
	}
}

func TestWord(t *testing.T) {
	cases := []struct {
		input string
		want1 string
		want2 Type
	}{
		{"Host", "Host", Section},
		{"User", "User", Keyword},
		{"foo", "foo", Illegal},
	}

	for i, tc := range cases {
		l := Lexer{input: tc.input}
		got := l.lexWord()
		assert.Equal(t, tc.want1, got.Value, "Test Case %d %v", i, tc)
		assert.Equal(t, tc.want2, got.Type, "Test Case %d %v", i, tc)
	}
}

func TestValues(t *testing.T) {
	cases := []struct {
		input string
		want1 string
		want2 Type
	}{
		{`foo
`, "foo", Value},
		{`foo # comment`, "foo", Value},
		{"\"foo # comment\"", "\"foo # comment\"", Value},
		{`bar foo
`, "bar foo", Value},
		{`"foo bar" foo, baz
`, "\"foo bar\" foo, baz", Value},
		{`"bar foo, baz`, ErrUnclosedQuote.Error(), Err},
	}

	for i, tc := range cases {
		l := Lexer{input: tc.input}
		got := l.lexValue()
		assert.Equal(t, tc.want1, got.Value, "Test Case %d %v", i, tc)
		assert.Equal(t, tc.want2, got.Type, "Test Case %d %v", i, tc)
	}
}

func TestLexing(t *testing.T) {
	cases := []struct {
		want *Token
	}{
		{&Token{Type: EOL, Value: string(eol)}},
		{&Token{Type: EOL, Value: string(eol)}},
		{&Token{Type: Whitespace, Value: "	"}},
		{&Token{Type: Section, Value: "host"}},
		{&Token{Type: Separator, Value: " "}},
		{&Token{Type: Value, Value: "*"}},
		{&Token{Type: Whitespace, Value: " "}},
		{&Token{Type: Comment, Value: "# here is the first comment"}},
		{&Token{Type: EOL, Value: string(eol)}},

		{&Token{Type: Whitespace, Value: "	"}},
		{&Token{Type: Keyword, Value: "VisualHostKey"}},
		{&Token{Type: Separator, Value: "="}},
		{&Token{Type: Value, Value: "foo"}},
		{&Token{Type: Whitespace, Value: " "}},
		{&Token{Type: Comment, Value: "# here is the second comment"}},
		{&Token{Type: EOL, Value: string(eol)}},

		{&Token{Type: Whitespace, Value: "	"}},
		{&Token{Type: Comment, Value: "# This is the third comment"}},
		{&Token{Type: EOL, Value: string(eol)}},

		{&Token{Type: Whitespace, Value: "  "}},
		{&Token{Type: Keyword, Value: "hostname"}},
		{&Token{Type: Separator, Value: " "}},
		{&Token{Type: Value, Value: "bar"}},
		{&Token{Type: EOL, Value: string(eol)}},

		{&Token{Type: EOL, Value: string(eol)}},

		{&Token{Type: Keyword, Value: "ServerAliveInterval"}},
		{&Token{Type: Illegal, Value: string(eol)}},

		{&Token{Type: Illegal, Value: string(eol)}},

		{&Token{Type: Illegal, Value: "foobar"}},
		{&Token{Type: Illegal, Value: string(eol)}},

		{&Token{Type: Illegal, Value: string(eol)}},

		{&Token{Type: Keyword, Value: "VerifyHostKeyDNS"}},
		{&Token{Type: Separator, Value: " = "}},
		{&Token{Type: Value, Value: "baz"}},
		{&Token{Type: EOF, Value: string(eof)}},
	}

	l := Lexer{
		input: `

	host * # here is the first comment
	VisualHostKey=foo # here is the second comment
	# This is the third comment
  hostname bar

ServerAliveInterval

foobar

VerifyHostKeyDNS = baz`}

	tokens := l.Lex()

	var i int
	for got := range tokens {
		tc := cases[i]
		assert.Equal(t, tc.want.Type, got.Type, "Test Case %d '%v'", i, tc.want.Value)
		assert.Equal(t, tc.want.Value, got.Value, "Test Case %d %v", i, tc.want.Value)
		i++
	}
}
