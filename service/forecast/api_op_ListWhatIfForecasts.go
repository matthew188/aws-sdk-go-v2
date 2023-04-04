// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of what-if forecasts created using the CreateWhatIfForecast
// operation. For each what-if forecast, this operation returns a summary of its
// properties, including its Amazon Resource Name (ARN). You can retrieve the
// complete set of properties by using the what-if forecast ARN with the
// DescribeWhatIfForecast operation.
func (c *Client) ListWhatIfForecasts(ctx context.Context, params *ListWhatIfForecastsInput, optFns ...func(*Options)) (*ListWhatIfForecastsOutput, error) {
	if params == nil {
		params = &ListWhatIfForecastsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWhatIfForecasts", params, optFns, c.addOperationListWhatIfForecastsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWhatIfForecastsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWhatIfForecastsInput struct {

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether to
	// include or exclude the what-if forecast export jobs that match the statement
	// from the list, respectively. The match statement consists of a key and a value.
	// Filter properties
	//
	// * Condition - The condition to apply. Valid values are IS and
	// IS_NOT. To include the forecast export jobs that match the statement, specify
	// IS. To exclude matching forecast export jobs, specify IS_NOT.
	//
	// * Key - The name
	// of the parameter to filter on. Valid values are WhatIfForecastArn and Status.
	//
	// *
	// Value - The value to match.
	//
	// For example, to list all jobs that export a
	// forecast named electricityWhatIfForecast, specify the following filter:
	// "Filters": [ { "Condition": "IS", "Key": "WhatIfForecastArn", "Value":
	// "arn:aws:forecast:us-west-2::forecast/electricityWhatIfForecast" } ]
	Filters []types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWhatIfForecastsOutput struct {

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	// An array of WhatIfForecasts objects that describe the matched forecasts.
	WhatIfForecasts []types.WhatIfForecastSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWhatIfForecastsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWhatIfForecasts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWhatIfForecasts{}, middleware.After)
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
	if err = addOpListWhatIfForecastsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWhatIfForecasts(options.Region), middleware.Before); err != nil {
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

// ListWhatIfForecastsAPIClient is a client that implements the ListWhatIfForecasts
// operation.
type ListWhatIfForecastsAPIClient interface {
	ListWhatIfForecasts(context.Context, *ListWhatIfForecastsInput, ...func(*Options)) (*ListWhatIfForecastsOutput, error)
}

var _ ListWhatIfForecastsAPIClient = (*Client)(nil)

// ListWhatIfForecastsPaginatorOptions is the paginator options for
// ListWhatIfForecasts
type ListWhatIfForecastsPaginatorOptions struct {
	// The number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWhatIfForecastsPaginator is a paginator for ListWhatIfForecasts
type ListWhatIfForecastsPaginator struct {
	options   ListWhatIfForecastsPaginatorOptions
	client    ListWhatIfForecastsAPIClient
	params    *ListWhatIfForecastsInput
	nextToken *string
	firstPage bool
}

// NewListWhatIfForecastsPaginator returns a new ListWhatIfForecastsPaginator
func NewListWhatIfForecastsPaginator(client ListWhatIfForecastsAPIClient, params *ListWhatIfForecastsInput, optFns ...func(*ListWhatIfForecastsPaginatorOptions)) *ListWhatIfForecastsPaginator {
	if params == nil {
		params = &ListWhatIfForecastsInput{}
	}

	options := ListWhatIfForecastsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWhatIfForecastsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWhatIfForecastsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWhatIfForecasts page.
func (p *ListWhatIfForecastsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWhatIfForecastsOutput, error) {
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

	result, err := p.client.ListWhatIfForecasts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWhatIfForecasts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListWhatIfForecasts",
	}
}
