// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of all log groups that are being monitored and tagged by DevOps
// Guru.
func (c *Client) ListMonitoredResources(ctx context.Context, params *ListMonitoredResourcesInput, optFns ...func(*Options)) (*ListMonitoredResourcesOutput, error) {
	if params == nil {
		params = &ListMonitoredResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMonitoredResources", params, optFns, c.addOperationListMonitoredResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMonitoredResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMonitoredResourcesInput struct {

	// Filters to determine which monitored resources you want to retrieve. You can
	// filter by resource type or resource permission status.
	Filters *types.ListMonitoredResourcesFilters

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMonitoredResourcesOutput struct {

	// Information about the resource that is being monitored, including the name of
	// the resource, the type of resource, and whether or not permission is given to
	// DevOps Guru to access that resource.
	//
	// This member is required.
	MonitoredResourceIdentifiers []types.MonitoredResourceIdentifier

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMonitoredResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMonitoredResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMonitoredResources{}, middleware.After)
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
	if err = addOpListMonitoredResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMonitoredResources(options.Region), middleware.Before); err != nil {
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

// ListMonitoredResourcesAPIClient is a client that implements the
// ListMonitoredResources operation.
type ListMonitoredResourcesAPIClient interface {
	ListMonitoredResources(context.Context, *ListMonitoredResourcesInput, ...func(*Options)) (*ListMonitoredResourcesOutput, error)
}

var _ ListMonitoredResourcesAPIClient = (*Client)(nil)

// ListMonitoredResourcesPaginatorOptions is the paginator options for
// ListMonitoredResources
type ListMonitoredResourcesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMonitoredResourcesPaginator is a paginator for ListMonitoredResources
type ListMonitoredResourcesPaginator struct {
	options   ListMonitoredResourcesPaginatorOptions
	client    ListMonitoredResourcesAPIClient
	params    *ListMonitoredResourcesInput
	nextToken *string
	firstPage bool
}

// NewListMonitoredResourcesPaginator returns a new ListMonitoredResourcesPaginator
func NewListMonitoredResourcesPaginator(client ListMonitoredResourcesAPIClient, params *ListMonitoredResourcesInput, optFns ...func(*ListMonitoredResourcesPaginatorOptions)) *ListMonitoredResourcesPaginator {
	if params == nil {
		params = &ListMonitoredResourcesInput{}
	}

	options := ListMonitoredResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMonitoredResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMonitoredResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMonitoredResources page.
func (p *ListMonitoredResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMonitoredResourcesOutput, error) {
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

	result, err := p.client.ListMonitoredResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMonitoredResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "ListMonitoredResources",
	}
}
