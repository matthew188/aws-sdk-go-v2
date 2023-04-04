// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the association between a specified service and the specific service
// network. This request will fail if an association is still in progress.
func (c *Client) DeleteServiceNetworkServiceAssociation(ctx context.Context, params *DeleteServiceNetworkServiceAssociationInput, optFns ...func(*Options)) (*DeleteServiceNetworkServiceAssociationOutput, error) {
	if params == nil {
		params = &DeleteServiceNetworkServiceAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteServiceNetworkServiceAssociation", params, optFns, c.addOperationDeleteServiceNetworkServiceAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteServiceNetworkServiceAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteServiceNetworkServiceAssociationInput struct {

	// The ID or Amazon Resource Name (ARN) of the association.
	//
	// This member is required.
	ServiceNetworkServiceAssociationIdentifier *string

	noSmithyDocumentSerde
}

type DeleteServiceNetworkServiceAssociationOutput struct {

	// The Amazon Resource Name (ARN) of the association.
	Arn *string

	// The ID of the association.
	Id *string

	// The operation's status. You can retry the operation if the status is
	// DELETE_FAILED. However, if you retry it when the status is DELETE_IN_PROGRESS,
	// there is no change in the status.
	Status types.ServiceNetworkServiceAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteServiceNetworkServiceAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteServiceNetworkServiceAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteServiceNetworkServiceAssociation{}, middleware.After)
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
	if err = addOpDeleteServiceNetworkServiceAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteServiceNetworkServiceAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteServiceNetworkServiceAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "DeleteServiceNetworkServiceAssociation",
	}
}
