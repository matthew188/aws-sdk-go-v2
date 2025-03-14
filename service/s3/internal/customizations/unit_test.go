package customizations_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/matthew188/aws-sdk-go-v2/aws"
	"github.com/matthew188/aws-sdk-go-v2/service/s3"

	"github.com/aws/smithy-go"
)

func Test_EmptyResponse(t *testing.T) {
	cases := map[string]struct {
		status       int
		responseBody []byte
		expectError  bool
	}{
		"success case with no response body": {
			status:       200,
			responseBody: []byte(``),
		},
		"error case with no response body": {
			status:       400,
			responseBody: []byte(``),
			expectError:  true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(c.status)
					w.Write(c.responseBody)
				}))
			defer server.Close()

			ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelFn()

			cfg := aws.Config{
				Region: "us-east-1",
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:         server.URL,
						SigningName: "s3",
					}, nil
				}),
				Retryer: func() aws.Retryer {
					return aws.NopRetryer{}
				},
			}

			client := s3.NewFromConfig(cfg, func(options *s3.Options) {
				options.UsePathStyle = true
			})

			params := &s3.HeadBucketInput{Bucket: aws.String("aws-sdk-go-data")}
			_, err := client.HeadBucket(ctx, params)
			if c.expectError {
				var apiErr smithy.APIError
				if !errors.As(err, &apiErr) {
					t.Fatalf("expect error to be API error, was not, %v", err)
				}
				if len(apiErr.ErrorCode()) == 0 {
					t.Errorf("expect non-empty error code")
				}
				if len(apiErr.ErrorMessage()) == 0 {
					t.Errorf("expect non-empty error message")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err.Error())
				}
			}
		})
	}
}

func TestBucketLocationPopulation(t *testing.T) {
	cases := map[string]struct {
		responseBody   string
		expectLocation string
		expectError    string
	}{
		"empty location": {
			responseBody:   `<?xml version="1.0" encoding="UTF-8"?><LocationConstraint xmlns="http://s3.amazonaws.com/doc/2006-03-01/"/>`,
			expectLocation: "",
		},
		"EU location": {
			responseBody:   `<?xml version="1.0" encoding="UTF-8"?><LocationConstraint xmlns="http://s3.amazonaws.com/doc/2006-03-01/">EU</LocationConstraint>`,
			expectLocation: "EU",
		},
		"AfSouth1 location": {
			responseBody:   `<?xml version="1.0" encoding="UTF-8"?><LocationConstraint xmlns="http://s3.amazonaws.com/doc/2006-03-01/">af-south-1</LocationConstraint>`,
			expectLocation: "af-south-1",
		},
		"IncompleteResponse": {
			responseBody: `<?xml version="1.0" encoding="UTF-8"?><LocationConstraint xmlns="http://s3.amazonaws.com/doc/2006-03-01/">`,
			expectError:  "unexpected EOF",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
					w.Write([]byte(c.responseBody))
				}))
			defer server.Close()

			ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelFn()

			cfg := aws.Config{
				Region: "us-east-1",
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:         server.URL,
						SigningName: "s3",
					}, nil
				}),
				Retryer: func() aws.Retryer { return aws.NopRetryer{} },
			}

			client := s3.NewFromConfig(cfg, func(options *s3.Options) {
				options.UsePathStyle = true
			})

			params := &s3.GetBucketLocationInput{
				Bucket: aws.String("aws-sdk-go-data"),
			}
			resp, err := client.GetBucketLocation(ctx, params)
			if len(c.expectError) != 0 && err == nil {
				t.Fatal("expect error, got none")
			}

			if err != nil && len(c.expectError) == 0 {
				t.Fatalf("expect no error, got %v", err)
			} else {
				if err != nil {
					if !strings.Contains(err.Error(), c.expectError) {
						t.Fatalf("expect error to be %v, got %v", err.Error(), c.expectError)
					}
					return
				}
			}

			if e, a := c.expectLocation, resp.LocationConstraint; !strings.EqualFold(e, string(a)) {
				t.Fatalf("expected location constraint to be deserialized as %v, got %v", e, a)
			}
		})
	}

}
