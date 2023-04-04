// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details about a connection.
func (c *Client) DescribeConnection(ctx context.Context, params *DescribeConnectionInput, optFns ...func(*Options)) (*DescribeConnectionOutput, error) {
	if params == nil {
		params = &DescribeConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnection", params, optFns, c.addOperationDescribeConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectionInput struct {

	// The name of the connection to retrieve.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeConnectionOutput struct {

	// The parameters to use for authorization for the connection.
	AuthParameters *types.ConnectionAuthResponseParameters

	// The type of authorization specified for the connection.
	AuthorizationType types.ConnectionAuthorizationType

	// The ARN of the connection retrieved.
	ConnectionArn *string

	// The state of the connection retrieved.
	ConnectionState types.ConnectionState

	// A time stamp for the time that the connection was created.
	CreationTime *time.Time

	// The description for the connection retrieved.
	Description *string

	// A time stamp for the time that the connection was last authorized.
	LastAuthorizedTime *time.Time

	// A time stamp for the time that the connection was last modified.
	LastModifiedTime *time.Time

	// The name of the connection retrieved.
	Name *string

	// The ARN of the secret created from the authorization parameters specified for
	// the connection.
	SecretArn *string

	// The reason that the connection is in the current connection state.
	StateReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConnection{}, middleware.After)
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
	if err = addOpDescribeConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeConnection",
	}
}
