// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the actions in your account and their properties.
func (c *Client) ListActions(ctx context.Context, params *ListActionsInput, optFns ...func(*Options)) (*ListActionsOutput, error) {
	if params == nil {
		params = &ListActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListActions", params, optFns, c.addOperationListActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListActionsInput struct {

	// A filter that returns only actions of the specified type.
	ActionType *string

	// A filter that returns only actions created on or after the specified time.
	CreatedAfter *time.Time

	// A filter that returns only actions created on or before the specified time.
	CreatedBefore *time.Time

	// The maximum number of actions to return in the response. The default value is
	// 10.
	MaxResults *int32

	// If the previous call to ListActions didn't return the full set of actions, the
	// call returns a token for getting the next set of actions.
	NextToken *string

	// The property used to sort results. The default value is CreationTime.
	SortBy types.SortActionsBy

	// The sort order. The default value is Descending.
	SortOrder types.SortOrder

	// A filter that returns only actions with the specified source URI.
	SourceUri *string

	noSmithyDocumentSerde
}

type ListActionsOutput struct {

	// A list of actions and their properties.
	ActionSummaries []types.ActionSummary

	// A token for getting the next set of actions, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListActions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListActions(options.Region), middleware.Before); err != nil {
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

// ListActionsAPIClient is a client that implements the ListActions operation.
type ListActionsAPIClient interface {
	ListActions(context.Context, *ListActionsInput, ...func(*Options)) (*ListActionsOutput, error)
}

var _ ListActionsAPIClient = (*Client)(nil)

// ListActionsPaginatorOptions is the paginator options for ListActions
type ListActionsPaginatorOptions struct {
	// The maximum number of actions to return in the response. The default value is
	// 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListActionsPaginator is a paginator for ListActions
type ListActionsPaginator struct {
	options   ListActionsPaginatorOptions
	client    ListActionsAPIClient
	params    *ListActionsInput
	nextToken *string
	firstPage bool
}

// NewListActionsPaginator returns a new ListActionsPaginator
func NewListActionsPaginator(client ListActionsAPIClient, params *ListActionsInput, optFns ...func(*ListActionsPaginatorOptions)) *ListActionsPaginator {
	if params == nil {
		params = &ListActionsInput{}
	}

	options := ListActionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListActions page.
func (p *ListActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListActionsOutput, error) {
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

	result, err := p.client.ListActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListActions",
	}
}
