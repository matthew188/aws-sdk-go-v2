// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds the specified list of principals and list of resources to a resource share.
// Principals that already have access to this resource share immediately receive
// access to the added resources. Newly added principals immediately receive access
// to the resources shared in this resource share.
func (c *Client) AssociateResourceShare(ctx context.Context, params *AssociateResourceShareInput, optFns ...func(*Options)) (*AssociateResourceShareOutput, error) {
	if params == nil {
		params = &AssociateResourceShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateResourceShare", params, optFns, c.addOperationAssociateResourceShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateResourceShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResourceShareInput struct {

	// Specifies the Amazon Resoure Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the resource share that you want to add principals or resources to.
	//
	// This member is required.
	ResourceShareArn *string

	// Specifies a unique, case-sensitive identifier that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value.
	// (https://wikipedia.org/wiki/Universally_unique_identifier). If you don't provide
	// this value, then Amazon Web Services generates a random one for you.
	ClientToken *string

	// Specifies a list of principals to whom you want to the resource share. This can
	// be null if you want to add only resources. What the principals can do with the
	// resources in the share is determined by the RAM permissions that you associate
	// with the resource share. See AssociateResourceSharePermission. You can include
	// the following values:
	//
	// * An Amazon Web Services account ID, for example:
	// 123456789012
	//
	// * An Amazon Resoure Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// an organization in Organizations, for example:
	// organizations::123456789012:organization/o-exampleorgid
	//
	// * An ARN of an
	// organizational unit (OU) in Organizations, for example:
	// organizations::123456789012:ou/o-exampleorgid/ou-examplerootid-exampleouid123
	//
	// *
	// An ARN of an IAM role, for example: iam::123456789012:role/rolename
	//
	// * An ARN of
	// an IAM user, for example: iam::123456789012user/username
	//
	// Not all resource types
	// can be shared with IAM roles and users. For more information, see Sharing with
	// IAM roles and users
	// (https://docs.aws.amazon.com/ram/latest/userguide/permissions.html#permissions-rbp-supported-resource-types)
	// in the Resource Access Manager User Guide.
	Principals []string

	// Specifies a list of Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the resources that you want to share. This can be null if you want to add only
	// principals.
	ResourceArns []string

	noSmithyDocumentSerde
}

type AssociateResourceShareOutput struct {

	// The idempotency identifier associated with this request. If you want to repeat
	// the same operation in an idempotent manner then you must include this value in
	// the clientToken request parameter of that later call. All other parameters must
	// also have the same values that you used in the first call.
	ClientToken *string

	// An array of objects that contain information about the associations.
	ResourceShareAssociations []types.ResourceShareAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateResourceShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateResourceShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateResourceShare{}, middleware.After)
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
	if err = addOpAssociateResourceShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResourceShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateResourceShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "AssociateResourceShare",
	}
}
