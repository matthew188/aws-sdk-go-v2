// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotroborunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Grants permission to list sites
func (c *Client) ListSites(ctx context.Context, params *ListSitesInput, optFns ...func(*Options)) (*ListSitesOutput, error) {
	if params == nil {
		params = &ListSitesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSites", params, optFns, c.addOperationListSitesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSitesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSitesInput struct {

	// Maximum number of results to retrieve in a single ListSites call.
	MaxResults *int32

	// Pagination token returned when another page of data exists. Provide it in your
	// next call to the API to receive the next page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSitesOutput struct {

	// Pagination token returned when another page of data exists. Provide it in your
	// next call to the API to receive the next page.
	NextToken *string

	// List of facilities.
	Sites []types.Site

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSitesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSites{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSites{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSites(options.Region), middleware.Before); err != nil {
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

// ListSitesAPIClient is a client that implements the ListSites operation.
type ListSitesAPIClient interface {
	ListSites(context.Context, *ListSitesInput, ...func(*Options)) (*ListSitesOutput, error)
}

var _ ListSitesAPIClient = (*Client)(nil)

// ListSitesPaginatorOptions is the paginator options for ListSites
type ListSitesPaginatorOptions struct {
	// Maximum number of results to retrieve in a single ListSites call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSitesPaginator is a paginator for ListSites
type ListSitesPaginator struct {
	options   ListSitesPaginatorOptions
	client    ListSitesAPIClient
	params    *ListSitesInput
	nextToken *string
	firstPage bool
}

// NewListSitesPaginator returns a new ListSitesPaginator
func NewListSitesPaginator(client ListSitesAPIClient, params *ListSitesInput, optFns ...func(*ListSitesPaginatorOptions)) *ListSitesPaginator {
	if params == nil {
		params = &ListSitesInput{}
	}

	options := ListSitesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSitesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSitesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSites page.
func (p *ListSitesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSitesOutput, error) {
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

	result, err := p.client.ListSites(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSites(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotroborunner",
		OperationName: "ListSites",
	}
}
