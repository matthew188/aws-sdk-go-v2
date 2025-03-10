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

// Searches contacts and lists the ones that meet a set of filter and sort
// criteria.
func (c *Client) SearchContacts(ctx context.Context, params *SearchContactsInput, optFns ...func(*Options)) (*SearchContactsOutput, error) {
	if params == nil {
		params = &SearchContactsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchContacts", params, optFns, c.addOperationSearchContactsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchContactsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchContactsInput struct {

	// The filters to use to list a specified set of address books. The supported
	// filter keys are DisplayName, FirstName, LastName, and AddressBookArns.
	Filters []types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response only
	// includes results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The sort order to use in listing the specified set of contacts. The supported
	// sort keys are DisplayName, FirstName, and LastName.
	SortCriteria []types.Sort

	noSmithyDocumentSerde
}

type SearchContactsOutput struct {

	// The contacts that meet the specified set of filter criteria, in sort order.
	Contacts []types.ContactData

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The total number of contacts returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchContactsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchContacts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchContacts{}, middleware.After)
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
	if err = addOpSearchContactsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchContacts(options.Region), middleware.Before); err != nil {
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

// SearchContactsAPIClient is a client that implements the SearchContacts
// operation.
type SearchContactsAPIClient interface {
	SearchContacts(context.Context, *SearchContactsInput, ...func(*Options)) (*SearchContactsOutput, error)
}

var _ SearchContactsAPIClient = (*Client)(nil)

// SearchContactsPaginatorOptions is the paginator options for SearchContacts
type SearchContactsPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchContactsPaginator is a paginator for SearchContacts
type SearchContactsPaginator struct {
	options   SearchContactsPaginatorOptions
	client    SearchContactsAPIClient
	params    *SearchContactsInput
	nextToken *string
	firstPage bool
}

// NewSearchContactsPaginator returns a new SearchContactsPaginator
func NewSearchContactsPaginator(client SearchContactsAPIClient, params *SearchContactsInput, optFns ...func(*SearchContactsPaginatorOptions)) *SearchContactsPaginator {
	if params == nil {
		params = &SearchContactsInput{}
	}

	options := SearchContactsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchContactsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchContactsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchContacts page.
func (p *SearchContactsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchContactsOutput, error) {
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

	result, err := p.client.SearchContacts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchContacts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchContacts",
	}
}
