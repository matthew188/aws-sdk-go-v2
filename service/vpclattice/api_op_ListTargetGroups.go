// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists your target groups. You can narrow your search by using the filters below
// in your request.
func (c *Client) ListTargetGroups(ctx context.Context, params *ListTargetGroupsInput, optFns ...func(*Options)) (*ListTargetGroupsOutput, error) {
	if params == nil {
		params = &ListTargetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTargetGroups", params, optFns, c.addOperationListTargetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetGroupsInput struct {

	// The maximum number of results to return.
	MaxResults *int32

	// A pagination token for the next page of results.
	NextToken *string

	// The target group type.
	TargetGroupType types.TargetGroupType

	// The ID or Amazon Resource Name (ARN) of the service.
	VpcIdentifier *string

	noSmithyDocumentSerde
}

type ListTargetGroupsOutput struct {

	// Information about the target groups.
	Items []types.TargetGroupSummary

	// If there are additional results, a pagination token for the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTargetGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTargetGroups(options.Region), middleware.Before); err != nil {
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

// ListTargetGroupsAPIClient is a client that implements the ListTargetGroups
// operation.
type ListTargetGroupsAPIClient interface {
	ListTargetGroups(context.Context, *ListTargetGroupsInput, ...func(*Options)) (*ListTargetGroupsOutput, error)
}

var _ ListTargetGroupsAPIClient = (*Client)(nil)

// ListTargetGroupsPaginatorOptions is the paginator options for ListTargetGroups
type ListTargetGroupsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTargetGroupsPaginator is a paginator for ListTargetGroups
type ListTargetGroupsPaginator struct {
	options   ListTargetGroupsPaginatorOptions
	client    ListTargetGroupsAPIClient
	params    *ListTargetGroupsInput
	nextToken *string
	firstPage bool
}

// NewListTargetGroupsPaginator returns a new ListTargetGroupsPaginator
func NewListTargetGroupsPaginator(client ListTargetGroupsAPIClient, params *ListTargetGroupsInput, optFns ...func(*ListTargetGroupsPaginatorOptions)) *ListTargetGroupsPaginator {
	if params == nil {
		params = &ListTargetGroupsInput{}
	}

	options := ListTargetGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTargetGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTargetGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTargetGroups page.
func (p *ListTargetGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTargetGroupsOutput, error) {
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

	result, err := p.client.ListTargetGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTargetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "ListTargetGroups",
	}
}
