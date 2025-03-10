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

// Deletes a data repository association on an Amazon FSx for Lustre file system.
// Deleting the data repository association unlinks the file system from the Amazon
// S3 bucket. When deleting a data repository association, you have the option of
// deleting the data in the file system that corresponds to the data repository
// association. Data repository associations are supported only for file systems
// with the Persistent_2 deployment type.
func (c *Client) DeleteDataRepositoryAssociation(ctx context.Context, params *DeleteDataRepositoryAssociationInput, optFns ...func(*Options)) (*DeleteDataRepositoryAssociationOutput, error) {
	if params == nil {
		params = &DeleteDataRepositoryAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataRepositoryAssociation", params, optFns, c.addOperationDeleteDataRepositoryAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDataRepositoryAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDataRepositoryAssociationInput struct {

	// The ID of the data repository association that you want to delete.
	//
	// This member is required.
	AssociationId *string

	// (Optional) An idempotency token for resource creation, in a string of up to 64
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// Set to true to delete the data in the file system that corresponds to the data
	// repository association.
	DeleteDataInFileSystem *bool

	noSmithyDocumentSerde
}

type DeleteDataRepositoryAssociationOutput struct {

	// The ID of the data repository association being deleted.
	AssociationId *string

	// Indicates whether data in the file system that corresponds to the data
	// repository association is being deleted. Default is false.
	DeleteDataInFileSystem *bool

	// Describes the lifecycle state of the data repository association being deleted.
	Lifecycle types.DataRepositoryLifecycle

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDataRepositoryAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDataRepositoryAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDataRepositoryAssociation{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteDataRepositoryAssociationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteDataRepositoryAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataRepositoryAssociation(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteDataRepositoryAssociation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteDataRepositoryAssociation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteDataRepositoryAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteDataRepositoryAssociationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteDataRepositoryAssociationInput ")
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
func addIdempotencyToken_opDeleteDataRepositoryAssociationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteDataRepositoryAssociation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteDataRepositoryAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DeleteDataRepositoryAssociation",
	}
}
