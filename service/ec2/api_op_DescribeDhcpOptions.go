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

// Describes one or more of your DHCP options sets. For more information, see DHCP
// options sets
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_DHCP_Options.html) in the
// Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeDhcpOptions(ctx context.Context, params *DescribeDhcpOptionsInput, optFns ...func(*Options)) (*DescribeDhcpOptionsOutput, error) {
	if params == nil {
		params = &DescribeDhcpOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDhcpOptions", params, optFns, c.addOperationDescribeDhcpOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDhcpOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDhcpOptionsInput struct {

	// The IDs of one or more DHCP options sets. Default: Describes all your DHCP
	// options sets.
	DhcpOptionsIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * dhcp-options-id - The ID of a DHCP options set.
	//
	// * key -
	// The key for one of the options (for example, domain-name).
	//
	// * value - The value
	// for one of the options.
	//
	// * owner-id - The ID of the Amazon Web Services account
	// that owns the DHCP options set.
	//
	// * tag: - The key/value combination of a tag
	// assigned to the resource. Use the tag key in the filter name and the tag value
	// as the filter value. For example, to find all resources that have a tag with the
	// key Owner and the value TeamA, specify tag:Owner for the filter name and TeamA
	// for the filter value.
	//
	// * tag-key - The key of a tag assigned to the resource.
	// Use this filter to find all resources assigned a tag with a specific key,
	// regardless of the tag value.
	Filters []types.Filter

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeDhcpOptionsOutput struct {

	// Information about one or more DHCP options sets.
	DhcpOptions []types.DhcpOptions

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDhcpOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeDhcpOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeDhcpOptions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDhcpOptions(options.Region), middleware.Before); err != nil {
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

// DescribeDhcpOptionsAPIClient is a client that implements the DescribeDhcpOptions
// operation.
type DescribeDhcpOptionsAPIClient interface {
	DescribeDhcpOptions(context.Context, *DescribeDhcpOptionsInput, ...func(*Options)) (*DescribeDhcpOptionsOutput, error)
}

var _ DescribeDhcpOptionsAPIClient = (*Client)(nil)

// DescribeDhcpOptionsPaginatorOptions is the paginator options for
// DescribeDhcpOptions
type DescribeDhcpOptionsPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDhcpOptionsPaginator is a paginator for DescribeDhcpOptions
type DescribeDhcpOptionsPaginator struct {
	options   DescribeDhcpOptionsPaginatorOptions
	client    DescribeDhcpOptionsAPIClient
	params    *DescribeDhcpOptionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDhcpOptionsPaginator returns a new DescribeDhcpOptionsPaginator
func NewDescribeDhcpOptionsPaginator(client DescribeDhcpOptionsAPIClient, params *DescribeDhcpOptionsInput, optFns ...func(*DescribeDhcpOptionsPaginatorOptions)) *DescribeDhcpOptionsPaginator {
	if params == nil {
		params = &DescribeDhcpOptionsInput{}
	}

	options := DescribeDhcpOptionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDhcpOptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDhcpOptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDhcpOptions page.
func (p *DescribeDhcpOptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDhcpOptionsOutput, error) {
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

	result, err := p.client.DescribeDhcpOptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDhcpOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeDhcpOptions",
	}
}
