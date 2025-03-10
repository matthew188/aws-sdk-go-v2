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

// Creates a set of DHCP options for your VPC. After creating the set, you must
// associate it with the VPC, causing all existing and new instances that you
// launch in the VPC to use this set of DHCP options. The following are the
// individual DHCP options you can specify. For more information about the options,
// see RFC 2132 (http://www.ietf.org/rfc/rfc2132.txt).
//
// * domain-name-servers - The
// IP addresses of up to four domain name servers, or AmazonProvidedDNS. The
// default DHCP option set specifies AmazonProvidedDNS. If specifying more than one
// domain name server, specify the IP addresses in a single parameter, separated by
// commas. To have your instance receive a custom DNS hostname as specified in
// domain-name, you must set domain-name-servers to a custom DNS server.
//
// *
// domain-name - If you're using AmazonProvidedDNS in us-east-1, specify
// ec2.internal. If you're using AmazonProvidedDNS in another Region, specify
// region.compute.internal (for example, ap-northeast-1.compute.internal).
// Otherwise, specify a domain name (for example, ExampleCompany.com). This value
// is used to complete unqualified DNS hostnames. Important: Some Linux operating
// systems accept multiple domain names separated by spaces. However, Windows and
// other Linux operating systems treat the value as a single domain, which results
// in unexpected behavior. If your DHCP options set is associated with a VPC that
// has instances with multiple operating systems, specify only one domain name.
//
// *
// ntp-servers - The IP addresses of up to four Network Time Protocol (NTP)
// servers.
//
// * netbios-name-servers - The IP addresses of up to four NetBIOS name
// servers.
//
// * netbios-node-type - The NetBIOS node type (1, 2, 4, or 8). We
// recommend that you specify 2 (broadcast and multicast are not currently
// supported). For more information about these node types, see RFC 2132
// (http://www.ietf.org/rfc/rfc2132.txt).
//
// Your VPC automatically starts out with a
// set of DHCP options that includes only a DNS server that we provide
// (AmazonProvidedDNS). If you create a set of options, and if your VPC has an
// internet gateway, make sure to set the domain-name-servers option either to
// AmazonProvidedDNS or to a domain name server of your choice. For more
// information, see DHCP options sets
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateDhcpOptions(ctx context.Context, params *CreateDhcpOptionsInput, optFns ...func(*Options)) (*CreateDhcpOptionsOutput, error) {
	if params == nil {
		params = &CreateDhcpOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDhcpOptions", params, optFns, c.addOperationCreateDhcpOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDhcpOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDhcpOptionsInput struct {

	// A DHCP configuration option.
	//
	// This member is required.
	DhcpConfigurations []types.NewDhcpConfiguration

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags to assign to the DHCP option.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateDhcpOptionsOutput struct {

	// A set of DHCP options.
	DhcpOptions *types.DhcpOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDhcpOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateDhcpOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateDhcpOptions{}, middleware.After)
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
	if err = addOpCreateDhcpOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDhcpOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDhcpOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateDhcpOptions",
	}
}
