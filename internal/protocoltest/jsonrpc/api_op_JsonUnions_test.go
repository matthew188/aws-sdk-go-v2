// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc

import (
	"bytes"
	"context"
	"github.com/matthew188/aws-sdk-go-v2/aws"
	awshttp "github.com/matthew188/aws-sdk-go-v2/aws/transport/http"
	"github.com/matthew188/aws-sdk-go-v2/internal/protocoltest/jsonrpc/types"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
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

func TestClient_JsonUnions_awsAwsjson11Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *JsonUnionsInput
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
		// Serializes a string union value
		"AwsJson11SerializeStringUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberStringValue{Value: "foo"},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "stringValue": "foo"
			    }
			}`))
			},
		},
		// Serializes a boolean union value
		"AwsJson11SerializeBooleanUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberBooleanValue{Value: true},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "booleanValue": true
			    }
			}`))
			},
		},
		// Serializes a number union value
		"AwsJson11SerializeNumberUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberNumberValue{Value: 1},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "numberValue": 1
			    }
			}`))
			},
		},
		// Serializes a blob union value
		"AwsJson11SerializeBlobUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberBlobValue{Value: []byte("foo")},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "blobValue": "Zm9v"
			    }
			}`))
			},
		},
		// Serializes a timestamp union value
		"AwsJson11SerializeTimestampUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberTimestampValue{Value: smithytime.ParseEpochSeconds(1398796238)},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "timestampValue": 1398796238
			    }
			}`))
			},
		},
		// Serializes an enum union value
		"AwsJson11SerializeEnumUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberEnumValue{Value: types.FooEnum("Foo")},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "enumValue": "Foo"
			    }
			}`))
			},
		},
		// Serializes a list union value
		"AwsJson11SerializeListUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberListValue{Value: []string{
					"foo",
					"bar",
				}},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "listValue": ["foo", "bar"]
			    }
			}`))
			},
		},
		// Serializes a map union value
		"AwsJson11SerializeMapUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberMapValue{Value: map[string]string{
					"foo":  "bar",
					"spam": "eggs",
				}},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "mapValue": {
			            "foo": "bar",
			            "spam": "eggs"
			        }
			    }
			}`))
			},
		},
		// Serializes a structure union value
		"AwsJson11SerializeStructureUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberStructureValue{Value: types.GreetingStruct{
					Hi: ptr.String("hello"),
				}},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
				"X-Amz-Target": []string{"JsonProtocol.JsonUnions"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "structureValue": {
			            "hi": "hello"
			        }
			    }
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
				HTTPClient: awshttp.NewBuildableClient(),
				Region:     "us-west-2",
			})
			result, err := client.JsonUnions(context.Background(), c.Params)
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

func TestClient_JsonUnions_awsAwsjson11Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *JsonUnionsOutput
	}{
		// Deserializes a string union value
		"AwsJson11DeserializeStringUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "stringValue": "foo"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberStringValue{Value: "foo"},
			},
		},
		// Deserializes a boolean union value
		"AwsJson11DeserializeBooleanUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "booleanValue": true
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberBooleanValue{Value: true},
			},
		},
		// Deserializes a number union value
		"AwsJson11DeserializeNumberUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "numberValue": 1
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberNumberValue{Value: 1},
			},
		},
		// Deserializes a blob union value
		"AwsJson11DeserializeBlobUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "blobValue": "Zm9v"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberBlobValue{Value: []byte("foo")},
			},
		},
		// Deserializes a timestamp union value
		"AwsJson11DeserializeTimestampUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "timestampValue": 1398796238
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberTimestampValue{Value: smithytime.ParseEpochSeconds(1398796238)},
			},
		},
		// Deserializes an enum union value
		"AwsJson11DeserializeEnumUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "enumValue": "Foo"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberEnumValue{Value: types.FooEnum("Foo")},
			},
		},
		// Deserializes a list union value
		"AwsJson11DeserializeListUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "listValue": ["foo", "bar"]
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberListValue{Value: []string{
					"foo",
					"bar",
				}},
			},
		},
		// Deserializes a map union value
		"AwsJson11DeserializeMapUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "mapValue": {
			            "foo": "bar",
			            "spam": "eggs"
			        }
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberMapValue{Value: map[string]string{
					"foo":  "bar",
					"spam": "eggs",
				}},
			},
		},
		// Deserializes a structure union value
		"AwsJson11DeserializeStructureUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/x-amz-json-1.1"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "structureValue": {
			            "hi": "hello"
			        }
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberStructureValue{Value: types.GreetingStruct{
					Hi: ptr.String("hello"),
				}},
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
				Region: "us-west-2",
			})
			var params JsonUnionsInput
			result, err := client.JsonUnions(context.Background(), &params)
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
