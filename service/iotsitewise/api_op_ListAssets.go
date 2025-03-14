// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of asset summaries. You can use this operation to do
// the following:
//
// * List assets based on a specific asset model.
//
// * List top-level
// assets.
//
// You can't use this operation to list all assets. To retrieve summaries
// for all of your assets, use ListAssetModels
// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_ListAssetModels.html)
// to get all of your asset model IDs. Then, use ListAssets to get all assets for
// each asset model.
func (c *Client) ListAssets(ctx context.Context, params *ListAssetsInput, optFns ...func(*Options)) (*ListAssetsOutput, error) {
	if params == nil {
		params = &ListAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssets", params, optFns, c.addOperationListAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssetsInput struct {

	// The ID of the asset model by which to filter the list of assets. This parameter
	// is required if you choose ALL for filter.
	AssetModelId *string

	// The filter for the requested list of assets. Choose one of the following
	// options:
	//
	// * ALL – The list includes all assets for a given asset model ID. The
	// assetModelId parameter is required if you filter by ALL.
	//
	// * TOP_LEVEL – The list
	// includes only top-level assets in the asset hierarchy tree.
	//
	// Default: ALL
	Filter types.ListAssetsFilter

	// The maximum number of results to return for each paginated request. Default: 50
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssetsOutput struct {

	// A list that summarizes each asset.
	//
	// This member is required.
	AssetSummaries []types.AssetSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssets{}, middleware.After)
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
	if err = addEndpointPrefix_opListAssetsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssets(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListAssetsMiddleware struct {
}

func (*endpointPrefix_opListAssetsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAssetsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListAssetsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListAssetsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListAssetsAPIClient is a client that implements the ListAssets operation.
type ListAssetsAPIClient interface {
	ListAssets(context.Context, *ListAssetsInput, ...func(*Options)) (*ListAssetsOutput, error)
}

var _ ListAssetsAPIClient = (*Client)(nil)

// ListAssetsPaginatorOptions is the paginator options for ListAssets
type ListAssetsPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. Default: 50
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssetsPaginator is a paginator for ListAssets
type ListAssetsPaginator struct {
	options   ListAssetsPaginatorOptions
	client    ListAssetsAPIClient
	params    *ListAssetsInput
	nextToken *string
	firstPage bool
}

// NewListAssetsPaginator returns a new ListAssetsPaginator
func NewListAssetsPaginator(client ListAssetsAPIClient, params *ListAssetsInput, optFns ...func(*ListAssetsPaginatorOptions)) *ListAssetsPaginator {
	if params == nil {
		params = &ListAssetsInput{}
	}

	options := ListAssetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssets page.
func (p *ListAssetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssetsOutput, error) {
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

	result, err := p.client.ListAssets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListAssets",
	}
}
