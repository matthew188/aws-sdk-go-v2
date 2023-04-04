// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns one or more snapshot objects, which contain metadata about your cluster
// snapshots. By default, this operation returns information about all snapshots of
// all clusters that are owned by your Amazon Web Services account. No information
// is returned for snapshots owned by inactive Amazon Web Services accounts. If you
// specify both tag keys and tag values in the same request, Amazon Redshift
// returns all snapshots that match any combination of the specified keys and
// values. For example, if you have owner and environment for tag keys, and admin
// and test for tag values, all snapshots that have any combination of those values
// are returned. Only snapshots that you own are returned in the response; shared
// snapshots are not returned with the tag key and tag value request parameters. If
// both tag keys and values are omitted from the request, snapshots are returned
// regardless of whether they have tag keys or values associated with them.
func (c *Client) DescribeClusterSnapshots(ctx context.Context, params *DescribeClusterSnapshotsInput, optFns ...func(*Options)) (*DescribeClusterSnapshotsOutput, error) {
	if params == nil {
		params = &DescribeClusterSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusterSnapshots", params, optFns, c.addOperationDescribeClusterSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterSnapshotsInput struct {

	// A value that indicates whether to return snapshots only for an existing cluster.
	// You can perform table-level restore only by using a snapshot of an existing
	// cluster, that is, a cluster that has not been deleted. Values for this parameter
	// work as follows:
	//
	// * If ClusterExists is set to true, ClusterIdentifier is
	// required.
	//
	// * If ClusterExists is set to false and ClusterIdentifier isn't
	// specified, all snapshots associated with deleted clusters (orphaned snapshots)
	// are returned.
	//
	// * If ClusterExists is set to false and ClusterIdentifier is
	// specified for a deleted cluster, snapshots associated with that cluster are
	// returned.
	//
	// * If ClusterExists is set to false and ClusterIdentifier is specified
	// for an existing cluster, no snapshots are returned.
	ClusterExists *bool

	// The identifier of the cluster which generated the requested snapshots.
	ClusterIdentifier *string

	// A time value that requests only snapshots created at or before the specified
	// time. The time value is specified in ISO 8601 format. For more information about
	// ISO 8601, go to the ISO8601 Wikipedia page.
	// (http://en.wikipedia.org/wiki/ISO_8601) Example: 2012-07-16T18:00:00Z
	EndTime *time.Time

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterSnapshots request exceed
	// the value specified in MaxRecords, Amazon Web Services returns a value in the
	// Marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 500.
	MaxRecords *int32

	// The Amazon Web Services account used to create or copy the snapshot. Use this
	// field to filter the results to snapshots owned by a particular account. To
	// describe snapshots you own, either specify your Amazon Web Services account, or
	// do not specify the parameter.
	OwnerAccount *string

	// The Amazon Resource Name (ARN) of the snapshot associated with the message to
	// describe cluster snapshots.
	SnapshotArn *string

	// The snapshot identifier of the snapshot about which to return information.
	SnapshotIdentifier *string

	// The type of snapshots for which you are requesting information. By default,
	// snapshots of all types are returned. Valid Values: automated | manual
	SnapshotType *string

	//
	SortingEntities []types.SnapshotSortingEntity

	// A value that requests only snapshots created at or after the specified time. The
	// time value is specified in ISO 8601 format. For more information about ISO 8601,
	// go to the ISO8601 Wikipedia page. (http://en.wikipedia.org/wiki/ISO_8601)
	// Example: 2012-07-16T18:00:00Z
	StartTime *time.Time

	// A tag key or keys for which you want to return all matching cluster snapshots
	// that are associated with the specified key or keys. For example, suppose that
	// you have snapshots that are tagged with keys called owner and environment. If
	// you specify both of these tag keys in the request, Amazon Redshift returns a
	// response with the snapshots that have either or both of these tag keys
	// associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching cluster
	// snapshots that are associated with the specified tag value or values. For
	// example, suppose that you have snapshots that are tagged with values called
	// admin and test. If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the snapshots that have either or both of these
	// tag values associated with them.
	TagValues []string

	noSmithyDocumentSerde
}

// Contains the output from the DescribeClusterSnapshots action.
type DescribeClusterSnapshotsOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of Snapshot instances.
	Snapshots []types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterSnapshots{}, middleware.After)
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
	if err = addOpDescribeClusterSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterSnapshots(options.Region), middleware.Before); err != nil {
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

// DescribeClusterSnapshotsAPIClient is a client that implements the
// DescribeClusterSnapshots operation.
type DescribeClusterSnapshotsAPIClient interface {
	DescribeClusterSnapshots(context.Context, *DescribeClusterSnapshotsInput, ...func(*Options)) (*DescribeClusterSnapshotsOutput, error)
}

var _ DescribeClusterSnapshotsAPIClient = (*Client)(nil)

// DescribeClusterSnapshotsPaginatorOptions is the paginator options for
// DescribeClusterSnapshots
type DescribeClusterSnapshotsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClusterSnapshotsPaginator is a paginator for DescribeClusterSnapshots
type DescribeClusterSnapshotsPaginator struct {
	options   DescribeClusterSnapshotsPaginatorOptions
	client    DescribeClusterSnapshotsAPIClient
	params    *DescribeClusterSnapshotsInput
	nextToken *string
	firstPage bool
}

// NewDescribeClusterSnapshotsPaginator returns a new
// DescribeClusterSnapshotsPaginator
func NewDescribeClusterSnapshotsPaginator(client DescribeClusterSnapshotsAPIClient, params *DescribeClusterSnapshotsInput, optFns ...func(*DescribeClusterSnapshotsPaginatorOptions)) *DescribeClusterSnapshotsPaginator {
	if params == nil {
		params = &DescribeClusterSnapshotsInput{}
	}

	options := DescribeClusterSnapshotsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClusterSnapshotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClusterSnapshotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClusterSnapshots page.
func (p *DescribeClusterSnapshotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClusterSnapshotsOutput, error) {
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

	result, err := p.client.DescribeClusterSnapshots(ctx, &params, optFns...)
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

// SnapshotAvailableWaiterOptions are waiter options for SnapshotAvailableWaiter
type SnapshotAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// SnapshotAvailableWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, SnapshotAvailableWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeClusterSnapshotsInput, *DescribeClusterSnapshotsOutput, error) (bool, error)
}

// SnapshotAvailableWaiter defines the waiters for SnapshotAvailable
type SnapshotAvailableWaiter struct {
	client DescribeClusterSnapshotsAPIClient

	options SnapshotAvailableWaiterOptions
}

// NewSnapshotAvailableWaiter constructs a SnapshotAvailableWaiter.
func NewSnapshotAvailableWaiter(client DescribeClusterSnapshotsAPIClient, optFns ...func(*SnapshotAvailableWaiterOptions)) *SnapshotAvailableWaiter {
	options := SnapshotAvailableWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = snapshotAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &SnapshotAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for SnapshotAvailable waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *SnapshotAvailableWaiter) Wait(ctx context.Context, params *DescribeClusterSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*SnapshotAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for SnapshotAvailable waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *SnapshotAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeClusterSnapshotsInput, maxWaitDur time.Duration, optFns ...func(*SnapshotAvailableWaiterOptions)) (*DescribeClusterSnapshotsOutput, error) {
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

		out, err := w.client.DescribeClusterSnapshots(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for SnapshotAvailable waiter")
}

func snapshotAvailableStateRetryable(ctx context.Context, input *DescribeClusterSnapshotsInput, output *DescribeClusterSnapshotsOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Snapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Snapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "failed"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Snapshots[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeClusterSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusterSnapshots",
	}
}
