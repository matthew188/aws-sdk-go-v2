// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/braket/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for tasks that match the specified filter values.
func (c *Client) SearchQuantumTasks(ctx context.Context, params *SearchQuantumTasksInput, optFns ...func(*Options)) (*SearchQuantumTasksOutput, error) {
	if params == nil {
		params = &SearchQuantumTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchQuantumTasks", params, optFns, c.addOperationSearchQuantumTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchQuantumTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchQuantumTasksInput struct {

	// Array of SearchQuantumTasksFilter objects.
	//
	// This member is required.
	Filters []types.SearchQuantumTasksFilter

	// Maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned in the response. Use the token
	// returned from the previous request continue results where the previous request
	// ended.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchQuantumTasksOutput struct {

	// An array of QuantumTaskSummary objects for tasks that match the specified
	// filters.
	//
	// This member is required.
	QuantumTasks []types.QuantumTaskSummary

	// A token used for pagination of results, or null if there are no additional
	// results. Use the token value in a subsequent request to continue results where
	// the previous request ended.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchQuantumTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchQuantumTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchQuantumTasks{}, middleware.After)
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
	if err = addOpSearchQuantumTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchQuantumTasks(options.Region), middleware.Before); err != nil {
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

// SearchQuantumTasksAPIClient is a client that implements the SearchQuantumTasks
// operation.
type SearchQuantumTasksAPIClient interface {
	SearchQuantumTasks(context.Context, *SearchQuantumTasksInput, ...func(*Options)) (*SearchQuantumTasksOutput, error)
}

var _ SearchQuantumTasksAPIClient = (*Client)(nil)

// SearchQuantumTasksPaginatorOptions is the paginator options for
// SearchQuantumTasks
type SearchQuantumTasksPaginatorOptions struct {
	// Maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchQuantumTasksPaginator is a paginator for SearchQuantumTasks
type SearchQuantumTasksPaginator struct {
	options   SearchQuantumTasksPaginatorOptions
	client    SearchQuantumTasksAPIClient
	params    *SearchQuantumTasksInput
	nextToken *string
	firstPage bool
}

// NewSearchQuantumTasksPaginator returns a new SearchQuantumTasksPaginator
func NewSearchQuantumTasksPaginator(client SearchQuantumTasksAPIClient, params *SearchQuantumTasksInput, optFns ...func(*SearchQuantumTasksPaginatorOptions)) *SearchQuantumTasksPaginator {
	if params == nil {
		params = &SearchQuantumTasksInput{}
	}

	options := SearchQuantumTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchQuantumTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchQuantumTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchQuantumTasks page.
func (p *SearchQuantumTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchQuantumTasksOutput, error) {
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

	result, err := p.client.SearchQuantumTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchQuantumTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "braket",
		OperationName: "SearchQuantumTasks",
	}
}
