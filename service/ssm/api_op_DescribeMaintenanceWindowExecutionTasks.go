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

// For a given maintenance window execution, lists the tasks that were run.
func (c *Client) DescribeMaintenanceWindowExecutionTasks(ctx context.Context, params *DescribeMaintenanceWindowExecutionTasksInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionTasksOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowExecutionTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindowExecutionTasks", params, optFns, c.addOperationDescribeMaintenanceWindowExecutionTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowExecutionTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowExecutionTasksInput struct {

	// The ID of the maintenance window execution whose task executions should be
	// retrieved.
	//
	// This member is required.
	WindowExecutionId *string

	// Optional filters used to scope down the returned tasks. The supported filter key
	// is STATUS with the corresponding values PENDING, IN_PROGRESS, SUCCESS, FAILED,
	// TIMED_OUT, CANCELLING, and CANCELLED.
	Filters []types.MaintenanceWindowFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeMaintenanceWindowExecutionTasksOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the task executions.
	WindowExecutionTaskIdentities []types.MaintenanceWindowExecutionTaskIdentity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMaintenanceWindowExecutionTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindowExecutionTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindowExecutionTasks{}, middleware.After)
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
	if err = addOpDescribeMaintenanceWindowExecutionTasksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutionTasks(options.Region), middleware.Before); err != nil {
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

// DescribeMaintenanceWindowExecutionTasksAPIClient is a client that implements the
// DescribeMaintenanceWindowExecutionTasks operation.
type DescribeMaintenanceWindowExecutionTasksAPIClient interface {
	DescribeMaintenanceWindowExecutionTasks(context.Context, *DescribeMaintenanceWindowExecutionTasksInput, ...func(*Options)) (*DescribeMaintenanceWindowExecutionTasksOutput, error)
}

var _ DescribeMaintenanceWindowExecutionTasksAPIClient = (*Client)(nil)

// DescribeMaintenanceWindowExecutionTasksPaginatorOptions is the paginator options
// for DescribeMaintenanceWindowExecutionTasks
type DescribeMaintenanceWindowExecutionTasksPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowExecutionTasksPaginator is a paginator for
// DescribeMaintenanceWindowExecutionTasks
type DescribeMaintenanceWindowExecutionTasksPaginator struct {
	options   DescribeMaintenanceWindowExecutionTasksPaginatorOptions
	client    DescribeMaintenanceWindowExecutionTasksAPIClient
	params    *DescribeMaintenanceWindowExecutionTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowExecutionTasksPaginator returns a new
// DescribeMaintenanceWindowExecutionTasksPaginator
func NewDescribeMaintenanceWindowExecutionTasksPaginator(client DescribeMaintenanceWindowExecutionTasksAPIClient, params *DescribeMaintenanceWindowExecutionTasksInput, optFns ...func(*DescribeMaintenanceWindowExecutionTasksPaginatorOptions)) *DescribeMaintenanceWindowExecutionTasksPaginator {
	if params == nil {
		params = &DescribeMaintenanceWindowExecutionTasksInput{}
	}

	options := DescribeMaintenanceWindowExecutionTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMaintenanceWindowExecutionTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowExecutionTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMaintenanceWindowExecutionTasks page.
func (p *DescribeMaintenanceWindowExecutionTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowExecutionTasksOutput, error) {
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

	result, err := p.client.DescribeMaintenanceWindowExecutionTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMaintenanceWindowExecutionTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindowExecutionTasks",
	}
}
