//go:build integration && perftest
// +build integration,perftest

package uploader

import (
	"flag"
	"net/http"
	"strings"
	"time"

	"github.com/matthew188/aws-sdk-go-v2/feature/s3/manager"
)

type SDKConfig struct {
	PartSize       int64
	Concurrency    int
	BufferProvider manager.ReadSeekerWriteToProvider
}

func (c *SDKConfig) SetupFlags(prefix string, flagset *flag.FlagSet) {
	prefix += "sdk."

	flagset.Int64Var(&c.PartSize, prefix+"part-size", manager.DefaultUploadPartSize,
		"Specifies the `size` of parts of the object to upload.")
	flagset.IntVar(&c.Concurrency, prefix+"concurrency", manager.DefaultUploadConcurrency,
		"Specifies the number of parts to upload `at once`.")
}

func (c *SDKConfig) Validate() error {
	return nil
}

type ClientConfig struct {
	KeepAlive bool
	Timeouts  Timeouts

	MaxIdleConns        int
	MaxIdleConnsPerHost int
}

func (c *ClientConfig) SetupFlags(prefix string, flagset *flag.FlagSet) {
	prefix += "client."

	flagset.BoolVar(&c.KeepAlive, prefix+"http-keep-alive", true,
		"Specifies if HTTP keep alive is enabled.")

	defTR := http.DefaultTransport.(*http.Transport)

	flagset.IntVar(&c.MaxIdleConns, prefix+"idle-conns", defTR.MaxIdleConns,
		"Specifies max idle connection pool size.")

	flagset.IntVar(&c.MaxIdleConnsPerHost, prefix+"idle-conns-host", http.DefaultMaxIdleConnsPerHost,
		"Specifies max idle connection pool per host, will be truncated by idle-conns.")

	c.Timeouts.SetupFlags(prefix, flagset)
}

func (c *ClientConfig) Validate() error {
	var errs Errors

	if err := c.Timeouts.Validate(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		return errs
	}
	return nil
}

type Timeouts struct {
	Connect        time.Duration
	TLSHandshake   time.Duration
	ExpectContinue time.Duration
	ResponseHeader time.Duration
}

func (c *Timeouts) SetupFlags(prefix string, flagset *flag.FlagSet) {
	prefix += "timeout."

	flagset.DurationVar(&c.Connect, prefix+"connect", 30*time.Second,
		"The `timeout` connecting to the remote host.")

	defTR := http.DefaultTransport.(*http.Transport)

	flagset.DurationVar(&c.TLSHandshake, prefix+"tls", defTR.TLSHandshakeTimeout,
		"The `timeout` waiting for the TLS handshake to complete.")

	flagset.DurationVar(&c.ExpectContinue, prefix+"expect-continue", defTR.ExpectContinueTimeout,
		"The `timeout` waiting for the TLS handshake to complete.")

	flagset.DurationVar(&c.ResponseHeader, prefix+"response-header", defTR.ResponseHeaderTimeout,
		"The `timeout` waiting for the TLS handshake to complete.")
}

func (c *Timeouts) Validate() error {
	return nil
}

type Errors []error

func (es Errors) Error() string {
	var buf strings.Builder
	for _, e := range es {
		buf.WriteString(e.Error())
	}

	return buf.String()
}
