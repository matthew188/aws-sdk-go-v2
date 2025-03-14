// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the bandwidth rate limits of a gateway. You can delete either the upload
// and download bandwidth rate limit, or you can delete both. If you delete only
// one of the limits, the other limit remains unchanged. To specify which gateway
// to work with, use the Amazon Resource Name (ARN) of the gateway in your request.
// This operation is supported only for the stored volume, cached volume, and tape
// gateway types.
func (c *Client) DeleteBandwidthRateLimit(ctx context.Context, params *DeleteBandwidthRateLimitInput, optFns ...func(*Options)) (*DeleteBandwidthRateLimitOutput, error) {
	if params == nil {
		params = &DeleteBandwidthRateLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBandwidthRateLimit", params, optFns, c.addOperationDeleteBandwidthRateLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBandwidthRateLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the following fields:
//
// *
// DeleteBandwidthRateLimitInput$BandwidthType
type DeleteBandwidthRateLimitInput struct {

	// One of the BandwidthType values that indicates the gateway bandwidth rate limit
	// to delete. Valid Values: UPLOAD | DOWNLOAD | ALL
	//
	// This member is required.
	BandwidthType *string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	noSmithyDocumentSerde
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway whose
// bandwidth rate information was deleted.
type DeleteBandwidthRateLimitOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBandwidthRateLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBandwidthRateLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBandwidthRateLimit{}, middleware.After)
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
	if err = addOpDeleteBandwidthRateLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBandwidthRateLimit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBandwidthRateLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteBandwidthRateLimit",
	}
}
