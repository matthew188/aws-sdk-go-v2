// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of sequence stores.
func (c *Client) ListSequenceStores(ctx context.Context, params *ListSequenceStoresInput, optFns ...func(*Options)) (*ListSequenceStoresOutput, error) {
	if params == nil {
		params = &ListSequenceStoresInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSequenceStores", params, optFns, c.addOperationListSequenceStoresMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSequenceStoresOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSequenceStoresInput struct {

	// A filter to apply to the list.
	Filter *types.SequenceStoreFilter

	// The maximum number of stores to return in one page of results.
	MaxResults *int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSequenceStoresOutput struct {

	// A list of sequence stores.
	//
	// This member is required.
	SequenceStores []types.SequenceStoreDetail

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSequenceStoresMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSequenceStores{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSequenceStores{}, middleware.After)
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
	if err = addEndpointPrefix_opListSequenceStoresMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSequenceStores(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListSequenceStoresMiddleware struct {
}

func (*endpointPrefix_opListSequenceStoresMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListSequenceStoresMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListSequenceStoresMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListSequenceStoresMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListSequenceStoresAPIClient is a client that implements the ListSequenceStores
// operation.
type ListSequenceStoresAPIClient interface {
	ListSequenceStores(context.Context, *ListSequenceStoresInput, ...func(*Options)) (*ListSequenceStoresOutput, error)
}

var _ ListSequenceStoresAPIClient = (*Client)(nil)

// ListSequenceStoresPaginatorOptions is the paginator options for
// ListSequenceStores
type ListSequenceStoresPaginatorOptions struct {
	// The maximum number of stores to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSequenceStoresPaginator is a paginator for ListSequenceStores
type ListSequenceStoresPaginator struct {
	options   ListSequenceStoresPaginatorOptions
	client    ListSequenceStoresAPIClient
	params    *ListSequenceStoresInput
	nextToken *string
	firstPage bool
}

// NewListSequenceStoresPaginator returns a new ListSequenceStoresPaginator
func NewListSequenceStoresPaginator(client ListSequenceStoresAPIClient, params *ListSequenceStoresInput, optFns ...func(*ListSequenceStoresPaginatorOptions)) *ListSequenceStoresPaginator {
	if params == nil {
		params = &ListSequenceStoresInput{}
	}

	options := ListSequenceStoresPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSequenceStoresPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSequenceStoresPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSequenceStores page.
func (p *ListSequenceStoresPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSequenceStoresOutput, error) {
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

	result, err := p.client.ListSequenceStores(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSequenceStores(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "ListSequenceStores",
	}
}
