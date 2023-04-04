// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves metadata for all crawlers defined in the customer account.
func (c *Client) GetCrawlers(ctx context.Context, params *GetCrawlersInput, optFns ...func(*Options)) (*GetCrawlersOutput, error) {
	if params == nil {
		params = &GetCrawlersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCrawlers", params, optFns, c.addOperationGetCrawlersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCrawlersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCrawlersInput struct {

	// The number of crawlers to return on each call.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCrawlersOutput struct {

	// A list of crawler metadata.
	Crawlers []types.Crawler

	// A continuation token, if the returned list has not reached the end of those
	// defined in this customer account.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCrawlersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCrawlers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCrawlers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCrawlers(options.Region), middleware.Before); err != nil {
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

// GetCrawlersAPIClient is a client that implements the GetCrawlers operation.
type GetCrawlersAPIClient interface {
	GetCrawlers(context.Context, *GetCrawlersInput, ...func(*Options)) (*GetCrawlersOutput, error)
}

var _ GetCrawlersAPIClient = (*Client)(nil)

// GetCrawlersPaginatorOptions is the paginator options for GetCrawlers
type GetCrawlersPaginatorOptions struct {
	// The number of crawlers to return on each call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCrawlersPaginator is a paginator for GetCrawlers
type GetCrawlersPaginator struct {
	options   GetCrawlersPaginatorOptions
	client    GetCrawlersAPIClient
	params    *GetCrawlersInput
	nextToken *string
	firstPage bool
}

// NewGetCrawlersPaginator returns a new GetCrawlersPaginator
func NewGetCrawlersPaginator(client GetCrawlersAPIClient, params *GetCrawlersInput, optFns ...func(*GetCrawlersPaginatorOptions)) *GetCrawlersPaginator {
	if params == nil {
		params = &GetCrawlersInput{}
	}

	options := GetCrawlersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCrawlersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCrawlersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCrawlers page.
func (p *GetCrawlersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCrawlersOutput, error) {
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

	result, err := p.client.GetCrawlers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCrawlers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetCrawlers",
	}
}
