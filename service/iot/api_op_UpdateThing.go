// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the data for a thing. Requires permission to access the UpdateThing
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) UpdateThing(ctx context.Context, params *UpdateThingInput, optFns ...func(*Options)) (*UpdateThingOutput, error) {
	if params == nil {
		params = &UpdateThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateThing", params, optFns, c.addOperationUpdateThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the UpdateThing operation.
type UpdateThingInput struct {

	// The name of the thing to update. You can't change a thing's name. To change a
	// thing's name, you must create a new thing, give it the new name, and then delete
	// the old thing.
	//
	// This member is required.
	ThingName *string

	// A list of thing attributes, a JSON string containing name-value pairs. For
	// example: {\"attributes\":{\"name1\":\"value2\"}} This data is used to add new
	// attributes or update existing attributes.
	AttributePayload *types.AttributePayload

	// The expected version of the thing record in the registry. If the version of the
	// record in the registry does not match the expected version specified in the
	// request, the UpdateThing request is rejected with a VersionConflictException.
	ExpectedVersion *int64

	// Remove a thing type association. If true, the association is removed.
	RemoveThingType bool

	// The name of the thing type.
	ThingTypeName *string

	noSmithyDocumentSerde
}

// The output from the UpdateThing operation.
type UpdateThingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThing{}, middleware.After)
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
	if err = addOpUpdateThingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThing(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateThing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateThing",
	}
}
