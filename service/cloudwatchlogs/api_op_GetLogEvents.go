// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists log events from the specified log stream. You can list all of the log
// events or filter using a time range. By default, this operation returns as many
// log events as can fit in a response size of 1MB (up to 10,000 log events). You
// can get additional log events by specifying one of the tokens in a subsequent
// call. This operation can return empty results while there are more log events
// available through the token. If you are using CloudWatch cross-account
// observability, you can use this operation in a monitoring account and view data
// from the linked source accounts. For more information, see CloudWatch
// cross-account observability
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html).
// You can specify the log group to search by using either logGroupIdentifier or
// logGroupName. You must include one of these two parameters, but you can't
// include both.
func (c *Client) GetLogEvents(ctx context.Context, params *GetLogEventsInput, optFns ...func(*Options)) (*GetLogEventsOutput, error) {
	if params == nil {
		params = &GetLogEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLogEvents", params, optFns, c.addOperationGetLogEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLogEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLogEventsInput struct {

	// The name of the log stream.
	//
	// This member is required.
	LogStreamName *string

	// The end of the time range, expressed as the number of milliseconds after Jan 1,
	// 1970 00:00:00 UTC. Events with a timestamp equal to or later than this time are
	// not included.
	EndTime *int64

	// The maximum number of log events returned. If you don't specify a limit, the
	// default is as many log events as can fit in a response size of 1 MB (up to
	// 10,000 log events).
	Limit *int32

	// Specify either the name or ARN of the log group to view events from. If the log
	// group is in a source account and you are using a monitoring account, you must
	// use the log group ARN. You must include either logGroupIdentifier or
	// logGroupName, but not both.
	LogGroupIdentifier *string

	// The name of the log group. You must include either logGroupIdentifier or
	// logGroupName, but not both.
	LogGroupName *string

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// If the value is true, the earliest log events are returned first. If the value
	// is false, the latest log events are returned first. The default value is false.
	// If you are using a previous nextForwardToken value as the nextToken in this
	// operation, you must specify true for startFromHead.
	StartFromHead *bool

	// The start of the time range, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC. Events with a timestamp equal to this time or later than
	// this time are included. Events with a timestamp earlier than this time are not
	// included.
	StartTime *int64

	// Specify true to display the log event fields with all sensitive data unmasked
	// and visible. The default is false. To use this operation with this parameter,
	// you must be signed into an account with the logs:Unmask permission.
	Unmask bool

	noSmithyDocumentSerde
}

type GetLogEventsOutput struct {

	// The events.
	Events []types.OutputLogEvent

	// The token for the next set of items in the backward direction. The token expires
	// after 24 hours. This token is not null. If you have reached the end of the
	// stream, it returns the same token you passed in.
	NextBackwardToken *string

	// The token for the next set of items in the forward direction. The token expires
	// after 24 hours. If you have reached the end of the stream, it returns the same
	// token you passed in.
	NextForwardToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLogEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLogEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLogEvents{}, middleware.After)
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
	if err = addOpGetLogEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLogEvents(options.Region), middleware.Before); err != nil {
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

// GetLogEventsAPIClient is a client that implements the GetLogEvents operation.
type GetLogEventsAPIClient interface {
	GetLogEvents(context.Context, *GetLogEventsInput, ...func(*Options)) (*GetLogEventsOutput, error)
}

var _ GetLogEventsAPIClient = (*Client)(nil)

// GetLogEventsPaginatorOptions is the paginator options for GetLogEvents
type GetLogEventsPaginatorOptions struct {
	// The maximum number of log events returned. If you don't specify a limit, the
	// default is as many log events as can fit in a response size of 1 MB (up to
	// 10,000 log events).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetLogEventsPaginator is a paginator for GetLogEvents
type GetLogEventsPaginator struct {
	options   GetLogEventsPaginatorOptions
	client    GetLogEventsAPIClient
	params    *GetLogEventsInput
	nextToken *string
	firstPage bool
}

// NewGetLogEventsPaginator returns a new GetLogEventsPaginator
func NewGetLogEventsPaginator(client GetLogEventsAPIClient, params *GetLogEventsInput, optFns ...func(*GetLogEventsPaginatorOptions)) *GetLogEventsPaginator {
	if params == nil {
		params = &GetLogEventsInput{}
	}

	options := GetLogEventsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetLogEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetLogEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetLogEvents page.
func (p *GetLogEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetLogEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.GetLogEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextForwardToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetLogEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "GetLogEvents",
	}
}
