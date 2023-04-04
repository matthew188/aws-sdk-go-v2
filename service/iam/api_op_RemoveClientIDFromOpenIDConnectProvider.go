// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified client ID (also known as audience) from the list of client
// IDs registered for the specified IAM OpenID Connect (OIDC) provider resource
// object. This operation is idempotent; it does not fail or return an error if you
// try to remove a client ID that does not exist.
func (c *Client) RemoveClientIDFromOpenIDConnectProvider(ctx context.Context, params *RemoveClientIDFromOpenIDConnectProviderInput, optFns ...func(*Options)) (*RemoveClientIDFromOpenIDConnectProviderOutput, error) {
	if params == nil {
		params = &RemoveClientIDFromOpenIDConnectProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveClientIDFromOpenIDConnectProvider", params, optFns, c.addOperationRemoveClientIDFromOpenIDConnectProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveClientIDFromOpenIDConnectProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveClientIDFromOpenIDConnectProviderInput struct {

	// The client ID (also known as audience) to remove from the IAM OIDC provider
	// resource. For more information about client IDs, see
	// CreateOpenIDConnectProvider.
	//
	// This member is required.
	ClientID *string

	// The Amazon Resource Name (ARN) of the IAM OIDC provider resource to remove the
	// client ID from. You can get a list of OIDC provider ARNs by using the
	// ListOpenIDConnectProviders operation. For more information about ARNs, see
	// Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference.
	//
	// This member is required.
	OpenIDConnectProviderArn *string

	noSmithyDocumentSerde
}

type RemoveClientIDFromOpenIDConnectProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveClientIDFromOpenIDConnectProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemoveClientIDFromOpenIDConnectProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveClientIDFromOpenIDConnectProvider{}, middleware.After)
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
	if err = addOpRemoveClientIDFromOpenIDConnectProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveClientIDFromOpenIDConnectProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveClientIDFromOpenIDConnectProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "RemoveClientIDFromOpenIDConnectProvider",
	}
}
