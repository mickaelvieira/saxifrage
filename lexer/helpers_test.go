package lexer

import (
	"fmt"
	"testing"
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isWhitespace(tt.input)
			if got != tt.want {
				t.Errorf("Failed for %v ...", tt.input)
			}
		})
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isHash(tt.input)
			if got != tt.want {
				t.Errorf("Failed for %v ...", tt.input)
			}
		})
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isDoubleQuote(tt.input)
			if got != tt.want {
				t.Errorf("Failed for %v ...", tt.input)
			}
		})
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isEOF(tt.input)
			if got != tt.want {
				t.Errorf("Failed for %v ...", tt.input)
			}
		})
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isEOL(tt.input)
			if got != tt.want {
				t.Errorf("Failed for %v ...", tt.input)
			}
		})
	}
}
