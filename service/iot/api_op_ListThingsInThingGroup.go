// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the things in the specified group. Requires permission to access the
// ListThingsInThingGroup
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListThingsInThingGroup(ctx context.Context, params *ListThingsInThingGroupInput, optFns ...func(*Options)) (*ListThingsInThingGroupOutput, error) {
	if params == nil {
		params = &ListThingsInThingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingsInThingGroup", params, optFns, c.addOperationListThingsInThingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingsInThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingsInThingGroupInput struct {

	// The thing group name.
	//
	// This member is required.
	ThingGroupName *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	// When true, list things in this thing group and in all child groups as well.
	Recursive bool

	noSmithyDocumentSerde
}

type ListThingsInThingGroupOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The things in the specified thing group.
	Things []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThingsInThingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingsInThingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingsInThingGroup{}, middleware.After)
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
	if err = addOpListThingsInThingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingsInThingGroup(options.Region), middleware.Before); err != nil {
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

// ListThingsInThingGroupAPIClient is a client that implements the
// ListThingsInThingGroup operation.
type ListThingsInThingGroupAPIClient interface {
	ListThingsInThingGroup(context.Context, *ListThingsInThingGroupInput, ...func(*Options)) (*ListThingsInThingGroupOutput, error)
}

var _ ListThingsInThingGroupAPIClient = (*Client)(nil)

// ListThingsInThingGroupPaginatorOptions is the paginator options for
// ListThingsInThingGroup
type ListThingsInThingGroupPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingsInThingGroupPaginator is a paginator for ListThingsInThingGroup
type ListThingsInThingGroupPaginator struct {
	options   ListThingsInThingGroupPaginatorOptions
	client    ListThingsInThingGroupAPIClient
	params    *ListThingsInThingGroupInput
	nextToken *string
	firstPage bool
}

// NewListThingsInThingGroupPaginator returns a new ListThingsInThingGroupPaginator
func NewListThingsInThingGroupPaginator(client ListThingsInThingGroupAPIClient, params *ListThingsInThingGroupInput, optFns ...func(*ListThingsInThingGroupPaginatorOptions)) *ListThingsInThingGroupPaginator {
	if params == nil {
		params = &ListThingsInThingGroupInput{}
	}

	options := ListThingsInThingGroupPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThingsInThingGroupPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingsInThingGroupPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThingsInThingGroup page.
func (p *ListThingsInThingGroupPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingsInThingGroupOutput, error) {
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

	result, err := p.client.ListThingsInThingGroup(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListThingsInThingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingsInThingGroup",
	}
}
