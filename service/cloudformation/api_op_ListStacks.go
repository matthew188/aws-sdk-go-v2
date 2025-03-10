// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the summary information for stacks whose status matches the specified
// StackStatusFilter. Summary information for stacks that have been deleted is kept
// for 90 days after the stack is deleted. If no StackStatusFilter is specified,
// summary information for all stacks is returned (including existing stacks and
// stacks that have been deleted).
func (c *Client) ListStacks(ctx context.Context, params *ListStacksInput, optFns ...func(*Options)) (*ListStacksOutput, error) {
	if params == nil {
		params = &ListStacksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStacks", params, optFns, c.addOperationListStacksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for ListStacks action.
type ListStacksInput struct {

	// A string that identifies the next page of stacks that you want to retrieve.
	NextToken *string

	// Stack status to use as a filter. Specify one or more stack status codes to list
	// only stacks with the specified status codes. For a complete list of stack status
	// codes, see the StackStatus parameter of the Stack data type.
	StackStatusFilter []types.StackStatus

	noSmithyDocumentSerde
}

// The output for ListStacks action.
type ListStacksOutput struct {

	// If the output exceeds 1 MB in size, a string that identifies the next page of
	// stacks. If no additional page exists, this value is null.
	NextToken *string

	// A list of StackSummary structures containing information about the specified
	// stacks.
	StackSummaries []types.StackSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStacksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListStacks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListStacks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStacks(options.Region), middleware.Before); err != nil {
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

// ListStacksAPIClient is a client that implements the ListStacks operation.
type ListStacksAPIClient interface {
	ListStacks(context.Context, *ListStacksInput, ...func(*Options)) (*ListStacksOutput, error)
}

var _ ListStacksAPIClient = (*Client)(nil)

// ListStacksPaginatorOptions is the paginator options for ListStacks
type ListStacksPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStacksPaginator is a paginator for ListStacks
type ListStacksPaginator struct {
	options   ListStacksPaginatorOptions
	client    ListStacksAPIClient
	params    *ListStacksInput
	nextToken *string
	firstPage bool
}

// NewListStacksPaginator returns a new ListStacksPaginator
func NewListStacksPaginator(client ListStacksAPIClient, params *ListStacksInput, optFns ...func(*ListStacksPaginatorOptions)) *ListStacksPaginator {
	if params == nil {
		params = &ListStacksInput{}
	}

	options := ListStacksPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStacksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStacksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStacks page.
func (p *ListStacksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStacksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListStacks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStacks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListStacks",
	}
}
