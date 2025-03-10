// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Retrieves information about a portal.
func (c *Client) DescribePortal(ctx context.Context, params *DescribePortalInput, optFns ...func(*Options)) (*DescribePortalOutput, error) {
	if params == nil {
		params = &DescribePortalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePortal", params, optFns, c.addOperationDescribePortalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePortalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePortalInput struct {

	// The ID of the portal.
	//
	// This member is required.
	PortalId *string

	noSmithyDocumentSerde
}

type DescribePortalOutput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the portal, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:portal/${PortalId}
	//
	// This member is required.
	PortalArn *string

	// The IAM Identity Center application generated client ID (used with IAM Identity
	// Center APIs). IoT SiteWise includes portalClientId for only portals that use IAM
	// Identity Center to authenticate users.
	//
	// This member is required.
	PortalClientId *string

	// The Amazon Web Services administrator's contact email address.
	//
	// This member is required.
	PortalContactEmail *string

	// The date the portal was created, in Unix epoch time.
	//
	// This member is required.
	PortalCreationDate *time.Time

	// The ID of the portal.
	//
	// This member is required.
	PortalId *string

	// The date the portal was last updated, in Unix epoch time.
	//
	// This member is required.
	PortalLastUpdateDate *time.Time

	// The name of the portal.
	//
	// This member is required.
	PortalName *string

	// The URL for the IoT SiteWise Monitor portal. You can use this URL to access
	// portals that use IAM Identity Center for authentication. For portals that use
	// IAM for authentication, you must use the IoT SiteWise console to get a URL that
	// you can use to access the portal.
	//
	// This member is required.
	PortalStartUrl *string

	// The current status of the portal, which contains a state and any error message.
	//
	// This member is required.
	PortalStatus *types.PortalStatus

	// Contains the configuration information of an alarm created in an IoT SiteWise
	// Monitor portal.
	Alarms *types.Alarms

	// The email address that sends alarm notifications.
	NotificationSenderEmail *string

	// The service to use to authenticate users to the portal.
	PortalAuthMode types.AuthMode

	// The portal's description.
	PortalDescription *string

	// The portal's logo image, which is available at a URL.
	PortalLogoImageLocation *types.ImageLocation

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the service role that allows the portal's users to access your IoT SiteWise
	// resources on your behalf. For more information, see Using service roles for IoT
	// SiteWise Monitor
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/monitor-service-role.html)
	// in the IoT SiteWise User Guide.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePortalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePortal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePortal{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribePortalMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribePortalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePortal(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribePortalMiddleware struct {
}

func (*endpointPrefix_opDescribePortalMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribePortalMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "monitor." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribePortalMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribePortalMiddleware{}, `OperationSerializer`, middleware.After)
}

// DescribePortalAPIClient is a client that implements the DescribePortal
// operation.
type DescribePortalAPIClient interface {
	DescribePortal(context.Context, *DescribePortalInput, ...func(*Options)) (*DescribePortalOutput, error)
}

var _ DescribePortalAPIClient = (*Client)(nil)

// PortalActiveWaiterOptions are waiter options for PortalActiveWaiter
type PortalActiveWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// PortalActiveWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, PortalActiveWaiter will use default max delay of 120 seconds. Note that
	// MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribePortalInput, *DescribePortalOutput, error) (bool, error)
}

// PortalActiveWaiter defines the waiters for PortalActive
type PortalActiveWaiter struct {
	client DescribePortalAPIClient

	options PortalActiveWaiterOptions
}

// NewPortalActiveWaiter constructs a PortalActiveWaiter.
func NewPortalActiveWaiter(client DescribePortalAPIClient, optFns ...func(*PortalActiveWaiterOptions)) *PortalActiveWaiter {
	options := PortalActiveWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = portalActiveStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &PortalActiveWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for PortalActive waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *PortalActiveWaiter) Wait(ctx context.Context, params *DescribePortalInput, maxWaitDur time.Duration, optFns ...func(*PortalActiveWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for PortalActive waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *PortalActiveWaiter) WaitForOutput(ctx context.Context, params *DescribePortalInput, maxWaitDur time.Duration, optFns ...func(*PortalActiveWaiterOptions)) (*DescribePortalOutput, error) {
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

		out, err := w.client.DescribePortal(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for PortalActive waiter")
}

func portalActiveStateRetryable(ctx context.Context, input *DescribePortalInput, output *DescribePortalOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("portalStatus.state", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.PortalState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.PortalState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

// PortalNotExistsWaiterOptions are waiter options for PortalNotExistsWaiter
type PortalNotExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// PortalNotExistsWaiter will use default minimum delay of 3 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, PortalNotExistsWaiter will use default max delay of 120 seconds. Note
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
	Retryable func(context.Context, *DescribePortalInput, *DescribePortalOutput, error) (bool, error)
}

// PortalNotExistsWaiter defines the waiters for PortalNotExists
type PortalNotExistsWaiter struct {
	client DescribePortalAPIClient

	options PortalNotExistsWaiterOptions
}

// NewPortalNotExistsWaiter constructs a PortalNotExistsWaiter.
func NewPortalNotExistsWaiter(client DescribePortalAPIClient, optFns ...func(*PortalNotExistsWaiterOptions)) *PortalNotExistsWaiter {
	options := PortalNotExistsWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = portalNotExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &PortalNotExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for PortalNotExists waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *PortalNotExistsWaiter) Wait(ctx context.Context, params *DescribePortalInput, maxWaitDur time.Duration, optFns ...func(*PortalNotExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for PortalNotExists waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *PortalNotExistsWaiter) WaitForOutput(ctx context.Context, params *DescribePortalInput, maxWaitDur time.Duration, optFns ...func(*PortalNotExistsWaiterOptions)) (*DescribePortalOutput, error) {
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

		out, err := w.client.DescribePortal(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for PortalNotExists waiter")
}

func portalNotExistsStateRetryable(ctx context.Context, input *DescribePortalInput, output *DescribePortalOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribePortal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribePortal",
	}
}
