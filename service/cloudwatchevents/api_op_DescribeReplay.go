// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudwatchevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves details about a replay. Use DescribeReplay to determine the progress
// of a running replay. A replay processes events to replay based on the time in
// the event, and replays them using 1 minute intervals. If you use StartReplay and
// specify an EventStartTime and an EventEndTime that covers a 20 minute time
// range, the events are replayed from the first minute of that 20 minute range
// first. Then the events from the second minute are replayed. You can use
// DescribeReplay to determine the progress of a replay. The value returned for
// EventLastReplayedTime indicates the time within the specified time range
// associated with the last event replayed.
func (c *Client) DescribeReplay(ctx context.Context, params *DescribeReplayInput, optFns ...func(*Options)) (*DescribeReplayOutput, error) {
	if params == nil {
		params = &DescribeReplayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplay", params, optFns, c.addOperationDescribeReplayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplayInput struct {

	// The name of the replay to retrieve.
	//
	// This member is required.
	ReplayName *string

	noSmithyDocumentSerde
}

type DescribeReplayOutput struct {

	// The description of the replay.
	Description *string

	// A ReplayDestination object that contains details about the replay.
	Destination *types.ReplayDestination

	// The time stamp for the last event that was replayed from the archive.
	EventEndTime *time.Time

	// The time that the event was last replayed.
	EventLastReplayedTime *time.Time

	// The ARN of the archive events were replayed from.
	EventSourceArn *string

	// The time stamp of the first event that was last replayed from the archive.
	EventStartTime *time.Time

	// The ARN of the replay.
	ReplayArn *string

	// A time stamp for the time that the replay stopped.
	ReplayEndTime *time.Time

	// The name of the replay.
	ReplayName *string

	// A time stamp for the time that the replay started.
	ReplayStartTime *time.Time

	// The current state of the replay.
	State types.ReplayState

	// The reason that the replay is in the current state.
	StateReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplay{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplay{}, middleware.After)
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
	if err = addOpDescribeReplayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplay(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReplay(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeReplay",
	}
}
