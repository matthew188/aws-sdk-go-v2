// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/backupgateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your virtual machines.
func (c *Client) ListVirtualMachines(ctx context.Context, params *ListVirtualMachinesInput, optFns ...func(*Options)) (*ListVirtualMachinesOutput, error) {
	if params == nil {
		params = &ListVirtualMachinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualMachines", params, optFns, c.addOperationListVirtualMachinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualMachinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVirtualMachinesInput struct {

	// The Amazon Resource Name (ARN) of the hypervisor connected to your virtual
	// machine.
	HypervisorArn *string

	// The maximum number of virtual machines to list.
	MaxResults *int32

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return maxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListVirtualMachinesOutput struct {

	// The next item following a partial list of returned resources. For example, if a
	// request is made to return maxResults number of resources, NextToken allows you
	// to return more items in your list starting at the location pointed to by the
	// next token.
	NextToken *string

	// A list of your VirtualMachine objects, ordered by their Amazon Resource Names
	// (ARNs).
	VirtualMachines []types.VirtualMachine

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVirtualMachinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListVirtualMachines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListVirtualMachines{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVirtualMachines(options.Region), middleware.Before); err != nil {
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

// ListVirtualMachinesAPIClient is a client that implements the ListVirtualMachines
// operation.
type ListVirtualMachinesAPIClient interface {
	ListVirtualMachines(context.Context, *ListVirtualMachinesInput, ...func(*Options)) (*ListVirtualMachinesOutput, error)
}

var _ ListVirtualMachinesAPIClient = (*Client)(nil)

// ListVirtualMachinesPaginatorOptions is the paginator options for
// ListVirtualMachines
type ListVirtualMachinesPaginatorOptions struct {
	// The maximum number of virtual machines to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVirtualMachinesPaginator is a paginator for ListVirtualMachines
type ListVirtualMachinesPaginator struct {
	options   ListVirtualMachinesPaginatorOptions
	client    ListVirtualMachinesAPIClient
	params    *ListVirtualMachinesInput
	nextToken *string
	firstPage bool
}

// NewListVirtualMachinesPaginator returns a new ListVirtualMachinesPaginator
func NewListVirtualMachinesPaginator(client ListVirtualMachinesAPIClient, params *ListVirtualMachinesInput, optFns ...func(*ListVirtualMachinesPaginatorOptions)) *ListVirtualMachinesPaginator {
	if params == nil {
		params = &ListVirtualMachinesInput{}
	}

	options := ListVirtualMachinesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVirtualMachinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVirtualMachinesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListVirtualMachines page.
func (p *ListVirtualMachinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVirtualMachinesOutput, error) {
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

	result, err := p.client.ListVirtualMachines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListVirtualMachines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-gateway",
		OperationName: "ListVirtualMachines",
	}
}
