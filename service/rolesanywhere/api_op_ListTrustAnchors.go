// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rolesanywhere/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the trust anchors in the authenticated account and Amazon Web Services
// Region. Required permissions: rolesanywhere:ListTrustAnchors.
func (c *Client) ListTrustAnchors(ctx context.Context, params *ListTrustAnchorsInput, optFns ...func(*Options)) (*ListTrustAnchorsOutput, error) {
	if params == nil {
		params = &ListTrustAnchorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrustAnchors", params, optFns, c.addOperationListTrustAnchorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrustAnchorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrustAnchorsInput struct {

	// A token that indicates where the output should continue from, if a previous
	// operation did not show all results. To get the next results, call the operation
	// again with this value.
	NextToken *string

	// The number of resources in the paginated list.
	PageSize *int32

	noSmithyDocumentSerde
}

type ListTrustAnchorsOutput struct {

	// A token that indicates where the output should continue from, if a previous
	// operation did not show all results. To get the next results, call the operation
	// again with this value.
	NextToken *string

	// A list of trust anchors.
	TrustAnchors []types.TrustAnchorDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrustAnchorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTrustAnchors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTrustAnchors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrustAnchors(options.Region), middleware.Before); err != nil {
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

// ListTrustAnchorsAPIClient is a client that implements the ListTrustAnchors
// operation.
type ListTrustAnchorsAPIClient interface {
	ListTrustAnchors(context.Context, *ListTrustAnchorsInput, ...func(*Options)) (*ListTrustAnchorsOutput, error)
}

var _ ListTrustAnchorsAPIClient = (*Client)(nil)

// ListTrustAnchorsPaginatorOptions is the paginator options for ListTrustAnchors
type ListTrustAnchorsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrustAnchorsPaginator is a paginator for ListTrustAnchors
type ListTrustAnchorsPaginator struct {
	options   ListTrustAnchorsPaginatorOptions
	client    ListTrustAnchorsAPIClient
	params    *ListTrustAnchorsInput
	nextToken *string
	firstPage bool
}

// NewListTrustAnchorsPaginator returns a new ListTrustAnchorsPaginator
func NewListTrustAnchorsPaginator(client ListTrustAnchorsAPIClient, params *ListTrustAnchorsInput, optFns ...func(*ListTrustAnchorsPaginatorOptions)) *ListTrustAnchorsPaginator {
	if params == nil {
		params = &ListTrustAnchorsInput{}
	}

	options := ListTrustAnchorsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrustAnchorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrustAnchorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrustAnchors page.
func (p *ListTrustAnchorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrustAnchorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListTrustAnchors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTrustAnchors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rolesanywhere",
		OperationName: "ListTrustAnchors",
	}
}
