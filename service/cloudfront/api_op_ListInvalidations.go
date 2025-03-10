// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists invalidation batches.
func (c *Client) ListInvalidations(ctx context.Context, params *ListInvalidationsInput, optFns ...func(*Options)) (*ListInvalidationsOutput, error) {
	if params == nil {
		params = &ListInvalidationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInvalidations", params, optFns, c.addOperationListInvalidationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInvalidationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list invalidations.
type ListInvalidationsInput struct {

	// The distribution's ID.
	//
	// This member is required.
	DistributionId *string

	// Use this parameter when paginating results to indicate where to begin in your
	// list of invalidation batches. Because the results are returned in decreasing
	// order from most recent to oldest, the most recent results are on the first page,
	// the second page will contain earlier results, and so on. To get the next page of
	// results, set Marker to the value of the NextMarker from the current page's
	// response. This value is the same as the ID of the last invalidation batch on
	// that page.
	Marker *string

	// The maximum number of invalidation batches that you want in the response body.
	MaxItems *int32

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type ListInvalidationsOutput struct {

	// Information about invalidation batches.
	InvalidationList *types.InvalidationList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInvalidationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListInvalidations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListInvalidations{}, middleware.After)
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
	if err = addOpListInvalidationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInvalidations(options.Region), middleware.Before); err != nil {
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

// ListInvalidationsAPIClient is a client that implements the ListInvalidations
// operation.
type ListInvalidationsAPIClient interface {
	ListInvalidations(context.Context, *ListInvalidationsInput, ...func(*Options)) (*ListInvalidationsOutput, error)
}

var _ ListInvalidationsAPIClient = (*Client)(nil)

// ListInvalidationsPaginatorOptions is the paginator options for ListInvalidations
type ListInvalidationsPaginatorOptions struct {
	// The maximum number of invalidation batches that you want in the response body.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInvalidationsPaginator is a paginator for ListInvalidations
type ListInvalidationsPaginator struct {
	options   ListInvalidationsPaginatorOptions
	client    ListInvalidationsAPIClient
	params    *ListInvalidationsInput
	nextToken *string
	firstPage bool
}

// NewListInvalidationsPaginator returns a new ListInvalidationsPaginator
func NewListInvalidationsPaginator(client ListInvalidationsAPIClient, params *ListInvalidationsInput, optFns ...func(*ListInvalidationsPaginatorOptions)) *ListInvalidationsPaginator {
	if params == nil {
		params = &ListInvalidationsInput{}
	}

	options := ListInvalidationsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInvalidationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInvalidationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInvalidations page.
func (p *ListInvalidationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInvalidationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListInvalidations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = nil
	if result.InvalidationList != nil {
		p.nextToken = result.InvalidationList.NextMarker
	}

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListInvalidations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListInvalidations",
	}
}
