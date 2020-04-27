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
