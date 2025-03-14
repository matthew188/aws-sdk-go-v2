// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the flag, ChangeProtection, which indicates whether it is possible to
// change the firewall. If the flag is set to TRUE, the firewall is protected from
// changes. This setting helps protect against accidentally changing a firewall
// that's in use.
func (c *Client) UpdateFirewallPolicyChangeProtection(ctx context.Context, params *UpdateFirewallPolicyChangeProtectionInput, optFns ...func(*Options)) (*UpdateFirewallPolicyChangeProtectionOutput, error) {
	if params == nil {
		params = &UpdateFirewallPolicyChangeProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFirewallPolicyChangeProtection", params, optFns, c.addOperationUpdateFirewallPolicyChangeProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFirewallPolicyChangeProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFirewallPolicyChangeProtectionInput struct {

	// A setting indicating whether the firewall is protected against a change to the
	// firewall policy association. Use this setting to protect against accidentally
	// modifying the firewall policy for a firewall that is in use. When you create a
	// firewall, the operation initializes this setting to TRUE.
	//
	// This member is required.
	FirewallPolicyChangeProtection bool

	// The Amazon Resource Name (ARN) of the firewall. You must specify the ARN or the
	// name, and you can specify both.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it. You must specify the ARN or the name, and you can specify
	// both.
	FirewallName *string

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request. To make an
	// unconditional change to the firewall, omit the token in your update request.
	// Without the token, Network Firewall performs your updates regardless of whether
	// the firewall has changed since you last retrieved it. To make a conditional
	// change to the firewall, provide the token in your update request. Network
	// Firewall uses the token to ensure that the firewall hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an
	// InvalidTokenException. If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	noSmithyDocumentSerde
}

type UpdateFirewallPolicyChangeProtectionOutput struct {

	// The Amazon Resource Name (ARN) of the firewall.
	FirewallArn *string

	// The descriptive name of the firewall. You can't change the name of a firewall
	// after you create it.
	FirewallName *string

	// A setting indicating whether the firewall is protected against a change to the
	// firewall policy association. Use this setting to protect against accidentally
	// modifying the firewall policy for a firewall that is in use. When you create a
	// firewall, the operation initializes this setting to TRUE.
	FirewallPolicyChangeProtection bool

	// An optional token that you can use for optimistic locking. Network Firewall
	// returns a token to your requests that access the firewall. The token marks the
	// state of the firewall resource at the time of the request. To make an
	// unconditional change to the firewall, omit the token in your update request.
	// Without the token, Network Firewall performs your updates regardless of whether
	// the firewall has changed since you last retrieved it. To make a conditional
	// change to the firewall, provide the token in your update request. Network
	// Firewall uses the token to ensure that the firewall hasn't changed since you
	// last retrieved it. If it has changed, the operation fails with an
	// InvalidTokenException. If this happens, retrieve the firewall again to get a
	// current copy of it with a new token. Reapply your changes as needed, then try
	// the operation again using the new token.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFirewallPolicyChangeProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateFirewallPolicyChangeProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateFirewallPolicyChangeProtection{}, middleware.After)
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
	if err = addOpUpdateFirewallPolicyChangeProtectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFirewallPolicyChangeProtection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFirewallPolicyChangeProtection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "UpdateFirewallPolicyChangeProtection",
	}
}
