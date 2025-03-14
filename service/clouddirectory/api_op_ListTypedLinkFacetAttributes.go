// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of all attribute definitions for a particular
// TypedLinkFacet. For more information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) ListTypedLinkFacetAttributes(ctx context.Context, params *ListTypedLinkFacetAttributesInput, optFns ...func(*Options)) (*ListTypedLinkFacetAttributesOutput, error) {
	if params == nil {
		params = &ListTypedLinkFacetAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTypedLinkFacetAttributes", params, optFns, c.addOperationListTypedLinkFacetAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTypedLinkFacetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTypedLinkFacetAttributesInput struct {

	// The unique name of the typed link facet.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// This member is required.
	SchemaArn *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTypedLinkFacetAttributesOutput struct {

	// An ordered set of attributes associate with the typed link.
	Attributes []types.TypedLinkAttributeDefinition

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTypedLinkFacetAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTypedLinkFacetAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTypedLinkFacetAttributes{}, middleware.After)
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
	if err = addOpListTypedLinkFacetAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTypedLinkFacetAttributes(options.Region), middleware.Before); err != nil {
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

// ListTypedLinkFacetAttributesAPIClient is a client that implements the
// ListTypedLinkFacetAttributes operation.
type ListTypedLinkFacetAttributesAPIClient interface {
	ListTypedLinkFacetAttributes(context.Context, *ListTypedLinkFacetAttributesInput, ...func(*Options)) (*ListTypedLinkFacetAttributesOutput, error)
}

var _ ListTypedLinkFacetAttributesAPIClient = (*Client)(nil)

// ListTypedLinkFacetAttributesPaginatorOptions is the paginator options for
// ListTypedLinkFacetAttributes
type ListTypedLinkFacetAttributesPaginatorOptions struct {
	// The maximum number of results to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTypedLinkFacetAttributesPaginator is a paginator for
// ListTypedLinkFacetAttributes
type ListTypedLinkFacetAttributesPaginator struct {
	options   ListTypedLinkFacetAttributesPaginatorOptions
	client    ListTypedLinkFacetAttributesAPIClient
	params    *ListTypedLinkFacetAttributesInput
	nextToken *string
	firstPage bool
}

// NewListTypedLinkFacetAttributesPaginator returns a new
// ListTypedLinkFacetAttributesPaginator
func NewListTypedLinkFacetAttributesPaginator(client ListTypedLinkFacetAttributesAPIClient, params *ListTypedLinkFacetAttributesInput, optFns ...func(*ListTypedLinkFacetAttributesPaginatorOptions)) *ListTypedLinkFacetAttributesPaginator {
	if params == nil {
		params = &ListTypedLinkFacetAttributesInput{}
	}

	options := ListTypedLinkFacetAttributesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTypedLinkFacetAttributesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTypedLinkFacetAttributesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTypedLinkFacetAttributes page.
func (p *ListTypedLinkFacetAttributesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTypedLinkFacetAttributesOutput, error) {
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

	result, err := p.client.ListTypedLinkFacetAttributes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTypedLinkFacetAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListTypedLinkFacetAttributes",
	}
}
