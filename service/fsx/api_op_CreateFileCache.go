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

// Creates a new Amazon File Cache resource. You can use this operation with a
// client request token in the request that Amazon File Cache uses to ensure
// idempotent creation. If a cache with the specified client request token exists
// and the parameters match, CreateFileCache returns the description of the
// existing cache. If a cache with the specified client request token exists and
// the parameters don't match, this call returns IncompatibleParameterError. If a
// file cache with the specified client request token doesn't exist,
// CreateFileCache does the following:
//
// * Creates a new, empty Amazon File Cache
// resourcewith an assigned ID, and an initial lifecycle state of CREATING.
//
// *
// Returns the description of the cache in JSON format.
//
// The CreateFileCache call
// returns while the cache's lifecycle state is still CREATING. You can check the
// cache creation status by calling the DescribeFileCaches
// (https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileCaches.html)
// operation, which returns the cache state along with other information.
func (c *Client) CreateFileCache(ctx context.Context, params *CreateFileCacheInput, optFns ...func(*Options)) (*CreateFileCacheOutput, error) {
	if params == nil {
		params = &CreateFileCacheInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFileCache", params, optFns, c.addOperationCreateFileCacheMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFileCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFileCacheInput struct {

	// The type of cache that you're creating, which must be LUSTRE.
	//
	// This member is required.
	FileCacheType types.FileCacheType

	// Sets the Lustre version for the cache that you're creating, which must be 2.12.
	//
	// This member is required.
	FileCacheTypeVersion *string

	// The storage capacity of the cache in gibibytes (GiB). Valid values are 1200 GiB,
	// 2400 GiB, and increments of 2400 GiB.
	//
	// This member is required.
	StorageCapacity *int32

	// A list of subnet IDs that the cache will be accessible from. You can specify
	// only one subnet ID in a call to the CreateFileCache operation.
	//
	// This member is required.
	SubnetIds []string

	// An idempotency token for resource creation, in a string of up to 64 ASCII
	// characters. This token is automatically filled on your behalf when you use the
	// Command Line Interface (CLI) or an Amazon Web Services SDK. By using the
	// idempotent operation, you can retry a CreateFileCache operation without the risk
	// of creating an extra cache. This approach can be useful when an initial call
	// fails in a way that makes it unclear whether a cache was created. Examples are
	// if a transport level timeout occurred, or your connection was reset. If you use
	// the same client request token and the initial call created a cache, the client
	// receives success as long as the parameters are the same.
	ClientRequestToken *string

	// A boolean flag indicating whether tags for the cache should be copied to data
	// repository associations. This value defaults to false.
	CopyTagsToDataRepositoryAssociations *bool

	// A list of up to 8 configurations for data repository associations (DRAs) to be
	// created during the cache creation. The DRAs link the cache to either an Amazon
	// S3 data repository or a Network File System (NFS) data repository that supports
	// the NFSv3 protocol. The DRA configurations must meet the following
	// requirements:
	//
	// * All configurations on the list must be of the same data
	// repository type, either all S3 or all NFS. A cache can't link to different data
	// repository types at the same time.
	//
	// * An NFS DRA must link to an NFS file system
	// that supports the NFSv3 protocol.
	//
	// DRA automatic import and automatic export is
	// not supported.
	DataRepositoryAssociations []types.FileCacheDataRepositoryAssociation

	// Specifies the ID of the Key Management Service (KMS) key to use for encrypting
	// data on an Amazon File Cache. If a KmsKeyId isn't specified, the Amazon
	// FSx-managed KMS key for your account is used. For more information, see Encrypt
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_Encrypt.html) in the
	// Key Management Service API Reference.
	KmsKeyId *string

	// The configuration for the Amazon File Cache resource being created.
	LustreConfiguration *types.CreateFileCacheLustreConfiguration

	// A list of IDs specifying the security groups to apply to all network interfaces
	// created for Amazon File Cache access. This list isn't returned in later requests
	// to describe the cache.
	SecurityGroupIds []string

	// A list of Tag values, with a maximum of 50 elements.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFileCacheOutput struct {

	// A description of the cache that was created.
	FileCache *types.FileCacheCreating

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFileCacheMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFileCache{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFileCache{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateFileCacheMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFileCacheValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFileCache(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFileCache struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFileCache) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFileCache) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFileCacheInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFileCacheInput ")
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
func addIdempotencyToken_opCreateFileCacheMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFileCache{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFileCache(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "CreateFileCache",
	}
}
