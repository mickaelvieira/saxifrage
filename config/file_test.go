package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUseConfig(t *testing.T) {
	f1 := &File{}
	assert.False(t, f1.IsUserConfig())
	f2 := &File{Path: os.Getenv("HOME") + "/.ssh/config"}
	assert.True(t, f2.IsUserConfig())
}

func TestGetUserConfig(t *testing.T) {
	f1 := &File{}
	f2 := &File{Path: os.Getenv("HOME") + "/.ssh/config"}

	f := make(Files, 2)
	f[0] = f1
	f[1] = f2

	got := f.GetUserConfig()

	assert.Equal(t, f2, got)
}
