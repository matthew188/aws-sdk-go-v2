// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkmeetings/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates AttendeeCapabilities except the capabilities listed in an
// ExcludedAttendeeIds table. You use the capabilities with a set of values that
// control what the capabilities can do, such as SendReceive data. For more
// information about those values, see . When using capabilities, be aware of these
// corner cases:
//
// * You can't set content capabilities to SendReceive or Receive
// unless you also set video capabilities to SendReceive or Receive. If you don't
// set the video capability to receive, the response will contain an HTTP 400 Bad
// Request status code. However, you can set your video capability to receive and
// you set your content capability to not receive.
//
// * When you change an audio
// capability from None or Receive to Send or SendReceive , and if the attendee
// left their microphone unmuted, audio will flow from the attendee to the other
// meeting participants.
//
// * When you change a video or content capability from None
// or Receive to Send or SendReceive , and if the attendee turned on their video or
// content streams, remote attendess can receive those streams, but only after
// media renegotiation between the client and the Amazon Chime back-end server.
func (c *Client) BatchUpdateAttendeeCapabilitiesExcept(ctx context.Context, params *BatchUpdateAttendeeCapabilitiesExceptInput, optFns ...func(*Options)) (*BatchUpdateAttendeeCapabilitiesExceptOutput, error) {
	if params == nil {
		params = &BatchUpdateAttendeeCapabilitiesExceptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateAttendeeCapabilitiesExcept", params, optFns, c.addOperationBatchUpdateAttendeeCapabilitiesExceptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateAttendeeCapabilitiesExceptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateAttendeeCapabilitiesExceptInput struct {

	// The capabilities (audio, video, or content) that you want to update.
	//
	// This member is required.
	Capabilities *types.AttendeeCapabilities

	// The AttendeeIDs that you want to exclude from one or more capabilities.
	//
	// This member is required.
	ExcludedAttendeeIds []types.AttendeeIdItem

	// The ID of the meeting associated with the update request.
	//
	// This member is required.
	MeetingId *string

	noSmithyDocumentSerde
}

type BatchUpdateAttendeeCapabilitiesExceptOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateAttendeeCapabilitiesExceptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateAttendeeCapabilitiesExcept{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateAttendeeCapabilitiesExcept{}, middleware.After)
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
	if err = addOpBatchUpdateAttendeeCapabilitiesExceptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateAttendeeCapabilitiesExcept(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateAttendeeCapabilitiesExcept(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchUpdateAttendeeCapabilitiesExcept",
	}
}
