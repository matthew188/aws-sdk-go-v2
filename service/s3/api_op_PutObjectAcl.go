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

// Uses the acl subresource to set the access control list (ACL) permissions for a
// new or existing object in an S3 bucket. You must have WRITE_ACP permission to
// set the ACL of an object. For more information, see What permissions can I
// grant?
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#permissions)
// in the Amazon S3 User Guide. This action is not supported by Amazon S3 on
// Outposts. Depending on your application needs, you can choose to set the ACL on
// an object using either the request body or the headers. For example, if you have
// an existing application that updates a bucket ACL using the request body, you
// can continue to use that approach. For more information, see Access Control List
// (ACL) Overview
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html) in the
// Amazon S3 User Guide. If your bucket uses the bucket owner enforced setting for
// S3 Object Ownership, ACLs are disabled and no longer affect permissions. You
// must use policies to grant access to your bucket and the objects in it. Requests
// to set ACLs or update ACLs fail and return the AccessControlListNotSupported
// error code. Requests to read ACLs are still supported. For more information, see
// Controlling object ownership
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html)
// in the Amazon S3 User Guide. Access Permissions You can set access permissions
// using one of the following methods:
//
// * Specify a canned ACL with the x-amz-acl
// request header. Amazon S3 supports a set of predefined ACLs, known as canned
// ACLs. Each canned ACL has a predefined set of grantees and permissions. Specify
// the canned ACL name as the value of x-amz-acl. If you use this header, you
// cannot use other access control-specific headers in your request. For more
// information, see Canned ACL
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
//
// *
// Specify access permissions explicitly with the x-amz-grant-read,
// x-amz-grant-read-acp, x-amz-grant-write-acp, and x-amz-grant-full-control
// headers. When using these headers, you specify explicit access permissions and
// grantees (Amazon Web Services accounts or Amazon S3 groups) who will receive the
// permission. If you use these ACL-specific headers, you cannot use x-amz-acl
// header to set a canned ACL. These parameters map to the set of permissions that
// Amazon S3 supports in an ACL. For more information, see Access Control List
// (ACL) Overview
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html). You specify
// each grantee as a type=value pair, where the type is one of the following:
//
// * id
// – if the value specified is the canonical user ID of an Amazon Web Services
// account
//
// * uri – if you are granting permissions to a predefined group
//
// *
// emailAddress – if the value specified is the email address of an Amazon Web
// Services account Using email addresses to specify a grantee is only supported in
// the following Amazon Web Services Regions:
//
// * US East (N. Virginia)
//
// * US West
// (N. California)
//
// * US West (Oregon)
//
// * Asia Pacific (Singapore)
//
// * Asia Pacific
// (Sydney)
//
// * Asia Pacific (Tokyo)
//
// * Europe (Ireland)
//
// * South America (São
// Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see
// Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) in the
// Amazon Web Services General Reference.
//
// For example, the following
// x-amz-grant-read header grants list objects permission to the two Amazon Web
// Services accounts identified by their email addresses. x-amz-grant-read:
// emailAddress="xyz@amazon.com", emailAddress="abc@amazon.com"
//
// You can use either
// a canned ACL or specify access permissions explicitly. You cannot do both.
// Grantee Values You can specify the person (grantee) to whom you're assigning
// access rights (using request elements) in the following ways:
//
// * By the person's
// ID: <>ID<><>GranteesEmail<>  DisplayName is optional and ignored in the
// request.
//
// * By URI:
// <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
// * By Email
// address: <>Grantees@email.com<>lt;/Grantee> The grantee is resolved to the
// CanonicalUser and, in a response to a GET Object acl request, appears as the
// CanonicalUser. Using email addresses to specify a grantee is only supported in
// the following Amazon Web Services Regions:
//
// * US East (N. Virginia)
//
// * US West
// (N. California)
//
// * US West (Oregon)
//
// * Asia Pacific (Singapore)
//
// * Asia Pacific
// (Sydney)
//
// * Asia Pacific (Tokyo)
//
// * Europe (Ireland)
//
// * South America (São
// Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see
// Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) in the
// Amazon Web Services General Reference.
//
// Versioning The ACL of an object is set
// at the object version level. By default, PUT sets the ACL of the current version
// of an object. To set the ACL of a different version, use the versionId
// subresource. Related Resources
//
// * CopyObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html)
//
// *
// GetObject (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
func (c *Client) PutObjectAcl(ctx context.Context, params *PutObjectAclInput, optFns ...func(*Options)) (*PutObjectAclOutput, error) {
	if params == nil {
		params = &PutObjectAclInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutObjectAcl", params, optFns, c.addOperationPutObjectAclMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutObjectAclOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutObjectAclInput struct {

	// The bucket name that contains the object to which you want to attach the ACL.
	// When using this action with an access point, you must direct requests to the
	// access point hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// Key for which the PUT action was initiated. When using this action with an
	// access point, you must direct requests to the access point hostname. The access
	// point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When you use this action with Amazon S3 on
	// Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on
	// Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
	// Amazon S3 User Guide.
	//
	// This member is required.
	Key *string

	// The canned ACL to apply to the object. For more information, see Canned ACL
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL).
	ACL types.ObjectCannedACL

	// Contains the elements that set the ACL permissions for an object per grantee.
	AccessControlPolicy *types.AccessControlPolicy

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

	// The base64-encoded 128-bit MD5 digest of the data. This header must be used as a
	// message integrity check to verify that the request body was not corrupted in
	// transit. For more information, go to RFC 1864.>
	// (http://www.ietf.org/rfc/rfc1864.txt) For requests made using the Amazon Web
	// Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is
	// calculated automatically.
	ContentMD5 *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket. This action is not supported by Amazon S3 on Outposts.
	GrantFullControl *string

	// Allows grantee to list the objects in the bucket. This action is not supported
	// by Amazon S3 on Outposts.
	GrantRead *string

	// Allows grantee to read the bucket ACL. This action is not supported by Amazon S3
	// on Outposts.
	GrantReadACP *string

	// Allows grantee to create new objects in the bucket. For the bucket and object
	// owners of existing objects, also allows deletions and overwrites of those
	// objects.
	GrantWrite *string

	// Allows grantee to write the ACL for the applicable bucket. This action is not
	// supported by Amazon S3 on Outposts.
	GrantWriteACP *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from Requester Pays buckets, see Downloading Objects
	// in Requester Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 User Guide.
	RequestPayer types.RequestPayer

	// VersionId used to reference a specific version of the object.
	VersionId *string

	noSmithyDocumentSerde
}

