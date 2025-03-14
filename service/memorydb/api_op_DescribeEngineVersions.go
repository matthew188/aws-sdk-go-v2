// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the available Redis engine versions.
func (c *Client) DescribeEngineVersions(ctx context.Context, params *DescribeEngineVersionsInput, optFns ...func(*Options)) (*DescribeEngineVersionsOutput, error) {
	if params == nil {
		params = &DescribeEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEngineVersions", params, optFns, c.addOperationDescribeEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEngineVersionsInput struct {

	// If true, specifies that only the default version of the specified engine or
	// engine and major version combination is to be returned.
	DefaultOnly bool

	// The Redis engine version
	EngineVersion *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// The name of a specific parameter group family to return details for.
	ParameterGroupFamily *string

	noSmithyDocumentSerde
}

type DescribeEngineVersionsOutput struct {

	// A list of engine version details. Each element in the list contains detailed
	// information about one engine version.
	EngineVersions []types.EngineVersionInfo

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEngineVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEngineVersions(options.Region), middleware.Before); err != nil {
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

// DescribeEngineVersionsAPIClient is a client that implements the
// DescribeEngineVersions operation.
type DescribeEngineVersionsAPIClient interface {
	DescribeEngineVersions(context.Context, *DescribeEngineVersionsInput, ...func(*Options)) (*DescribeEngineVersionsOutput, error)
}

var _ DescribeEngineVersionsAPIClient = (*Client)(nil)

// DescribeEngineVersionsPaginatorOptions is the paginator options for
// DescribeEngineVersions
type DescribeEngineVersionsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEngineVersionsPaginator is a paginator for DescribeEngineVersions
type DescribeEngineVersionsPaginator struct {
	options   DescribeEngineVersionsPaginatorOptions
	client    DescribeEngineVersionsAPIClient
	params    *DescribeEngineVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeEngineVersionsPaginator returns a new DescribeEngineVersionsPaginator
func NewDescribeEngineVersionsPaginator(client DescribeEngineVersionsAPIClient, params *DescribeEngineVersionsInput, optFns ...func(*DescribeEngineVersionsPaginatorOptions)) *DescribeEngineVersionsPaginator {
	if params == nil {
		params = &DescribeEngineVersionsInput{}
	}

	options := DescribeEngineVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEngineVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEngineVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEngineVersions page.
func (p *DescribeEngineVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEngineVersionsOutput, error) {
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

	result, err := p.client.DescribeEngineVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEngineVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "DescribeEngineVersions",
	}
}
