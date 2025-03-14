// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about all the managed data identifiers that Amazon Macie
// currently provides.
func (c *Client) ListManagedDataIdentifiers(ctx context.Context, params *ListManagedDataIdentifiersInput, optFns ...func(*Options)) (*ListManagedDataIdentifiersOutput, error) {
	if params == nil {
		params = &ListManagedDataIdentifiersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListManagedDataIdentifiers", params, optFns, c.addOperationListManagedDataIdentifiersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListManagedDataIdentifiersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListManagedDataIdentifiersInput struct {

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListManagedDataIdentifiersOutput struct {

	// An array of objects, one for each managed data identifier.
	Items []types.ManagedDataIdentifierSummary

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListManagedDataIdentifiersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListManagedDataIdentifiers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListManagedDataIdentifiers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListManagedDataIdentifiers(options.Region), middleware.Before); err != nil {
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

// ListManagedDataIdentifiersAPIClient is a client that implements the
// ListManagedDataIdentifiers operation.
type ListManagedDataIdentifiersAPIClient interface {
	ListManagedDataIdentifiers(context.Context, *ListManagedDataIdentifiersInput, ...func(*Options)) (*ListManagedDataIdentifiersOutput, error)
}

var _ ListManagedDataIdentifiersAPIClient = (*Client)(nil)

// ListManagedDataIdentifiersPaginatorOptions is the paginator options for
// ListManagedDataIdentifiers
type ListManagedDataIdentifiersPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListManagedDataIdentifiersPaginator is a paginator for
// ListManagedDataIdentifiers
type ListManagedDataIdentifiersPaginator struct {
	options   ListManagedDataIdentifiersPaginatorOptions
	client    ListManagedDataIdentifiersAPIClient
	params    *ListManagedDataIdentifiersInput
	nextToken *string
	firstPage bool
}

// NewListManagedDataIdentifiersPaginator returns a new
// ListManagedDataIdentifiersPaginator
func NewListManagedDataIdentifiersPaginator(client ListManagedDataIdentifiersAPIClient, params *ListManagedDataIdentifiersInput, optFns ...func(*ListManagedDataIdentifiersPaginatorOptions)) *ListManagedDataIdentifiersPaginator {
	if params == nil {
		params = &ListManagedDataIdentifiersInput{}
	}

	options := ListManagedDataIdentifiersPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListManagedDataIdentifiersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListManagedDataIdentifiersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListManagedDataIdentifiers page.
func (p *ListManagedDataIdentifiersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListManagedDataIdentifiersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListManagedDataIdentifiers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListManagedDataIdentifiers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "ListManagedDataIdentifiers",
	}
}
