// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified tags. You can use filters to limit the results. For
// example, you can query for the tags for a specific Auto Scaling group. You can
// specify multiple values for a filter. A tag must match at least one of the
// specified values for it to be included in the results. You can also specify
// multiple filters. The result includes information for a particular tag only if
// it matches all the filters. If there's no match, no special message is returned.
// For more information, see Tag Auto Scaling groups and instances
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-tagging.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) {
	if params == nil {
		params = &DescribeTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTags", params, optFns, c.addOperationDescribeTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTagsInput struct {

	// One or more filters to scope the tags to return. The maximum number of filters
	// per filter type (for example, auto-scaling-group) is 1000.
	Filters []types.Filter

	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100.
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeTagsOutput struct {

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// One or more tags.
	Tags []types.TagDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTags{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTags(options.Region), middleware.Before); err != nil {
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

// DescribeTagsAPIClient is a client that implements the DescribeTags operation.
type DescribeTagsAPIClient interface {
	DescribeTags(context.Context, *DescribeTagsInput, ...func(*Options)) (*DescribeTagsOutput, error)
}

var _ DescribeTagsAPIClient = (*Client)(nil)

// DescribeTagsPaginatorOptions is the paginator options for DescribeTags
type DescribeTagsPaginatorOptions struct {
	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTagsPaginator is a paginator for DescribeTags
type DescribeTagsPaginator struct {
	options   DescribeTagsPaginatorOptions
	client    DescribeTagsAPIClient
	params    *DescribeTagsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTagsPaginator returns a new DescribeTagsPaginator
func NewDescribeTagsPaginator(client DescribeTagsAPIClient, params *DescribeTagsInput, optFns ...func(*DescribeTagsPaginatorOptions)) *DescribeTagsPaginator {
	if params == nil {
		params = &DescribeTagsInput{}
	}

	options := DescribeTagsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTags page.
func (p *DescribeTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTagsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeTags(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeTags",
	}
}
