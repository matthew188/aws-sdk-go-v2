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

// We are retiring EC2-Classic. We recommend that you migrate from EC2-Classic to a
// VPC. For more information, see Migrate from EC2-Classic to a VPC
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html) in the
// Amazon Elastic Compute Cloud User Guide. Describes the ClassicLink DNS support
// status of one or more VPCs. If enabled, the DNS hostname of a linked EC2-Classic
// instance resolves to its private IP address when addressed from an instance in
// the VPC to which it's linked. Similarly, the DNS hostname of an instance in a
// VPC resolves to its private IP address when addressed from a linked EC2-Classic
// instance. For more information, see ClassicLink
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVpcClassicLinkDnsSupport(ctx context.Context, params *DescribeVpcClassicLinkDnsSupportInput, optFns ...func(*Options)) (*DescribeVpcClassicLinkDnsSupportOutput, error) {
	if params == nil {
		params = &DescribeVpcClassicLinkDnsSupportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcClassicLinkDnsSupport", params, optFns, c.addOperationDescribeVpcClassicLinkDnsSupportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcClassicLinkDnsSupportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcClassicLinkDnsSupportInput struct {

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	// One or more VPC IDs.
	VpcIds []string

	noSmithyDocumentSerde
}

type DescribeVpcClassicLinkDnsSupportOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Information about the ClassicLink DNS support status of the VPCs.
	Vpcs []types.ClassicLinkDnsSupport

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcClassicLinkDnsSupportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcClassicLinkDnsSupport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcClassicLinkDnsSupport{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcClassicLinkDnsSupport(options.Region), middleware.Before); err != nil {
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

// DescribeVpcClassicLinkDnsSupportAPIClient is a client that implements the
// DescribeVpcClassicLinkDnsSupport operation.
type DescribeVpcClassicLinkDnsSupportAPIClient interface {
	DescribeVpcClassicLinkDnsSupport(context.Context, *DescribeVpcClassicLinkDnsSupportInput, ...func(*Options)) (*DescribeVpcClassicLinkDnsSupportOutput, error)
}

var _ DescribeVpcClassicLinkDnsSupportAPIClient = (*Client)(nil)

// DescribeVpcClassicLinkDnsSupportPaginatorOptions is the paginator options for
// DescribeVpcClassicLinkDnsSupport
type DescribeVpcClassicLinkDnsSupportPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVpcClassicLinkDnsSupportPaginator is a paginator for
// DescribeVpcClassicLinkDnsSupport
type DescribeVpcClassicLinkDnsSupportPaginator struct {
	options   DescribeVpcClassicLinkDnsSupportPaginatorOptions
	client    DescribeVpcClassicLinkDnsSupportAPIClient
	params    *DescribeVpcClassicLinkDnsSupportInput
	nextToken *string
	firstPage bool
}

// NewDescribeVpcClassicLinkDnsSupportPaginator returns a new
// DescribeVpcClassicLinkDnsSupportPaginator
func NewDescribeVpcClassicLinkDnsSupportPaginator(client DescribeVpcClassicLinkDnsSupportAPIClient, params *DescribeVpcClassicLinkDnsSupportInput, optFns ...func(*DescribeVpcClassicLinkDnsSupportPaginatorOptions)) *DescribeVpcClassicLinkDnsSupportPaginator {
	if params == nil {
		params = &DescribeVpcClassicLinkDnsSupportInput{}
	}

	options := DescribeVpcClassicLinkDnsSupportPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVpcClassicLinkDnsSupportPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVpcClassicLinkDnsSupportPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVpcClassicLinkDnsSupport page.
func (p *DescribeVpcClassicLinkDnsSupportPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVpcClassicLinkDnsSupportOutput, error) {
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

	result, err := p.client.DescribeVpcClassicLinkDnsSupport(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeVpcClassicLinkDnsSupport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcClassicLinkDnsSupport",
	}
}
