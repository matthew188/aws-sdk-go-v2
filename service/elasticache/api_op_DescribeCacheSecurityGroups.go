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

// Returns a list of cache security group descriptions. If a cache security group
// name is specified, the list contains only the description of that group. This
// applicable only when you have ElastiCache in Classic setup
func (c *Client) DescribeCacheSecurityGroups(ctx context.Context, params *DescribeCacheSecurityGroupsInput, optFns ...func(*Options)) (*DescribeCacheSecurityGroupsOutput, error) {
	if params == nil {
		params = &DescribeCacheSecurityGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCacheSecurityGroups", params, optFns, c.addOperationDescribeCacheSecurityGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCacheSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheSecurityGroups operation.
type DescribeCacheSecurityGroupsInput struct {

	// The name of the cache security group to return details for.
	CacheSecurityGroupName *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Represents the output of a DescribeCacheSecurityGroups operation.
type DescribeCacheSecurityGroupsOutput struct {

	// A list of cache security groups. Each element in the list contains detailed
	// information about one group.
	CacheSecurityGroups []types.CacheSecurityGroup

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCacheSecurityGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheSecurityGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheSecurityGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheSecurityGroups(options.Region), middleware.Before); err != nil {
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

// DescribeCacheSecurityGroupsAPIClient is a client that implements the
// DescribeCacheSecurityGroups operation.
type DescribeCacheSecurityGroupsAPIClient interface {
	DescribeCacheSecurityGroups(context.Context, *DescribeCacheSecurityGroupsInput, ...func(*Options)) (*DescribeCacheSecurityGroupsOutput, error)
}

var _ DescribeCacheSecurityGroupsAPIClient = (*Client)(nil)

// DescribeCacheSecurityGroupsPaginatorOptions is the paginator options for
// DescribeCacheSecurityGroups
type DescribeCacheSecurityGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCacheSecurityGroupsPaginator is a paginator for
// DescribeCacheSecurityGroups
type DescribeCacheSecurityGroupsPaginator struct {
	options   DescribeCacheSecurityGroupsPaginatorOptions
	client    DescribeCacheSecurityGroupsAPIClient
	params    *DescribeCacheSecurityGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeCacheSecurityGroupsPaginator returns a new
// DescribeCacheSecurityGroupsPaginator
func NewDescribeCacheSecurityGroupsPaginator(client DescribeCacheSecurityGroupsAPIClient, params *DescribeCacheSecurityGroupsInput, optFns ...func(*DescribeCacheSecurityGroupsPaginatorOptions)) *DescribeCacheSecurityGroupsPaginator {
	if params == nil {
		params = &DescribeCacheSecurityGroupsInput{}
	}

	options := DescribeCacheSecurityGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCacheSecurityGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCacheSecurityGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCacheSecurityGroups page.
func (p *DescribeCacheSecurityGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCacheSecurityGroupsOutput, error) {
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

	result, err := p.client.DescribeCacheSecurityGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeCacheSecurityGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheSecurityGroups",
	}
}
