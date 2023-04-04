// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the outbound cross-cluster connections for a local (source) Amazon
// OpenSearch Service domain. For more information, see Cross-cluster search for
// Amazon OpenSearch Service
// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html).
func (c *Client) DescribeOutboundConnections(ctx context.Context, params *DescribeOutboundConnectionsInput, optFns ...func(*Options)) (*DescribeOutboundConnectionsOutput, error) {
	if params == nil {
		params = &DescribeOutboundConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOutboundConnections", params, optFns, c.addOperationDescribeOutboundConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOutboundConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeOutboundConnections operation.
type DescribeOutboundConnectionsInput struct {

	// List of filter names and values that you can use for requests.
	Filters []types.Filter

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	MaxResults int32

	// If your initial DescribeOutboundConnections operation returns a nextToken, you
	// can include the returned nextToken in subsequent DescribeOutboundConnections
	// operations, which returns results in the next page.
	NextToken *string

	noSmithyDocumentSerde
}

// Contains a list of connections matching the filter criteria.
type DescribeOutboundConnectionsOutput struct {

	// List of outbound connections that match the filter criteria.
	Connections []types.OutboundConnection

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOutboundConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOutboundConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOutboundConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOutboundConnections(options.Region), middleware.Before); err != nil {
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

// DescribeOutboundConnectionsAPIClient is a client that implements the
// DescribeOutboundConnections operation.
type DescribeOutboundConnectionsAPIClient interface {
	DescribeOutboundConnections(context.Context, *DescribeOutboundConnectionsInput, ...func(*Options)) (*DescribeOutboundConnectionsOutput, error)
}

var _ DescribeOutboundConnectionsAPIClient = (*Client)(nil)

// DescribeOutboundConnectionsPaginatorOptions is the paginator options for
// DescribeOutboundConnections
type DescribeOutboundConnectionsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOutboundConnectionsPaginator is a paginator for
// DescribeOutboundConnections
type DescribeOutboundConnectionsPaginator struct {
	options   DescribeOutboundConnectionsPaginatorOptions
	client    DescribeOutboundConnectionsAPIClient
	params    *DescribeOutboundConnectionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeOutboundConnectionsPaginator returns a new
// DescribeOutboundConnectionsPaginator
func NewDescribeOutboundConnectionsPaginator(client DescribeOutboundConnectionsAPIClient, params *DescribeOutboundConnectionsInput, optFns ...func(*DescribeOutboundConnectionsPaginatorOptions)) *DescribeOutboundConnectionsPaginator {
	if params == nil {
		params = &DescribeOutboundConnectionsInput{}
	}

	options := DescribeOutboundConnectionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOutboundConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOutboundConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOutboundConnections page.
func (p *DescribeOutboundConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOutboundConnectionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeOutboundConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOutboundConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeOutboundConnections",
	}
}
