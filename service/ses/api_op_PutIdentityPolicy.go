// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates a sending authorization policy for the specified identity (an
// email address or a domain). This API is for the identity owner only. If you have
// not verified the identity, this API will return an error. Sending authorization
// is a feature that enables an identity owner to authorize other senders to use
// its identities. For information about using sending authorization, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
// You can execute this operation no more than once per second.
func (c *Client) PutIdentityPolicy(ctx context.Context, params *PutIdentityPolicyInput, optFns ...func(*Options)) (*PutIdentityPolicyOutput, error) {
	if params == nil {
		params = &PutIdentityPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutIdentityPolicy", params, optFns, c.addOperationPutIdentityPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutIdentityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to add or update a sending authorization policy for an
// identity. Sending authorization is an Amazon SES feature that enables you to
// authorize other senders to use your identities. For information, see the Amazon
// SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
type PutIdentityPolicyInput struct {

	// The identity that the policy will apply to. You can specify an identity by using
	// its name or by using its Amazon Resource Name (ARN). Examples: user@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com. To
	// successfully call this API, you must own the identity.
	//
	// This member is required.
	Identity *string

	// The text of the policy in JSON format. The policy cannot exceed 4 KB. For
	// information about the syntax of sending authorization policies, see the Amazon
	// SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).
	//
	// This member is required.
	Policy *string

	// The name of the policy. The policy name cannot exceed 64 characters and can only
	// include alphanumeric characters, dashes, and underscores.
	//
	// This member is required.
	PolicyName *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type PutIdentityPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutIdentityPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutIdentityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutIdentityPolicy{}, middleware.After)
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
	if err = addOpPutIdentityPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutIdentityPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutIdentityPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutIdentityPolicy",
	}
}
