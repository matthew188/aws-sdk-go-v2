// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mediapackagevod/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a collection of MediaPackage VOD PackagingGroup resources.
func (c *Client) ListPackagingGroups(ctx context.Context, params *ListPackagingGroupsInput, optFns ...func(*Options)) (*ListPackagingGroupsOutput, error) {
	if params == nil {
		params = &ListPackagingGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackagingGroups", params, optFns, c.addOperationListPackagingGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackagingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackagingGroupsInput struct {

	// Upper bound on number of records to return.
	MaxResults int32

	// A token used to resume pagination from the end of a previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPackagingGroupsOutput struct {

	// A token that can be used to resume pagination from the end of the collection.
	NextToken *string

	// A list of MediaPackage VOD PackagingGroup resources.
	PackagingGroups []types.PackagingGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackagingGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackagingGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackagingGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackagingGroups(options.Region), middleware.Before); err != nil {
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

// ListPackagingGroupsAPIClient is a client that implements the ListPackagingGroups
// operation.
type ListPackagingGroupsAPIClient interface {
	ListPackagingGroups(context.Context, *ListPackagingGroupsInput, ...func(*Options)) (*ListPackagingGroupsOutput, error)
}

var _ ListPackagingGroupsAPIClient = (*Client)(nil)

// ListPackagingGroupsPaginatorOptions is the paginator options for
// ListPackagingGroups
type ListPackagingGroupsPaginatorOptions struct {
	// Upper bound on number of records to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackagingGroupsPaginator is a paginator for ListPackagingGroups
type ListPackagingGroupsPaginator struct {
	options   ListPackagingGroupsPaginatorOptions
	client    ListPackagingGroupsAPIClient
	params    *ListPackagingGroupsInput
	nextToken *string
	firstPage bool
}

// NewListPackagingGroupsPaginator returns a new ListPackagingGroupsPaginator
func NewListPackagingGroupsPaginator(client ListPackagingGroupsAPIClient, params *ListPackagingGroupsInput, optFns ...func(*ListPackagingGroupsPaginatorOptions)) *ListPackagingGroupsPaginator {
	if params == nil {
		params = &ListPackagingGroupsInput{}
	}

	options := ListPackagingGroupsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPackagingGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackagingGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPackagingGroups page.
func (p *ListPackagingGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackagingGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPackagingGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPackagingGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "ListPackagingGroups",
	}
}
