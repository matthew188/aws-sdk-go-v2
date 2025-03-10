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

// Use this operation to get raster data collections.
func (c *Client) ListRasterDataCollections(ctx context.Context, params *ListRasterDataCollectionsInput, optFns ...func(*Options)) (*ListRasterDataCollectionsOutput, error) {
	if params == nil {
		params = &ListRasterDataCollectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRasterDataCollections", params, optFns, c.addOperationListRasterDataCollectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRasterDataCollectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRasterDataCollectionsInput struct {

	// The total number of items to return.
	MaxResults *int32

	// If the previous response was truncated, you receive this token. Use it in your
	// next request to receive the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRasterDataCollectionsOutput struct {

	// Contains summary information about the raster data collection.
	//
	// This member is required.
	RasterDataCollectionSummaries []types.RasterDataCollectionMetadata

	// If the previous response was truncated, you receive this token. Use it in your
	// next request to receive the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRasterDataCollectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRasterDataCollections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRasterDataCollections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRasterDataCollections(options.Region), middleware.Before); err != nil {
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

// ListRasterDataCollectionsAPIClient is a client that implements the
// ListRasterDataCollections operation.
type ListRasterDataCollectionsAPIClient interface {
	ListRasterDataCollections(context.Context, *ListRasterDataCollectionsInput, ...func(*Options)) (*ListRasterDataCollectionsOutput, error)
}

var _ ListRasterDataCollectionsAPIClient = (*Client)(nil)

// ListRasterDataCollectionsPaginatorOptions is the paginator options for
// ListRasterDataCollections
type ListRasterDataCollectionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRasterDataCollectionsPaginator is a paginator for ListRasterDataCollections
type ListRasterDataCollectionsPaginator struct {
	options   ListRasterDataCollectionsPaginatorOptions
	client    ListRasterDataCollectionsAPIClient
	params    *ListRasterDataCollectionsInput
	nextToken *string
	firstPage bool
}

// NewListRasterDataCollectionsPaginator returns a new
// ListRasterDataCollectionsPaginator
func NewListRasterDataCollectionsPaginator(client ListRasterDataCollectionsAPIClient, params *ListRasterDataCollectionsInput, optFns ...func(*ListRasterDataCollectionsPaginatorOptions)) *ListRasterDataCollectionsPaginator {
	if params == nil {
		params = &ListRasterDataCollectionsInput{}
	}

	options := ListRasterDataCollectionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRasterDataCollectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRasterDataCollectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRasterDataCollections page.
func (p *ListRasterDataCollectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRasterDataCollectionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListRasterDataCollections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRasterDataCollections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker-geospatial",
		OperationName: "ListRasterDataCollections",
	}
}
