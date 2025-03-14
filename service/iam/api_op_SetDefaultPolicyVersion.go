// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the specified version of the specified policy as the policy's default
// (operative) version. This operation affects all users, groups, and roles that
// the policy is attached to. To list the users, groups, and roles that the policy
// is attached to, use ListEntitiesForPolicy. For information about managed
// policies, see Managed policies and inline policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
func (c *Client) SetDefaultPolicyVersion(ctx context.Context, params *SetDefaultPolicyVersionInput, optFns ...func(*Options)) (*SetDefaultPolicyVersionOutput, error) {
	if params == nil {
		params = &SetDefaultPolicyVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetDefaultPolicyVersion", params, optFns, c.addOperationSetDefaultPolicyVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetDefaultPolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetDefaultPolicyVersionInput struct {

	// The Amazon Resource Name (ARN) of the IAM policy whose default version you want
	// to set. For more information about ARNs, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference.
	//
	// This member is required.
	PolicyArn *string

	// The version of the policy to set as the default (operative) version. For more
	// information about managed policy versions, see Versioning for managed policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
	// in the IAM User Guide.
	//
	// This member is required.
	VersionId *string

	noSmithyDocumentSerde
}

type SetDefaultPolicyVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetDefaultPolicyVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetDefaultPolicyVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetDefaultPolicyVersion{}, middleware.After)
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
	if err = addOpSetDefaultPolicyVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetDefaultPolicyVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetDefaultPolicyVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "SetDefaultPolicyVersion",
	}
}
