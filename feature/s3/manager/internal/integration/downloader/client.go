//go:build integration && perftest
// +build integration,perftest

package downloader

import (
	"net"
	"net/http"
	"time"

	"github.com/matthew188/aws-sdk-go-v2/aws"
	awshttp "github.com/matthew188/aws-sdk-go-v2/aws/transport/http"
)

func NewHTTPClient(cfg ClientConfig) aws.HTTPClient {
	return awshttp.NewBuildableClient().WithTransportOptions(func(tr *http.Transport) {
		*tr = http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   cfg.Timeouts.Connect,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        cfg.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
			IdleConnTimeout:     90 * time.Second,

			DisableKeepAlives:     !cfg.KeepAlive,
			TLSHandshakeTimeout:   cfg.Timeouts.TLSHandshake,
			ExpectContinueTimeout: cfg.Timeouts.ExpectContinue,
			ResponseHeaderTimeout: cfg.Timeouts.ResponseHeader,

			ReadBufferSize:  cfg.ReadBufferSize,
			WriteBufferSize: cfg.WriteBufferSize,
		}
	})
}
