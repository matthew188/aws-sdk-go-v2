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

// Describes a Redshift-managed VPC endpoint.
func (c *Client) DescribeEndpointAccess(ctx context.Context, params *DescribeEndpointAccessInput, optFns ...func(*Options)) (*DescribeEndpointAccessOutput, error) {
	if params == nil {
		params = &DescribeEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpointAccess", params, optFns, c.addOperationDescribeEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointAccessInput struct {

	// The cluster identifier associated with the described endpoint.
	ClusterIdentifier *string

	// The name of the endpoint to be described.
	EndpointName *string

	// An optional pagination token provided by a previous DescribeEndpointAccess
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by the MaxRecords parameter.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a Marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32

	// The Amazon Web Services account ID of the owner of the cluster.
	ResourceOwner *string

	// The virtual private cloud (VPC) identifier with access to the cluster.
	VpcId *string

	noSmithyDocumentSerde
}

type DescribeEndpointAccessOutput struct {

	// The list of endpoints with access to the cluster.
	EndpointAccessList []types.EndpointAccess

	// An optional pagination token provided by a previous DescribeEndpointAccess
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by the MaxRecords parameter.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEndpointAccess{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpointAccess(options.Region), middleware.Before); err != nil {
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

// DescribeEndpointAccessAPIClient is a client that implements the
// DescribeEndpointAccess operation.
type DescribeEndpointAccessAPIClient interface {
	DescribeEndpointAccess(context.Context, *DescribeEndpointAccessInput, ...func(*Options)) (*DescribeEndpointAccessOutput, error)
}

var _ DescribeEndpointAccessAPIClient = (*Client)(nil)

// DescribeEndpointAccessPaginatorOptions is the paginator options for
// DescribeEndpointAccess
type DescribeEndpointAccessPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a Marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEndpointAccessPaginator is a paginator for DescribeEndpointAccess
type DescribeEndpointAccessPaginator struct {
	options   DescribeEndpointAccessPaginatorOptions
	client    DescribeEndpointAccessAPIClient
	params    *DescribeEndpointAccessInput
	nextToken *string
	firstPage bool
}

// NewDescribeEndpointAccessPaginator returns a new DescribeEndpointAccessPaginator
func NewDescribeEndpointAccessPaginator(client DescribeEndpointAccessAPIClient, params *DescribeEndpointAccessInput, optFns ...func(*DescribeEndpointAccessPaginatorOptions)) *DescribeEndpointAccessPaginator {
	if params == nil {
		params = &DescribeEndpointAccessInput{}
	}

	options := DescribeEndpointAccessPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEndpointAccessPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEndpointAccessPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEndpointAccess page.
func (p *DescribeEndpointAccessPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEndpointAccessOutput, error) {
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

	result, err := p.client.DescribeEndpointAccess(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeEndpointAccess",
	}
}
