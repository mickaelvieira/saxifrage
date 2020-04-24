package lexer

import (
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
		got := isSection(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}
