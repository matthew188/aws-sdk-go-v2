// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Jobs. Use the JobsID and fromDate and toData filters to limit
// which jobs are returned. The response is sorted by creationDataTime - latest
// date first. Jobs are normally created by the StartTest, StartCutover, and
// TerminateTargetInstances APIs. Jobs are also created by DiagnosticLaunch and
// TerminateDiagnosticInstances, which are APIs available only to *Support* and
// only used in response to relevant support tickets.
func (c *Client) DescribeJobs(ctx context.Context, params *DescribeJobsInput, optFns ...func(*Options)) (*DescribeJobsOutput, error) {
	if params == nil {
		params = &DescribeJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobs", params, optFns, c.addOperationDescribeJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobsInput struct {

	// Request to describe Job log filters.
	Filters *types.DescribeJobsRequestFilters

	// Request to describe job log items by max results.
	MaxResults int32

	// Request to describe job log items by next token.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeJobsOutput struct {

	// Request to describe Job log items.
	Items []types.Job

	// Request to describe Job response by next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobs(options.Region), middleware.Before); err != nil {
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

// DescribeJobsAPIClient is a client that implements the DescribeJobs operation.
type DescribeJobsAPIClient interface {
	DescribeJobs(context.Context, *DescribeJobsInput, ...func(*Options)) (*DescribeJobsOutput, error)
}

var _ DescribeJobsAPIClient = (*Client)(nil)

// DescribeJobsPaginatorOptions is the paginator options for DescribeJobs
type DescribeJobsPaginatorOptions struct {
	// Request to describe job log items by max results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeJobsPaginator is a paginator for DescribeJobs
type DescribeJobsPaginator struct {
	options   DescribeJobsPaginatorOptions
	client    DescribeJobsAPIClient
	params    *DescribeJobsInput
	nextToken *string
	firstPage bool
}

// NewDescribeJobsPaginator returns a new DescribeJobsPaginator
func NewDescribeJobsPaginator(client DescribeJobsAPIClient, params *DescribeJobsInput, optFns ...func(*DescribeJobsPaginatorOptions)) *DescribeJobsPaginator {
	if params == nil {
		params = &DescribeJobsInput{}
	}

	options := DescribeJobsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeJobs page.
func (p *DescribeJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "DescribeJobs",
	}
}
