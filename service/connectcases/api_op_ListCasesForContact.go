// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists cases for a given contact.
func (c *Client) ListCasesForContact(ctx context.Context, params *ListCasesForContactInput, optFns ...func(*Options)) (*ListCasesForContactOutput, error) {
	if params == nil {
		params = &ListCasesForContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCasesForContact", params, optFns, c.addOperationListCasesForContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCasesForContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCasesForContactInput struct {

	// A unique identifier of a contact in Amazon Connect.
	//
	// This member is required.
	ContactArn *string

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCasesForContactOutput struct {

	// A list of Case summary information.
	//
	// This member is required.
	Cases []types.CaseSummary

	// The token for the next set of results. This is null if there are no more results
	// to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCasesForContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCasesForContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCasesForContact{}, middleware.After)
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
	if err = addOpListCasesForContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCasesForContact(options.Region), middleware.Before); err != nil {
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

// ListCasesForContactAPIClient is a client that implements the ListCasesForContact
// operation.
type ListCasesForContactAPIClient interface {
	ListCasesForContact(context.Context, *ListCasesForContactInput, ...func(*Options)) (*ListCasesForContactOutput, error)
}

var _ ListCasesForContactAPIClient = (*Client)(nil)

// ListCasesForContactPaginatorOptions is the paginator options for
// ListCasesForContact
type ListCasesForContactPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCasesForContactPaginator is a paginator for ListCasesForContact
type ListCasesForContactPaginator struct {
	options   ListCasesForContactPaginatorOptions
	client    ListCasesForContactAPIClient
	params    *ListCasesForContactInput
	nextToken *string
	firstPage bool
}

// NewListCasesForContactPaginator returns a new ListCasesForContactPaginator
func NewListCasesForContactPaginator(client ListCasesForContactAPIClient, params *ListCasesForContactInput, optFns ...func(*ListCasesForContactPaginatorOptions)) *ListCasesForContactPaginator {
	if params == nil {
		params = &ListCasesForContactInput{}
	}

	options := ListCasesForContactPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCasesForContactPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCasesForContactPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCasesForContact page.
func (p *ListCasesForContactPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCasesForContactOutput, error) {
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

	result, err := p.client.ListCasesForContact(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCasesForContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cases",
		OperationName: "ListCasesForContact",
	}
}
