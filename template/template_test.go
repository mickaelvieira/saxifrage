package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawLine(t *testing.T) {

	cases := []struct {
		input int
		want  string
	}{
		{0, ""},
		{1, "-"},
		{10, "----------"},
	}

	for i, tc := range cases {
		got := Line(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}

}
