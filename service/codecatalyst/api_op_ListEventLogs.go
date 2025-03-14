// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/service/codecatalyst/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves a list of events that occurred during a specified time period in a
// space. You can use these events to audit user and system activity in a space.
func (c *Client) ListEventLogs(ctx context.Context, params *ListEventLogsInput, optFns ...func(*Options)) (*ListEventLogsOutput, error) {
	if params == nil {
		params = &ListEventLogsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventLogs", params, optFns, c.addOperationListEventLogsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventLogsInput struct {

	// The time after which you do not want any events retrieved, in coordinated
	// universal time (UTC) timestamp format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	//
	// This member is required.
	EndTime *time.Time

	// The name of the space.
	//
	// This member is required.
	SpaceName *string

	// The date and time when you want to start retrieving events, in coordinated
	// universal time (UTC) timestamp format as specified in RFC 3339
	// (https://www.rfc-editor.org/rfc/rfc3339#section-5.6).
	//
	// This member is required.
	StartTime *time.Time

	// The name of the event.
	EventName *string

	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	MaxResults *int32

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEventLogsOutput struct {

	// Information about each event retrieved in the list.
	//
	// This member is required.
	Items []types.EventLogEntry

	// A token returned from a call to this API to indicate the next batch of results
	// to return, if any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventLogsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEventLogs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventLogs{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addBearerAuthSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListEventLogsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventLogs(options.Region), middleware.Before); err != nil {
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

// ListEventLogsAPIClient is a client that implements the ListEventLogs operation.
type ListEventLogsAPIClient interface {
	ListEventLogs(context.Context, *ListEventLogsInput, ...func(*Options)) (*ListEventLogsOutput, error)
}

var _ ListEventLogsAPIClient = (*Client)(nil)

// ListEventLogsPaginatorOptions is the paginator options for ListEventLogs
type ListEventLogsPaginatorOptions struct {
	// The maximum number of results to show in a single call to this API. If the
	// number of results is larger than the number you specified, the response will
	// include a NextToken element, which you can use to obtain additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventLogsPaginator is a paginator for ListEventLogs
type ListEventLogsPaginator struct {
	options   ListEventLogsPaginatorOptions
	client    ListEventLogsAPIClient
	params    *ListEventLogsInput
	nextToken *string
	firstPage bool
}

// NewListEventLogsPaginator returns a new ListEventLogsPaginator
func NewListEventLogsPaginator(client ListEventLogsAPIClient, params *ListEventLogsInput, optFns ...func(*ListEventLogsPaginatorOptions)) *ListEventLogsPaginator {
	if params == nil {
		params = &ListEventLogsInput{}
	}

	options := ListEventLogsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventLogsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventLogsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventLogs page.
func (p *ListEventLogsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventLogsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListEventLogs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListEventLogs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEventLogs",
	}
}
