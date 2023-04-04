// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about cluster or replication group snapshots. By default,
// DescribeSnapshots lists all of your snapshots; it can optionally describe a
// single snapshot, or just the snapshots associated with a particular cache
// cluster. This operation is valid for Redis only.
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

// Represents the input of a DescribeSnapshotsMessage operation.
type DescribeSnapshotsInput struct {

	// A user-supplied cluster identifier. If this parameter is specified, only
	// snapshots associated with that specific cluster are described.
	CacheClusterId *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 50 Constraints: minimum
	// 20; maximum 50.
	MaxRecords *int32

	// A user-supplied replication group identifier. If this parameter is specified,
	// only snapshots associated with that specific replication group are described.
	ReplicationGroupId *string

	// A Boolean value which if true, the node group (shard) configuration is included
	// in the snapshot description.
	ShowNodeGroupConfig *bool

	// A user-supplied name of the snapshot. If this parameter is specified, only this
	// snapshot are described.
	SnapshotName *string

	// If set to system, the output shows snapshots that were automatically created by
	// ElastiCache. If set to user the output shows snapshots that were manually
	// created. If omitted, the output shows both automatically and manually created
	// snapshots.
	SnapshotSource *string

	noSmithyDocumentSerde
}

// Represents the output of a DescribeSnapshots operation.
type DescribeSnapshotsOutput struct {

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// A list of snapshots. Each item in the list contains detailed information about
	// one snapshot.
	Snapshots []types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSnapshots{}, middleware.After)
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
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 50 Constraints: minimum
	// 20; maximum 50.
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
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
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
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeSnapshots(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

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
		SigningName:   "elasticache",
		OperationName: "DescribeSnapshots",
	}
}
