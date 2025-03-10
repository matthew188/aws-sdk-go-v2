// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates connectivity information for a Greengrass core device. Connectivity
// information includes endpoints and ports where client devices can connect to an
// MQTT broker on the core device. When a client device calls the IoT Greengrass
// discovery API
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/greengrass-discover-api.html),
// IoT Greengrass returns connectivity information for all of the core devices
// where the client device can connect. For more information, see Connect client
// devices to core devices
// (https://docs.aws.amazon.com/greengrass/v2/developerguide/connect-client-devices.html)
// in the IoT Greengrass Version 2 Developer Guide.
func (c *Client) UpdateConnectivityInfo(ctx context.Context, params *UpdateConnectivityInfoInput, optFns ...func(*Options)) (*UpdateConnectivityInfoOutput, error) {
	if params == nil {
		params = &UpdateConnectivityInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConnectivityInfo", params, optFns, c.addOperationUpdateConnectivityInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConnectivityInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConnectivityInfoInput struct {

	// The connectivity information for the core device.
	//
	// This member is required.
	ConnectivityInfo []types.ConnectivityInfo

	// The name of the core device. This is also the name of the IoT thing.
	//
	// This member is required.
	ThingName *string

	noSmithyDocumentSerde
}

type UpdateConnectivityInfoOutput struct {

	// A message about the connectivity information update request.
	Message *string

	// The new version of the connectivity information for the core device.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConnectivityInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConnectivityInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConnectivityInfo{}, middleware.After)
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
	if err = addOpUpdateConnectivityInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConnectivityInfo(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConnectivityInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "UpdateConnectivityInfo",
	}
}
