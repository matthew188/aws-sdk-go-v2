// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List source server post migration custom actions.
func (c *Client) ListSourceServerActions(ctx context.Context, params *ListSourceServerActionsInput, optFns ...func(*Options)) (*ListSourceServerActionsOutput, error) {
	if params == nil {
		params = &ListSourceServerActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSourceServerActions", params, optFns, c.addOperationListSourceServerActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSourceServerActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSourceServerActionsInput struct {

	// Source server ID.
	//
	// This member is required.
	SourceServerID *string

	// Filters to apply when listing source server post migration custom actions.
	Filters *types.SourceServerActionsRequestFilters

	// Maximum amount of items to return when listing source server post migration
	// custom actions.
	MaxResults int32

	// Next token to use when listing source server post migration custom actions.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSourceServerActionsOutput struct {

	// List of source server post migration custom actions.
	Items []types.SourceServerActionDocument

	// Next token returned when listing source server post migration custom actions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSourceServerActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSourceServerActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSourceServerActions{}, middleware.After)
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
	if err = addOpListSourceServerActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSourceServerActions(options.Region), middleware.Before); err != nil {
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

// ListSourceServerActionsAPIClient is a client that implements the
// ListSourceServerActions operation.
type ListSourceServerActionsAPIClient interface {
	ListSourceServerActions(context.Context, *ListSourceServerActionsInput, ...func(*Options)) (*ListSourceServerActionsOutput, error)
}

var _ ListSourceServerActionsAPIClient = (*Client)(nil)

// ListSourceServerActionsPaginatorOptions is the paginator options for
// ListSourceServerActions
type ListSourceServerActionsPaginatorOptions struct {
	// Maximum amount of items to return when listing source server post migration
	// custom actions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSourceServerActionsPaginator is a paginator for ListSourceServerActions
type ListSourceServerActionsPaginator struct {
	options   ListSourceServerActionsPaginatorOptions
	client    ListSourceServerActionsAPIClient
	params    *ListSourceServerActionsInput
	nextToken *string
	firstPage bool
}

// NewListSourceServerActionsPaginator returns a new
// ListSourceServerActionsPaginator
func NewListSourceServerActionsPaginator(client ListSourceServerActionsAPIClient, params *ListSourceServerActionsInput, optFns ...func(*ListSourceServerActionsPaginatorOptions)) *ListSourceServerActionsPaginator {
	if params == nil {
		params = &ListSourceServerActionsInput{}
	}

	options := ListSourceServerActionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSourceServerActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSourceServerActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSourceServerActions page.
func (p *ListSourceServerActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSourceServerActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListSourceServerActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSourceServerActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "ListSourceServerActions",
	}
}
