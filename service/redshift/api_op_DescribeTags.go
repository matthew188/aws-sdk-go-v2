// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of tags. You can return tags from a specific resource by
// specifying an ARN, or you can return all tags for a given type of resource, such
// as clusters, snapshots, and so on. The following are limitations for
// DescribeTags:
//
// * You cannot specify an ARN and a resource-type value together in
// the same request.
//
// * You cannot use the MaxRecords and Marker parameters
// together with the ARN parameter.
//
// * The MaxRecords parameter can be a range from
// 10 to 50 results to return in a request.
//
// If you specify both tag keys and tag
// values in the same request, Amazon Redshift returns all resources that match any
// combination of the specified keys and values. For example, if you have owner and
// environment for tag keys, and admin and test for tag values, all resources that
// have any combination of those values are returned. If both tag keys and values
// are omitted from the request, resources are returned regardless of whether they
// have tag keys or values associated with them.
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

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// The maximum number or response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32

	// The Amazon Resource Name (ARN) for which you want to describe the tag or tags.
	// For example, arn:aws:redshift:us-east-2:123456789:cluster:t1.
	ResourceName *string

	// The type of resource with which you want to view tags. Valid resource types
	// are:
	//
	// * Cluster
	//
	// * CIDR/IP
	//
	// * EC2 security group
	//
	// * Snapshot
	//
	// * Cluster security
	// group
	//
	// * Subnet group
	//
	// * HSM connection
	//
	// * HSM certificate
	//
	// * Parameter group
	//
	// *
	// Snapshot copy grant
	//
	// For more information about Amazon Redshift resource types
	// and constructing ARNs, go to Specifying Policy Elements: Actions, Effects,
	// Resources, and Principals
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-overview.html#redshift-iam-access-control-specify-actions)
	// in the Amazon Redshift Cluster Management Guide.
	ResourceType *string

	// A tag key or keys for which you want to return all matching resources that are
	// associated with the specified key or keys. For example, suppose that you have
	// resources tagged with keys called owner and environment. If you specify both of
	// these tag keys in the request, Amazon Redshift returns a response with all
	// resources that have either or both of these tag keys associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching resources that
	// are associated with the specified value or values. For example, suppose that you
	// have resources tagged with values called admin and test. If you specify both of
	// these tag values in the request, Amazon Redshift returns a response with all
	// resources that have either or both of these tag values associated with them.
	TagValues []string

	noSmithyDocumentSerde
}

type DescribeTagsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of tags with their associated resources.
	TaggedResources []types.TaggedResource

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
	// The maximum number or response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
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
		nextToken: params.Marker,
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
	params.Marker = p.nextToken

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
	p.nextToken = result.Marker

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
		SigningName:   "redshift",
		OperationName: "DescribeTags",
	}
}
