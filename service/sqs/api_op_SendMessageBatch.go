// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delivers up to ten messages to the specified queue. This is a batch version of
// SendMessage. For a FIFO queue, multiple messages within a single batch are
// enqueued in the order they are sent. The result of sending each message is
// reported individually in the response. Because the batch request can result in a
// combination of successful and unsuccessful actions, you should check for batch
// errors even when the call returns an HTTP status code of 200. The maximum
// allowed individual message size and the maximum total payload size (the sum of
// the individual lengths of all of the batched messages) are both 256 KB (262,144
// bytes). A message can include only XML, JSON, and unformatted text. The
// following Unicode characters are allowed: #x9 | #xA | #xD | #x20 to #xD7FF |
// #xE000 to #xFFFD | #x10000 to #x10FFFF Any characters not included in this list
// will be rejected. For more information, see the W3C specification for characters
// (http://www.w3.org/TR/REC-xml/#charsets). If you don't specify the DelaySeconds
// parameter for an entry, Amazon SQS uses the default value for the queue. Some
// actions take lists of parameters. These lists are specified using the param.n
// notation. Values of n are integers starting from 1. For example, a parameter
// list with two elements looks like this: &AttributeName.1=first
//
// &AttributeName.2=second
func (c *Client) SendMessageBatch(ctx context.Context, params *SendMessageBatchInput, optFns ...func(*Options)) (*SendMessageBatchOutput, error) {
	if params == nil {
		params = &SendMessageBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendMessageBatch", params, optFns, c.addOperationSendMessageBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendMessageBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendMessageBatchInput struct {

	// A list of SendMessageBatchRequestEntry items.
	//
	// This member is required.
	Entries []types.SendMessageBatchRequestEntry

	// The URL of the Amazon SQS queue to which batched messages are sent. Queue URLs
	// and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	noSmithyDocumentSerde
}

// For each message in the batch, the response contains a
// SendMessageBatchResultEntry tag if the message succeeds or a
// BatchResultErrorEntry tag if the message fails.
type SendMessageBatchOutput struct {

	// A list of BatchResultErrorEntry items with error details about each message that
	// can't be enqueued.
	//
	// This member is required.
	Failed []types.BatchResultErrorEntry

	// A list of SendMessageBatchResultEntry items.
	//
	// This member is required.
	Successful []types.SendMessageBatchResultEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendMessageBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSendMessageBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSendMessageBatch{}, middleware.After)
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
	if err = addValidateSendMessageBatchChecksum(stack, options); err != nil {
		return err
	}
	if err = addOpSendMessageBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendMessageBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendMessageBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "SendMessageBatch",
	}
}
