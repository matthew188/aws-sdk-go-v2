// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the entries for a specified managed prefix list.
func (c *Client) GetManagedPrefixListEntries(ctx context.Context, params *GetManagedPrefixListEntriesInput, optFns ...func(*Options)) (*GetManagedPrefixListEntriesOutput, error) {
	if params == nil {
		params = &GetManagedPrefixListEntriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetManagedPrefixListEntries", params, optFns, c.addOperationGetManagedPrefixListEntriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetManagedPrefixListEntriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetManagedPrefixListEntriesInput struct {

	// The ID of the prefix list.
	//
	// This member is required.
	PrefixListId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The version of the prefix list for which to return the entries. The default is
	// the current version.
	TargetVersion *int64

	noSmithyDocumentSerde
}

type GetManagedPrefixListEntriesOutput struct {

	// Information about the prefix list entries.
	Entries []types.PrefixListEntry

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetManagedPrefixListEntriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetManagedPrefixListEntries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetManagedPrefixListEntries{}, middleware.After)
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
	if err = addOpGetManagedPrefixListEntriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetManagedPrefixListEntries(options.Region), middleware.Before); err != nil {
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

// GetManagedPrefixListEntriesAPIClient is a client that implements the
// GetManagedPrefixListEntries operation.
type GetManagedPrefixListEntriesAPIClient interface {
	GetManagedPrefixListEntries(context.Context, *GetManagedPrefixListEntriesInput, ...func(*Options)) (*GetManagedPrefixListEntriesOutput, error)
}

var _ GetManagedPrefixListEntriesAPIClient = (*Client)(nil)

// GetManagedPrefixListEntriesPaginatorOptions is the paginator options for
// GetManagedPrefixListEntries
type GetManagedPrefixListEntriesPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetManagedPrefixListEntriesPaginator is a paginator for
// GetManagedPrefixListEntries
type GetManagedPrefixListEntriesPaginator struct {
	options   GetManagedPrefixListEntriesPaginatorOptions
	client    GetManagedPrefixListEntriesAPIClient
	params    *GetManagedPrefixListEntriesInput
	nextToken *string
	firstPage bool
}

// NewGetManagedPrefixListEntriesPaginator returns a new
// GetManagedPrefixListEntriesPaginator
func NewGetManagedPrefixListEntriesPaginator(client GetManagedPrefixListEntriesAPIClient, params *GetManagedPrefixListEntriesInput, optFns ...func(*GetManagedPrefixListEntriesPaginatorOptions)) *GetManagedPrefixListEntriesPaginator {
	if params == nil {
		params = &GetManagedPrefixListEntriesInput{}
	}

	options := GetManagedPrefixListEntriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetManagedPrefixListEntriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetManagedPrefixListEntriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetManagedPrefixListEntries page.
func (p *GetManagedPrefixListEntriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetManagedPrefixListEntriesOutput, error) {
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

	result, err := p.client.GetManagedPrefixListEntries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetManagedPrefixListEntries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetManagedPrefixListEntries",
	}
}
