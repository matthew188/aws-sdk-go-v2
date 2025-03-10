// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Calls back Chime SDK Messaging with a processing response message. This should
// be invoked from the processor Lambda. This is a developer API. You can return
// one of the following processing responses:
//
// * Update message content or
// metadata
//
// * Deny a message
//
// * Make no changes to the message
func (c *Client) ChannelFlowCallback(ctx context.Context, params *ChannelFlowCallbackInput, optFns ...func(*Options)) (*ChannelFlowCallbackOutput, error) {
	if params == nil {
		params = &ChannelFlowCallbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChannelFlowCallback", params, optFns, c.addOperationChannelFlowCallbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChannelFlowCallbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ChannelFlowCallbackInput struct {

	// The identifier passed to the processor by the service when invoked. Use the
	// identifier to call back the service.
	//
	// This member is required.
	CallbackId *string

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// Stores information about the processed message.
	//
	// This member is required.
	ChannelMessage *types.ChannelMessageCallback

	// When a processor determines that a message needs to be DENIED, pass this
	// parameter with a value of true.
	DeleteResource bool

	noSmithyDocumentSerde
}

type ChannelFlowCallbackOutput struct {

	// The call back ID passed in the request.
	CallbackId *string

	// The ARN of the channel.
	ChannelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationChannelFlowCallbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpChannelFlowCallback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpChannelFlowCallback{}, middleware.After)
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
	if err = addIdempotencyToken_opChannelFlowCallbackMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpChannelFlowCallbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChannelFlowCallback(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpChannelFlowCallback struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpChannelFlowCallback) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpChannelFlowCallback) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ChannelFlowCallbackInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ChannelFlowCallbackInput ")
	}

	if input.CallbackId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CallbackId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opChannelFlowCallbackMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpChannelFlowCallback{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opChannelFlowCallback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ChannelFlowCallback",
	}
}
