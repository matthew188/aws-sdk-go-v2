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

// Retrieves a paginated list of associated assets. You can use this operation to
// do the following:
//
// * List child assets associated to a parent asset by a
// hierarchy that you specify.
//
// * List an asset's parent asset.
func (c *Client) ListAssociatedAssets(ctx context.Context, params *ListAssociatedAssetsInput, optFns ...func(*Options)) (*ListAssociatedAssetsOutput, error) {
	if params == nil {
		params = &ListAssociatedAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssociatedAssets", params, optFns, c.addOperationListAssociatedAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssociatedAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociatedAssetsInput struct {

	// The ID of the asset to query.
	//
	// This member is required.
	AssetId *string

	// The ID of the hierarchy by which child assets are associated to the asset. To
	// find a hierarchy ID, use the DescribeAsset
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAsset.html)
	// or DescribeAssetModel
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeAssetModel.html)
	// operations. This parameter is required if you choose CHILD for
	// traversalDirection. For more information, see Asset hierarchies
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the IoT SiteWise User Guide.
	HierarchyId *string

	// The maximum number of results to return for each paginated request. Default: 50
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	// The direction to list associated assets. Choose one of the following options:
	//
	// *
	// CHILD – The list includes all child assets associated to the asset. The
	// hierarchyId parameter is required if you choose CHILD.
	//
	// * PARENT – The list
	// includes the asset's parent asset.
	//
	// Default: CHILD
	TraversalDirection types.TraversalDirection

	noSmithyDocumentSerde
}

type ListAssociatedAssetsOutput struct {

	// A list that summarizes the associated assets.
	//
	// This member is required.
	AssetSummaries []types.AssociatedAssetsSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssociatedAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssociatedAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssociatedAssets{}, middleware.After)
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
	if err = addEndpointPrefix_opListAssociatedAssetsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListAssociatedAssetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociatedAssets(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListAssociatedAssetsMiddleware struct {
}

func (*endpointPrefix_opListAssociatedAssetsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListAssociatedAssetsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addEndpointPrefix_opListAssociatedAssetsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListAssociatedAssetsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListAssociatedAssetsAPIClient is a client that implements the
// ListAssociatedAssets operation.
type ListAssociatedAssetsAPIClient interface {
	ListAssociatedAssets(context.Context, *ListAssociatedAssetsInput, ...func(*Options)) (*ListAssociatedAssetsOutput, error)
}

var _ ListAssociatedAssetsAPIClient = (*Client)(nil)

// ListAssociatedAssetsPaginatorOptions is the paginator options for
// ListAssociatedAssets
type ListAssociatedAssetsPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. Default: 50
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssociatedAssetsPaginator is a paginator for ListAssociatedAssets
type ListAssociatedAssetsPaginator struct {
	options   ListAssociatedAssetsPaginatorOptions
	client    ListAssociatedAssetsAPIClient
	params    *ListAssociatedAssetsInput
	nextToken *string
	firstPage bool
}

// NewListAssociatedAssetsPaginator returns a new ListAssociatedAssetsPaginator
func NewListAssociatedAssetsPaginator(client ListAssociatedAssetsAPIClient, params *ListAssociatedAssetsInput, optFns ...func(*ListAssociatedAssetsPaginatorOptions)) *ListAssociatedAssetsPaginator {
	if params == nil {
		params = &ListAssociatedAssetsInput{}
	}

	options := ListAssociatedAssetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssociatedAssetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssociatedAssetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssociatedAssets page.
func (p *ListAssociatedAssetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssociatedAssetsOutput, error) {
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

	result, err := p.client.ListAssociatedAssets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssociatedAssets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListAssociatedAssets",
	}
}
