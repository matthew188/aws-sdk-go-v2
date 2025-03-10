// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/servicecatalogappregistry/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing attribute group with new details.
func (c *Client) UpdateAttributeGroup(ctx context.Context, params *UpdateAttributeGroupInput, optFns ...func(*Options)) (*UpdateAttributeGroupOutput, error) {
	if params == nil {
		params = &UpdateAttributeGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAttributeGroup", params, optFns, c.addOperationUpdateAttributeGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAttributeGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAttributeGroupInput struct {

	// The name, ID, or ARN of the attribute group that holds the attributes to
	// describe the application.
	//
	// This member is required.
	AttributeGroup *string

	// A JSON string in the form of nested key-value pairs that represent the
	// attributes in the group and describes an application and its components.
	Attributes *string

	// The description of the attribute group that the user provides.
	Description *string

	// Deprecated: The new name of the attribute group. The name must be unique in the
	// region in which you are updating the attribute group. Please do not use this
	// field as we have stopped supporting name updates.
	//
	// Deprecated: Name update for attribute group is deprecated.
	Name *string

	noSmithyDocumentSerde
}

type UpdateAttributeGroupOutput struct {

	// The updated information of the attribute group.
	AttributeGroup *types.AttributeGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAttributeGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAttributeGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAttributeGroup{}, middleware.After)
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
	if err = addOpUpdateAttributeGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAttributeGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAttributeGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "UpdateAttributeGroup",
	}
}
