// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the custom action targets in Security Hub in your account.
func (c *Client) DescribeActionTargets(ctx context.Context, params *DescribeActionTargetsInput, optFns ...func(*Options)) (*DescribeActionTargetsOutput, error) {
	if params == nil {
		params = &DescribeActionTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeActionTargets", params, optFns, c.addOperationDescribeActionTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeActionTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActionTargetsInput struct {

	// A list of custom action target ARNs for the custom action targets to retrieve.
	ActionTargetArns []string

	// The maximum number of results to return.
	MaxResults int32

	// The token that is required for pagination. On your first call to the
	// DescribeActionTargets operation, set the value of this parameter to NULL. For
	// subsequent calls to the operation, to continue listing data, set the value of
	// this parameter to the value returned from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeActionTargetsOutput struct {

	// A list of ActionTarget objects. Each object includes the ActionTargetArn,
	// Description, and Name of a custom action target available in Security Hub.
	//
	// This member is required.
	ActionTargets []types.ActionTarget

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeActionTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeActionTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeActionTargets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActionTargets(options.Region), middleware.Before); err != nil {
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

// DescribeActionTargetsAPIClient is a client that implements the
// DescribeActionTargets operation.
type DescribeActionTargetsAPIClient interface {
	DescribeActionTargets(context.Context, *DescribeActionTargetsInput, ...func(*Options)) (*DescribeActionTargetsOutput, error)
}

var _ DescribeActionTargetsAPIClient = (*Client)(nil)

// DescribeActionTargetsPaginatorOptions is the paginator options for
// DescribeActionTargets
type DescribeActionTargetsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeActionTargetsPaginator is a paginator for DescribeActionTargets
type DescribeActionTargetsPaginator struct {
	options   DescribeActionTargetsPaginatorOptions
	client    DescribeActionTargetsAPIClient
	params    *DescribeActionTargetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeActionTargetsPaginator returns a new DescribeActionTargetsPaginator
func NewDescribeActionTargetsPaginator(client DescribeActionTargetsAPIClient, params *DescribeActionTargetsInput, optFns ...func(*DescribeActionTargetsPaginatorOptions)) *DescribeActionTargetsPaginator {
	if params == nil {
		params = &DescribeActionTargetsInput{}
	}

	options := DescribeActionTargetsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeActionTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeActionTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeActionTargets page.
func (p *DescribeActionTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeActionTargetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeActionTargets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeActionTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DescribeActionTargets",
	}
}
