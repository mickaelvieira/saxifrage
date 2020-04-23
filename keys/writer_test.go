package keys

import (
	"fmt"
	"os"
	"testing"
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := GetDir(tt.input)
			if got != tt.want {
				t.Errorf("want '%s', got '%s'", tt.want, got)
			}
		})
	}
}

func TestDefaultFilenames(t *testing.T) {
	cases := []struct {
		input1 Type
		want1  string
		want2  string
	}{
		{RSA, "id_rsa", "id_rsa.pub"},
		{DSA, "id_dsa", "id_dsa.pub"},
		{ECDSA, "id_ecdsa", "id_ecdsa.pub"},
		{ED25519, "id_ed25519", "id_ed25519.pub"},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got1, got2 := GetFilenamesFromType(tt.input1)
			if got1 != tt.want1 {
				t.Errorf("want '%s', got '%s'", tt.want1, got1)
			}
			if got2 != tt.want2 {
				t.Errorf("want '%s', got '%s'", tt.want2, got2)
			}
		})
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

	for i, tt := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got1, got2 := GetFilenamesFromString(tt.input1)
			if got1 != tt.want1 {
				t.Errorf("want '%s', got '%s'", tt.want1, got1)
			}
			if got2 != tt.want2 {
				t.Errorf("want '%s', got '%s'", tt.want2, got2)
			}
		})
	}
}
