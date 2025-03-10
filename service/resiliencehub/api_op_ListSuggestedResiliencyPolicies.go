// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the suggested resiliency policies for the Resilience Hub applications.
func (c *Client) ListSuggestedResiliencyPolicies(ctx context.Context, params *ListSuggestedResiliencyPoliciesInput, optFns ...func(*Options)) (*ListSuggestedResiliencyPoliciesOutput, error) {
	if params == nil {
		params = &ListSuggestedResiliencyPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSuggestedResiliencyPolicies", params, optFns, c.addOperationListSuggestedResiliencyPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSuggestedResiliencyPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSuggestedResiliencyPoliciesInput struct {

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// Null, or the token from a previous call to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSuggestedResiliencyPoliciesOutput struct {

	// The suggested resiliency policies for the Resilience Hub applications.
	//
	// This member is required.
	ResiliencyPolicies []types.ResiliencyPolicy

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSuggestedResiliencyPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSuggestedResiliencyPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSuggestedResiliencyPolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSuggestedResiliencyPolicies(options.Region), middleware.Before); err != nil {
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

// ListSuggestedResiliencyPoliciesAPIClient is a client that implements the
// ListSuggestedResiliencyPolicies operation.
type ListSuggestedResiliencyPoliciesAPIClient interface {
	ListSuggestedResiliencyPolicies(context.Context, *ListSuggestedResiliencyPoliciesInput, ...func(*Options)) (*ListSuggestedResiliencyPoliciesOutput, error)
}

var _ ListSuggestedResiliencyPoliciesAPIClient = (*Client)(nil)

// ListSuggestedResiliencyPoliciesPaginatorOptions is the paginator options for
// ListSuggestedResiliencyPolicies
type ListSuggestedResiliencyPoliciesPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSuggestedResiliencyPoliciesPaginator is a paginator for
// ListSuggestedResiliencyPolicies
type ListSuggestedResiliencyPoliciesPaginator struct {
	options   ListSuggestedResiliencyPoliciesPaginatorOptions
	client    ListSuggestedResiliencyPoliciesAPIClient
	params    *ListSuggestedResiliencyPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListSuggestedResiliencyPoliciesPaginator returns a new
// ListSuggestedResiliencyPoliciesPaginator
func NewListSuggestedResiliencyPoliciesPaginator(client ListSuggestedResiliencyPoliciesAPIClient, params *ListSuggestedResiliencyPoliciesInput, optFns ...func(*ListSuggestedResiliencyPoliciesPaginatorOptions)) *ListSuggestedResiliencyPoliciesPaginator {
	if params == nil {
		params = &ListSuggestedResiliencyPoliciesInput{}
	}

	options := ListSuggestedResiliencyPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSuggestedResiliencyPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSuggestedResiliencyPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSuggestedResiliencyPolicies page.
func (p *ListSuggestedResiliencyPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSuggestedResiliencyPoliciesOutput, error) {
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

	result, err := p.client.ListSuggestedResiliencyPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSuggestedResiliencyPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "ListSuggestedResiliencyPolicies",
	}
}
