// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns events related to clusters, security groups, snapshots, and parameter
// groups for the past 14 days. Events specific to a particular cluster, security
// group, snapshot or parameter group can be obtained by providing the name as a
// parameter. By default, the past hour of events are returned.
func (c *Client) DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) {
	if params == nil {
		params = &DescribeEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEvents", params, optFns, c.addOperationDescribeEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventsInput struct {

	// The number of minutes prior to the time of the request for which to retrieve
	// events. For example, if the request is sent at 18:00 and you specify a duration
	// of 60, then only events which have occurred after 17:00 will be returned.
	// Default: 60
	Duration *int32

	// The end of the time interval for which to retrieve events, specified in ISO 8601
	// format. For more information about ISO 8601, go to the ISO8601 Wikipedia page.
	// (http://en.wikipedia.org/wiki/ISO_8601) Example: 2009-07-08T18:00Z
	EndTime *time.Time

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeEvents request exceed the value
	// specified in MaxRecords, Amazon Web Services returns a value in the Marker field
	// of the response. You can retrieve the next set of response records by providing
	// the returned marker value in the Marker parameter and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The identifier of the event source for which events will be returned. If this
	// parameter is not specified, then all sources are included in the response.
	// Constraints: If SourceIdentifier is supplied, SourceType must also be
	// provided.
	//
	// * Specify a cluster identifier when SourceType is cluster.
	//
	// * Specify
	// a cluster security group name when SourceType is cluster-security-group.
	//
	// *
	// Specify a cluster parameter group name when SourceType is
	// cluster-parameter-group.
	//
	// * Specify a cluster snapshot identifier when
	// SourceType is cluster-snapshot.
	SourceIdentifier *string

	// The event source to retrieve events for. If no value is specified, all events
	// are returned. Constraints: If SourceType is supplied, SourceIdentifier must also
	// be provided.
	//
	// * Specify cluster when SourceIdentifier is a cluster
	// identifier.
	//
	// * Specify cluster-security-group when SourceIdentifier is a cluster
	// security group name.
	//
	// * Specify cluster-parameter-group when SourceIdentifier is
	// a cluster parameter group name.
	//
	// * Specify cluster-snapshot when
	// SourceIdentifier is a cluster snapshot identifier.
	SourceType types.SourceType

	// The beginning of the time interval to retrieve events for, specified in ISO 8601
	// format. For more information about ISO 8601, go to the ISO8601 Wikipedia page.
	// (http://en.wikipedia.org/wiki/ISO_8601) Example: 2009-07-08T18:00Z
	StartTime *time.Time

	noSmithyDocumentSerde
}

type DescribeEventsOutput struct {

	// A list of Event instances.
	Events []types.Event

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEvents{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEvents(options.Region), middleware.Before); err != nil {
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

// DescribeEventsAPIClient is a client that implements the DescribeEvents
// operation.
type DescribeEventsAPIClient interface {
	DescribeEvents(context.Context, *DescribeEventsInput, ...func(*Options)) (*DescribeEventsOutput, error)
}

var _ DescribeEventsAPIClient = (*Client)(nil)

// DescribeEventsPaginatorOptions is the paginator options for DescribeEvents
type DescribeEventsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEventsPaginator is a paginator for DescribeEvents
type DescribeEventsPaginator struct {
	options   DescribeEventsPaginatorOptions
	client    DescribeEventsAPIClient
	params    *DescribeEventsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEventsPaginator returns a new DescribeEventsPaginator
func NewDescribeEventsPaginator(client DescribeEventsAPIClient, params *DescribeEventsInput, optFns ...func(*DescribeEventsPaginatorOptions)) *DescribeEventsPaginator {
	if params == nil {
		params = &DescribeEventsInput{}
	}

	options := DescribeEventsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEvents page.
func (p *DescribeEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeEvents",
	}
}
