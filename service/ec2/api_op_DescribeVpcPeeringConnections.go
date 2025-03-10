// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes one or more of your VPC peering connections.
func (c *Client) DescribeVpcPeeringConnections(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, optFns ...func(*Options)) (*DescribeVpcPeeringConnectionsOutput, error) {
	if params == nil {
		params = &DescribeVpcPeeringConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcPeeringConnections", params, optFns, c.addOperationDescribeVpcPeeringConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcPeeringConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcPeeringConnectionsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * accepter-vpc-info.cidr-block - The IPv4 CIDR block of
	// the accepter VPC.
	//
	// * accepter-vpc-info.owner-id - The ID of the Amazon Web
	// Services account that owns the accepter VPC.
	//
	// * accepter-vpc-info.vpc-id - The
	// ID of the accepter VPC.
	//
	// * expiration-time - The expiration date and time for
	// the VPC peering connection.
	//
	// * requester-vpc-info.cidr-block - The IPv4 CIDR
	// block of the requester's VPC.
	//
	// * requester-vpc-info.owner-id - The ID of the
	// Amazon Web Services account that owns the requester VPC.
	//
	// *
	// requester-vpc-info.vpc-id - The ID of the requester VPC.
	//
	// * status-code - The
	// status of the VPC peering connection (pending-acceptance | failed | expired |
	// provisioning | active | deleting | deleted | rejected).
	//
	// * status-message - A
	// message that provides more information about the status of the VPC peering
	// connection, if applicable.
	//
	// * tag: - The key/value combination of a tag assigned
	// to the resource. Use the tag key in the filter name and the tag value as the
	// filter value. For example, to find all resources that have a tag with the key
	// Owner and the value TeamA, specify tag:Owner for the filter name and TeamA for
	// the filter value.
	//
	// * tag-key - The key of a tag assigned to the resource. Use
	// this filter to find all resources assigned a tag with a specific key, regardless
	// of the tag value.
	//
	// * vpc-peering-connection-id - The ID of the VPC peering
	// connection.
	Filters []types.Filter

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	// One or more VPC peering connection IDs. Default: Describes all your VPC peering
	// connections.
	VpcPeeringConnectionIds []string

	noSmithyDocumentSerde
}

type DescribeVpcPeeringConnectionsOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// Information about the VPC peering connections.
	VpcPeeringConnections []types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcPeeringConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcPeeringConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcPeeringConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcPeeringConnections(options.Region), middleware.Before); err != nil {
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

// DescribeVpcPeeringConnectionsAPIClient is a client that implements the
// DescribeVpcPeeringConnections operation.
type DescribeVpcPeeringConnectionsAPIClient interface {
	DescribeVpcPeeringConnections(context.Context, *DescribeVpcPeeringConnectionsInput, ...func(*Options)) (*DescribeVpcPeeringConnectionsOutput, error)
}

var _ DescribeVpcPeeringConnectionsAPIClient = (*Client)(nil)

// DescribeVpcPeeringConnectionsPaginatorOptions is the paginator options for
// DescribeVpcPeeringConnections
type DescribeVpcPeeringConnectionsPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVpcPeeringConnectionsPaginator is a paginator for
// DescribeVpcPeeringConnections
type DescribeVpcPeeringConnectionsPaginator struct {
	options   DescribeVpcPeeringConnectionsPaginatorOptions
	client    DescribeVpcPeeringConnectionsAPIClient
	params    *DescribeVpcPeeringConnectionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVpcPeeringConnectionsPaginator returns a new
// DescribeVpcPeeringConnectionsPaginator
func NewDescribeVpcPeeringConnectionsPaginator(client DescribeVpcPeeringConnectionsAPIClient, params *DescribeVpcPeeringConnectionsInput, optFns ...func(*DescribeVpcPeeringConnectionsPaginatorOptions)) *DescribeVpcPeeringConnectionsPaginator {
	if params == nil {
		params = &DescribeVpcPeeringConnectionsInput{}
	}

	options := DescribeVpcPeeringConnectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVpcPeeringConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVpcPeeringConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVpcPeeringConnections page.
func (p *DescribeVpcPeeringConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVpcPeeringConnectionsOutput, error) {
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

	result, err := p.client.DescribeVpcPeeringConnections(ctx, &params, optFns...)
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

// VpcPeeringConnectionDeletedWaiterOptions are waiter options for
// VpcPeeringConnectionDeletedWaiter
type VpcPeeringConnectionDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VpcPeeringConnectionDeletedWaiter will use default minimum delay of 15 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, VpcPeeringConnectionDeletedWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeVpcPeeringConnectionsInput, *DescribeVpcPeeringConnectionsOutput, error) (bool, error)
}

// VpcPeeringConnectionDeletedWaiter defines the waiters for
// VpcPeeringConnectionDeleted
type VpcPeeringConnectionDeletedWaiter struct {
	client DescribeVpcPeeringConnectionsAPIClient

	options VpcPeeringConnectionDeletedWaiterOptions
}

// NewVpcPeeringConnectionDeletedWaiter constructs a
// VpcPeeringConnectionDeletedWaiter.
func NewVpcPeeringConnectionDeletedWaiter(client DescribeVpcPeeringConnectionsAPIClient, optFns ...func(*VpcPeeringConnectionDeletedWaiterOptions)) *VpcPeeringConnectionDeletedWaiter {
	options := VpcPeeringConnectionDeletedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = vpcPeeringConnectionDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VpcPeeringConnectionDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VpcPeeringConnectionDeleted waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *VpcPeeringConnectionDeletedWaiter) Wait(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, maxWaitDur time.Duration, optFns ...func(*VpcPeeringConnectionDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for VpcPeeringConnectionDeleted waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *VpcPeeringConnectionDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, maxWaitDur time.Duration, optFns ...func(*VpcPeeringConnectionDeletedWaiterOptions)) (*DescribeVpcPeeringConnectionsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeVpcPeeringConnections(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for VpcPeeringConnectionDeleted waiter")
}

func vpcPeeringConnectionDeletedStateRetryable(ctx context.Context, input *DescribeVpcPeeringConnectionsInput, output *DescribeVpcPeeringConnectionsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("VpcPeeringConnections[].Status.Code", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.VpcPeeringConnectionStateReasonCode)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.VpcPeeringConnectionStateReasonCode value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidVpcPeeringConnectionID.NotFound" == apiErr.ErrorCode() {
			return false, nil
		}
	}

	return true, nil
}

// VpcPeeringConnectionExistsWaiterOptions are waiter options for
// VpcPeeringConnectionExistsWaiter
type VpcPeeringConnectionExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// VpcPeeringConnectionExistsWaiter will use default minimum delay of 15 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, VpcPeeringConnectionExistsWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeVpcPeeringConnectionsInput, *DescribeVpcPeeringConnectionsOutput, error) (bool, error)
}

// VpcPeeringConnectionExistsWaiter defines the waiters for
// VpcPeeringConnectionExists
type VpcPeeringConnectionExistsWaiter struct {
	client DescribeVpcPeeringConnectionsAPIClient

	options VpcPeeringConnectionExistsWaiterOptions
}

// NewVpcPeeringConnectionExistsWaiter constructs a
// VpcPeeringConnectionExistsWaiter.
func NewVpcPeeringConnectionExistsWaiter(client DescribeVpcPeeringConnectionsAPIClient, optFns ...func(*VpcPeeringConnectionExistsWaiterOptions)) *VpcPeeringConnectionExistsWaiter {
	options := VpcPeeringConnectionExistsWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = vpcPeeringConnectionExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &VpcPeeringConnectionExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for VpcPeeringConnectionExists waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *VpcPeeringConnectionExistsWaiter) Wait(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, maxWaitDur time.Duration, optFns ...func(*VpcPeeringConnectionExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for VpcPeeringConnectionExists waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *VpcPeeringConnectionExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, maxWaitDur time.Duration, optFns ...func(*VpcPeeringConnectionExistsWaiterOptions)) (*DescribeVpcPeeringConnectionsOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeVpcPeeringConnections(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for VpcPeeringConnectionExists waiter")
}

func vpcPeeringConnectionExistsStateRetryable(ctx context.Context, input *DescribeVpcPeeringConnectionsInput, output *DescribeVpcPeeringConnectionsOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidVpcPeeringConnectionID.NotFound" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeVpcPeeringConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcPeeringConnections",
	}
}
