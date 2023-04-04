// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details about an API destination.
func (c *Client) DescribeApiDestination(ctx context.Context, params *DescribeApiDestinationInput, optFns ...func(*Options)) (*DescribeApiDestinationOutput, error) {
	if params == nil {
		params = &DescribeApiDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApiDestination", params, optFns, c.addOperationDescribeApiDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApiDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeApiDestinationInput struct {

	// The name of the API destination to retrieve.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeApiDestinationOutput struct {

	// The ARN of the API destination retrieved.
	ApiDestinationArn *string

	// The state of the API destination retrieved.
	ApiDestinationState types.ApiDestinationState

	// The ARN of the connection specified for the API destination retrieved.
	ConnectionArn *string

	// A time stamp for the time that the API destination was created.
	CreationTime *time.Time

	// The description for the API destination retrieved.
	Description *string

	// The method to use to connect to the HTTP endpoint.
	HttpMethod types.ApiDestinationHttpMethod

	// The URL to use to connect to the HTTP endpoint.
	InvocationEndpoint *string

	// The maximum number of invocations per second to specified for the API
	// destination. Note that if you set the invocation rate maximum to a value lower
	// the rate necessary to send all events received on to the destination HTTP
	// endpoint, some events may not be delivered within the 24-hour retry window. If
	// you plan to set the rate lower than the rate necessary to deliver all events,
	// consider using a dead-letter queue to catch events that are not delivered within
	// 24 hours.
	InvocationRateLimitPerSecond *int32

	// A time stamp for the time that the API destination was last modified.
	LastModifiedTime *time.Time

	// The name of the API destination retrieved.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeApiDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApiDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApiDestination{}, middleware.After)
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
	if err = addOpDescribeApiDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApiDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeApiDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeApiDestination",
	}
}
