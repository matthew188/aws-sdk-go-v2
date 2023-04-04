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

// Updates a configured table association.
func (c *Client) UpdateConfiguredTableAssociation(ctx context.Context, params *UpdateConfiguredTableAssociationInput, optFns ...func(*Options)) (*UpdateConfiguredTableAssociationOutput, error) {
	if params == nil {
		params = &UpdateConfiguredTableAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfiguredTableAssociation", params, optFns, c.addOperationUpdateConfiguredTableAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfiguredTableAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConfiguredTableAssociationInput struct {

	// The unique identifier for the configured table association to update. Currently
	// accepts the configured table association ID.
	//
	// This member is required.
	ConfiguredTableAssociationIdentifier *string

	// The unique ID for the membership that the configured table association belongs
	// to.
	//
	// This member is required.
	MembershipIdentifier *string

	// A new description for the configured table association.
	Description *string

	// The service will assume this role to access catalog metadata and query the
	// table.
	RoleArn *string

	noSmithyDocumentSerde
}

type UpdateConfiguredTableAssociationOutput struct {

	// The entire updated configured table association.
	//
	// This member is required.
	ConfiguredTableAssociation *types.ConfiguredTableAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConfiguredTableAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConfiguredTableAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConfiguredTableAssociation{}, middleware.After)
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
	if err = addOpUpdateConfiguredTableAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfiguredTableAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConfiguredTableAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "UpdateConfiguredTableAssociation",
	}
}
