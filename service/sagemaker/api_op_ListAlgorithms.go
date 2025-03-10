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

// Lists the machine learning algorithms that have been created.
func (c *Client) ListAlgorithms(ctx context.Context, params *ListAlgorithmsInput, optFns ...func(*Options)) (*ListAlgorithmsOutput, error) {
	if params == nil {
		params = &ListAlgorithmsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAlgorithms", params, optFns, c.addOperationListAlgorithmsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAlgorithmsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAlgorithmsInput struct {

	// A filter that returns only algorithms created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only algorithms created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// The maximum number of algorithms to return in the response.
	MaxResults *int32

	// A string in the algorithm name. This filter returns only algorithms whose name
	// contains the specified string.
	NameContains *string

	// If the response to a previous ListAlgorithms request was truncated, the response
	// includes a NextToken. To retrieve the next set of algorithms, use the token in
	// the next request.
	NextToken *string

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy types.AlgorithmSortBy

	// The sort order for the results. The default is Ascending.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListAlgorithmsOutput struct {

	// >An array of AlgorithmSummary objects, each of which lists an algorithm.
	//
	// This member is required.
	AlgorithmSummaryList []types.AlgorithmSummary

	// If the response is truncated, SageMaker returns this token. To retrieve the next
	// set of algorithms, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAlgorithmsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAlgorithms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAlgorithms{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAlgorithms(options.Region), middleware.Before); err != nil {
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

// ListAlgorithmsAPIClient is a client that implements the ListAlgorithms
// operation.
type ListAlgorithmsAPIClient interface {
	ListAlgorithms(context.Context, *ListAlgorithmsInput, ...func(*Options)) (*ListAlgorithmsOutput, error)
}

var _ ListAlgorithmsAPIClient = (*Client)(nil)

// ListAlgorithmsPaginatorOptions is the paginator options for ListAlgorithms
type ListAlgorithmsPaginatorOptions struct {
	// The maximum number of algorithms to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAlgorithmsPaginator is a paginator for ListAlgorithms
type ListAlgorithmsPaginator struct {
	options   ListAlgorithmsPaginatorOptions
	client    ListAlgorithmsAPIClient
	params    *ListAlgorithmsInput
	nextToken *string
	firstPage bool
}

// NewListAlgorithmsPaginator returns a new ListAlgorithmsPaginator
func NewListAlgorithmsPaginator(client ListAlgorithmsAPIClient, params *ListAlgorithmsInput, optFns ...func(*ListAlgorithmsPaginatorOptions)) *ListAlgorithmsPaginator {
	if params == nil {
		params = &ListAlgorithmsInput{}
	}

	options := ListAlgorithmsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAlgorithmsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAlgorithmsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAlgorithms page.
func (p *ListAlgorithmsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAlgorithmsOutput, error) {
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

	result, err := p.client.ListAlgorithms(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAlgorithms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListAlgorithms",
	}
}
