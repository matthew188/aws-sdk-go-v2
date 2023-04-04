// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified Client VPN endpoint. Modifying the DNS server resets
// existing client connections.
func (c *Client) ModifyClientVpnEndpoint(ctx context.Context, params *ModifyClientVpnEndpointInput, optFns ...func(*Options)) (*ModifyClientVpnEndpointOutput, error) {
	if params == nil {
		params = &ModifyClientVpnEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClientVpnEndpoint", params, optFns, c.addOperationModifyClientVpnEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClientVpnEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClientVpnEndpointInput struct {

	// The ID of the Client VPN endpoint to modify.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The options for managing connection authorization for new client connections.
	ClientConnectOptions *types.ClientConnectOptions

	// Options for enabling a customizable text banner that will be displayed on Amazon
	// Web Services provided clients when a VPN session is established.
	ClientLoginBannerOptions *types.ClientLoginBannerOptions

	// Information about the client connection logging options. If you enable client
	// connection logging, data about client connections is sent to a Cloudwatch Logs
	// log stream. The following information is logged:
	//
	// * Client connection
	// requests
	//
	// * Client connection results (successful and unsuccessful)
	//
	// * Reasons
	// for unsuccessful client connection requests
	//
	// * Client connection termination
	// time
	ConnectionLogOptions *types.ConnectionLogOptions

	// A brief description of the Client VPN endpoint.
	Description *string

	// Information about the DNS servers to be used by Client VPN connections. A Client
	// VPN endpoint can have up to two DNS servers.
	DnsServers *types.DnsServersOptionsModifyStructure

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The IDs of one or more security groups to apply to the target network.
	SecurityGroupIds []string

	// Specify whether to enable the self-service portal for the Client VPN endpoint.
	SelfServicePortal types.SelfServicePortal

	// The ARN of the server certificate to be used. The server certificate must be
	// provisioned in Certificate Manager (ACM).
	ServerCertificateArn *string

	// The maximum VPN session duration time in hours. Valid values: 8 | 10 | 12 | 24
	// Default value: 24
	SessionTimeoutHours *int32

	// Indicates whether the VPN is split-tunnel. For information about split-tunnel
	// VPN endpoints, see Split-tunnel Client VPN endpoint
	// (https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/split-tunnel-vpn.html)
	// in the Client VPN Administrator Guide.
	SplitTunnel *bool

	// The ID of the VPC to associate with the Client VPN endpoint.
	VpcId *string

	// The port number to assign to the Client VPN endpoint for TCP and UDP traffic.
	// Valid Values: 443 | 1194 Default Value: 443
	VpnPort *int32

	noSmithyDocumentSerde
}

type ModifyClientVpnEndpointOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyClientVpnEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyClientVpnEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyClientVpnEndpoint{}, middleware.After)
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
	if err = addOpModifyClientVpnEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClientVpnEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyClientVpnEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyClientVpnEndpoint",
	}
}
