// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists world export jobs.
func (c *Client) ListWorldExportJobs(ctx context.Context, params *ListWorldExportJobsInput, optFns ...func(*Options)) (*ListWorldExportJobsOutput, error) {
	if params == nil {
		params = &ListWorldExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorldExportJobs", params, optFns, c.addOperationListWorldExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorldExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorldExportJobsInput struct {

	// Optional filters to limit results. You can use generationJobId and templateId.
	Filters []types.Filter

	// When this parameter is used, ListWorldExportJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListWorldExportJobs
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter is not used, then ListWorldExportJobs returns up to 100
	// results and a nextToken value if applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorldExportJobs again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorldExportJobsOutput struct {

	// Summary information for world export jobs.
	//
	// This member is required.
	WorldExportJobSummaries []types.WorldExportJobSummary

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorldExportJobsRequest again and assign that
	// token to the request object's nextToken parameter. If there are no remaining
	// results, the previous response object's NextToken parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorldExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorldExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorldExportJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorldExportJobs(options.Region), middleware.Before); err != nil {
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

// ListWorldExportJobsAPIClient is a client that implements the ListWorldExportJobs
// operation.
type ListWorldExportJobsAPIClient interface {
	ListWorldExportJobs(context.Context, *ListWorldExportJobsInput, ...func(*Options)) (*ListWorldExportJobsOutput, error)
}

var _ ListWorldExportJobsAPIClient = (*Client)(nil)

// ListWorldExportJobsPaginatorOptions is the paginator options for
// ListWorldExportJobs
type ListWorldExportJobsPaginatorOptions struct {
	// When this parameter is used, ListWorldExportJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListWorldExportJobs
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter is not used, then ListWorldExportJobs returns up to 100
	// results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorldExportJobsPaginator is a paginator for ListWorldExportJobs
type ListWorldExportJobsPaginator struct {
	options   ListWorldExportJobsPaginatorOptions
	client    ListWorldExportJobsAPIClient
	params    *ListWorldExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListWorldExportJobsPaginator returns a new ListWorldExportJobsPaginator
func NewListWorldExportJobsPaginator(client ListWorldExportJobsAPIClient, params *ListWorldExportJobsInput, optFns ...func(*ListWorldExportJobsPaginatorOptions)) *ListWorldExportJobsPaginator {
	if params == nil {
		params = &ListWorldExportJobsInput{}
	}

	options := ListWorldExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorldExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorldExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorldExportJobs page.
func (p *ListWorldExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorldExportJobsOutput, error) {
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

	result, err := p.client.ListWorldExportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorldExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListWorldExportJobs",
	}
}
