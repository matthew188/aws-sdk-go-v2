// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the targets registered with the maintenance window.
func (c *Client) DescribeMaintenanceWindowTargets(ctx context.Context, params *DescribeMaintenanceWindowTargetsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowTargetsOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowTargets", params, optFns, c.addOperationDescribeMaintenanceWindowTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowTargetsInput struct {

	// The ID of the maintenance window whose targets should be retrieved.
	//
	// This member is required.
	WindowId *string

	// Optional filters that can be used to narrow down the scope of the returned
	// window targets. The supported filter keys are Type, WindowTargetId, and
	// OwnerInformation.
	Filters []types.MaintenanceWindowFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeMaintenanceWindowTargetsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the targets in the maintenance window.
	Targets []types.MaintenanceWindowTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMaintenanceWindowTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowTargets{}, middleware.After)
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
	if err = addOpDescribeMaintenanceWindowTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowTargets(options.Region), middleware.Before); err != nil {
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

// DescribeMaintenanceWindowTargetsAPIClient is a client that implements the
// DescribeMaintenanceWindowTargets operation.
type DescribeMaintenanceWindowTargetsAPIClient interface {
	DescribeMaintenanceWindowTargets(context.Context, *DescribeMaintenanceWindowTargetsInput, ...func(*Options)) (*DescribeMaintenanceWindowTargetsOutput, error)
}

var _ DescribeMaintenanceWindowTargetsAPIClient = (*Client)(nil)

// DescribeMaintenanceWindowTargetsPaginatorOptions is the paginator options for
// DescribeMaintenanceWindowTargets
type DescribeMaintenanceWindowTargetsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowTargetsPaginator is a paginator for
// DescribeMaintenanceWindowTargets
type DescribeMaintenanceWindowTargetsPaginator struct {
	options   DescribeMaintenanceWindowTargetsPaginatorOptions
	client    DescribeMaintenanceWindowTargetsAPIClient
	params    *DescribeMaintenanceWindowTargetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowTargetsPaginator returns a new
// DescribeMaintenanceWindowTargetsPaginator
func NewDescribeMaintenanceWindowTargetsPaginator(client DescribeMaintenanceWindowTargetsAPIClient, params *DescribeMaintenanceWindowTargetsInput, optFns ...func(*DescribeMaintenanceWindowTargetsPaginatorOptions)) *DescribeMaintenanceWindowTargetsPaginator {
	if params == nil {
		params = &DescribeMaintenanceWindowTargetsInput{}
	}

	options := DescribeMaintenanceWindowTargetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMaintenanceWindowTargetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowTargetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMaintenanceWindowTargets page.
func (p *DescribeMaintenanceWindowTargetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowTargetsOutput, error) {
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

	result, err := p.client.DescribeMaintenanceWindowTargets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowTargets",
	}
}
