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

// Authorizes the Amazon Web Services account that created a specified VPC to
// submit an AssociateVPCWithHostedZone request to associate the VPC with a
// specified hosted zone that was created by a different account. To submit a
// CreateVPCAssociationAuthorization request, you must use the account that created
// the hosted zone. After you authorize the association, use the account that
// created the VPC to submit an AssociateVPCWithHostedZone request. If you want to
// associate multiple VPCs that you created by using one account with a hosted zone
// that you created by using a different account, you must submit one authorization
// request for each VPC.
func (c *Client) CreateVPCAssociationAuthorization(ctx context.Context, params *CreateVPCAssociationAuthorizationInput, optFns ...func(*Options)) (*CreateVPCAssociationAuthorizationOutput, error) {
	if params == nil {
		params = &CreateVPCAssociationAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVPCAssociationAuthorization", params, optFns, c.addOperationCreateVPCAssociationAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVPCAssociationAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to authorize
// associating a VPC with your private hosted zone. Authorization is only required
// when a private hosted zone and a VPC were created by using different accounts.
type CreateVPCAssociationAuthorizationInput struct {

	// The ID of the private hosted zone that you want to authorize associating a VPC
	// with.
	//
	// This member is required.
	HostedZoneId *string

	// A complex type that contains the VPC ID and region for the VPC that you want to
	// authorize associating with your hosted zone.
	//
	// This member is required.
	VPC *types.VPC

	noSmithyDocumentSerde
}

// A complex type that contains the response information from a
// CreateVPCAssociationAuthorization request.
type CreateVPCAssociationAuthorizationOutput struct {

	// The ID of the hosted zone that you authorized associating a VPC with.
	//
	// This member is required.
	HostedZoneId *string

	// The VPC that you authorized associating with a hosted zone.
	//
	// This member is required.
	VPC *types.VPC

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVPCAssociationAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateVPCAssociationAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateVPCAssociationAuthorization{}, middleware.After)
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
	if err = addOpCreateVPCAssociationAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVPCAssociationAuthorization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVPCAssociationAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateVPCAssociationAuthorization",
	}
}
