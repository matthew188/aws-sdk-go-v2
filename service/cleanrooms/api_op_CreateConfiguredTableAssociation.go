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

// Creates a configured table association. A configured table association links a
// configured table with a collaboration.
func (c *Client) CreateConfiguredTableAssociation(ctx context.Context, params *CreateConfiguredTableAssociationInput, optFns ...func(*Options)) (*CreateConfiguredTableAssociationOutput, error) {
	if params == nil {
		params = &CreateConfiguredTableAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfiguredTableAssociation", params, optFns, c.addOperationCreateConfiguredTableAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfiguredTableAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConfiguredTableAssociationInput struct {

	// A unique identifier for the configured table to be associated to. Currently
	// accepts a configured table ID.
	//
	// This member is required.
	ConfiguredTableIdentifier *string

	// A unique identifier for one of your memberships for a collaboration. The
	// configured table is associated to the collaboration that this membership belongs
	// to. Currently accepts a membership ID.
	//
	// This member is required.
	MembershipIdentifier *string

	// The name of the configured table association. This name is used to query the
	// underlying configured table.
	//
	// This member is required.
	Name *string

	// The service will assume this role to access catalog metadata and query the
	// table.
	//
	// This member is required.
	RoleArn *string

	// A description for the configured table association.
	Description *string

	// An optional label that you can assign to a resource when you create it. Each tag
	// consists of a key and an optional value, both of which you define. When you use
	// tagging, you can also use tag-based access control in IAM policies to control
	// access to this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateConfiguredTableAssociationOutput struct {

	// The entire configured table association object.
	//
	// This member is required.
	ConfiguredTableAssociation *types.ConfiguredTableAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfiguredTableAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfiguredTableAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfiguredTableAssociation{}, middleware.After)
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
	if err = addOpCreateConfiguredTableAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfiguredTableAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfiguredTableAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "CreateConfiguredTableAssociation",
	}
}
