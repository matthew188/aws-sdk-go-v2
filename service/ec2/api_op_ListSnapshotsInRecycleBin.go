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

// Lists one or more snapshots that are currently in the Recycle Bin.
func (c *Client) ListSnapshotsInRecycleBin(ctx context.Context, params *ListSnapshotsInRecycleBinInput, optFns ...func(*Options)) (*ListSnapshotsInRecycleBinOutput, error) {
	if params == nil {
		params = &ListSnapshotsInRecycleBinInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSnapshotsInRecycleBin", params, optFns, c.addOperationListSnapshotsInRecycleBinMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSnapshotsInRecycleBinOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSnapshotsInRecycleBinInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	// The IDs of the snapshots to list. Omit this parameter to list all of the
	// snapshots that are in the Recycle Bin.
	SnapshotIds []string

	noSmithyDocumentSerde
}

type ListSnapshotsInRecycleBinOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Information about the snapshots.
	Snapshots []types.SnapshotRecycleBinInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSnapshotsInRecycleBinMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpListSnapshotsInRecycleBin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpListSnapshotsInRecycleBin{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSnapshotsInRecycleBin(options.Region), middleware.Before); err != nil {
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

// ListSnapshotsInRecycleBinAPIClient is a client that implements the
// ListSnapshotsInRecycleBin operation.
type ListSnapshotsInRecycleBinAPIClient interface {
	ListSnapshotsInRecycleBin(context.Context, *ListSnapshotsInRecycleBinInput, ...func(*Options)) (*ListSnapshotsInRecycleBinOutput, error)
}

var _ ListSnapshotsInRecycleBinAPIClient = (*Client)(nil)

// ListSnapshotsInRecycleBinPaginatorOptions is the paginator options for
// ListSnapshotsInRecycleBin
type ListSnapshotsInRecycleBinPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSnapshotsInRecycleBinPaginator is a paginator for ListSnapshotsInRecycleBin
type ListSnapshotsInRecycleBinPaginator struct {
	options   ListSnapshotsInRecycleBinPaginatorOptions
	client    ListSnapshotsInRecycleBinAPIClient
	params    *ListSnapshotsInRecycleBinInput
	nextToken *string
	firstPage bool
}

// NewListSnapshotsInRecycleBinPaginator returns a new
// ListSnapshotsInRecycleBinPaginator
func NewListSnapshotsInRecycleBinPaginator(client ListSnapshotsInRecycleBinAPIClient, params *ListSnapshotsInRecycleBinInput, optFns ...func(*ListSnapshotsInRecycleBinPaginatorOptions)) *ListSnapshotsInRecycleBinPaginator {
	if params == nil {
		params = &ListSnapshotsInRecycleBinInput{}
	}

	options := ListSnapshotsInRecycleBinPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSnapshotsInRecycleBinPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSnapshotsInRecycleBinPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSnapshotsInRecycleBin page.
func (p *ListSnapshotsInRecycleBinPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSnapshotsInRecycleBinOutput, error) {
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

	result, err := p.client.ListSnapshotsInRecycleBin(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSnapshotsInRecycleBin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ListSnapshotsInRecycleBin",
	}
}
