// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	internalChecksum "github.com/matthew188/aws-sdk-go-v2/service/internal/checksum"
	s3cust "github.com/matthew188/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/matthew188/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set the logging parameters for a bucket and to specify permissions for who can
// view and modify the logging parameters. All logs are saved to buckets in the
// same Amazon Web Services Region as the source bucket. To set the logging status
// of a bucket, you must be the bucket owner. The bucket owner is automatically
// granted FULL_CONTROL to all logs. You use the Grantee request element to grant
// access to other people. The Permissions request element specifies the kind of
// access the grantee has to the logs. If the target bucket for log delivery uses
// the bucket owner enforced setting for S3 Object Ownership, you can't use the
// Grantee request element to grant access to others. Permissions can only be
// granted using policies. For more information, see Permissions for server access
// log delivery
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html#grant-log-delivery-permissions-general)
// in the Amazon S3 User Guide. Grantee Values You can specify the person (grantee)
// to whom you're assigning access rights (using request elements) in the following
// ways:
//
// * By the person's ID: <>ID<><>GranteesEmail<>  DisplayName is optional
// and ignored in the request.
//
// * By Email address:  <>Grantees@email.com<> The
// grantee is resolved to the CanonicalUser and, in a response to a GET Object acl
// request, appears as the CanonicalUser.
//
// * By URI:
// <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
// To enable
// logging, you use LoggingEnabled and its children request elements. To disable
// logging, you use an empty BucketLoggingStatus request element:  For more
// information about server access logging, see Server Access Logging
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerLogs.html) in the
// Amazon S3 User Guide. For more information about creating a bucket, see
// CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html). For
// more information about returning the logging status of a bucket, see
// GetBucketLogging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html). The
// following operations are related to PutBucketLogging:
//
// * PutObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//
// *
// DeleteBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
//
// *
// CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
//
// *
// GetBucketLogging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html)
func (c *Client) PutBucketLogging(ctx context.Context, params *PutBucketLoggingInput, optFns ...func(*Options)) (*PutBucketLoggingOutput, error) {
	if params == nil {
		params = &PutBucketLoggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketLogging", params, optFns, c.addOperationPutBucketLoggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketLoggingInput struct {

	// The name of the bucket for which to set the logging parameters.
	//
	// This member is required.
	Bucket *string

	// Container for logging status information.
	//
	// This member is required.
	BucketLoggingStatus *types.BucketLoggingStatus

	// Indicates the algorithm used to create the checksum for the object when using
	// the SDK. This header will not provide any additional functionality if not using
	// the SDK. When sending this header, there must be a corresponding x-amz-checksum
	// or x-amz-trailer header sent. Otherwise, Amazon S3 fails the request with the
	// HTTP status code 400 Bad Request. For more information, see Checking object
	// integrity
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/checking-object-integrity.html)
	// in the Amazon S3 User Guide. If you provide an individual checksum, Amazon S3
	// ignores any provided ChecksumAlgorithm parameter.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The MD5 hash of the PutBucketLogging request body. For requests made using the
	// Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs,
	// this field is calculated automatically.
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

type PutBucketLoggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketLoggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketLogging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketLogging{}, middleware.After)
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
	if err = addOpPutBucketLoggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketLogging(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketLoggingInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addPutBucketLoggingUpdateEndpoint(stack, options); err != nil {
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

func newServiceMetadataMiddleware_opPutBucketLogging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketLogging",
	}
}

// getPutBucketLoggingRequestAlgorithmMember gets the request checksum algorithm
// value provided as input.
func getPutBucketLoggingRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutBucketLoggingInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutBucketLoggingInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutBucketLoggingRequestAlgorithmMember,
		RequireChecksum:                  true,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutBucketLoggingBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutBucketLoggingBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketLoggingInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketLoggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketLoggingBucketMember,
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
