// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List all of the dedicated IP pools that exist in your Amazon Web Services
// account in the current Region.
func (c *Client) ListDedicatedIpPools(ctx context.Context, params *ListDedicatedIpPoolsInput, optFns ...func(*Options)) (*ListDedicatedIpPoolsOutput, error) {
	if params == nil {
		params = &ListDedicatedIpPoolsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDedicatedIpPools", params, optFns, c.addOperationListDedicatedIpPoolsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDedicatedIpPoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain a list of dedicated IP pools.
type ListDedicatedIpPoolsInput struct {

	// A token returned from a previous call to ListDedicatedIpPools to indicate the
	// position in the list of dedicated IP pools.
	NextToken *string

	// The number of results to show in a single call to ListDedicatedIpPools. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32

	noSmithyDocumentSerde
}

// A list of dedicated IP pools.
type ListDedicatedIpPoolsOutput struct {

	// A list of all of the dedicated IP pools that are associated with your Amazon Web
	// Services account in the current Region.
	DedicatedIpPools []string

	// A token that indicates that there are additional IP pools to list. To view
	// additional IP pools, issue another request to ListDedicatedIpPools, passing this
	// token in the NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDedicatedIpPoolsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDedicatedIpPools{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDedicatedIpPools{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDedicatedIpPools(options.Region), middleware.Before); err != nil {
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

// ListDedicatedIpPoolsAPIClient is a client that implements the
// ListDedicatedIpPools operation.
type ListDedicatedIpPoolsAPIClient interface {
	ListDedicatedIpPools(context.Context, *ListDedicatedIpPoolsInput, ...func(*Options)) (*ListDedicatedIpPoolsOutput, error)
}

var _ ListDedicatedIpPoolsAPIClient = (*Client)(nil)

// ListDedicatedIpPoolsPaginatorOptions is the paginator options for
// ListDedicatedIpPools
type ListDedicatedIpPoolsPaginatorOptions struct {
	// The number of results to show in a single call to ListDedicatedIpPools. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDedicatedIpPoolsPaginator is a paginator for ListDedicatedIpPools
type ListDedicatedIpPoolsPaginator struct {
	options   ListDedicatedIpPoolsPaginatorOptions
	client    ListDedicatedIpPoolsAPIClient
	params    *ListDedicatedIpPoolsInput
	nextToken *string
	firstPage bool
}

// NewListDedicatedIpPoolsPaginator returns a new ListDedicatedIpPoolsPaginator
func NewListDedicatedIpPoolsPaginator(client ListDedicatedIpPoolsAPIClient, params *ListDedicatedIpPoolsInput, optFns ...func(*ListDedicatedIpPoolsPaginatorOptions)) *ListDedicatedIpPoolsPaginator {
	if params == nil {
		params = &ListDedicatedIpPoolsInput{}
	}

	options := ListDedicatedIpPoolsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDedicatedIpPoolsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDedicatedIpPoolsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDedicatedIpPools page.
func (p *ListDedicatedIpPoolsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDedicatedIpPoolsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListDedicatedIpPools(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDedicatedIpPools(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListDedicatedIpPools",
	}
}
