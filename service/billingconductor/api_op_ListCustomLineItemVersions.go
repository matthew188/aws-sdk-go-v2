// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A paginated call to get a list of all custom line item versions.
func (c *Client) ListCustomLineItemVersions(ctx context.Context, params *ListCustomLineItemVersionsInput, optFns ...func(*Options)) (*ListCustomLineItemVersionsOutput, error) {
	if params == nil {
		params = &ListCustomLineItemVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomLineItemVersions", params, optFns, c.addOperationListCustomLineItemVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomLineItemVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCustomLineItemVersionsInput struct {

	// The Amazon Resource Name (ARN) for the custom line item.
	//
	// This member is required.
	Arn *string

	// A ListCustomLineItemVersionsFilter that specifies the billing period range in
	// which the custom line item versions are applied.
	Filters *types.ListCustomLineItemVersionsFilter

	// The maximum number of custom line item versions to retrieve.
	MaxResults *int32

	// The pagination token that's used on subsequent calls to retrieve custom line
	// item versions.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCustomLineItemVersionsOutput struct {

	// A list of CustomLineItemVersionListElements that are received.
	CustomLineItemVersions []types.CustomLineItemVersionListElement

	// The pagination token that's used on subsequent calls to retrieve custom line
	// item versions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomLineItemVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomLineItemVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomLineItemVersions{}, middleware.After)
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
	if err = addOpListCustomLineItemVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomLineItemVersions(options.Region), middleware.Before); err != nil {
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

// ListCustomLineItemVersionsAPIClient is a client that implements the
// ListCustomLineItemVersions operation.
type ListCustomLineItemVersionsAPIClient interface {
	ListCustomLineItemVersions(context.Context, *ListCustomLineItemVersionsInput, ...func(*Options)) (*ListCustomLineItemVersionsOutput, error)
}

var _ ListCustomLineItemVersionsAPIClient = (*Client)(nil)

// ListCustomLineItemVersionsPaginatorOptions is the paginator options for
// ListCustomLineItemVersions
type ListCustomLineItemVersionsPaginatorOptions struct {
	// The maximum number of custom line item versions to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomLineItemVersionsPaginator is a paginator for
// ListCustomLineItemVersions
type ListCustomLineItemVersionsPaginator struct {
	options   ListCustomLineItemVersionsPaginatorOptions
	client    ListCustomLineItemVersionsAPIClient
	params    *ListCustomLineItemVersionsInput
	nextToken *string
	firstPage bool
}

// NewListCustomLineItemVersionsPaginator returns a new
// ListCustomLineItemVersionsPaginator
func NewListCustomLineItemVersionsPaginator(client ListCustomLineItemVersionsAPIClient, params *ListCustomLineItemVersionsInput, optFns ...func(*ListCustomLineItemVersionsPaginatorOptions)) *ListCustomLineItemVersionsPaginator {
	if params == nil {
		params = &ListCustomLineItemVersionsInput{}
	}

	options := ListCustomLineItemVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomLineItemVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomLineItemVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCustomLineItemVersions page.
func (p *ListCustomLineItemVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomLineItemVersionsOutput, error) {
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

	result, err := p.client.ListCustomLineItemVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCustomLineItemVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "ListCustomLineItemVersions",
	}
}
