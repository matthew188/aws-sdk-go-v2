// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a list of queries and query statuses for the past seven days. You must
// specify an ARN value for EventDataStore. Optionally, to shorten the list of
// results, you can specify a time range, formatted as timestamps, by adding
// StartTime and EndTime parameters, and a QueryStatus value. Valid values for
// QueryStatus include QUEUED, RUNNING, FINISHED, FAILED, TIMED_OUT, or CANCELLED.
func (c *Client) ListQueries(ctx context.Context, params *ListQueriesInput, optFns ...func(*Options)) (*ListQueriesOutput, error) {
	if params == nil {
		params = &ListQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueries", params, optFns, c.addOperationListQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueriesInput struct {

	// The ARN (or the ID suffix of the ARN) of an event data store on which queries
	// were run.
	//
	// This member is required.
	EventDataStore *string

	// Use with StartTime to bound a ListQueries request, and limit its results to only
	// those queries run within a specified time period.
	EndTime *time.Time

	// The maximum number of queries to show on a page.
	MaxResults *int32

	// A token you can use to get the next page of results.
	NextToken *string

	// The status of queries that you want to return in results. Valid values for
	// QueryStatus include QUEUED, RUNNING, FINISHED, FAILED, TIMED_OUT, or CANCELLED.
	QueryStatus types.QueryStatus

	// Use with EndTime to bound a ListQueries request, and limit its results to only
	// those queries run within a specified time period.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type ListQueriesOutput struct {

	// A token you can use to get the next page of results.
	NextToken *string

	// Lists matching query results, and shows query ID, status, and creation time of
	// each query.
	Queries []types.Query

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListQueries{}, middleware.After)
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
	if err = addOpListQueriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueries(options.Region), middleware.Before); err != nil {
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

// ListQueriesAPIClient is a client that implements the ListQueries operation.
type ListQueriesAPIClient interface {
	ListQueries(context.Context, *ListQueriesInput, ...func(*Options)) (*ListQueriesOutput, error)
}

var _ ListQueriesAPIClient = (*Client)(nil)

// ListQueriesPaginatorOptions is the paginator options for ListQueries
type ListQueriesPaginatorOptions struct {
	// The maximum number of queries to show on a page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueriesPaginator is a paginator for ListQueries
type ListQueriesPaginator struct {
	options   ListQueriesPaginatorOptions
	client    ListQueriesAPIClient
	params    *ListQueriesInput
	nextToken *string
	firstPage bool
}

// NewListQueriesPaginator returns a new ListQueriesPaginator
func NewListQueriesPaginator(client ListQueriesAPIClient, params *ListQueriesInput, optFns ...func(*ListQueriesPaginatorOptions)) *ListQueriesPaginator {
	if params == nil {
		params = &ListQueriesInput{}
	}

	options := ListQueriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueries page.
func (p *ListQueriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueriesOutput, error) {
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

	result, err := p.client.ListQueries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "ListQueries",
	}
}
