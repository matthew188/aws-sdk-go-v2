// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the endpoint groups that are associated with a listener.
func (c *Client) ListEndpointGroups(ctx context.Context, params *ListEndpointGroupsInput, optFns ...func(*Options)) (*ListEndpointGroupsOutput, error) {
	if params == nil {
		params = &ListEndpointGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpointGroups", params, optFns, c.addOperationListEndpointGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointGroupsInput struct {

	// The Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerArn *string

	// The number of endpoint group objects that you want to return with this call. The
	// default value is 10.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEndpointGroupsOutput struct {

	// The list of the endpoint groups associated with a listener.
	EndpointGroups []types.EndpointGroup

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEndpointGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpointGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpointGroups{}, middleware.After)
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
	if err = addOpListEndpointGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpointGroups(options.Region), middleware.Before); err != nil {
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

// ListEndpointGroupsAPIClient is a client that implements the ListEndpointGroups
// operation.
type ListEndpointGroupsAPIClient interface {
	ListEndpointGroups(context.Context, *ListEndpointGroupsInput, ...func(*Options)) (*ListEndpointGroupsOutput, error)
}

var _ ListEndpointGroupsAPIClient = (*Client)(nil)

// ListEndpointGroupsPaginatorOptions is the paginator options for
// ListEndpointGroups
type ListEndpointGroupsPaginatorOptions struct {
	// The number of endpoint group objects that you want to return with this call. The
	// default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEndpointGroupsPaginator is a paginator for ListEndpointGroups
type ListEndpointGroupsPaginator struct {
	options   ListEndpointGroupsPaginatorOptions
	client    ListEndpointGroupsAPIClient
	params    *ListEndpointGroupsInput
	nextToken *string
	firstPage bool
}

// NewListEndpointGroupsPaginator returns a new ListEndpointGroupsPaginator
func NewListEndpointGroupsPaginator(client ListEndpointGroupsAPIClient, params *ListEndpointGroupsInput, optFns ...func(*ListEndpointGroupsPaginatorOptions)) *ListEndpointGroupsPaginator {
	if params == nil {
		params = &ListEndpointGroupsInput{}
	}

	options := ListEndpointGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEndpointGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEndpointGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEndpointGroups page.
func (p *ListEndpointGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEndpointGroupsOutput, error) {
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

	result, err := p.client.ListEndpointGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEndpointGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListEndpointGroups",
	}
}
