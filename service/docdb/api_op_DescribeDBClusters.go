// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about provisioned Amazon DocumentDB clusters. This API
// operation supports pagination. For certain management features such as cluster
// and instance lifecycle management, Amazon DocumentDB leverages operational
// technology that is shared with Amazon RDS and Amazon Neptune. Use the
// filterName=engine,Values=docdb filter parameter to return only Amazon DocumentDB
// clusters.
func (c *Client) DescribeDBClusters(ctx context.Context, params *DescribeDBClustersInput, optFns ...func(*Options)) (*DescribeDBClustersOutput, error) {
	if params == nil {
		params = &DescribeDBClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusters", params, optFns, c.addOperationDescribeDBClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBClusters.
type DescribeDBClustersInput struct {

	// The user-provided cluster identifier. If this parameter is specified,
	// information from only the specific cluster is returned. This parameter isn't
	// case sensitive. Constraints:
	//
	// * If provided, must match an existing
	// DBClusterIdentifier.
	DBClusterIdentifier *string

	// A filter that specifies one or more clusters to describe. Supported filters:
	//
	// *
	// db-cluster-id - Accepts cluster identifiers and cluster Amazon Resource Names
	// (ARNs). The results list only includes information about the clusters identified
	// by these ARNs.
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Represents the output of DescribeDBClusters.
type DescribeDBClustersOutput struct {

	// A list of clusters.
	DBClusters []types.DBCluster

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusters{}, middleware.After)
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
	if err = addOpDescribeDBClustersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusters(options.Region), middleware.Before); err != nil {
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

// DescribeDBClustersAPIClient is a client that implements the DescribeDBClusters
// operation.
type DescribeDBClustersAPIClient interface {
	DescribeDBClusters(context.Context, *DescribeDBClustersInput, ...func(*Options)) (*DescribeDBClustersOutput, error)
}

var _ DescribeDBClustersAPIClient = (*Client)(nil)

// DescribeDBClustersPaginatorOptions is the paginator options for
// DescribeDBClusters
type DescribeDBClustersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClustersPaginator is a paginator for DescribeDBClusters
type DescribeDBClustersPaginator struct {
	options   DescribeDBClustersPaginatorOptions
	client    DescribeDBClustersAPIClient
	params    *DescribeDBClustersInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClustersPaginator returns a new DescribeDBClustersPaginator
func NewDescribeDBClustersPaginator(client DescribeDBClustersAPIClient, params *DescribeDBClustersInput, optFns ...func(*DescribeDBClustersPaginatorOptions)) *DescribeDBClustersPaginator {
	if params == nil {
		params = &DescribeDBClustersInput{}
	}

	options := DescribeDBClustersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClustersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBClusters page.
func (p *DescribeDBClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClustersOutput, error) {
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

	result, err := p.client.DescribeDBClusters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDBClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusters",
	}
}
