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

// Retrieves all waves or multiple waves by ID.
func (c *Client) ListWaves(ctx context.Context, params *ListWavesInput, optFns ...func(*Options)) (*ListWavesOutput, error) {
	if params == nil {
		params = &ListWavesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWaves", params, optFns, c.addOperationListWavesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWavesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWavesInput struct {

	// Waves list filters.
	Filters *types.ListWavesRequestFilters

	// Maximum results to return when listing waves.
	MaxResults int32

	// Request next token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWavesOutput struct {

	// Waves list.
	Items []types.Wave

	// Response next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWavesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWaves{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWaves{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWaves(options.Region), middleware.Before); err != nil {
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

// ListWavesAPIClient is a client that implements the ListWaves operation.
type ListWavesAPIClient interface {
	ListWaves(context.Context, *ListWavesInput, ...func(*Options)) (*ListWavesOutput, error)
}

var _ ListWavesAPIClient = (*Client)(nil)

// ListWavesPaginatorOptions is the paginator options for ListWaves
type ListWavesPaginatorOptions struct {
	// Maximum results to return when listing waves.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWavesPaginator is a paginator for ListWaves
type ListWavesPaginator struct {
	options   ListWavesPaginatorOptions
	client    ListWavesAPIClient
	params    *ListWavesInput
	nextToken *string
	firstPage bool
}

// NewListWavesPaginator returns a new ListWavesPaginator
func NewListWavesPaginator(client ListWavesAPIClient, params *ListWavesInput, optFns ...func(*ListWavesPaginatorOptions)) *ListWavesPaginator {
	if params == nil {
		params = &ListWavesInput{}
	}

	options := ListWavesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWavesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWavesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWaves page.
func (p *ListWavesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWavesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListWaves(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWaves(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "ListWaves",
	}
}
