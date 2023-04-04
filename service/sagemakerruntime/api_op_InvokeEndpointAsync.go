// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerruntime

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// After you deploy a model into production using Amazon SageMaker hosting
// services, your client applications use this API to get inferences from the model
// hosted at the specified endpoint in an asynchronous manner. Inference requests
// sent to this API are enqueued for asynchronous processing. The processing of the
// inference request may or may not complete before you receive a response from
// this API. The response from this API will not contain the result of the
// inference request but contain information about where you can locate it. Amazon
// SageMaker strips all POST headers except those supported by the API. Amazon
// SageMaker might add additional headers. You should not rely on the behavior of
// headers outside those enumerated in the request syntax. Calls to
// InvokeEndpointAsync are authenticated by using Amazon Web Services Signature
// Version 4. For information, see Authenticating Requests (Amazon Web Services
// Signature Version 4)
// (https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html)
// in the Amazon S3 API Reference.
func (c *Client) InvokeEndpointAsync(ctx context.Context, params *InvokeEndpointAsyncInput, optFns ...func(*Options)) (*InvokeEndpointAsyncOutput, error) {
	if params == nil {
		params = &InvokeEndpointAsyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeEndpointAsync", params, optFns, c.addOperationInvokeEndpointAsyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeEndpointAsyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeEndpointAsyncInput struct {

	// The name of the endpoint that you specified when you created the endpoint using
	// the CreateEndpoint
	// (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html)
	// API.
	//
	// This member is required.
	EndpointName *string

	// The Amazon S3 URI where the inference request payload is stored.
	//
	// This member is required.
	InputLocation *string

	// The desired MIME type of the inference in the response.
	Accept *string

	// The MIME type of the input data in the request body.
	ContentType *string

	// Provides additional information about a request for an inference submitted to a
	// model hosted at an Amazon SageMaker endpoint. The information is an opaque value
	// that is forwarded verbatim. You could use this value, for example, to provide an
	// ID that you can use to track a request or to provide other metadata that a
	// service endpoint was programmed to process. The value must consist of no more
	// than 1024 visible US-ASCII characters as specified in Section 3.3.6. Field Value
	// Components (https://datatracker.ietf.org/doc/html/rfc7230#section-3.2.6) of the
	// Hypertext Transfer Protocol (HTTP/1.1). The code in your model is responsible
	// for setting or updating any custom attributes in the response. If your code does
	// not set this value in the response, an empty value is returned. For example, if
	// a custom attribute represents the trace ID, your model can prepend the custom
	// attribute with Trace ID: in your post-processing function. This feature is
	// currently supported in the Amazon Web Services SDKs but not in the Amazon
	// SageMaker Python SDK.
	CustomAttributes *string

	// The identifier for the inference request. Amazon SageMaker will generate an
	// identifier for you if none is specified.
	InferenceId *string

	// Maximum amount of time in seconds a request can be processed before it is marked
	// as expired. The default is 15 minutes, or 900 seconds.
	InvocationTimeoutSeconds *int32

	// Maximum age in seconds a request can be in the queue before it is marked as
	// expired. The default is 6 hours, or 21,600 seconds.
	RequestTTLSeconds *int32

	noSmithyDocumentSerde
}

type InvokeEndpointAsyncOutput struct {

	// Identifier for an inference request. This will be the same as the InferenceId
	// specified in the input. Amazon SageMaker will generate an identifier for you if
	// you do not specify one.
	InferenceId *string

	// The Amazon S3 URI where the inference response payload is stored.
	OutputLocation *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInvokeEndpointAsyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeEndpointAsync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeEndpointAsync{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpInvokeEndpointAsyncValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeEndpointAsync(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opInvokeEndpointAsync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "InvokeEndpointAsync",
	}
}
