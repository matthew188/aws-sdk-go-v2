// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsecuretunneling

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotsecuretunneling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new tunnel, and returns two client access tokens for clients to use to
// connect to the IoT Secure Tunneling proxy server. Requires permission to access
// the OpenTunnel
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) OpenTunnel(ctx context.Context, params *OpenTunnelInput, optFns ...func(*Options)) (*OpenTunnelOutput, error) {
	if params == nil {
		params = &OpenTunnelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "OpenTunnel", params, optFns, c.addOperationOpenTunnelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*OpenTunnelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type OpenTunnelInput struct {

	// A short text description of the tunnel.
	Description *string

	// The destination configuration for the OpenTunnel request.
	DestinationConfig *types.DestinationConfig

	// A collection of tag metadata.
	Tags []types.Tag

	// Timeout configuration for a tunnel.
	TimeoutConfig *types.TimeoutConfig

	noSmithyDocumentSerde
}

type OpenTunnelOutput struct {

	// The access token the destination local proxy uses to connect to IoT Secure
	// Tunneling.
	DestinationAccessToken *string

	// The access token the source local proxy uses to connect to IoT Secure Tunneling.
	SourceAccessToken *string

	// The Amazon Resource Name for the tunnel.
	TunnelArn *string

	// A unique alpha-numeric tunnel ID.
	TunnelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationOpenTunnelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpOpenTunnel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpOpenTunnel{}, middleware.After)
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
	if err = addOpOpenTunnelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opOpenTunnel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opOpenTunnel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "IoTSecuredTunneling",
		OperationName: "OpenTunnel",
	}
}
