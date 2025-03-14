// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about cluster snapshots. By default, DescribeSnapshots lists
// all of your snapshots; it can optionally describe a single snapshot, or just the
// snapshots associated with a particular cluster.
func (c *Client) DescribeSnapshots(ctx context.Context, params *DescribeSnapshotsInput, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshots", params, optFns, c.addOperationDescribeSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSnapshotsInput struct {

	// A user-supplied cluster identifier. If this parameter is specified, only
	// snapshots associated with that specific cluster are described.
	ClusterName *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// A Boolean value which if true, the shard configuration is included in the
	// snapshot description.
	ShowDetail *bool

	// A user-supplied name of the snapshot. If this parameter is specified, only this
	// named snapshot is described.
	SnapshotName *string

	// If set to system, the output shows snapshots that were automatically created by
	// MemoryDB. If set to user the output shows snapshots that were manually created.
	// If omitted, the output shows both automatically and manually created snapshots.
	Source *string

	noSmithyDocumentSerde
}

type DescribeSnapshotsOutput struct {

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// A list of snapshots. Each item in the list contains detailed information about
	// one snapshot.
	Snapshots []types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSnapshots{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshots(options.Region), middleware.Before); err != nil {
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

// DescribeSnapshotsAPIClient is a client that implements the DescribeSnapshots
// operation.
type DescribeSnapshotsAPIClient interface {
	DescribeSnapshots(context.Context, *DescribeSnapshotsInput, ...func(*Options)) (*DescribeSnapshotsOutput, error)
}

var _ DescribeSnapshotsAPIClient = (*Client)(nil)

// DescribeSnapshotsPaginatorOptions is the paginator options for DescribeSnapshots
type DescribeSnapshotsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSnapshotsPaginator is a paginator for DescribeSnapshots
type DescribeSnapshotsPaginator struct {
	options   DescribeSnapshotsPaginatorOptions
	client    DescribeSnapshotsAPIClient
	params    *DescribeSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewDescribeSnapshotsPaginator returns a new DescribeSnapshotsPaginator
func NewDescribeSnapshotsPaginator(client DescribeSnapshotsAPIClient, params *DescribeSnapshotsInput, optFns ...func(*DescribeSnapshotsPaginatorOptions)) *DescribeSnapshotsPaginator {
	if params == nil {
		params = &DescribeSnapshotsInput{}
	}

	options := DescribeSnapshotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSnapshots page.
func (p *DescribeSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSnapshotsOutput, error) {
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

	result, err := p.client.DescribeSnapshots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "DescribeSnapshots",
	}
}
