// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads segment documents to Amazon Web Services X-Ray. The X-Ray SDK
// (https://docs.aws.amazon.com/xray/index.html) generates segment documents and
// sends them to the X-Ray daemon, which uploads them in batches. A segment
// document can be a completed segment, an in-progress segment, or an array of
// subsegments. Segments must include the following fields. For the full segment
// document schema, see Amazon Web Services X-Ray Segment Documents
// (https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html)
// in the Amazon Web Services X-Ray Developer Guide. Required segment document
// fields
//
// * name - The name of the service that handled the request.
//
// * id - A
// 64-bit identifier for the segment, unique among segments in the same trace, in
// 16 hexadecimal digits.
//
// * trace_id - A unique identifier that connects all
// segments and subsegments originating from a single client request.
//
// * start_time
// - Time the segment or subsegment was created, in floating point seconds in epoch
// time, accurate to milliseconds. For example, 1480615200.010 or
// 1.480615200010E9.
//
// * end_time - Time the segment or subsegment was closed. For
// example, 1480615200.090 or 1.480615200090E9. Specify either an end_time or
// in_progress.
//
// * in_progress - Set to true instead of specifying an end_time to
// record that a segment has been started, but is not complete. Send an in-progress
// segment when your application receives a request that will take a long time to
// serve, to trace that the request was received. When the response is sent, send
// the complete segment to overwrite the in-progress segment.
//
// A trace_id consists
// of three numbers separated by hyphens. For example,
// 1-58406520-a006649127e371903a2de979. This includes: Trace ID Format
//
// * The
// version number, for instance, 1.
//
// * The time of the original request, in Unix
// epoch time, in 8 hexadecimal digits. For example, 10:00AM December 2nd, 2016 PST
// in epoch time is 1480615200 seconds, or 58406520 in hexadecimal.
//
// * A 96-bit
// identifier for the trace, globally unique, in 24 hexadecimal digits.
func (c *Client) PutTraceSegments(ctx context.Context, params *PutTraceSegmentsInput, optFns ...func(*Options)) (*PutTraceSegmentsOutput, error) {
	if params == nil {
		params = &PutTraceSegmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutTraceSegments", params, optFns, c.addOperationPutTraceSegmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutTraceSegmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutTraceSegmentsInput struct {

	// A string containing a JSON document defining one or more segments or
	// subsegments.
	//
	// This member is required.
	TraceSegmentDocuments []string

	noSmithyDocumentSerde
}

type PutTraceSegmentsOutput struct {

	// Segments that failed processing.
	UnprocessedTraceSegments []types.UnprocessedTraceSegment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutTraceSegmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutTraceSegments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutTraceSegments{}, middleware.After)
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
	if err = addOpPutTraceSegmentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutTraceSegments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutTraceSegments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "PutTraceSegments",
	}
}
