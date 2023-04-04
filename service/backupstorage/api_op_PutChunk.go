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

// Upload chunk.
func (c *Client) PutChunk(ctx context.Context, params *PutChunkInput, optFns ...func(*Options)) (*PutChunkOutput, error) {
	if params == nil {
		params = &PutChunkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutChunk", params, optFns, c.addOperationPutChunkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutChunkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutChunkInput struct {

	// Backup job Id for the in-progress backup.
	//
	// This member is required.
	BackupJobId *string

	// Data checksum
	//
	// This member is required.
	Checksum *string

	// Checksum algorithm
	//
	// This member is required.
	ChecksumAlgorithm types.DataChecksumAlgorithm

	// Describes this chunk's position relative to the other chunks
	//
	// This member is required.
	ChunkIndex int64

	// Data to be uploaded
	//
	// This member is required.
	Data io.Reader

	// Data length
	//
	// This member is required.
	Length int64

	// Upload Id for the in-progress upload.
	//
	// This member is required.
	UploadId *string

	noSmithyDocumentSerde
}

type PutChunkOutput struct {

	// Chunk checksum
	//
	// This member is required.
	ChunkChecksum *string

	// Checksum algorithm
	//
	// This member is required.
	ChunkChecksumAlgorithm types.DataChecksumAlgorithm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutChunkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutChunk{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutChunk{}, middleware.After)
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
	if err = v4.AddUnsignedPayloadMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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
	if err = addOpPutChunkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutChunk(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutChunk(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-storage",
		OperationName: "PutChunk",
	}
}
