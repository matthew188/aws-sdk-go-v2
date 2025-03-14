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

// Returns a list of all instance types offered. The results can be filtered by
// location (Region or Availability Zone). If no location is specified, the
// instance types offered in the current Region are returned.
func (c *Client) DescribeInstanceTypeOfferings(ctx context.Context, params *DescribeInstanceTypeOfferingsInput, optFns ...func(*Options)) (*DescribeInstanceTypeOfferingsOutput, error) {
	if params == nil {
		params = &DescribeInstanceTypeOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceTypeOfferings", params, optFns, c.addOperationDescribeInstanceTypeOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceTypeOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceTypeOfferingsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. Filter names and values are case-sensitive.
	//
	// * location -
	// This depends on the location type. For example, if the location type is region
	// (default), the location is the Region code (for example, us-east-2.)
	//
	// *
	// instance-type - The instance type. For example, c5.2xlarge.
	Filters []types.Filter

	// The location type.
	LocationType types.LocationType

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

type DescribeInstanceTypeOfferingsOutput struct {

	// The instance types offered.
	InstanceTypeOfferings []types.InstanceTypeOffering

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceTypeOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceTypeOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceTypeOfferings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceTypeOfferings(options.Region), middleware.Before); err != nil {
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

// DescribeInstanceTypeOfferingsAPIClient is a client that implements the
// DescribeInstanceTypeOfferings operation.
type DescribeInstanceTypeOfferingsAPIClient interface {
	DescribeInstanceTypeOfferings(context.Context, *DescribeInstanceTypeOfferingsInput, ...func(*Options)) (*DescribeInstanceTypeOfferingsOutput, error)
}

var _ DescribeInstanceTypeOfferingsAPIClient = (*Client)(nil)

// DescribeInstanceTypeOfferingsPaginatorOptions is the paginator options for
// DescribeInstanceTypeOfferings
type DescribeInstanceTypeOfferingsPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstanceTypeOfferingsPaginator is a paginator for
// DescribeInstanceTypeOfferings
type DescribeInstanceTypeOfferingsPaginator struct {
	options   DescribeInstanceTypeOfferingsPaginatorOptions
	client    DescribeInstanceTypeOfferingsAPIClient
	params    *DescribeInstanceTypeOfferingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstanceTypeOfferingsPaginator returns a new
// DescribeInstanceTypeOfferingsPaginator
func NewDescribeInstanceTypeOfferingsPaginator(client DescribeInstanceTypeOfferingsAPIClient, params *DescribeInstanceTypeOfferingsInput, optFns ...func(*DescribeInstanceTypeOfferingsPaginatorOptions)) *DescribeInstanceTypeOfferingsPaginator {
	if params == nil {
		params = &DescribeInstanceTypeOfferingsInput{}
	}

	options := DescribeInstanceTypeOfferingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstanceTypeOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstanceTypeOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstanceTypeOfferings page.
func (p *DescribeInstanceTypeOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstanceTypeOfferingsOutput, error) {
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

	result, err := p.client.DescribeInstanceTypeOfferings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInstanceTypeOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceTypeOfferings",
	}
}
