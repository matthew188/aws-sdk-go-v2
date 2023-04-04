// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists map resources in your Amazon Web Services account.
func (c *Client) ListMaps(ctx context.Context, params *ListMapsInput, optFns ...func(*Options)) (*ListMapsOutput, error) {
	if params == nil {
		params = &ListMapsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMaps", params, optFns, c.addOperationListMapsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMapsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMapsInput struct {

	// An optional limit for the number of resources returned in a single call. Default
	// value: 100
	MaxResults *int32

	// The pagination token specifying which page of results to return in the response.
	// If no token is provided, the default page is the first page. Default value: null
	NextToken *string

	noSmithyDocumentSerde
}

type ListMapsOutput struct {

	// Contains a list of maps in your Amazon Web Services account
	//
	// This member is required.
	Entries []types.ListMapsResponseEntry

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMapsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMaps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMaps{}, middleware.After)
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
	if err = addEndpointPrefix_opListMapsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMaps(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListMapsMiddleware struct {
}

func (*endpointPrefix_opListMapsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListMapsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "maps." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListMapsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListMapsMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListMapsAPIClient is a client that implements the ListMaps operation.
type ListMapsAPIClient interface {
	ListMaps(context.Context, *ListMapsInput, ...func(*Options)) (*ListMapsOutput, error)
}

var _ ListMapsAPIClient = (*Client)(nil)

// ListMapsPaginatorOptions is the paginator options for ListMaps
type ListMapsPaginatorOptions struct {
	// An optional limit for the number of resources returned in a single call. Default
	// value: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMapsPaginator is a paginator for ListMaps
type ListMapsPaginator struct {
	options   ListMapsPaginatorOptions
	client    ListMapsAPIClient
	params    *ListMapsInput
	nextToken *string
	firstPage bool
}

// NewListMapsPaginator returns a new ListMapsPaginator
func NewListMapsPaginator(client ListMapsAPIClient, params *ListMapsInput, optFns ...func(*ListMapsPaginatorOptions)) *ListMapsPaginator {
	if params == nil {
		params = &ListMapsInput{}
	}

	options := ListMapsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMapsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMapsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMaps page.
func (p *ListMapsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMapsOutput, error) {
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

	result, err := p.client.ListMaps(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMaps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "ListMaps",
	}
}
