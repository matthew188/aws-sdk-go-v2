// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/matthew188/aws-sdk-go-v2/aws"
	awshttp "github.com/matthew188/aws-sdk-go-v2/aws/transport/http"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestClient_JsonTimestamps_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *JsonTimestampsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		Host          *url.URL
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Tests how normal timestamps are serialized
		"RestJsonJsonTimestamps": {
			Params: &JsonTimestampsInput{
				Normal: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "normal": 1398796238
			}`))
			},
		},
		// Ensures that the timestampFormat of date-time works like normal timestamps
		"RestJsonJsonTimestampsWithDateTimeFormat": {
			Params: &JsonTimestampsInput{
				DateTime: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "dateTime": "2014-04-29T18:30:38Z"
			}`))
			},
		},
		// Ensures that the timestampFormat of date-time on the target shape works like
		// normal timestamps
		"RestJsonJsonTimestampsWithDateTimeOnTargetFormat": {
			Params: &JsonTimestampsInput{
				DateTimeOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "dateTimeOnTarget": "2014-04-29T18:30:38Z"
			}`))
			},
		},
		// Ensures that the timestampFormat of epoch-seconds works
		"RestJsonJsonTimestampsWithEpochSecondsFormat": {
			Params: &JsonTimestampsInput{
				EpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "epochSeconds": 1398796238
			}`))
			},
		},
		// Ensures that the timestampFormat of epoch-seconds on the target shape works
		"RestJsonJsonTimestampsWithEpochSecondsOnTargetFormat": {
			Params: &JsonTimestampsInput{
				EpochSecondsOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "epochSecondsOnTarget": 1398796238
			}`))
			},
		},
		// Ensures that the timestampFormat of http-date works
		"RestJsonJsonTimestampsWithHttpDateFormat": {
			Params: &JsonTimestampsInput{
				HttpDate: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "httpDate": "Tue, 29 Apr 2014 18:30:38 GMT"
			}`))
			},
		},
		// Ensures that the timestampFormat of http-date on the target shape works
		"RestJsonJsonTimestampsWithHttpDateOnTargetFormat": {
			Params: &JsonTimestampsInput{
				HttpDateOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonTimestamps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "httpDateOnTarget": "Tue, 29 Apr 2014 18:30:38 GMT"
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			serverURL := server.URL
			if c.Host != nil {
				u, err := url.Parse(serverURL)
				if err != nil {
					t.Fatalf("expect no error, got %v", err)
				}
				u.Path = c.Host.Path
				u.RawPath = c.Host.RawPath
				u.RawQuery = c.Host.RawQuery
				serverURL = u.String()
			}
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               awshttp.NewBuildableClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.JsonTimestamps(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_JsonTimestamps_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *JsonTimestampsOutput
	}{
		// Tests how normal timestamps are serialized
		"RestJsonJsonTimestamps": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "normal": 1398796238
			}`),
			ExpectResult: &JsonTimestampsOutput{
				Normal: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of date-time works like normal timestamps
		"RestJsonJsonTimestampsWithDateTimeFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "dateTime": "2014-04-29T18:30:38Z"
			}`),
			ExpectResult: &JsonTimestampsOutput{
				DateTime: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of date-time on the target shape works like
		// normal timestamps
		"RestJsonJsonTimestampsWithDateTimeOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "dateTimeOnTarget": "2014-04-29T18:30:38Z"
			}`),
			ExpectResult: &JsonTimestampsOutput{
				DateTimeOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of epoch-seconds works
		"RestJsonJsonTimestampsWithEpochSecondsFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "epochSeconds": 1398796238
			}`),
			ExpectResult: &JsonTimestampsOutput{
				EpochSeconds: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of epoch-seconds on the target shape works
		"RestJsonJsonTimestampsWithEpochSecondsOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "epochSecondsOnTarget": 1398796238
			}`),
			ExpectResult: &JsonTimestampsOutput{
				EpochSecondsOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of http-date works
		"RestJsonJsonTimestampsWithHttpDateFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "httpDate": "Tue, 29 Apr 2014 18:30:38 GMT"
			}`),
			ExpectResult: &JsonTimestampsOutput{
				HttpDate: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
		// Ensures that the timestampFormat of http-date on the target shape works
		"RestJsonJsonTimestampsWithHttpDateOnTargetFormat": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "httpDateOnTarget": "Tue, 29 Apr 2014 18:30:38 GMT"
			}`),
			ExpectResult: &JsonTimestampsOutput{
				HttpDateOnTarget: ptr.Time(smithytime.ParseEpochSeconds(1398796238)),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params JsonTimestampsInput
			result, err := client.JsonTimestamps(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