type PutObjectAclOutput struct {

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutObjectAclMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutObjectAcl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutObjectAcl{}, middleware.After)
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
	if err = addOpPutObjectAclValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutObjectAcl(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutObjectAclInputChecksumMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addPutObjectAclUpdateEndpoint(stack, options); err != nil {
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

func newServiceMetadataMiddleware_opPutObjectAcl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutObjectAcl",
	}
}

// getPutObjectAclRequestAlgorithmMember gets the request checksum algorithm value
// provided as input.
func getPutObjectAclRequestAlgorithmMember(input interface{}) (string, bool) {
	in := input.(*PutObjectAclInput)
	if len(in.ChecksumAlgorithm) == 0 {
		return "", false
	}
	return string(in.ChecksumAlgorithm), true
}

func addPutObjectAclInputChecksumMiddlewares(stack *middleware.Stack, options Options) error {
	return internalChecksum.AddInputMiddleware(stack, internalChecksum.InputMiddlewareOptions{
		GetAlgorithm:                     getPutObjectAclRequestAlgorithmMember,
		RequireChecksum:                  true,
		EnableTrailingChecksum:           false,
		EnableComputeSHA256PayloadHash:   true,
		EnableDecodedContentLengthHeader: true,
	})
}

// getPutObjectAclBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getPutObjectAclBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutObjectAclInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutObjectAclUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutObjectAclBucketMember,
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
