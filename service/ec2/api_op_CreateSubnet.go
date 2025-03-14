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

// Creates a subnet in the specified VPC. For an IPv4 only subnet, specify an IPv4
// CIDR block. If the VPC has an IPv6 CIDR block, you can create an IPv6 only
// subnet or a dual stack subnet instead. For an IPv6 only subnet, specify an IPv6
// CIDR block. For a dual stack subnet, specify both an IPv4 CIDR block and an IPv6
// CIDR block. A subnet CIDR block must not overlap the CIDR block of an existing
// subnet in the VPC. After you create a subnet, you can't change its CIDR block.
// The allowed size for an IPv4 subnet is between a /28 netmask (16 IP addresses)
// and a /16 netmask (65,536 IP addresses). Amazon Web Services reserves both the
// first four and the last IPv4 address in each subnet's CIDR block. They're not
// available for your use. If you've associated an IPv6 CIDR block with your VPC,
// you can associate an IPv6 CIDR block with a subnet when you create it. The
// allowed block size for an IPv6 subnet is a /64 netmask. If you add more than one
// subnet to a VPC, they're set up in a star topology with a logical router in the
// middle. When you stop an instance in a subnet, it retains its private IPv4
// address. It's therefore possible to have a subnet with no running instances
// (they're all stopped), but no remaining IP addresses available. For more
// information, see Subnets
// (https://docs.aws.amazon.com/vpc/latest/userguide/configure-subnets.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateSubnet(ctx context.Context, params *CreateSubnetInput, optFns ...func(*Options)) (*CreateSubnetOutput, error) {
	if params == nil {
		params = &CreateSubnetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubnet", params, optFns, c.addOperationCreateSubnetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubnetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubnetInput struct {

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// The Availability Zone or Local Zone for the subnet. Default: Amazon Web Services
	// selects one for you. If you create more than one subnet in your VPC, we do not
	// necessarily select a different zone for each subnet. To create a subnet in a
	// Local Zone, set this value to the Local Zone ID, for example us-west-2-lax-1a.
	// For information about the Regions that support Local Zones, see Available
	// Regions
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions)
	// in the Amazon Elastic Compute Cloud User Guide. To create a subnet in an
	// Outpost, set this value to the Availability Zone for the Outpost and specify the
	// Outpost ARN.
	AvailabilityZone *string

	// The AZ ID or the Local Zone ID of the subnet.
	AvailabilityZoneId *string

	// The IPv4 network range for the subnet, in CIDR notation. For example,
	// 10.0.0.0/24. We modify the specified CIDR block to its canonical form; for
	// example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18. This
	// parameter is not supported for an IPv6 only subnet.
	CidrBlock *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The IPv6 network range for the subnet, in CIDR notation. The subnet size must
	// use a /64 prefix length. This parameter is required for an IPv6 only subnet.
	Ipv6CidrBlock *string

	// Indicates whether to create an IPv6 only subnet.
	Ipv6Native *bool

	// The Amazon Resource Name (ARN) of the Outpost. If you specify an Outpost ARN,
	// you must also specify the Availability Zone of the Outpost subnet.
	OutpostArn *string

	// The tags to assign to the subnet.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateSubnetOutput struct {

	// Information about the subnet.
	Subnet *types.Subnet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubnetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateSubnet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateSubnet{}, middleware.After)
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
	if err = addOpCreateSubnetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubnet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSubnet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateSubnet",
	}
}
