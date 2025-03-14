// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of App Runner connections that are associated with your Amazon
// Web Services account.
func (c *Client) ListConnections(ctx context.Context, params *ListConnectionsInput, optFns ...func(*Options)) (*ListConnectionsOutput, error) {
	if params == nil {
		params = &ListConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConnections", params, optFns, c.addOperationListConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectionsInput struct {

	// If specified, only this connection is returned. If not specified, the result
	// isn't filtered by name.
	ConnectionName *string

	// The maximum number of results to include in each response (result page). Used
	// for a paginated request. If you don't specify MaxResults, the request retrieves
	// all available results in a single response.
	MaxResults *int32

	// A token from a previous result page. Used for a paginated request. The request
	// retrieves the next result page. All other parameter values must be identical to
	// the ones specified in the initial request. If you don't specify NextToken, the
	// request retrieves the first result page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConnectionsOutput struct {

	// A list of summary information records for connections. In a paginated request,
	// the request returns up to MaxResults records for each call.
	//
	// This member is required.
	ConnectionSummaryList []types.ConnectionSummary

	// The token that you can pass in a subsequent request to get the next result page.
	// Returned in a paginated request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConnections(options.Region), middleware.Before); err != nil {
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

// ListConnectionsAPIClient is a client that implements the ListConnections
// operation.
type ListConnectionsAPIClient interface {
	ListConnections(context.Context, *ListConnectionsInput, ...func(*Options)) (*ListConnectionsOutput, error)
}

var _ ListConnectionsAPIClient = (*Client)(nil)

// ListConnectionsPaginatorOptions is the paginator options for ListConnections
type ListConnectionsPaginatorOptions struct {
	// The maximum number of results to include in each response (result page). Used
	// for a paginated request. If you don't specify MaxResults, the request retrieves
	// all available results in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConnectionsPaginator is a paginator for ListConnections
type ListConnectionsPaginator struct {
	options   ListConnectionsPaginatorOptions
	client    ListConnectionsAPIClient
	params    *ListConnectionsInput
	nextToken *string
	firstPage bool
}

// NewListConnectionsPaginator returns a new ListConnectionsPaginator
func NewListConnectionsPaginator(client ListConnectionsAPIClient, params *ListConnectionsInput, optFns ...func(*ListConnectionsPaginatorOptions)) *ListConnectionsPaginator {
	if params == nil {
		params = &ListConnectionsInput{}
	}

	options := ListConnectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConnections page.
func (p *ListConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConnectionsOutput, error) {
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

	result, err := p.client.ListConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "ListConnections",
	}
}
