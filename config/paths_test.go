package config

import (
	"os"
	"testing"
)

func TestUserConfigPath(t *testing.T) {
	want := os.Getenv("HOME") + "/.ssh/config"
	got := GetUserConfigPath()

	if got != want {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestGlobalConfigPath(t *testing.T) {
	want := "/etc/ssh/ssh_config"
	got := GetGlobalConfigPath()

	if got != want {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}
