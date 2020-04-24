package keys

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultType(t *testing.T) {
	want := RSA
	got := GetDefaultType()
	assert.Equal(t, want, got, "Default types don't match")
}

func TestTypesToString(t *testing.T) {
	want := "rsa, dsa, ecdsa, ed25519"
	got := TypesToString()
	assert.Equal(t, want, got, "Strings don't match")
}

func TestGetKeyType(t *testing.T) {
	cases := []struct {
		input string
		want  Type
	}{
		{"rsa", RSA},
		{"dsa", DSA},
		{"ecdsa", ECDSA},
		{"ed25519", ED25519},
		{"", INVALID},
		{"foo", INVALID},
	}

	for i, tc := range cases {
		got := GetKeyType(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}
