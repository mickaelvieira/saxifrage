package lexer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSection(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"foo", false},
		{"Host", true},
		{"Match", true},
	}

	for i, tc := range cases {
		uc := tc.input
		lc := strings.ToLower(uc)

		got := isSection(uc)
		assert.Equal(t, tc.want, got, "Test Case [Uppercase] %d %v", i, tc)

		got = isSection(lc)
		assert.Equal(t, tc.want, got, "Test Case [Lowercase] %d %v", i, tc)
	}
}

func TestGetSection(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"Host"},
		{"Match"},
	}

	for i, tc := range cases {
		uc := tc.input
		got := getSection(uc)
		assert.Equal(t, uc, got, "Test Case [Uppercase] %d %v", i, tc)
	}
}
