// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets details about an existing OriginEndpoint.
func (c *Client) DescribeOriginEndpoint(ctx context.Context, params *DescribeOriginEndpointInput, optFns ...func(*Options)) (*DescribeOriginEndpointOutput, error) {
	if params == nil {
		params = &DescribeOriginEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOriginEndpoint", params, optFns, c.addOperationDescribeOriginEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOriginEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOriginEndpointInput struct {

	// The ID of the OriginEndpoint.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DescribeOriginEndpointOutput struct {

	// The Amazon Resource Name (ARN) assigned to the OriginEndpoint.
	Arn *string

	// CDN Authorization credentials
	Authorization *types.Authorization

	// The ID of the Channel the OriginEndpoint is associated with.
	ChannelId *string

	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *types.CmafPackage

	// The date and time the OriginEndpoint was created.
	CreatedAt *string

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *types.DashPackage

	// A short text description of the OriginEndpoint.
	Description *string

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *types.HlsPackage

	// The ID of the OriginEndpoint.
	Id *string

	// A short string appended to the end of the OriginEndpoint URL.
	ManifestName *string

	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *types.MssPackage

	// Control whether origination of video is allowed for this OriginEndpoint. If set
	// to ALLOW, the OriginEndpoint may by requested, pursuant to any other form of
	// access control. If set to DENY, the OriginEndpoint may not be requested. This
	// can be helpful for Live to VOD harvesting, or for temporarily disabling
	// origination
	Origination types.Origination

	// Maximum duration (seconds) of content to retain for startover playback. If not
	// specified, startover playback will be disabled for the OriginEndpoint.
	StartoverWindowSeconds int32

	// A collection of tags associated with a resource
	Tags map[string]string

	// Amount of delay (seconds) to enforce on the playback of live content. If not
	// specified, there will be no time delay in effect for the OriginEndpoint.
	TimeDelaySeconds int32

	// The URL of the packaged OriginEndpoint for consumption.
	Url *string

	// A list of source IP CIDR blocks that will be allowed to access the
	// OriginEndpoint.
	Whitelist []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOriginEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOriginEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOriginEndpoint{}, middleware.After)
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
	if err = addOpDescribeOriginEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOriginEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOriginEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "DescribeOriginEndpoint",
	}
}
