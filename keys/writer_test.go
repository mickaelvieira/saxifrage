package keys

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var home = os.Getenv("HOME")

func TestDestinationDir(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"", home + "/.ssh"},
		{"foo", home + "/.ssh/foo"},
	}

	for i, tc := range cases {
		got := GetDir(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestDefaultFilenames(t *testing.T) {
	cases := []struct {
		input1 Type
		want1  string
		want2  string
	}{
		{RSA, "id_rsa", "id_rsa.pub"},
		{ECDSA, "id_ecdsa", "id_ecdsa.pub"},
		{ED25519, "id_ed25519", "id_ed25519.pub"},
	}

	for i, tc := range cases {
		got1, got2 := GetFilenamesFromType(tc.input1)
		assert.Equal(t, tc.want1, got1, "Test case %d %v", i, tc)
		assert.Equal(t, tc.want2, got2, "Test case %d %v", i, tc)
	}
}

func TestUserFilenames(t *testing.T) {
	cases := []struct {
		input1 string
		want1  string
		want2  string
	}{
		{"baz", "baz", "baz.pub"},
	}

	for i, tc := range cases {
		got1, got2 := GetFilenamesFromString(tc.input1)
		assert.Equal(t, tc.want1, got1, "Test case %d %v", i, tc)
		assert.Equal(t, tc.want2, got2, "Test case %d %v", i, tc)
	}
}
