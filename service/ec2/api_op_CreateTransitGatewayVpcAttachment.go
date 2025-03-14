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

// Attaches the specified VPC to the specified transit gateway. If you attach a VPC
// with a CIDR range that overlaps the CIDR range of a VPC that is already
// attached, the new VPC CIDR range is not propagated to the default propagation
// route table. To send VPC traffic to an attached transit gateway, add a route to
// the VPC route table using CreateRoute.
func (c *Client) CreateTransitGatewayVpcAttachment(ctx context.Context, params *CreateTransitGatewayVpcAttachmentInput, optFns ...func(*Options)) (*CreateTransitGatewayVpcAttachmentOutput, error) {
	if params == nil {
		params = &CreateTransitGatewayVpcAttachmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransitGatewayVpcAttachment", params, optFns, c.addOperationCreateTransitGatewayVpcAttachmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransitGatewayVpcAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitGatewayVpcAttachmentInput struct {

	// The IDs of one or more subnets. You can specify only one subnet per Availability
	// Zone. You must specify at least one subnet, but we recommend that you specify
	// two subnets for better availability. The transit gateway uses one IP address
	// from each specified subnet.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the transit gateway.
	//
	// This member is required.
	TransitGatewayId *string

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The VPC attachment options.
	Options *types.CreateTransitGatewayVpcAttachmentRequestOptions

	// The tags to apply to the VPC attachment.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateTransitGatewayVpcAttachmentOutput struct {

	// Information about the VPC attachment.
	TransitGatewayVpcAttachment *types.TransitGatewayVpcAttachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTransitGatewayVpcAttachmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTransitGatewayVpcAttachment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTransitGatewayVpcAttachment{}, middleware.After)
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
	if err = addOpCreateTransitGatewayVpcAttachmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitGatewayVpcAttachment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTransitGatewayVpcAttachment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTransitGatewayVpcAttachment",
	}
}
