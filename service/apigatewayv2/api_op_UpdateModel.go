// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a Model.
func (c *Client) UpdateModel(ctx context.Context, params *UpdateModelInput, optFns ...func(*Options)) (*UpdateModelOutput, error) {
	if params == nil {
		params = &UpdateModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateModel", params, optFns, c.addOperationUpdateModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a Model.
type UpdateModelInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The model ID.
	//
	// This member is required.
	ModelId *string

	// The content-type for the model, for example, "application/json".
	ContentType *string

	// The description of the model.
	Description *string

	// The name of the model.
	Name *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 model.
	Schema *string

	noSmithyDocumentSerde
}

type UpdateModelOutput struct {

	// The content-type for the model, for example, "application/json".
	ContentType *string

	// The description of the model.
	Description *string

	// The model identifier.
	ModelId *string

	// The name of the model. Must be alphanumeric.
	Name *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 model.
	Schema *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateModel{}, middleware.After)
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
	if err = addOpUpdateModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateModel",
	}
}
