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

// Provides information to Amazon Web Services about your customer gateway device.
// The customer gateway device is the appliance at your end of the VPN connection.
// You must provide the IP address of the customer gateway device’s external
// interface. The IP address must be static and can be behind a device performing
// network address translation (NAT). For devices that use Border Gateway Protocol
// (BGP), you can also provide the device's BGP Autonomous System Number (ASN). You
// can use an existing ASN assigned to your network. If you don't have an ASN
// already, you can use a private ASN. For more information, see Customer gateway
// options for your Site-to-Site VPN connection
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/cgw-options.html) in the Amazon
// Web Services Site-to-Site VPN User Guide. To create more than one customer
// gateway with the same VPN type, IP address, and BGP ASN, specify a unique device
// name for each customer gateway. An identical request returns information about
// the existing customer gateway; it doesn't create a new customer gateway.
func (c *Client) CreateCustomerGateway(ctx context.Context, params *CreateCustomerGatewayInput, optFns ...func(*Options)) (*CreateCustomerGatewayOutput, error) {
	if params == nil {
		params = &CreateCustomerGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomerGateway", params, optFns, c.addOperationCreateCustomerGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomerGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateCustomerGateway.
type CreateCustomerGatewayInput struct {

	// The type of VPN connection that this customer gateway supports (ipsec.1).
	//
	// This member is required.
	Type types.GatewayType

	// For devices that support BGP, the customer gateway's BGP ASN. Default: 65000
	BgpAsn *int32

	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	CertificateArn *string

	// A name for the customer gateway device. Length Constraints: Up to 255
	// characters.
	DeviceName *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// IPv4 address for the customer gateway device's outside interface. The address
	// must be static.
	IpAddress *string

	// This member has been deprecated. The Internet-routable IP address for the
	// customer gateway's outside interface. The address must be static.
	PublicIp *string

	// The tags to apply to the customer gateway.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

// Contains the output of CreateCustomerGateway.
type CreateCustomerGatewayOutput struct {

	// Information about the customer gateway.
	CustomerGateway *types.CustomerGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomerGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateCustomerGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateCustomerGateway{}, middleware.After)
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
	if err = addOpCreateCustomerGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomerGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCustomerGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateCustomerGateway",
	}
}
