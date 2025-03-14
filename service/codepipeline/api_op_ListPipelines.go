// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a summary of all of the pipelines associated with your account.
func (c *Client) ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) {
	if params == nil {
		params = &ListPipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPipelines", params, optFns, c.addOperationListPipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListPipelines action.
type ListPipelinesInput struct {

	// The maximum number of pipelines to return in a single call. To retrieve the
	// remaining pipelines, make another call with the returned nextToken value. The
	// minimum value you can specify is 1. The maximum accepted value is 1000.
	MaxResults *int32

	// An identifier that was returned from the previous list pipelines call. It can be
	// used to return the next set of pipelines in the list.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the output of a ListPipelines action.
type ListPipelinesOutput struct {

	// If the amount of returned information is significantly large, an identifier is
	// also returned. It can be used in a subsequent list pipelines call to return the
	// next set of pipelines in the list.
	NextToken *string

	// The list of pipelines.
	Pipelines []types.PipelineSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPipelines{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPipelines(options.Region), middleware.Before); err != nil {
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

// ListPipelinesAPIClient is a client that implements the ListPipelines operation.
type ListPipelinesAPIClient interface {
	ListPipelines(context.Context, *ListPipelinesInput, ...func(*Options)) (*ListPipelinesOutput, error)
}

var _ ListPipelinesAPIClient = (*Client)(nil)

// ListPipelinesPaginatorOptions is the paginator options for ListPipelines
type ListPipelinesPaginatorOptions struct {
	// The maximum number of pipelines to return in a single call. To retrieve the
	// remaining pipelines, make another call with the returned nextToken value. The
	// minimum value you can specify is 1. The maximum accepted value is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPipelinesPaginator is a paginator for ListPipelines
type ListPipelinesPaginator struct {
	options   ListPipelinesPaginatorOptions
	client    ListPipelinesAPIClient
	params    *ListPipelinesInput
	nextToken *string
	firstPage bool
}

// NewListPipelinesPaginator returns a new ListPipelinesPaginator
func NewListPipelinesPaginator(client ListPipelinesAPIClient, params *ListPipelinesInput, optFns ...func(*ListPipelinesPaginatorOptions)) *ListPipelinesPaginator {
	if params == nil {
		params = &ListPipelinesInput{}
	}

	options := ListPipelinesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPipelinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPipelinesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPipelines page.
func (p *ListPipelinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPipelinesOutput, error) {
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

	result, err := p.client.ListPipelines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPipelines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "ListPipelines",
	}
}
