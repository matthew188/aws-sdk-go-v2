// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing Amazon FSx for ONTAP storage virtual machine (SVM). Prior to
// deleting an SVM, you must delete all non-root volumes in the SVM, otherwise the
// operation will fail.
func (c *Client) DeleteStorageVirtualMachine(ctx context.Context, params *DeleteStorageVirtualMachineInput, optFns ...func(*Options)) (*DeleteStorageVirtualMachineOutput, error) {
	if params == nil {
		params = &DeleteStorageVirtualMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStorageVirtualMachine", params, optFns, c.addOperationDeleteStorageVirtualMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStorageVirtualMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteStorageVirtualMachineInput struct {

	// The ID of the SVM that you want to delete.
	//
	// This member is required.
	StorageVirtualMachineId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 64
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type DeleteStorageVirtualMachineOutput struct {

	// Describes the lifecycle state of the SVM being deleted.
	Lifecycle types.StorageVirtualMachineLifecycle

	// The ID of the SVM Amazon FSx is deleting.
	StorageVirtualMachineId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteStorageVirtualMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteStorageVirtualMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteStorageVirtualMachine{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteStorageVirtualMachineMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteStorageVirtualMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStorageVirtualMachine(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteStorageVirtualMachine struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteStorageVirtualMachine) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteStorageVirtualMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteStorageVirtualMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteStorageVirtualMachineInput ")
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
func addIdempotencyToken_opDeleteStorageVirtualMachineMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteStorageVirtualMachine{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteStorageVirtualMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DeleteStorageVirtualMachine",
	}
}
