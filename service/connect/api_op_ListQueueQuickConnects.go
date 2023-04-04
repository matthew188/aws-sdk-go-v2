// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Lists the quick connects associated with a queue.
func (c *Client) ListQueueQuickConnects(ctx context.Context, params *ListQueueQuickConnectsInput, optFns ...func(*Options)) (*ListQueueQuickConnectsOutput, error) {
	if params == nil {
		params = &ListQueueQuickConnectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueueQuickConnects", params, optFns, c.addOperationListQueueQuickConnectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueueQuickConnectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueueQuickConnectsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier for the queue.
	//
	// This member is required.
	QueueId *string

	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListQueueQuickConnectsOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the quick connects.
	QuickConnectSummaryList []types.QuickConnectSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueueQuickConnectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListQueueQuickConnects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListQueueQuickConnects{}, middleware.After)
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
	if err = addOpListQueueQuickConnectsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueueQuickConnects(options.Region), middleware.Before); err != nil {
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

// ListQueueQuickConnectsAPIClient is a client that implements the
// ListQueueQuickConnects operation.
type ListQueueQuickConnectsAPIClient interface {
	ListQueueQuickConnects(context.Context, *ListQueueQuickConnectsInput, ...func(*Options)) (*ListQueueQuickConnectsOutput, error)
}

var _ ListQueueQuickConnectsAPIClient = (*Client)(nil)

// ListQueueQuickConnectsPaginatorOptions is the paginator options for
// ListQueueQuickConnects
type ListQueueQuickConnectsPaginatorOptions struct {
	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueueQuickConnectsPaginator is a paginator for ListQueueQuickConnects
type ListQueueQuickConnectsPaginator struct {
	options   ListQueueQuickConnectsPaginatorOptions
	client    ListQueueQuickConnectsAPIClient
	params    *ListQueueQuickConnectsInput
	nextToken *string
	firstPage bool
}

// NewListQueueQuickConnectsPaginator returns a new ListQueueQuickConnectsPaginator
func NewListQueueQuickConnectsPaginator(client ListQueueQuickConnectsAPIClient, params *ListQueueQuickConnectsInput, optFns ...func(*ListQueueQuickConnectsPaginatorOptions)) *ListQueueQuickConnectsPaginator {
	if params == nil {
		params = &ListQueueQuickConnectsInput{}
	}

	options := ListQueueQuickConnectsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueueQuickConnectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueueQuickConnectsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueueQuickConnects page.
func (p *ListQueueQuickConnectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueueQuickConnectsOutput, error) {
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

	result, err := p.client.ListQueueQuickConnects(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQueueQuickConnects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListQueueQuickConnects",
	}
}
