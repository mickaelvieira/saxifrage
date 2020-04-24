package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhitespaces(t *testing.T) {
	cases := []struct {
		input rune
		want  bool
	}{
		{'a', false},
		{'\n', false},
		{' ', true},
		{'\t', true},
	}

	for i, tc := range cases {
		got := isWhitespace(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestHash(t *testing.T) {
	cases := []struct {
		input rune
		want  bool
	}{
		{'a', false},
		{'#', true},
	}

	for i, tc := range cases {
		got := isHash(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestDoubleQuote(t *testing.T) {
	cases := []struct {
		input rune
		want  bool
	}{
		{'\'', false},
		{'`', false},
		{'"', true},
	}

	for i, tc := range cases {
		got := isDoubleQuote(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestEOF(t *testing.T) {
	cases := []struct {
		input rune
		want  bool
	}{
		{eof, true},
		{'\n', false},
	}

	for i, tc := range cases {
		got := isEOF(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestEOL(t *testing.T) {
	cases := []struct {
		input rune
		want  bool
	}{
		{eof, false},
		{'\n', true},
	}

	for i, tc := range cases {
		got := isEOL(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}
