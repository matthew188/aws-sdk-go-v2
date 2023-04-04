// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches the specified managed policy to the specified IAM group. You use this
// operation to attach a managed policy to a group. To embed an inline policy in a
// group, use PutGroupPolicy. As a best practice, you can validate your IAM
// policies. To learn more, see Validating IAM policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_policy-validator.html)
// in the IAM User Guide. For more information about policies, see Managed policies
// and inline policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
func (c *Client) AttachGroupPolicy(ctx context.Context, params *AttachGroupPolicyInput, optFns ...func(*Options)) (*AttachGroupPolicyOutput, error) {
	if params == nil {
		params = &AttachGroupPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachGroupPolicy", params, optFns, c.addOperationAttachGroupPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachGroupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachGroupPolicyInput struct {

	// The name (friendly name, not ARN) of the group to attach the policy to. This
	// parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a
	// string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	GroupName *string

	// The Amazon Resource Name (ARN) of the IAM policy you want to attach. For more
	// information about ARNs, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference.
	//
	// This member is required.
	PolicyArn *string

	noSmithyDocumentSerde
}

type AttachGroupPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachGroupPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAttachGroupPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachGroupPolicy{}, middleware.After)
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
	if err = addOpAttachGroupPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachGroupPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachGroupPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "AttachGroupPolicy",
	}
}
