// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of JobListEntry objects of the specified length. Each
// JobListEntry object is for a job in the specified cluster and contains a job's
// state, a job's ID, and other information.
func (c *Client) ListClusterJobs(ctx context.Context, params *ListClusterJobsInput, optFns ...func(*Options)) (*ListClusterJobsOutput, error) {
	if params == nil {
		params = &ListClusterJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListClusterJobs", params, optFns, c.addOperationListClusterJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListClusterJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListClusterJobsInput struct {

	// The 39-character ID for the cluster that you want to list, for example
	// CID123e4567-e89b-12d3-a456-426655440000.
	//
	// This member is required.
	ClusterId *string

	// The number of JobListEntry objects to return.
	MaxResults *int32

	// HTTP requests are stateless. To identify what object comes "next" in the list of
	// JobListEntry objects, you have the option of specifying NextToken as the
	// starting point for your returned list.
	NextToken *string

	noSmithyDocumentSerde
}

type ListClusterJobsOutput struct {

	// Each JobListEntry object contains a job's state, a job's ID, and a value that
	// indicates whether the job is a job part, in the case of export jobs.
	JobListEntries []types.JobListEntry

	// HTTP requests are stateless. If you use the automatically generated NextToken
	// value in your next ListClusterJobsResult call, your list of returned jobs will
	// start from this point in the array.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListClusterJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListClusterJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListClusterJobs{}, middleware.After)
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
	if err = addOpListClusterJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListClusterJobs(options.Region), middleware.Before); err != nil {
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

// ListClusterJobsAPIClient is a client that implements the ListClusterJobs
// operation.
type ListClusterJobsAPIClient interface {
	ListClusterJobs(context.Context, *ListClusterJobsInput, ...func(*Options)) (*ListClusterJobsOutput, error)
}

var _ ListClusterJobsAPIClient = (*Client)(nil)

// ListClusterJobsPaginatorOptions is the paginator options for ListClusterJobs
type ListClusterJobsPaginatorOptions struct {
	// The number of JobListEntry objects to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListClusterJobsPaginator is a paginator for ListClusterJobs
type ListClusterJobsPaginator struct {
	options   ListClusterJobsPaginatorOptions
	client    ListClusterJobsAPIClient
	params    *ListClusterJobsInput
	nextToken *string
	firstPage bool
}

// NewListClusterJobsPaginator returns a new ListClusterJobsPaginator
func NewListClusterJobsPaginator(client ListClusterJobsAPIClient, params *ListClusterJobsInput, optFns ...func(*ListClusterJobsPaginatorOptions)) *ListClusterJobsPaginator {
	if params == nil {
		params = &ListClusterJobsInput{}
	}

	options := ListClusterJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListClusterJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListClusterJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListClusterJobs page.
func (p *ListClusterJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListClusterJobsOutput, error) {
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

	result, err := p.client.ListClusterJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListClusterJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "ListClusterJobs",
	}
}
