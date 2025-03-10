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

// Starts a bulk disassociatin of all qualifying wireless devices from a multicast
// group.
func (c *Client) StartBulkDisassociateWirelessDeviceFromMulticastGroup(ctx context.Context, params *StartBulkDisassociateWirelessDeviceFromMulticastGroupInput, optFns ...func(*Options)) (*StartBulkDisassociateWirelessDeviceFromMulticastGroupOutput, error) {
	if params == nil {
		params = &StartBulkDisassociateWirelessDeviceFromMulticastGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBulkDisassociateWirelessDeviceFromMulticastGroup", params, optFns, c.addOperationStartBulkDisassociateWirelessDeviceFromMulticastGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBulkDisassociateWirelessDeviceFromMulticastGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBulkDisassociateWirelessDeviceFromMulticastGroupInput struct {

	// The ID of the multicast group.
	//
	// This member is required.
	Id *string

	// Query string used to search for wireless devices as part of the bulk associate
	// and disassociate process.
	QueryString *string

	// The tag to attach to the specified resource. Tags are metadata that you can use
	// to manage a resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartBulkDisassociateWirelessDeviceFromMulticastGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartBulkDisassociateWirelessDeviceFromMulticastGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartBulkDisassociateWirelessDeviceFromMulticastGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBulkDisassociateWirelessDeviceFromMulticastGroup{}, middleware.After)
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
	if err = addOpStartBulkDisassociateWirelessDeviceFromMulticastGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBulkDisassociateWirelessDeviceFromMulticastGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBulkDisassociateWirelessDeviceFromMulticastGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "StartBulkDisassociateWirelessDeviceFromMulticastGroup",
	}
}
