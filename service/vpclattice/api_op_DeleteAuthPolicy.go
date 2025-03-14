// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified auth policy. If an auth is set to Amazon Web Services_IAM
// and the auth policy is deleted, all requests will be denied by default. If you
// are trying to remove the auth policy completely, you must set the auth_type to
// NONE. If auth is enabled on the resource, but no auth policy is set, all
// requests will be denied.
func (c *Client) DeleteAuthPolicy(ctx context.Context, params *DeleteAuthPolicyInput, optFns ...func(*Options)) (*DeleteAuthPolicyOutput, error) {
	if params == nil {
		params = &DeleteAuthPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAuthPolicy", params, optFns, c.addOperationDeleteAuthPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAuthPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAuthPolicyInput struct {

	// The ID or Amazon Resource Name (ARN) of the resource.
	//
	// This member is required.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type DeleteAuthPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAuthPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAuthPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAuthPolicy{}, middleware.After)
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
	if err = addOpDeleteAuthPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAuthPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAuthPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "DeleteAuthPolicy",
	}
}
