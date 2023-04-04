// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all component types in a workspace.
func (c *Client) ListComponentTypes(ctx context.Context, params *ListComponentTypesInput, optFns ...func(*Options)) (*ListComponentTypesOutput, error) {
	if params == nil {
		params = &ListComponentTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponentTypes", params, optFns, c.addOperationListComponentTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentTypesInput struct {

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	// A list of objects that filter the request.
	Filters []types.ListComponentTypesFilter

	// The maximum number of results to return at one time. The default is 25. Valid
	// Range: Minimum value of 1. Maximum value of 250.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComponentTypesOutput struct {

	// A list of objects that contain information about the component types.
	//
	// This member is required.
	ComponentTypeSummaries []types.ComponentTypeSummary

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	// Specifies the maximum number of results to display.
	MaxResults *int32

	// The string that specifies the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComponentTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListComponentTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListComponentTypes{}, middleware.After)
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
	if err = addEndpointPrefix_opListComponentTypesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListComponentTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponentTypes(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListComponentTypesMiddleware struct {
}

func (*endpointPrefix_opListComponentTypesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListComponentTypesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListComponentTypesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListComponentTypesMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListComponentTypesAPIClient is a client that implements the ListComponentTypes
// operation.
type ListComponentTypesAPIClient interface {
	ListComponentTypes(context.Context, *ListComponentTypesInput, ...func(*Options)) (*ListComponentTypesOutput, error)
}

var _ ListComponentTypesAPIClient = (*Client)(nil)

// ListComponentTypesPaginatorOptions is the paginator options for
// ListComponentTypes
type ListComponentTypesPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25. Valid
	// Range: Minimum value of 1. Maximum value of 250.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComponentTypesPaginator is a paginator for ListComponentTypes
type ListComponentTypesPaginator struct {
	options   ListComponentTypesPaginatorOptions
	client    ListComponentTypesAPIClient
	params    *ListComponentTypesInput
	nextToken *string
	firstPage bool
}

// NewListComponentTypesPaginator returns a new ListComponentTypesPaginator
func NewListComponentTypesPaginator(client ListComponentTypesAPIClient, params *ListComponentTypesInput, optFns ...func(*ListComponentTypesPaginatorOptions)) *ListComponentTypesPaginator {
	if params == nil {
		params = &ListComponentTypesInput{}
	}

	options := ListComponentTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComponentTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComponentTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComponentTypes page.
func (p *ListComponentTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComponentTypesOutput, error) {
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

	result, err := p.client.ListComponentTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComponentTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "ListComponentTypes",
	}
}
