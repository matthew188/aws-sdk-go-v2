// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the description of a specific Amazon EFS access point if the
// AccessPointId is provided. If you provide an EFS FileSystemId, it returns
// descriptions of all access points for that file system. You can provide either
// an AccessPointId or a FileSystemId in the request, but not both. This operation
// requires permissions for the elasticfilesystem:DescribeAccessPoints action.
func (c *Client) DescribeAccessPoints(ctx context.Context, params *DescribeAccessPointsInput, optFns ...func(*Options)) (*DescribeAccessPointsOutput, error) {
	if params == nil {
		params = &DescribeAccessPointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccessPoints", params, optFns, c.addOperationDescribeAccessPointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccessPointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccessPointsInput struct {

	// (Optional) Specifies an EFS access point to describe in the response; mutually
	// exclusive with FileSystemId.
	AccessPointId *string

	// (Optional) If you provide a FileSystemId, EFS returns all access points for that
	// file system; mutually exclusive with AccessPointId.
	FileSystemId *string

	// (Optional) When retrieving all access points for a file system, you can
	// optionally specify the MaxItems parameter to limit the number of objects
	// returned in a response. The default value is 100.
	MaxResults *int32

	// NextToken is present if the response is paginated. You can use NextMarker in the
	// subsequent request to fetch the next page of access point descriptions.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAccessPointsOutput struct {

	// An array of access point descriptions.
	AccessPoints []types.AccessPointDescription

	// Present if there are more access points than returned in the response. You can
	// use the NextMarker in the subsequent request to fetch the additional
	// descriptions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccessPointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccessPoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccessPoints{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccessPoints(options.Region), middleware.Before); err != nil {
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

// DescribeAccessPointsAPIClient is a client that implements the
// DescribeAccessPoints operation.
type DescribeAccessPointsAPIClient interface {
	DescribeAccessPoints(context.Context, *DescribeAccessPointsInput, ...func(*Options)) (*DescribeAccessPointsOutput, error)
}

var _ DescribeAccessPointsAPIClient = (*Client)(nil)

// DescribeAccessPointsPaginatorOptions is the paginator options for
// DescribeAccessPoints
type DescribeAccessPointsPaginatorOptions struct {
	// (Optional) When retrieving all access points for a file system, you can
	// optionally specify the MaxItems parameter to limit the number of objects
	// returned in a response. The default value is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAccessPointsPaginator is a paginator for DescribeAccessPoints
type DescribeAccessPointsPaginator struct {
	options   DescribeAccessPointsPaginatorOptions
	client    DescribeAccessPointsAPIClient
	params    *DescribeAccessPointsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAccessPointsPaginator returns a new DescribeAccessPointsPaginator
func NewDescribeAccessPointsPaginator(client DescribeAccessPointsAPIClient, params *DescribeAccessPointsInput, optFns ...func(*DescribeAccessPointsPaginatorOptions)) *DescribeAccessPointsPaginator {
	if params == nil {
		params = &DescribeAccessPointsInput{}
	}

	options := DescribeAccessPointsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAccessPointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAccessPointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAccessPoints page.
func (p *DescribeAccessPointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAccessPointsOutput, error) {
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

	result, err := p.client.DescribeAccessPoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAccessPoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeAccessPoints",
	}
}
