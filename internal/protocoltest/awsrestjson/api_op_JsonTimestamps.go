// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This tests how timestamps are serialized, including using the default format of
// date-time and various @timestampFormat trait values.
func (c *Client) JsonTimestamps(ctx context.Context, params *JsonTimestampsInput, optFns ...func(*Options)) (*JsonTimestampsOutput, error) {
	if params == nil {
		params = &JsonTimestampsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "JsonTimestamps", params, optFns, c.addOperationJsonTimestampsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*JsonTimestampsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type JsonTimestampsInput struct {
	DateTime *time.Time

	DateTimeOnTarget *time.Time

	EpochSeconds *time.Time

	EpochSecondsOnTarget *time.Time

	HttpDate *time.Time

	HttpDateOnTarget *time.Time

	Normal *time.Time

	noSmithyDocumentSerde
}

type JsonTimestampsOutput struct {
	DateTime *time.Time

	DateTimeOnTarget *time.Time

	EpochSeconds *time.Time

	EpochSecondsOnTarget *time.Time

	HttpDate *time.Time

	HttpDateOnTarget *time.Time

	Normal *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationJsonTimestampsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpJsonTimestamps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpJsonTimestamps{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opJsonTimestamps(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opJsonTimestamps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "JsonTimestamps",
	}
}
