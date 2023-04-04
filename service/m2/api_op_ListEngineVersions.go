// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the available engine versions.
func (c *Client) ListEngineVersions(ctx context.Context, params *ListEngineVersionsInput, optFns ...func(*Options)) (*ListEngineVersionsOutput, error) {
	if params == nil {
		params = &ListEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEngineVersions", params, optFns, c.addOperationListEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEngineVersionsInput struct {

	// The type of target platform.
	EngineType types.EngineType

	// The maximum number of objects to return.
	MaxResults *int32

	// A pagination token returned from a previous call to this operation. This
	// specifies the next item to return. To return to the beginning of the list,
	// exclude this parameter.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEngineVersionsOutput struct {

	// Returns the engine versions.
	//
	// This member is required.
	EngineVersions []types.EngineVersionsSummary

	// If there are more items to return, this contains a token that is passed to a
	// subsequent call to this operation to retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEngineVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEngineVersions(options.Region), middleware.Before); err != nil {
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

// ListEngineVersionsAPIClient is a client that implements the ListEngineVersions
// operation.
type ListEngineVersionsAPIClient interface {
	ListEngineVersions(context.Context, *ListEngineVersionsInput, ...func(*Options)) (*ListEngineVersionsOutput, error)
}

var _ ListEngineVersionsAPIClient = (*Client)(nil)

// ListEngineVersionsPaginatorOptions is the paginator options for
// ListEngineVersions
type ListEngineVersionsPaginatorOptions struct {
	// The maximum number of objects to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEngineVersionsPaginator is a paginator for ListEngineVersions
type ListEngineVersionsPaginator struct {
	options   ListEngineVersionsPaginatorOptions
	client    ListEngineVersionsAPIClient
	params    *ListEngineVersionsInput
	nextToken *string
	firstPage bool
}

// NewListEngineVersionsPaginator returns a new ListEngineVersionsPaginator
func NewListEngineVersionsPaginator(client ListEngineVersionsAPIClient, params *ListEngineVersionsInput, optFns ...func(*ListEngineVersionsPaginatorOptions)) *ListEngineVersionsPaginator {
	if params == nil {
		params = &ListEngineVersionsInput{}
	}

	options := ListEngineVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEngineVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEngineVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEngineVersions page.
func (p *ListEngineVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEngineVersionsOutput, error) {
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

	result, err := p.client.ListEngineVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEngineVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "ListEngineVersions",
	}
}
