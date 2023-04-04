// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information, including current status, about a game session placement
// request. To get game session placement details, specify the placement ID. This
// operation is not designed to be continually called to track game session status.
// This practice can cause you to exceed your API limit, which results in errors.
// Instead, you must configure configure an Amazon Simple Notification Service
// (SNS) topic to receive notifications from FlexMatch or queues. Continuously
// polling with DescribeGameSessionPlacement should only be used for games in
// development with low game session usage.
func (c *Client) DescribeGameSessionPlacement(ctx context.Context, params *DescribeGameSessionPlacementInput, optFns ...func(*Options)) (*DescribeGameSessionPlacementOutput, error) {
	if params == nil {
		params = &DescribeGameSessionPlacementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGameSessionPlacement", params, optFns, c.addOperationDescribeGameSessionPlacementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGameSessionPlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGameSessionPlacementInput struct {

	// A unique identifier for a game session placement to retrieve.
	//
	// This member is required.
	PlacementId *string

	noSmithyDocumentSerde
}

type DescribeGameSessionPlacementOutput struct {

	// Object that describes the requested game session placement.
	GameSessionPlacement *types.GameSessionPlacement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeGameSessionPlacementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeGameSessionPlacement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeGameSessionPlacement{}, middleware.After)
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
	if err = addOpDescribeGameSessionPlacementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGameSessionPlacement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeGameSessionPlacement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeGameSessionPlacement",
	}
}
