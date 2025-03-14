// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists conference providers under a specific AWS account.
func (c *Client) ListConferenceProviders(ctx context.Context, params *ListConferenceProvidersInput, optFns ...func(*Options)) (*ListConferenceProvidersOutput, error) {
	if params == nil {
		params = &ListConferenceProvidersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConferenceProviders", params, optFns, c.addOperationListConferenceProvidersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConferenceProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConferenceProvidersInput struct {

	// The maximum number of conference providers to be returned, per paginated calls.
	MaxResults *int32

	// The tokens used for pagination.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConferenceProvidersOutput struct {

	// The conference providers.
	ConferenceProviders []types.ConferenceProvider

	// The tokens used for pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConferenceProvidersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListConferenceProviders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConferenceProviders{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConferenceProviders(options.Region), middleware.Before); err != nil {
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

// ListConferenceProvidersAPIClient is a client that implements the
// ListConferenceProviders operation.
type ListConferenceProvidersAPIClient interface {
	ListConferenceProviders(context.Context, *ListConferenceProvidersInput, ...func(*Options)) (*ListConferenceProvidersOutput, error)
}

var _ ListConferenceProvidersAPIClient = (*Client)(nil)

// ListConferenceProvidersPaginatorOptions is the paginator options for
// ListConferenceProviders
type ListConferenceProvidersPaginatorOptions struct {
	// The maximum number of conference providers to be returned, per paginated calls.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConferenceProvidersPaginator is a paginator for ListConferenceProviders
type ListConferenceProvidersPaginator struct {
	options   ListConferenceProvidersPaginatorOptions
	client    ListConferenceProvidersAPIClient
	params    *ListConferenceProvidersInput
	nextToken *string
	firstPage bool
}

// NewListConferenceProvidersPaginator returns a new
// ListConferenceProvidersPaginator
func NewListConferenceProvidersPaginator(client ListConferenceProvidersAPIClient, params *ListConferenceProvidersInput, optFns ...func(*ListConferenceProvidersPaginatorOptions)) *ListConferenceProvidersPaginator {
	if params == nil {
		params = &ListConferenceProvidersInput{}
	}

	options := ListConferenceProvidersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConferenceProvidersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConferenceProvidersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConferenceProviders page.
func (p *ListConferenceProvidersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConferenceProvidersOutput, error) {
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

	result, err := p.client.ListConferenceProviders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConferenceProviders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListConferenceProviders",
	}
}
