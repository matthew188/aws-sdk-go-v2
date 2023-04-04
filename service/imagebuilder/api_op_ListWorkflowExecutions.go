// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of workflow runtime instance metadata objects for a specific
// image build version.
func (c *Client) ListWorkflowExecutions(ctx context.Context, params *ListWorkflowExecutionsInput, optFns ...func(*Options)) (*ListWorkflowExecutionsOutput, error) {
	if params == nil {
		params = &ListWorkflowExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkflowExecutions", params, optFns, c.addOperationListWorkflowExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkflowExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkflowExecutionsInput struct {

	// List all workflow runtime instances for the specified image build version
	// resource ARN.
	//
	// This member is required.
	ImageBuildVersionArn *string

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkflowExecutionsOutput struct {

	// The resource ARN of the image build version for which you requested a list of
	// workflow runtime details.
	ImageBuildVersionArn *string

	// The output message from the list action, if applicable.
	Message *string

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service has'ot included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Contains an array of runtime details that represents each time a workflow ran
	// for the requested image build version.
	WorkflowExecutions []types.WorkflowExecutionMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkflowExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkflowExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkflowExecutions{}, middleware.After)
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
	if err = addOpListWorkflowExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkflowExecutions(options.Region), middleware.Before); err != nil {
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

// ListWorkflowExecutionsAPIClient is a client that implements the
// ListWorkflowExecutions operation.
type ListWorkflowExecutionsAPIClient interface {
	ListWorkflowExecutions(context.Context, *ListWorkflowExecutionsInput, ...func(*Options)) (*ListWorkflowExecutionsOutput, error)
}

var _ ListWorkflowExecutionsAPIClient = (*Client)(nil)

// ListWorkflowExecutionsPaginatorOptions is the paginator options for
// ListWorkflowExecutions
type ListWorkflowExecutionsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkflowExecutionsPaginator is a paginator for ListWorkflowExecutions
type ListWorkflowExecutionsPaginator struct {
	options   ListWorkflowExecutionsPaginatorOptions
	client    ListWorkflowExecutionsAPIClient
	params    *ListWorkflowExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListWorkflowExecutionsPaginator returns a new ListWorkflowExecutionsPaginator
func NewListWorkflowExecutionsPaginator(client ListWorkflowExecutionsAPIClient, params *ListWorkflowExecutionsInput, optFns ...func(*ListWorkflowExecutionsPaginatorOptions)) *ListWorkflowExecutionsPaginator {
	if params == nil {
		params = &ListWorkflowExecutionsInput{}
	}

	options := ListWorkflowExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkflowExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkflowExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkflowExecutions page.
func (p *ListWorkflowExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkflowExecutionsOutput, error) {
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

	result, err := p.client.ListWorkflowExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkflowExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "ListWorkflowExecutions",
	}
}
