// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/braket/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for Amazon Braket jobs that match the specified filter values.
func (c *Client) SearchJobs(ctx context.Context, params *SearchJobsInput, optFns ...func(*Options)) (*SearchJobsOutput, error) {
	if params == nil {
		params = &SearchJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchJobs", params, optFns, c.addOperationSearchJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchJobsInput struct {

	// The filter values to use when searching for a job.
	//
	// This member is required.
	Filters []types.SearchJobsFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned in the response. Use the token
	// returned from the previous request to continue results where the previous
	// request ended.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchJobsOutput struct {

	// An array of JobSummary objects for devices that match the specified filter
	// values.
	//
	// This member is required.
	Jobs []types.JobSummary

	// A token used for pagination of results, or null if there are no additional
	// results. Use the token value in a subsequent request to continue results where
	// the previous request ended.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchJobs{}, middleware.After)
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
	if err = addOpSearchJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchJobs(options.Region), middleware.Before); err != nil {
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

// SearchJobsAPIClient is a client that implements the SearchJobs operation.
type SearchJobsAPIClient interface {
	SearchJobs(context.Context, *SearchJobsInput, ...func(*Options)) (*SearchJobsOutput, error)
}

var _ SearchJobsAPIClient = (*Client)(nil)

// SearchJobsPaginatorOptions is the paginator options for SearchJobs
type SearchJobsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchJobsPaginator is a paginator for SearchJobs
type SearchJobsPaginator struct {
	options   SearchJobsPaginatorOptions
	client    SearchJobsAPIClient
	params    *SearchJobsInput
	nextToken *string
	firstPage bool
}

// NewSearchJobsPaginator returns a new SearchJobsPaginator
func NewSearchJobsPaginator(client SearchJobsAPIClient, params *SearchJobsInput, optFns ...func(*SearchJobsPaginatorOptions)) *SearchJobsPaginator {
	if params == nil {
		params = &SearchJobsInput{}
	}

	options := SearchJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchJobs page.
func (p *SearchJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchJobsOutput, error) {
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

	result, err := p.client.SearchJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "braket",
		OperationName: "SearchJobs",
	}
}
