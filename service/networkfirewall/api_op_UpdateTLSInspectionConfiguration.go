// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the TLS inspection configuration settings for the specified TLS
// inspection configuration. You use a TLS inspection configuration by reference in
// one or more firewall policies. When you modify a TLS inspection configuration,
// you modify all firewall policies that use the TLS inspection configuration. To
// update a TLS inspection configuration, first call
// DescribeTLSInspectionConfiguration to retrieve the current
// TLSInspectionConfiguration object, update the object as needed, and then provide
// the updated object to this call.
func (c *Client) UpdateTLSInspectionConfiguration(ctx context.Context, params *UpdateTLSInspectionConfigurationInput, optFns ...func(*Options)) (*UpdateTLSInspectionConfigurationOutput, error) {
	if params == nil {
		params = &UpdateTLSInspectionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTLSInspectionConfiguration", params, optFns, c.addOperationUpdateTLSInspectionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTLSInspectionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTLSInspectionConfigurationInput struct {

	// The object that defines a TLS inspection configuration. This, along with
	// TLSInspectionConfigurationResponse, define the TLS inspection configuration. You
	// can retrieve all objects for a TLS inspection configuration by calling
	// DescribeTLSInspectionConfiguration. Network Firewall uses a TLS inspection
	// configuration to decrypt traffic. Network Firewall re-encrypts the traffic
	// before sending it to its destination. To use a TLS inspection configuration, you
	// add it to a Network Firewall firewall policy, then you apply the firewall policy
	// to a firewall. Network Firewall acts as a proxy service to decrypt and inspect
	// inbound traffic. You can reference a TLS inspection configuration from more than
	// one firewall policy, and you can use a firewall policy in more than one
	// firewall. For more information about using TLS inspection configurations, see
	// Decrypting SSL/TLS traffic with TLS inspection configurations
	// (https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html)
	// in the Network Firewall Developer Guide.
	//
	// This member is required.
	TLSInspectionConfiguration *types.TLSInspectionConfiguration

	// A token used for optimistic locking. Network Firewall returns a token to your
	// requests that access the TLS inspection configuration. The token marks the state
	// of the TLS inspection configuration resource at the time of the request. To make
	// changes to the TLS inspection configuration, you provide the token in your
	// request. Network Firewall uses the token to ensure that the TLS inspection
	// configuration hasn't changed since you last retrieved it. If it has changed, the
	// operation fails with an InvalidTokenException. If this happens, retrieve the TLS
	// inspection configuration again to get a current copy of it with a current token.
	// Reapply your changes as needed, then try the operation again using the new
	// token.
	//
	// This member is required.
	UpdateToken *string

	// A description of the TLS inspection configuration.
	Description *string

	// A complex type that contains the Amazon Web Services KMS encryption
	// configuration settings for your TLS inspection configuration.
	EncryptionConfiguration *types.EncryptionConfiguration

	// The Amazon Resource Name (ARN) of the TLS inspection configuration.
	TLSInspectionConfigurationArn *string

	// The descriptive name of the TLS inspection configuration. You can't change the
	// name of a TLS inspection configuration after you create it.
	TLSInspectionConfigurationName *string

	noSmithyDocumentSerde
}

type UpdateTLSInspectionConfigurationOutput struct {

	// The high-level properties of a TLS inspection configuration. This, along with
	// the TLSInspectionConfiguration, define the TLS inspection configuration. You can
	// retrieve all objects for a TLS inspection configuration by calling
	// DescribeTLSInspectionConfiguration.
	//
	// This member is required.
	TLSInspectionConfigurationResponse *types.TLSInspectionConfigurationResponse

	// A token used for optimistic locking. Network Firewall returns a token to your
	// requests that access the TLS inspection configuration. The token marks the state
	// of the TLS inspection configuration resource at the time of the request. To make
	// changes to the TLS inspection configuration, you provide the token in your
	// request. Network Firewall uses the token to ensure that the TLS inspection
	// configuration hasn't changed since you last retrieved it. If it has changed, the
	// operation fails with an InvalidTokenException. If this happens, retrieve the TLS
	// inspection configuration again to get a current copy of it with a current token.
	// Reapply your changes as needed, then try the operation again using the new
	// token.
	//
	// This member is required.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTLSInspectionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTLSInspectionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTLSInspectionConfiguration{}, middleware.After)
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
	if err = addOpUpdateTLSInspectionConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTLSInspectionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTLSInspectionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "UpdateTLSInspectionConfiguration",
	}
}
