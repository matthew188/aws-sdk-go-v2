// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of all data ingestion jobs, including dataset name and ARN, S3
// location of the input data, status, and so on.
func (c *Client) ListDataIngestionJobs(ctx context.Context, params *ListDataIngestionJobsInput, optFns ...func(*Options)) (*ListDataIngestionJobsOutput, error) {
	if params == nil {
		params = &ListDataIngestionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataIngestionJobs", params, optFns, c.addOperationListDataIngestionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataIngestionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataIngestionJobsInput struct {

	// The name of the dataset being used for the data ingestion job.
	DatasetName *string

	// Specifies the maximum number of data ingestion jobs to list.
	MaxResults *int32

	// An opaque pagination token indicating where to continue the listing of data
	// ingestion jobs.
	NextToken *string

	// Indicates the status of the data ingestion job.
	Status types.IngestionJobStatus

	noSmithyDocumentSerde
}

type ListDataIngestionJobsOutput struct {

	// Specifies information about the specific data ingestion job, including dataset
	// name and status.
	DataIngestionJobSummaries []types.DataIngestionJobSummary

	// An opaque pagination token indicating where to continue the listing of data
	// ingestion jobs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataIngestionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListDataIngestionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListDataIngestionJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataIngestionJobs(options.Region), middleware.Before); err != nil {
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

// ListDataIngestionJobsAPIClient is a client that implements the
// ListDataIngestionJobs operation.
type ListDataIngestionJobsAPIClient interface {
	ListDataIngestionJobs(context.Context, *ListDataIngestionJobsInput, ...func(*Options)) (*ListDataIngestionJobsOutput, error)
}

var _ ListDataIngestionJobsAPIClient = (*Client)(nil)

// ListDataIngestionJobsPaginatorOptions is the paginator options for
// ListDataIngestionJobs
type ListDataIngestionJobsPaginatorOptions struct {
	// Specifies the maximum number of data ingestion jobs to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataIngestionJobsPaginator is a paginator for ListDataIngestionJobs
type ListDataIngestionJobsPaginator struct {
	options   ListDataIngestionJobsPaginatorOptions
	client    ListDataIngestionJobsAPIClient
	params    *ListDataIngestionJobsInput
	nextToken *string
	firstPage bool
}

// NewListDataIngestionJobsPaginator returns a new ListDataIngestionJobsPaginator
func NewListDataIngestionJobsPaginator(client ListDataIngestionJobsAPIClient, params *ListDataIngestionJobsInput, optFns ...func(*ListDataIngestionJobsPaginatorOptions)) *ListDataIngestionJobsPaginator {
	if params == nil {
		params = &ListDataIngestionJobsInput{}
	}

	options := ListDataIngestionJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataIngestionJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataIngestionJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataIngestionJobs page.
func (p *ListDataIngestionJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataIngestionJobsOutput, error) {
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

	result, err := p.client.ListDataIngestionJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataIngestionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "ListDataIngestionJobs",
	}
}
