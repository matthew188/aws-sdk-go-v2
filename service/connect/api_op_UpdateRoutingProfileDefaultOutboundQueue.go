// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the default outbound queue of a routing profile.
func (c *Client) UpdateRoutingProfileDefaultOutboundQueue(ctx context.Context, params *UpdateRoutingProfileDefaultOutboundQueueInput, optFns ...func(*Options)) (*UpdateRoutingProfileDefaultOutboundQueueOutput, error) {
	if params == nil {
		params = &UpdateRoutingProfileDefaultOutboundQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoutingProfileDefaultOutboundQueue", params, optFns, c.addOperationUpdateRoutingProfileDefaultOutboundQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoutingProfileDefaultOutboundQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoutingProfileDefaultOutboundQueueInput struct {

	// The identifier for the default outbound queue.
	//
	// This member is required.
	DefaultOutboundQueueId *string

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier of the routing profile.
	//
	// This member is required.
	RoutingProfileId *string

	noSmithyDocumentSerde
}

type UpdateRoutingProfileDefaultOutboundQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRoutingProfileDefaultOutboundQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRoutingProfileDefaultOutboundQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRoutingProfileDefaultOutboundQueue{}, middleware.After)
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
	if err = addOpUpdateRoutingProfileDefaultOutboundQueueValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoutingProfileDefaultOutboundQueue(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRoutingProfileDefaultOutboundQueue(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateRoutingProfileDefaultOutboundQueue",
	}
}
