package awstesting

import (
	"context"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/matthew188/aws-sdk-go-v2/aws"
)

// ZeroReader is a io.Reader which will always write zeros to the byte slice provided.
type ZeroReader struct{}

// Read fills the provided byte slice with zeros returning the number of bytes written.
func (r *ZeroReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
	return len(b), nil
}

// ReadCloser is a io.ReadCloser for unit testing.
// Designed to test for leaks and whether a handle has
// been closed
type ReadCloser struct {
	Size     int
	Closed   bool
	set      bool
	FillData func(bool, []byte, int, int)
}

// Read will call FillData and fill it with whatever data needed.
// Decrements the size until zero, then return io.EOF.
func (r *ReadCloser) Read(b []byte) (int, error) {
	if r.Closed {
		return 0, io.EOF
	}

	delta := len(b)
	if delta > r.Size {
		delta = r.Size
	}
	r.Size -= delta

	for i := 0; i < delta; i++ {
		b[i] = 'a'
	}

	if r.FillData != nil {
		r.FillData(r.set, b, r.Size, delta)
	}
	r.set = true

	if r.Size > 0 {
		return delta, nil
	}
	return delta, io.EOF
}

// Close sets Closed to true and returns no error
func (r *ReadCloser) Close() error {
	r.Closed = true
	return nil
}

// A FakeContext provides a simple stub implementation of a Context
type FakeContext struct {
	Error  error
	DoneCh chan struct{}
}

// Deadline always will return not set
func (c *FakeContext) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Done returns a read channel for listening to the Done event
func (c *FakeContext) Done() <-chan struct{} {
	return c.DoneCh
}

// Err returns the error, is nil if not set.
func (c *FakeContext) Err() error {
	return c.Error
}

// Value ignores the Value and always returns nil
func (c *FakeContext) Value(key interface{}) interface{} {
	return nil
}

// StashEnv stashes the current environment variables except variables listed in envToKeepx
// Returns an function to pop out old environment
func StashEnv(envToKeep ...string) []string {
	if runtime.GOOS == "windows" {
		envToKeep = append(envToKeep, "ComSpec")
		envToKeep = append(envToKeep, "SYSTEM32")
		envToKeep = append(envToKeep, "SYSTEMROOT")
	}
	envToKeep = append(envToKeep, "PATH", "HOME", "USERPROFILE")
	extraEnv := getEnvs(envToKeep)
	originalEnv := os.Environ()
	os.Clearenv() // clear env
	for key, val := range extraEnv {
		os.Setenv(key, val)
	}
	return originalEnv
}

// PopEnv takes the list of the environment values and injects them into the
// process's environment variable data. Clears any existing environment values
// that may already exist.
func PopEnv(env []string) {
	os.Clearenv()

	for _, e := range env {
		p := strings.SplitN(e, "=", 2)
		k, v := p[0], ""
		if len(p) > 1 {
			v = p[1]
		}
		os.Setenv(k, v)
	}
}

// MockCredentialsProvider is a type that can be used to mock out credentials
// providers
type MockCredentialsProvider struct {
	RetrieveFn   func(ctx context.Context) (aws.Credentials, error)
	InvalidateFn func()
}

// Retrieve calls the RetrieveFn
func (p MockCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return p.RetrieveFn(ctx)
}

// Invalidate calls the InvalidateFn
func (p MockCredentialsProvider) Invalidate() {
	p.InvalidateFn()
}

func getEnvs(envs []string) map[string]string {
	extraEnvs := make(map[string]string)
	for _, env := range envs {
		if val, ok := os.LookupEnv(env); ok && len(val) > 0 {
			extraEnvs[env] = val
		}
	}
	return extraEnvs
}
