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

// Lists transform jobs.
func (c *Client) ListTransformJobs(ctx context.Context, params *ListTransformJobsInput, optFns ...func(*Options)) (*ListTransformJobsOutput, error) {
	if params == nil {
		params = &ListTransformJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTransformJobs", params, optFns, c.addOperationListTransformJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTransformJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTransformJobsInput struct {

	// A filter that returns only transform jobs created after the specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only transform jobs created before the specified time.
	CreationTimeBefore *time.Time

	// A filter that returns only transform jobs modified after the specified time.
	LastModifiedTimeAfter *time.Time

	// A filter that returns only transform jobs modified before the specified time.
	LastModifiedTimeBefore *time.Time

	// The maximum number of transform jobs to return in the response. The default
	// value is 10.
	MaxResults *int32

	// A string in the transform job name. This filter returns only transform jobs
	// whose name contains the specified string.
	NameContains *string

	// If the result of the previous ListTransformJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of transform jobs, use
	// the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime.
	SortBy types.SortBy

	// The sort order for results. The default is Descending.
	SortOrder types.SortOrder

	// A filter that retrieves only transform jobs with a specific status.
	StatusEquals types.TransformJobStatus

	noSmithyDocumentSerde
}

type ListTransformJobsOutput struct {

	// An array of TransformJobSummary objects.
	//
	// This member is required.
	TransformJobSummaries []types.TransformJobSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of transform jobs, use it in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTransformJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTransformJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTransformJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTransformJobs(options.Region), middleware.Before); err != nil {
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

// ListTransformJobsAPIClient is a client that implements the ListTransformJobs
// operation.
type ListTransformJobsAPIClient interface {
	ListTransformJobs(context.Context, *ListTransformJobsInput, ...func(*Options)) (*ListTransformJobsOutput, error)
}

var _ ListTransformJobsAPIClient = (*Client)(nil)

// ListTransformJobsPaginatorOptions is the paginator options for ListTransformJobs
type ListTransformJobsPaginatorOptions struct {
	// The maximum number of transform jobs to return in the response. The default
	// value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTransformJobsPaginator is a paginator for ListTransformJobs
type ListTransformJobsPaginator struct {
	options   ListTransformJobsPaginatorOptions
	client    ListTransformJobsAPIClient
	params    *ListTransformJobsInput
	nextToken *string
	firstPage bool
}

// NewListTransformJobsPaginator returns a new ListTransformJobsPaginator
func NewListTransformJobsPaginator(client ListTransformJobsAPIClient, params *ListTransformJobsInput, optFns ...func(*ListTransformJobsPaginatorOptions)) *ListTransformJobsPaginator {
	if params == nil {
		params = &ListTransformJobsInput{}
	}

	options := ListTransformJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTransformJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTransformJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTransformJobs page.
func (p *ListTransformJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTransformJobsOutput, error) {
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

	result, err := p.client.ListTransformJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTransformJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTransformJobs",
	}
}
