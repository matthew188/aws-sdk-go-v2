// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of tasks. You can filter the results by cluster, task definition
// family, container instance, launch type, what IAM principal started the task, or
// by the desired status of the task. Recently stopped tasks might appear in the
// returned results. Currently, stopped tasks appear in the returned results for at
// least one hour.
func (c *Client) ListTasks(ctx context.Context, params *ListTasksInput, optFns ...func(*Options)) (*ListTasksOutput, error) {
	if params == nil {
		params = &ListTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTasks", params, optFns, c.addOperationListTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTasksInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster to use when
	// filtering the ListTasks results. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string

	// The container instance ID or full ARN of the container instance to use when
	// filtering the ListTasks results. Specifying a containerInstance limits the
	// results to tasks that belong to that container instance.
	ContainerInstance *string

	// The task desired status to use when filtering the ListTasks results. Specifying
	// a desiredStatus of STOPPED limits the results to tasks that Amazon ECS has set
	// the desired status to STOPPED. This can be useful for debugging tasks that
	// aren't starting properly or have died or finished. The default status filter is
	// RUNNING, which shows tasks that Amazon ECS has set the desired status to
	// RUNNING. Although you can filter results based on a desired status of PENDING,
	// this doesn't return any results. Amazon ECS never sets the desired status of a
	// task to that value (only a task's lastStatus may have a value of PENDING).
	DesiredStatus types.DesiredStatus

	// The name of the task definition family to use when filtering the ListTasks
	// results. Specifying a family limits the results to tasks that belong to that
	// family.
	Family *string

	// The launch type to use when filtering the ListTasks results.
	LaunchType types.LaunchType

	// The maximum number of task results that ListTasks returned in paginated output.
	// When this parameter is used, ListTasks only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another ListTasks request with the
	// returned nextToken value. This value can be between 1 and 100. If this parameter
	// isn't used, then ListTasks returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a ListTasks request indicating that more
	// results are available to fulfill the request and further calls will be needed.
	// If maxResults was provided, it's possible the number of results to be fewer than
	// maxResults. This token should be treated as an opaque identifier that is only
	// used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string

	// The name of the service to use when filtering the ListTasks results. Specifying
	// a serviceName limits the results to tasks that belong to that service.
	ServiceName *string

	// The startedBy value to filter the task results with. Specifying a startedBy
	// value limits the results to tasks that were started with that value. When you
	// specify startedBy as the filter, it must be the only filter that you use.
	StartedBy *string

	noSmithyDocumentSerde
}

type ListTasksOutput struct {

	// The nextToken value to include in a future ListTasks request. When the results
	// of a ListTasks request exceed maxResults, this value can be used to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// The list of task ARN entries for the ListTasks request.
	TaskArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTasks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTasks(options.Region), middleware.Before); err != nil {
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

// ListTasksAPIClient is a client that implements the ListTasks operation.
type ListTasksAPIClient interface {
	ListTasks(context.Context, *ListTasksInput, ...func(*Options)) (*ListTasksOutput, error)
}

var _ ListTasksAPIClient = (*Client)(nil)

// ListTasksPaginatorOptions is the paginator options for ListTasks
type ListTasksPaginatorOptions struct {
	// The maximum number of task results that ListTasks returned in paginated output.
	// When this parameter is used, ListTasks only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another ListTasks request with the
	// returned nextToken value. This value can be between 1 and 100. If this parameter
	// isn't used, then ListTasks returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTasksPaginator is a paginator for ListTasks
type ListTasksPaginator struct {
	options   ListTasksPaginatorOptions
	client    ListTasksAPIClient
	params    *ListTasksInput
	nextToken *string
	firstPage bool
}

// NewListTasksPaginator returns a new ListTasksPaginator
func NewListTasksPaginator(client ListTasksAPIClient, params *ListTasksInput, optFns ...func(*ListTasksPaginatorOptions)) *ListTasksPaginator {
	if params == nil {
		params = &ListTasksInput{}
	}

	options := ListTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTasks page.
func (p *ListTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTasksOutput, error) {
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

	result, err := p.client.ListTasks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListTasks",
	}
}
