// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/matthew188/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This implementation of the DELETE action resets the default encryption for the
// bucket as server-side encryption with Amazon S3 managed keys (SSE-S3). For
// information about the bucket default encryption feature, see Amazon S3 Bucket
// Default Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the
// Amazon S3 User Guide. To use this operation, you must have permissions to
// perform the s3:PutEncryptionConfiguration action. The bucket owner has this
// permission by default. The bucket owner can grant this permission to others. For
// more information about permissions, see Permissions Related to Bucket
// Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html)
// in the Amazon S3 User Guide. Related Resources
//
// * PutBucketEncryption
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html)
//
// *
// GetBucketEncryption
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html)
func (c *Client) DeleteBucketEncryption(ctx context.Context, params *DeleteBucketEncryptionInput, optFns ...func(*Options)) (*DeleteBucketEncryptionOutput, error) {
	if params == nil {
		params = &DeleteBucketEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketEncryption", params, optFns, c.addOperationDeleteBucketEncryptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketEncryptionInput struct {

	// The name of the bucket containing the server-side encryption configuration to
	// delete.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type DeleteBucketEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBucketEncryptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketEncryption{}, middleware.After)
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteBucketEncryptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketEncryption(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addDeleteBucketEncryptionUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteBucketEncryption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketEncryption",
	}
}

// getDeleteBucketEncryptionBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getDeleteBucketEncryptionBucketMember(input interface{}) (*string, bool) {
	in := input.(*DeleteBucketEncryptionInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addDeleteBucketEncryptionUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getDeleteBucketEncryptionBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
