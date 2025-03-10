// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Resolver endpoint. There are two types of Resolver endpoints, inbound
// and outbound:
//
// * An inbound Resolver endpoint forwards DNS queries to the DNS
// service for a VPC from your network.
//
// * An outbound Resolver endpoint forwards
// DNS queries from the DNS service for a VPC to your network.
func (c *Client) CreateResolverEndpoint(ctx context.Context, params *CreateResolverEndpointInput, optFns ...func(*Options)) (*CreateResolverEndpointOutput, error) {
	if params == nil {
		params = &CreateResolverEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResolverEndpoint", params, optFns, c.addOperationCreateResolverEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResolverEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResolverEndpointInput struct {

	// A unique string that identifies the request and that allows failed requests to
	// be retried without the risk of running the operation twice. CreatorRequestId can
	// be any unique string, for example, a date/time stamp.
	//
	// This member is required.
	CreatorRequestId *string

	// Specify the applicable value:
	//
	// * INBOUND: Resolver forwards DNS queries to the
	// DNS service for a VPC from your network
	//
	// * OUTBOUND: Resolver forwards DNS
	// queries from the DNS service for a VPC to your network
	//
	// This member is required.
	Direction types.ResolverEndpointDirection

	// The subnets and IP addresses in your VPC that DNS queries originate from (for
	// outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
	// The subnet ID uniquely identifies a VPC.
	//
	// This member is required.
	IpAddresses []types.IpAddressRequest

	// The ID of one or more security groups that you want to use to control access to
	// this VPC. The security group that you specify must include one or more inbound
	// rules (for inbound Resolver endpoints) or outbound rules (for outbound Resolver
	// endpoints). Inbound and outbound rules must allow TCP and UDP access. For
	// inbound access, open port 53. For outbound access, open the port that you're
	// using for DNS queries on your network.
	//
	// This member is required.
	SecurityGroupIds []string

	// A friendly name that lets you easily find a configuration in the Resolver
	// dashboard in the Route 53 console.
	Name *string

	// For the endpoint type you can choose either IPv4, IPv6. or dual-stack. A
	// dual-stack endpoint means that it will resolve via both IPv4 and IPv6. This
	// endpoint type is applied to all IP addresses.
	ResolverEndpointType types.ResolverEndpointType

	// A list of the tag keys and values that you want to associate with the endpoint.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateResolverEndpointOutput struct {

	// Information about the CreateResolverEndpoint request, including the status of
	// the request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResolverEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateResolverEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateResolverEndpoint{}, middleware.After)
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
	if err = addOpCreateResolverEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResolverEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResolverEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "CreateResolverEndpoint",
	}
}
