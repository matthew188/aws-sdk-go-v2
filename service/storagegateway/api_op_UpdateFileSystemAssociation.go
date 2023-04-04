// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a file system association. This operation is only supported in the FSx
// File Gateways.
func (c *Client) UpdateFileSystemAssociation(ctx context.Context, params *UpdateFileSystemAssociationInput, optFns ...func(*Options)) (*UpdateFileSystemAssociationOutput, error) {
	if params == nil {
		params = &UpdateFileSystemAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFileSystemAssociation", params, optFns, c.addOperationUpdateFileSystemAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFileSystemAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFileSystemAssociationInput struct {

	// The Amazon Resource Name (ARN) of the file system association that you want to
	// update.
	//
	// This member is required.
	FileSystemAssociationARN *string

	// The Amazon Resource Name (ARN) of the storage used for the audit logs.
	AuditDestinationARN *string

	// The refresh cache information for the file share or FSx file systems.
	CacheAttributes *types.CacheAttributes

	// The password of the user credential.
	Password *string

	// The user name of the user credential that has permission to access the root
	// share D$ of the Amazon FSx file system. The user account must belong to the
	// Amazon FSx delegated admin user group.
	UserName *string

	noSmithyDocumentSerde
}

type UpdateFileSystemAssociationOutput struct {

	// The ARN of the updated file system association.
	FileSystemAssociationARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFileSystemAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFileSystemAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFileSystemAssociation{}, middleware.After)
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
	if err = addOpUpdateFileSystemAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFileSystemAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFileSystemAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateFileSystemAssociation",
	}
}
