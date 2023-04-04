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

// Creates a VPC with the specified CIDR blocks. For more information, see VPC CIDR
// blocks
// (https://docs.aws.amazon.com/vpc/latest/userguide/configure-your-vpc.html#vpc-cidr-blocks)
// in the Amazon Virtual Private Cloud User Guide. You can optionally request an
// IPv6 CIDR block for the VPC. You can request an Amazon-provided IPv6 CIDR block
// from Amazon's pool of IPv6 addresses, or an IPv6 CIDR block from an IPv6 address
// pool that you provisioned through bring your own IP addresses (BYOIP
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html)). By
// default, each instance that you launch in the VPC has the default DHCP options,
// which include only a default DNS server that we provide (AmazonProvidedDNS). For
// more information, see DHCP option sets
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the
// Amazon Virtual Private Cloud User Guide. You can specify the instance tenancy
// value for the VPC when you create it. You can't change this value for the VPC
// after you create it. For more information, see Dedicated Instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CreateVpc(ctx context.Context, params *CreateVpcInput, optFns ...func(*Options)) (*CreateVpcOutput, error) {
	if params == nil {
		params = &CreateVpcInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpc", params, optFns, c.addOperationCreateVpcMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVpcInput struct {

	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the
	// VPC. You cannot specify the range of IP addresses, or the size of the CIDR
	// block.
	AmazonProvidedIpv6CidrBlock *bool

	// The IPv4 network range for the VPC, in CIDR notation. For example, 10.0.0.0/16.
	// We modify the specified CIDR block to its canonical form; for example, if you
	// specify 100.68.0.18/18, we modify it to 100.68.0.0/18.
	CidrBlock *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tenancy options for instances launched into the VPC. For default, instances
	// are launched with shared tenancy by default. You can launch instances with any
	// tenancy into a shared tenancy VPC. For dedicated, instances are launched as
	// dedicated tenancy instances by default. You can only launch instances with a
	// tenancy of dedicated or host into a dedicated tenancy VPC. Important: The host
	// value cannot be used with this parameter. Use the default or dedicated values
	// only. Default: default
	InstanceTenancy types.Tenancy

	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For
	// more information, see What is IPAM?
	// (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon
	// VPC IPAM User Guide.
	Ipv4IpamPoolId *string

	// The netmask length of the IPv4 CIDR you want to allocate to this VPC from an
	// Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see
	// What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	Ipv4NetmaskLength *int32

	// The IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool
	// in the request. To let Amazon choose the IPv6 CIDR block for you, omit this
	// parameter.
	Ipv6CidrBlock *string

	// The name of the location from which we advertise the IPV6 CIDR block. Use this
	// parameter to limit the address to this location. You must set
	// AmazonProvidedIpv6CidrBlock to true to use this parameter.
	Ipv6CidrBlockNetworkBorderGroup *string

	// The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6
	// CIDR. IPAM is a VPC feature that you can use to automate your IP address
	// management workflows including assigning, tracking, troubleshooting, and
	// auditing IP addresses across Amazon Web Services Regions and accounts throughout
	// your Amazon Web Services Organization. For more information, see What is IPAM?
	// (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon
	// VPC IPAM User Guide.
	Ipv6IpamPoolId *string

	// The netmask length of the IPv6 CIDR you want to allocate to this VPC from an
	// Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see
	// What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html)
	// in the Amazon VPC IPAM User Guide.
	Ipv6NetmaskLength *int32

	// The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.
	Ipv6Pool *string

	// The tags to assign to the VPC.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateVpcOutput struct {

	// Information about the VPC.
	Vpc *types.Vpc

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVpcMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpc{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpc{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpc(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpc(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVpc",
	}
}
