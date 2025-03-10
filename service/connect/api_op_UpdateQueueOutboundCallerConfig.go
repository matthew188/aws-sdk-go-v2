// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Updates the outbound caller ID name, number, and outbound whisper flow for a
// specified queue. If the number being used in the input is claimed to a traffic
// distribution group, and you are calling this API using an instance in the Amazon
// Web Services Region where the traffic distribution group was created, you can
// use either a full phone number ARN or UUID value for the
// OutboundCallerIdNumberId value of the OutboundCallerConfig
// (https://docs.aws.amazon.com/connect/latest/APIReference/API_OutboundCallerConfig)
// request body parameter. However, if the number is claimed to a traffic
// distribution group and you are calling this API using an instance in the
// alternate Amazon Web Services Region associated with the traffic distribution
// group, you must provide a full phone number ARN. If a UUID is provided in this
// scenario, you will receive a ResourceNotFoundException.
func (c *Client) UpdateQueueOutboundCallerConfig(ctx context.Context, params *UpdateQueueOutboundCallerConfigInput, optFns ...func(*Options)) (*UpdateQueueOutboundCallerConfigOutput, error) {
	if params == nil {
		params = &UpdateQueueOutboundCallerConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateQueueOutboundCallerConfig", params, optFns, c.addOperationUpdateQueueOutboundCallerConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateQueueOutboundCallerConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateQueueOutboundCallerConfigInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The outbound caller ID name, number, and outbound whisper flow.
	//
	// This member is required.
	OutboundCallerConfig *types.OutboundCallerConfig

	// The identifier for the queue.
	//
	// This member is required.
	QueueId *string

	noSmithyDocumentSerde
}

type UpdateQueueOutboundCallerConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateQueueOutboundCallerConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateQueueOutboundCallerConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateQueueOutboundCallerConfig{}, middleware.After)
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
	if err = addOpUpdateQueueOutboundCallerConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateQueueOutboundCallerConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateQueueOutboundCallerConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateQueueOutboundCallerConfig",
	}
}
