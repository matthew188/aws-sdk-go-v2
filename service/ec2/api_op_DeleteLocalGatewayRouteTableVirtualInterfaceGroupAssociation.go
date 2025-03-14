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

// Deletes a local gateway route table virtual interface group association.
func (c *Client) DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(ctx context.Context, params *DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput, optFns ...func(*Options)) (*DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput, error) {
	if params == nil {
		params = &DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation", params, optFns, c.addOperationDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationInput struct {

	// The ID of the local gateway route table virtual interface group association.
	//
	// This member is required.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociationId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationOutput struct {

	// Information about the association.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociation *types.LocalGatewayRouteTableVirtualInterfaceGroupAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation{}, middleware.After)
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
	if err = addOpDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation",
	}
}
