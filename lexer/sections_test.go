package lexer

import (
	"fmt"
	"testing"
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := isSection(tt.input)
			if got != tt.want {
				t.Errorf("Failed for %v ...", tt.input)
			}
		})
	}
}
