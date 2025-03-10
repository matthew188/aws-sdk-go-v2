// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a configured table association.
func (c *Client) GetConfiguredTableAssociation(ctx context.Context, params *GetConfiguredTableAssociationInput, optFns ...func(*Options)) (*GetConfiguredTableAssociationOutput, error) {
	if params == nil {
		params = &GetConfiguredTableAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConfiguredTableAssociation", params, optFns, c.addOperationGetConfiguredTableAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConfiguredTableAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConfiguredTableAssociationInput struct {

	// The unique ID for the configured table association to retrieve. Currently
	// accepts the configured table ID.
	//
	// This member is required.
	ConfiguredTableAssociationIdentifier *string

	// A unique identifier for the membership that the configured table association
	// belongs to. Currently accepts the membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	noSmithyDocumentSerde
}

type GetConfiguredTableAssociationOutput struct {

	// The entire configured table association object.
	//
	// This member is required.
	ConfiguredTableAssociation *types.ConfiguredTableAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConfiguredTableAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetConfiguredTableAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetConfiguredTableAssociation{}, middleware.After)
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
	if err = addOpGetConfiguredTableAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConfiguredTableAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConfiguredTableAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "GetConfiguredTableAssociation",
	}
}
