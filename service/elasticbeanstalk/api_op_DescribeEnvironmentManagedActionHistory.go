// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists an environment's completed and failed managed actions.
func (c *Client) DescribeEnvironmentManagedActionHistory(ctx context.Context, params *DescribeEnvironmentManagedActionHistoryInput, optFns ...func(*Options)) (*DescribeEnvironmentManagedActionHistoryOutput, error) {
	if params == nil {
		params = &DescribeEnvironmentManagedActionHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEnvironmentManagedActionHistory", params, optFns, c.addOperationDescribeEnvironmentManagedActionHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEnvironmentManagedActionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list completed and failed managed actions.
type DescribeEnvironmentManagedActionHistoryInput struct {

	// The environment ID of the target environment.
	EnvironmentId *string

	// The name of the target environment.
	EnvironmentName *string

	// The maximum number of items to return for a single request.
	MaxItems *int32

	// The pagination token returned by a previous request.
	NextToken *string

	noSmithyDocumentSerde
}

// A result message containing a list of completed and failed managed actions.
type DescribeEnvironmentManagedActionHistoryOutput struct {

	// A list of completed and failed managed actions.
	ManagedActionHistoryItems []types.ManagedActionHistoryItem

	// A pagination token that you pass to DescribeEnvironmentManagedActionHistory to
	// get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEnvironmentManagedActionHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEnvironmentManagedActionHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEnvironmentManagedActionHistory{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEnvironmentManagedActionHistory(options.Region), middleware.Before); err != nil {
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

// DescribeEnvironmentManagedActionHistoryAPIClient is a client that implements the
// DescribeEnvironmentManagedActionHistory operation.
type DescribeEnvironmentManagedActionHistoryAPIClient interface {
	DescribeEnvironmentManagedActionHistory(context.Context, *DescribeEnvironmentManagedActionHistoryInput, ...func(*Options)) (*DescribeEnvironmentManagedActionHistoryOutput, error)
}

var _ DescribeEnvironmentManagedActionHistoryAPIClient = (*Client)(nil)

// DescribeEnvironmentManagedActionHistoryPaginatorOptions is the paginator options
// for DescribeEnvironmentManagedActionHistory
type DescribeEnvironmentManagedActionHistoryPaginatorOptions struct {
	// The maximum number of items to return for a single request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEnvironmentManagedActionHistoryPaginator is a paginator for
// DescribeEnvironmentManagedActionHistory
type DescribeEnvironmentManagedActionHistoryPaginator struct {
	options   DescribeEnvironmentManagedActionHistoryPaginatorOptions
	client    DescribeEnvironmentManagedActionHistoryAPIClient
	params    *DescribeEnvironmentManagedActionHistoryInput
	nextToken *string
	firstPage bool
}

// NewDescribeEnvironmentManagedActionHistoryPaginator returns a new
// DescribeEnvironmentManagedActionHistoryPaginator
func NewDescribeEnvironmentManagedActionHistoryPaginator(client DescribeEnvironmentManagedActionHistoryAPIClient, params *DescribeEnvironmentManagedActionHistoryInput, optFns ...func(*DescribeEnvironmentManagedActionHistoryPaginatorOptions)) *DescribeEnvironmentManagedActionHistoryPaginator {
	if params == nil {
		params = &DescribeEnvironmentManagedActionHistoryInput{}
	}

	options := DescribeEnvironmentManagedActionHistoryPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEnvironmentManagedActionHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEnvironmentManagedActionHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEnvironmentManagedActionHistory page.
func (p *DescribeEnvironmentManagedActionHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEnvironmentManagedActionHistoryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.DescribeEnvironmentManagedActionHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeEnvironmentManagedActionHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeEnvironmentManagedActionHistory",
	}
}
