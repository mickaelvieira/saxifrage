package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserConfigPath(t *testing.T) {
	want := os.Getenv("HOME") + "/.ssh/config"
	got := GetUserConfigPath()
	assert.Equal(t, want, got, "Paths don't match")
}

func TestGlobalConfigPath(t *testing.T) {
	want := "/etc/ssh/ssh_config"
	got := GetGlobalConfigPath()
	assert.Equal(t, want, got, "Paths don't match")
}

func TestIsBaseSSHDirectory(t *testing.T) {
	home := os.Getenv("HOME")
	cases := []struct {
		input string
		want  bool
	}{
		{"~/.ssh", true},
		{home + "/.ssh", true},
		{"~/foo/bar", false},
	}

	for i, tc := range cases {
		got := IsBaseSSHDirectory(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestConvertToRelativePath(t *testing.T) {
	home := os.Getenv("HOME")
	cases := []struct {
		input string
		want  string
	}{
		{home + "/id_rsa", "~/id_rsa"},
		{home + "/foo/ed25519", "~/foo/ed25519"},
		{home + "/foo/bar", "~/foo/bar"},
	}

	for i, tc := range cases {
		got := ToRelativePath(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}

func TestConvertToAbsolute(t *testing.T) {
	home := os.Getenv("HOME")
	cases := []struct {
		input string
		want  string
	}{
		{"~/id_rsa", home + "/id_rsa"},
		{"~/foo/ed25519", home + "/foo/ed25519"},
		{"~/foo/bar", home + "/foo/bar"},
	}

	for i, tc := range cases {
		got := ToAbsolutePath(tc.input)
		assert.Equal(t, tc.want, got, "Test Case %d %v", i, tc)
	}
}
