// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the delegates associated with a resource. Users and groups can be resource
// delegates and answer requests on behalf of the resource.
func (c *Client) ListResourceDelegates(ctx context.Context, params *ListResourceDelegatesInput, optFns ...func(*Options)) (*ListResourceDelegatesOutput, error) {
	if params == nil {
		params = &ListResourceDelegatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceDelegates", params, optFns, c.addOperationListResourceDelegatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceDelegatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceDelegatesInput struct {

	// The identifier for the organization that contains the resource for which
	// delegates are listed.
	//
	// This member is required.
	OrganizationId *string

	// The identifier for the resource whose delegates are listed.
	//
	// This member is required.
	ResourceId *string

	// The number of maximum results in a page.
	MaxResults *int32

	// The token used to paginate through the delegates associated with a resource.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourceDelegatesOutput struct {

	// One page of the resource's delegates.
	Delegates []types.Delegate

	// The token used to paginate through the delegates associated with a resource.
	// While results are still available, it has an associated value. When the last
	// page is reached, the token is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceDelegatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourceDelegates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourceDelegates{}, middleware.After)
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
	if err = addOpListResourceDelegatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceDelegates(options.Region), middleware.Before); err != nil {
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

// ListResourceDelegatesAPIClient is a client that implements the
// ListResourceDelegates operation.
type ListResourceDelegatesAPIClient interface {
	ListResourceDelegates(context.Context, *ListResourceDelegatesInput, ...func(*Options)) (*ListResourceDelegatesOutput, error)
}

var _ ListResourceDelegatesAPIClient = (*Client)(nil)

// ListResourceDelegatesPaginatorOptions is the paginator options for
// ListResourceDelegates
type ListResourceDelegatesPaginatorOptions struct {
	// The number of maximum results in a page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceDelegatesPaginator is a paginator for ListResourceDelegates
type ListResourceDelegatesPaginator struct {
	options   ListResourceDelegatesPaginatorOptions
	client    ListResourceDelegatesAPIClient
	params    *ListResourceDelegatesInput
	nextToken *string
	firstPage bool
}

// NewListResourceDelegatesPaginator returns a new ListResourceDelegatesPaginator
func NewListResourceDelegatesPaginator(client ListResourceDelegatesAPIClient, params *ListResourceDelegatesInput, optFns ...func(*ListResourceDelegatesPaginatorOptions)) *ListResourceDelegatesPaginator {
	if params == nil {
		params = &ListResourceDelegatesInput{}
	}

	options := ListResourceDelegatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceDelegatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceDelegatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourceDelegates page.
func (p *ListResourceDelegatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceDelegatesOutput, error) {
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

	result, err := p.client.ListResourceDelegates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourceDelegates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "ListResourceDelegates",
	}
}
