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
