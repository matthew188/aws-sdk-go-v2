// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the bandwidth rate limits of a gateway. By default, these limits are not
// set, which means no bandwidth rate limiting is in effect. This operation is
// supported only for the stored volume, cached volume, and tape gateway types. To
// describe bandwidth rate limits for S3 file gateways, use
// DescribeBandwidthRateLimitSchedule. This operation returns a value for a
// bandwidth rate limit only if the limit is set. If no limits are set for the
// gateway, then this operation returns only the gateway ARN in the response body.
// To specify which gateway to describe, use the Amazon Resource Name (ARN) of the
// gateway in your request.
func (c *Client) DescribeBandwidthRateLimit(ctx context.Context, params *DescribeBandwidthRateLimitInput, optFns ...func(*Options)) (*DescribeBandwidthRateLimitOutput, error) {
	if params == nil {
		params = &DescribeBandwidthRateLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBandwidthRateLimit", params, optFns, c.addOperationDescribeBandwidthRateLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBandwidthRateLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway.
type DescribeBandwidthRateLimitInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
type DescribeBandwidthRateLimitOutput struct {

	// The average download bandwidth rate limit in bits per second. This field does
	// not appear in the response if the download rate limit is not set.
	AverageDownloadRateLimitInBitsPerSec *int64

	// The average upload bandwidth rate limit in bits per second. This field does not
	// appear in the response if the upload rate limit is not set.
	AverageUploadRateLimitInBitsPerSec *int64

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBandwidthRateLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBandwidthRateLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBandwidthRateLimit{}, middleware.After)
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
	if err = addOpDescribeBandwidthRateLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBandwidthRateLimit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeBandwidthRateLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeBandwidthRateLimit",
	}
}
