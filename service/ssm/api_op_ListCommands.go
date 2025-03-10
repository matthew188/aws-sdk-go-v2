// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the commands requested by users of the Amazon Web Services account.
func (c *Client) ListCommands(ctx context.Context, params *ListCommandsInput, optFns ...func(*Options)) (*ListCommandsOutput, error) {
	if params == nil {
		params = &ListCommandsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCommands", params, optFns, c.addOperationListCommandsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCommandsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCommandsInput struct {

	// (Optional) If provided, lists only the specified command.
	CommandId *string

	// (Optional) One or more filters. Use a filter to return a more specific list of
	// results.
	Filters []types.CommandFilter

	// (Optional) Lists commands issued against this managed node ID. You can't specify
	// a managed node ID in the same command that you specify Status = Pending. This is
	// because the command hasn't reached the managed node yet.
	InstanceId *string

	// (Optional) The maximum number of items to return for this call. The call also
	// returns a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int32

	// (Optional) The token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type ListCommandsOutput struct {

	// (Optional) The list of commands requested by the user.
	Commands []types.Command

	// (Optional) The token for the next set of items to return. (You received this
	// token from a previous call.)
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCommandsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCommands{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCommands{}, middleware.After)
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
	if err = addOpListCommandsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCommands(options.Region), middleware.Before); err != nil {
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

// ListCommandsAPIClient is a client that implements the ListCommands operation.
type ListCommandsAPIClient interface {
	ListCommands(context.Context, *ListCommandsInput, ...func(*Options)) (*ListCommandsOutput, error)
}

var _ ListCommandsAPIClient = (*Client)(nil)

// ListCommandsPaginatorOptions is the paginator options for ListCommands
type ListCommandsPaginatorOptions struct {
	// (Optional) The maximum number of items to return for this call. The call also
	// returns a token that you can specify in a subsequent call to get the next set of
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCommandsPaginator is a paginator for ListCommands
type ListCommandsPaginator struct {
	options   ListCommandsPaginatorOptions
	client    ListCommandsAPIClient
	params    *ListCommandsInput
	nextToken *string
	firstPage bool
}

// NewListCommandsPaginator returns a new ListCommandsPaginator
func NewListCommandsPaginator(client ListCommandsAPIClient, params *ListCommandsInput, optFns ...func(*ListCommandsPaginatorOptions)) *ListCommandsPaginator {
	if params == nil {
		params = &ListCommandsInput{}
	}

	options := ListCommandsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCommandsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCommandsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCommands page.
func (p *ListCommandsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCommandsOutput, error) {
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

	result, err := p.client.ListCommands(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCommands(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListCommands",
	}
}
