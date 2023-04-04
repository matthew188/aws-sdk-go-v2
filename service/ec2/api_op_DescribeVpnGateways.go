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

// Describes one or more of your virtual private gateways. For more information,
// see Amazon Web Services Site-to-Site VPN
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the Amazon Web
// Services Site-to-Site VPN User Guide.
func (c *Client) DescribeVpnGateways(ctx context.Context, params *DescribeVpnGatewaysInput, optFns ...func(*Options)) (*DescribeVpnGatewaysOutput, error) {
	if params == nil {
		params = &DescribeVpnGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpnGateways", params, optFns, c.addOperationDescribeVpnGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpnGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeVpnGateways.
type DescribeVpnGatewaysInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * amazon-side-asn - The Autonomous System Number (ASN) for
	// the Amazon side of the gateway.
	//
	// * attachment.state - The current state of the
	// attachment between the gateway and the VPC (attaching | attached | detaching |
	// detached).
	//
	// * attachment.vpc-id - The ID of an attached VPC.
	//
	// *
	// availability-zone - The Availability Zone for the virtual private gateway (if
	// applicable).
	//
	// * state - The state of the virtual private gateway (pending |
	// available | deleting | deleted).
	//
	// * tag: - The key/value combination of a tag
	// assigned to the resource. Use the tag key in the filter name and the tag value
	// as the filter value. For example, to find all resources that have a tag with the
	// key Owner and the value TeamA, specify tag:Owner for the filter name and TeamA
	// for the filter value.
	//
	// * tag-key - The key of a tag assigned to the resource.
	// Use this filter to find all resources assigned a tag with a specific key,
	// regardless of the tag value.
	//
	// * type - The type of virtual private gateway.
	// Currently the only supported type is ipsec.1.
	//
	// * vpn-gateway-id - The ID of the
	// virtual private gateway.
	Filters []types.Filter

	// One or more virtual private gateway IDs. Default: Describes all your virtual
	// private gateways.
	VpnGatewayIds []string

	noSmithyDocumentSerde
}

// Contains the output of DescribeVpnGateways.
type DescribeVpnGatewaysOutput struct {

	// Information about one or more virtual private gateways.
	VpnGateways []types.VpnGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpnGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpnGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpnGateways{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpnGateways(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpnGateways(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpnGateways",
	}
}
