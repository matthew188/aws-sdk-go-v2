// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakergeospatial

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemakergeospatial/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to get a list of the Earth Observation jobs associated with
// the calling Amazon Web Services account.
func (c *Client) ListEarthObservationJobs(ctx context.Context, params *ListEarthObservationJobsInput, optFns ...func(*Options)) (*ListEarthObservationJobsOutput, error) {
	if params == nil {
		params = &ListEarthObservationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEarthObservationJobs", params, optFns, c.addOperationListEarthObservationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEarthObservationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEarthObservationJobsInput struct {

	// The total number of items to return.
	MaxResults *int32

	// If the previous response was truncated, you receive this token. Use it in your
	// next request to receive the next set of results.
	NextToken *string

	// The parameter by which to sort the results.
	SortBy *string

	// An optional value that specifies whether you want the results sorted in
	// Ascending or Descending order.
	SortOrder types.SortOrder

	// A filter that retrieves only jobs with a specific status.
	StatusEquals types.EarthObservationJobStatus

	noSmithyDocumentSerde
}

type ListEarthObservationJobsOutput struct {

	// Contains summary information about the Earth Observation jobs.
	//
	// This member is required.
	EarthObservationJobSummaries []types.ListEarthObservationJobOutputConfig

	// If the previous response was truncated, you receive this token. Use it in your
	// next request to receive the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEarthObservationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEarthObservationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEarthObservationJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEarthObservationJobs(options.Region), middleware.Before); err != nil {
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

// ListEarthObservationJobsAPIClient is a client that implements the
// ListEarthObservationJobs operation.
type ListEarthObservationJobsAPIClient interface {
	ListEarthObservationJobs(context.Context, *ListEarthObservationJobsInput, ...func(*Options)) (*ListEarthObservationJobsOutput, error)
}

var _ ListEarthObservationJobsAPIClient = (*Client)(nil)

// ListEarthObservationJobsPaginatorOptions is the paginator options for
// ListEarthObservationJobs
type ListEarthObservationJobsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEarthObservationJobsPaginator is a paginator for ListEarthObservationJobs
type ListEarthObservationJobsPaginator struct {
	options   ListEarthObservationJobsPaginatorOptions
	client    ListEarthObservationJobsAPIClient
	params    *ListEarthObservationJobsInput
	nextToken *string
	firstPage bool
}

// NewListEarthObservationJobsPaginator returns a new
// ListEarthObservationJobsPaginator
func NewListEarthObservationJobsPaginator(client ListEarthObservationJobsAPIClient, params *ListEarthObservationJobsInput, optFns ...func(*ListEarthObservationJobsPaginatorOptions)) *ListEarthObservationJobsPaginator {
	if params == nil {
		params = &ListEarthObservationJobsInput{}
	}

	options := ListEarthObservationJobsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEarthObservationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEarthObservationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEarthObservationJobs page.
func (p *ListEarthObservationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEarthObservationJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListEarthObservationJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEarthObservationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker-geospatial",
		OperationName: "ListEarthObservationJobs",
	}
}
