//go:build windows
// +build windows

package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matthew188/aws-sdk-go-v2/internal/awstesting"
)

func TestSharedCredsFilename(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	os.Setenv("HOME", "home_dir")
	os.Setenv("USERPROFILE", "profile_dir")

	expect := filepath.Join("profile_dir", ".aws", "credentials")

	name := DefaultSharedCredentialsFilename()
	if e, a := expect, name; e != a {
		t.Errorf("expect %q shared creds filename, got %q", e, a)
	}
}

func TestSharedConfigFilename(t *testing.T) {
	restoreEnv := awstesting.StashEnv()
	defer awstesting.PopEnv(restoreEnv)

	os.Setenv("HOME", "home_dir")
	os.Setenv("USERPROFILE", "profile_dir")

	expect := filepath.Join("profile_dir", ".aws", "config")

	name := DefaultSharedConfigFilename()
	if e, a := expect, name; e != a {
		t.Errorf("expect %q shared config filename, got %q", e, a)
	}
}
