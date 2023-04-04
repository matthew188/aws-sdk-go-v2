// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideowebrtcstorage

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Join the ongoing one way-video and/or multi-way audio WebRTC session as a video
// producing device for an input channel. If there’s no existing session for the
// channel, a new streaming session needs to be created, and the Amazon Resource
// Name (ARN) of the signaling channel must be provided. Currently for the
// SINGLE_MASTER type, a video producing device is able to ingest both audio and
// video media into a stream, while viewers can only ingest audio. Both a video
// producing device and viewers can join the session first, and wait for other
// participants. While participants are having peer to peer conversations through
// webRTC, the ingested media session will be stored into the Kinesis Video Stream.
// Multiple viewers are able to playback real-time media. Customers can also use
// existing Kinesis Video Streams features like HLS or DASH playback, Image
// generation, and more with ingested WebRTC media. Assume that only one video
// producing device client can be associated with a session for the channel. If
// more than one client joins the session of a specific channel as a video
// producing device, the most recent client request takes precedence.
func (c *Client) JoinStorageSession(ctx context.Context, params *JoinStorageSessionInput, optFns ...func(*Options)) (*JoinStorageSessionOutput, error) {
	if params == nil {
		params = &JoinStorageSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "JoinStorageSession", params, optFns, c.addOperationJoinStorageSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*JoinStorageSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type JoinStorageSessionInput struct {

	// The Amazon Resource Name (ARN) of the signaling channel.
	//
	// This member is required.
	ChannelArn *string

	noSmithyDocumentSerde
}

type JoinStorageSessionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationJoinStorageSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpJoinStorageSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpJoinStorageSession{}, middleware.After)
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
	if err = addOpJoinStorageSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opJoinStorageSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opJoinStorageSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "JoinStorageSession",
	}
}
