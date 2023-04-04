// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This example operation serializes a request without an HTTP body. These tests
// are to ensure we do not attach a body or related headers (Content-Length,
// Content-Type) to operations that semantically cannot produce an HTTP body.
func (c *Client) TestNoPayload(ctx context.Context, params *TestNoPayloadInput, optFns ...func(*Options)) (*TestNoPayloadOutput, error) {
	if params == nil {
		params = &TestNoPayloadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestNoPayload", params, optFns, c.addOperationTestNoPayloadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestNoPayloadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestNoPayloadInput struct {
	TestId *string

	noSmithyDocumentSerde
}

type TestNoPayloadOutput struct {
	TestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestNoPayloadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestNoPayload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestNoPayload{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestNoPayload(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestNoPayload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "TestNoPayload",
	}
}
