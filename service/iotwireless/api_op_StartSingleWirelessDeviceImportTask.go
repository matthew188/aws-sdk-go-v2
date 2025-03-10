// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Start import task for a single wireless device.
func (c *Client) StartSingleWirelessDeviceImportTask(ctx context.Context, params *StartSingleWirelessDeviceImportTaskInput, optFns ...func(*Options)) (*StartSingleWirelessDeviceImportTaskOutput, error) {
	if params == nil {
		params = &StartSingleWirelessDeviceImportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSingleWirelessDeviceImportTask", params, optFns, c.addOperationStartSingleWirelessDeviceImportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSingleWirelessDeviceImportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSingleWirelessDeviceImportTaskInput struct {

	// The name of the Sidewalk destination that describes the IoT rule to route
	// messages from the device in the import task that will be onboarded to AWS IoT
	// Wireless.
	//
	// This member is required.
	DestinationName *string

	// The Sidewalk-related parameters for importing a single wireless device.
	//
	// This member is required.
	Sidewalk *types.SidewalkSingleStartImportInfo

	// Each resource must have a unique client request token. If you try to create a
	// new resource with the same token as a resource that already exists, an exception
	// occurs. If you omit this value, AWS SDKs will automatically generate a unique
	// client request.
	ClientRequestToken *string

	// The name of the wireless device for which an import task is being started.
	DeviceName *string

	// The tag to attach to the specified resource. Tags are metadata that you can use
	// to manage a resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartSingleWirelessDeviceImportTaskOutput struct {

	// The ARN (Amazon Resource Name) of the import task.
	Arn *string

	// The import task ID.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSingleWirelessDeviceImportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSingleWirelessDeviceImportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSingleWirelessDeviceImportTask{}, middleware.After)
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
	if err = addIdempotencyToken_opStartSingleWirelessDeviceImportTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSingleWirelessDeviceImportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSingleWirelessDeviceImportTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartSingleWirelessDeviceImportTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSingleWirelessDeviceImportTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSingleWirelessDeviceImportTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSingleWirelessDeviceImportTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSingleWirelessDeviceImportTaskInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartSingleWirelessDeviceImportTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartSingleWirelessDeviceImportTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSingleWirelessDeviceImportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "StartSingleWirelessDeviceImportTask",
	}
}
