// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the type of endpoints available.
func (c *Client) DescribeEndpointTypes(ctx context.Context, params *DescribeEndpointTypesInput, optFns ...func(*Options)) (*DescribeEndpointTypesOutput, error) {
	if params == nil {
		params = &DescribeEndpointTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpointTypes", params, optFns, c.addOperationDescribeEndpointTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointTypesInput struct {

	// Filters applied to the endpoint types. Valid filter names: engine-name |
	// endpoint-type
	Filters []types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeEndpointTypesOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The types of endpoints that are supported.
	SupportedEndpointTypes []types.SupportedEndpointType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpointTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpointTypes{}, middleware.After)
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
	if err = addOpDescribeEndpointTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpointTypes(options.Region), middleware.Before); err != nil {
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

// DescribeEndpointTypesAPIClient is a client that implements the
// DescribeEndpointTypes operation.
type DescribeEndpointTypesAPIClient interface {
	DescribeEndpointTypes(context.Context, *DescribeEndpointTypesInput, ...func(*Options)) (*DescribeEndpointTypesOutput, error)
}

var _ DescribeEndpointTypesAPIClient = (*Client)(nil)

// DescribeEndpointTypesPaginatorOptions is the paginator options for
// DescribeEndpointTypes
type DescribeEndpointTypesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEndpointTypesPaginator is a paginator for DescribeEndpointTypes
type DescribeEndpointTypesPaginator struct {
	options   DescribeEndpointTypesPaginatorOptions
	client    DescribeEndpointTypesAPIClient
	params    *DescribeEndpointTypesInput
	nextToken *string
	firstPage bool
}

// NewDescribeEndpointTypesPaginator returns a new DescribeEndpointTypesPaginator
func NewDescribeEndpointTypesPaginator(client DescribeEndpointTypesAPIClient, params *DescribeEndpointTypesInput, optFns ...func(*DescribeEndpointTypesPaginatorOptions)) *DescribeEndpointTypesPaginator {
	if params == nil {
		params = &DescribeEndpointTypesInput{}
	}

	options := DescribeEndpointTypesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEndpointTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEndpointTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEndpointTypes page.
func (p *DescribeEndpointTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEndpointTypesOutput, error) {
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

	result, err := p.client.DescribeEndpointTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEndpointTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeEndpointTypes",
	}
}
