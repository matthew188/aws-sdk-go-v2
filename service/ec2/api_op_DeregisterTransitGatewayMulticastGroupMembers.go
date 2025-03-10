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

// Deregisters the specified members (network interfaces) from the transit gateway
// multicast group.
func (c *Client) DeregisterTransitGatewayMulticastGroupMembers(ctx context.Context, params *DeregisterTransitGatewayMulticastGroupMembersInput, optFns ...func(*Options)) (*DeregisterTransitGatewayMulticastGroupMembersOutput, error) {
	if params == nil {
		params = &DeregisterTransitGatewayMulticastGroupMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterTransitGatewayMulticastGroupMembers", params, optFns, c.addOperationDeregisterTransitGatewayMulticastGroupMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterTransitGatewayMulticastGroupMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterTransitGatewayMulticastGroupMembersInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress *string

	// The IDs of the group members' network interfaces.
	NetworkInterfaceIds []string

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string

	noSmithyDocumentSerde
}

type DeregisterTransitGatewayMulticastGroupMembersOutput struct {

	// Information about the deregistered members.
	DeregisteredMulticastGroupMembers *types.TransitGatewayMulticastDeregisteredGroupMembers

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterTransitGatewayMulticastGroupMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeregisterTransitGatewayMulticastGroupMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeregisterTransitGatewayMulticastGroupMembers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterTransitGatewayMulticastGroupMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterTransitGatewayMulticastGroupMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeregisterTransitGatewayMulticastGroupMembers",
	}
}
