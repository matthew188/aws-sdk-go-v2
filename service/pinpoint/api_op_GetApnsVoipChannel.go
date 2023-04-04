// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the status and settings of the APNs VoIP channel for
// an application.
func (c *Client) GetApnsVoipChannel(ctx context.Context, params *GetApnsVoipChannelInput, optFns ...func(*Options)) (*GetApnsVoipChannelOutput, error) {
	if params == nil {
		params = &GetApnsVoipChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApnsVoipChannel", params, optFns, c.addOperationGetApnsVoipChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApnsVoipChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApnsVoipChannelInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	noSmithyDocumentSerde
}

type GetApnsVoipChannelOutput struct {

	// Provides information about the status and settings of the APNs (Apple Push
	// Notification service) VoIP channel for an application.
	//
	// This member is required.
	APNSVoipChannelResponse *types.APNSVoipChannelResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApnsVoipChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApnsVoipChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApnsVoipChannel{}, middleware.After)
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
	if err = addOpGetApnsVoipChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApnsVoipChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApnsVoipChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "GetApnsVoipChannel",
	}
}
