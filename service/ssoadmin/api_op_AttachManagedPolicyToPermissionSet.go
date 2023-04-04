// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches an AWS managed policy ARN to a permission set. If the permission set is
// already referenced by one or more account assignments, you will need to call
// ProvisionPermissionSet after this operation. Calling ProvisionPermissionSet
// applies the corresponding IAM policy updates to all assigned accounts.
func (c *Client) AttachManagedPolicyToPermissionSet(ctx context.Context, params *AttachManagedPolicyToPermissionSetInput, optFns ...func(*Options)) (*AttachManagedPolicyToPermissionSetOutput, error) {
	if params == nil {
		params = &AttachManagedPolicyToPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachManagedPolicyToPermissionSet", params, optFns, c.addOperationAttachManagedPolicyToPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachManagedPolicyToPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachManagedPolicyToPermissionSetInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// AWS Service Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The AWS managed policy ARN to be attached to a permission set.
	//
	// This member is required.
	ManagedPolicyArn *string

	// The ARN of the PermissionSet that the managed policy should be attached to.
	//
	// This member is required.
	PermissionSetArn *string

	noSmithyDocumentSerde
}

type AttachManagedPolicyToPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachManagedPolicyToPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachManagedPolicyToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachManagedPolicyToPermissionSet{}, middleware.After)
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
	if err = addOpAttachManagedPolicyToPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachManagedPolicyToPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachManagedPolicyToPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "AttachManagedPolicyToPermissionSet",
	}
}
