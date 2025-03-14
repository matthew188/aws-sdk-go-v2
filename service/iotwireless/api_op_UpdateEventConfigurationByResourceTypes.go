// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the event configuration based on resource types.
func (c *Client) UpdateEventConfigurationByResourceTypes(ctx context.Context, params *UpdateEventConfigurationByResourceTypesInput, optFns ...func(*Options)) (*UpdateEventConfigurationByResourceTypesOutput, error) {
	if params == nil {
		params = &UpdateEventConfigurationByResourceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEventConfigurationByResourceTypes", params, optFns, c.addOperationUpdateEventConfigurationByResourceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEventConfigurationByResourceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEventConfigurationByResourceTypesInput struct {

	// Connection status resource type event configuration object for enabling and
	// disabling wireless gateway topic.
	ConnectionStatus *types.ConnectionStatusResourceTypeEventConfiguration

	// Device registration state resource type event configuration object for enabling
	// and disabling wireless gateway topic.
	DeviceRegistrationState *types.DeviceRegistrationStateResourceTypeEventConfiguration

	// Join resource type event configuration object for enabling and disabling
	// wireless device topic.
	Join *types.JoinResourceTypeEventConfiguration

	// Message delivery status resource type event configuration object for enabling
	// and disabling wireless device topic.
	MessageDeliveryStatus *types.MessageDeliveryStatusResourceTypeEventConfiguration

	// Proximity resource type event configuration object for enabling and disabling
	// wireless gateway topic.
	Proximity *types.ProximityResourceTypeEventConfiguration

	noSmithyDocumentSerde
}

type UpdateEventConfigurationByResourceTypesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEventConfigurationByResourceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEventConfigurationByResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEventConfigurationByResourceTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEventConfigurationByResourceTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEventConfigurationByResourceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "UpdateEventConfigurationByResourceTypes",
	}
}
