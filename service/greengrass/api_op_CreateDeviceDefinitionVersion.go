// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a version of a device definition that has already been defined.
func (c *Client) CreateDeviceDefinitionVersion(ctx context.Context, params *CreateDeviceDefinitionVersionInput, optFns ...func(*Options)) (*CreateDeviceDefinitionVersionOutput, error) {
	if params == nil {
		params = &CreateDeviceDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeviceDefinitionVersion", params, optFns, c.addOperationCreateDeviceDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeviceDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeviceDefinitionVersionInput struct {

	// The ID of the device definition.
	//
	// This member is required.
	DeviceDefinitionId *string

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// A list of devices in the definition version.
	Devices []types.Device

	noSmithyDocumentSerde
}

type CreateDeviceDefinitionVersionOutput struct {

	// The ARN of the version.
	Arn *string

	// The time, in milliseconds since the epoch, when the version was created.
	CreationTimestamp *string

	// The ID of the parent definition that the version is associated with.
	Id *string

	// The ID of the version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeviceDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeviceDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeviceDefinitionVersion{}, middleware.After)
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
	if err = addOpCreateDeviceDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeviceDefinitionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDeviceDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateDeviceDefinitionVersion",
	}
}
