// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists processing jobs that satisfy various filters.
func (c *Client) ListProcessingJobs(ctx context.Context, params *ListProcessingJobsInput, optFns ...func(*Options)) (*ListProcessingJobsOutput, error) {
	if params == nil {
		params = &ListProcessingJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProcessingJobs", params, optFns, c.addOperationListProcessingJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProcessingJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProcessingJobsInput struct {

	// A filter that returns only processing jobs created after the specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only processing jobs created after the specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only processing jobs modified after the specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only processing jobs modified before the specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of processing jobs to return in the response.
	MaxResults *int32

	// A string in the processing job name. This filter returns only processing jobs
	// whose name contains the specified string.
	NameContains *string

	// If the result of the previous ListProcessingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of processing jobs, use
	// the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.SortBy

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that retrieves only processing jobs with a specific status.
	StatusEquals types.ProcessingJobStatus

	noSmithyDocumentSerde
}

type ListProcessingJobsOutput struct {

	// An array of ProcessingJobSummary objects, each listing a processing job.
	//
	// This member is required.
	ProcessingJobSummaries []types.ProcessingJobSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of processing jobs, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProcessingJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProcessingJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProcessingJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProcessingJobs(options.Region), middleware.Before); err != nil {
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

// ListProcessingJobsAPIClient is a client that implements the ListProcessingJobs
// operation.
type ListProcessingJobsAPIClient interface {
	ListProcessingJobs(context.Context, *ListProcessingJobsInput, ...func(*Options)) (*ListProcessingJobsOutput, error)
}

var _ ListProcessingJobsAPIClient = (*Client)(nil)

// ListProcessingJobsPaginatorOptions is the paginator options for
// ListProcessingJobs
type ListProcessingJobsPaginatorOptions struct {
	// The maximum number of processing jobs to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProcessingJobsPaginator is a paginator for ListProcessingJobs
type ListProcessingJobsPaginator struct {
	options   ListProcessingJobsPaginatorOptions
	client    ListProcessingJobsAPIClient
	params    *ListProcessingJobsInput
	nextToken *string
	firstPage bool
}

// NewListProcessingJobsPaginator returns a new ListProcessingJobsPaginator
func NewListProcessingJobsPaginator(client ListProcessingJobsAPIClient, params *ListProcessingJobsInput, optFns ...func(*ListProcessingJobsPaginatorOptions)) *ListProcessingJobsPaginator {
	if params == nil {
		params = &ListProcessingJobsInput{}
	}

	options := ListProcessingJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProcessingJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProcessingJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProcessingJobs page.
func (p *ListProcessingJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProcessingJobsOutput, error) {
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

	result, err := p.client.ListProcessingJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProcessingJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListProcessingJobs",
	}
}
