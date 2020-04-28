package lexer

import (
	"fmt"
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

	l := tokenizer{input: "Hello, 世界"}

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

	l := tokenizer{}

	for i, tc := range cases {
		got := l.next()
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestRewind(t *testing.T) {
	l := tokenizer{input: "世界"}

	got := l.next()
	assert.Equal(t, '世', got, "Runes don't match")

	l.rewind()

	got = l.next()
	assert.Equal(t, '世', got, "Runes don't match")
}

func TestPeek(t *testing.T) {

	i := `世
界`
	l := tokenizer{input: i, line: 1, column: 1}

	cases := []struct {
		r        rune
		position int
		column   int
	}{
		{'世', 0, 1},
		{'世', 3, 4},
		{eol, 3, 4},
		{eol, 4, 5},
		{'界', 4, 5},
		{'界', 7, 8},
	}

	for i, tc := range cases {
		var r rune
		if i%2 == 0 {
			r = l.peek()
		} else {
			r = l.next()
		}

		assert.Equal(t, tc.r, r, "Test Case %d [rune] %v", i, tc)
		assert.Equal(t, tc.position, l.position, "Test Case %d [position] %v", i, tc)
		assert.Equal(t, tc.column, l.column, "Test Case %d [column] %v", i, tc)
	}
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
		l := tokenizer{input: tc.input}
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
		l := tokenizer{input: tc.input}
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
		{"foo", fmt.Sprintf(msgIllegalCharacter, "foo", 0, 0), Illegal},
	}

	for i, tc := range cases {
		l := tokenizer{input: tc.input}
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
		{`"bar foo, baz`, fmt.Sprintf(msgMissingClosingQuote, 0, 0), Err},
	}

	for i, tc := range cases {
		l := tokenizer{input: tc.input}
		got := l.lexValue()
		assert.Equal(t, tc.want1, got.Value, "Test Case %d %v", i, tc)
		assert.Equal(t, tc.want2, got.Type, "Test Case %d %v", i, tc)
	}
}

func TestLexing(t *testing.T) {
	cases := []struct {
		token *Token
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
		{&Token{Type: Illegal, Value: fmt.Sprintf(msgIllegalCharacter, string(eol), 8, 20)}},

		{&Token{Type: Illegal, Value: fmt.Sprintf(msgIllegalCharacter, string(eol), 9, 1)}},

		{&Token{Type: Illegal, Value: fmt.Sprintf(msgIllegalCharacter, "foobar", 10, 1)}},
		{&Token{Type: Illegal, Value: fmt.Sprintf(msgIllegalCharacter, string(eol), 10, 7)}},

		{&Token{Type: Illegal, Value: fmt.Sprintf(msgIllegalCharacter, string(eol), 11, 1)}},

		{&Token{Type: Keyword, Value: "VerifyHostKeyDNS"}},
		{&Token{Type: Separator, Value: " = "}},
		{&Token{Type: Value, Value: "baz"}},
		{&Token{Type: EOF, Value: string(eof)}},
	}

	l := New(`

	host * # here is the first comment
	VisualHostKey=foo # here is the second comment
	# This is the third comment
  hostname bar

ServerAliveInterval

foobar

VerifyHostKeyDNS = baz`)

	tokens := l.Run()

	var i int
	for got := range tokens {
		tc := cases[i]
		assert.Equal(t, tc.token.Type, got.Type, "Test Case %d [type] '%v'", i, tc.token)
		assert.Equal(t, tc.token.Value, got.Value, "Test Case %d [value] '%v'", i, tc.token)
		i++
	}
}
