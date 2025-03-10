// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes authorization to submit an AssociateVPCWithHostedZone request to
// associate a specified VPC with a hosted zone that was created by a different
// account. You must use the account that created the hosted zone to submit a
// DeleteVPCAssociationAuthorization request. Sending this request only prevents
// the Amazon Web Services account that created the VPC from associating the VPC
// with the Amazon Route 53 hosted zone in the future. If the VPC is already
// associated with the hosted zone, DeleteVPCAssociationAuthorization won't
// disassociate the VPC from the hosted zone. If you want to delete an existing
// association, use DisassociateVPCFromHostedZone.
func (c *Client) DeleteVPCAssociationAuthorization(ctx context.Context, params *DeleteVPCAssociationAuthorizationInput, optFns ...func(*Options)) (*DeleteVPCAssociationAuthorizationOutput, error) {
	if params == nil {
		params = &DeleteVPCAssociationAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVPCAssociationAuthorization", params, optFns, c.addOperationDeleteVPCAssociationAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVPCAssociationAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to remove
// authorization to associate a VPC that was created by one Amazon Web Services
// account with a hosted zone that was created with a different Amazon Web Services
// account.
type DeleteVPCAssociationAuthorizationInput struct {

	// When removing authorization to associate a VPC that was created by one Amazon
	// Web Services account with a hosted zone that was created with a different Amazon
	// Web Services account, the ID of the hosted zone.
	//
	// This member is required.
	HostedZoneId *string

	// When removing authorization to associate a VPC that was created by one Amazon
	// Web Services account with a hosted zone that was created with a different Amazon
	// Web Services account, a complex type that includes the ID and region of the VPC.
	//
	// This member is required.
	VPC *types.VPC

	noSmithyDocumentSerde
}

// Empty response for the request.
type DeleteVPCAssociationAuthorizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVPCAssociationAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteVPCAssociationAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteVPCAssociationAuthorization{}, middleware.After)
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
	if err = addOpDeleteVPCAssociationAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVPCAssociationAuthorization(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteVPCAssociationAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteVPCAssociationAuthorization",
	}
}
