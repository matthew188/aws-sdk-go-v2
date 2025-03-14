// Code generated by smithy-go-codegen DO NOT EDIT.

package backupstorage

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/backupstorage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Get metadata associated with an Object.
func (c *Client) GetObjectMetadata(ctx context.Context, params *GetObjectMetadataInput, optFns ...func(*Options)) (*GetObjectMetadataOutput, error) {
	if params == nil {
		params = &GetObjectMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetObjectMetadata", params, optFns, c.addOperationGetObjectMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetObjectMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectMetadataInput struct {

	// Object token.
	//
	// This member is required.
	ObjectToken *string

	// Backup job id for the in-progress backup.
	//
	// This member is required.
	StorageJobId *string

	noSmithyDocumentSerde
}

type GetObjectMetadataOutput struct {

	// Metadata blob.
	MetadataBlob io.ReadCloser

	// MetadataBlob checksum.
	MetadataBlobChecksum *string

	// Checksum algorithm.
	MetadataBlobChecksumAlgorithm types.DataChecksumAlgorithm

	// The size of MetadataBlob.
	MetadataBlobLength int64

	// Metadata string.
	MetadataString *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetObjectMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetObjectMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetObjectMetadata{}, middleware.After)
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
	if err = addOpGetObjectMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetObjectMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-storage",
		OperationName: "GetObjectMetadata",
	}
}
