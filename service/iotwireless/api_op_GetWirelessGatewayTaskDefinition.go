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

// Gets information about a wireless gateway task definition.
func (c *Client) GetWirelessGatewayTaskDefinition(ctx context.Context, params *GetWirelessGatewayTaskDefinitionInput, optFns ...func(*Options)) (*GetWirelessGatewayTaskDefinitionOutput, error) {
	if params == nil {
		params = &GetWirelessGatewayTaskDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWirelessGatewayTaskDefinition", params, optFns, c.addOperationGetWirelessGatewayTaskDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWirelessGatewayTaskDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWirelessGatewayTaskDefinitionInput struct {

	// The ID of the resource to get.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetWirelessGatewayTaskDefinitionOutput struct {

	// The Amazon Resource Name of the resource.
	Arn *string

	// Whether to automatically create tasks using this task definition for all
	// gateways with the specified current version. If false, the task must me created
	// by calling CreateWirelessGatewayTask.
	AutoCreateTasks bool

	// The name of the resource.
	Name *string

	// Information about the gateways to update.
	Update *types.UpdateWirelessGatewayTaskCreate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWirelessGatewayTaskDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWirelessGatewayTaskDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWirelessGatewayTaskDefinition{}, middleware.After)
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
	if err = addOpGetWirelessGatewayTaskDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWirelessGatewayTaskDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWirelessGatewayTaskDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetWirelessGatewayTaskDefinition",
	}
}
