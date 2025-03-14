// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes a version of a SageMaker image.
func (c *Client) DescribeImageVersion(ctx context.Context, params *DescribeImageVersionInput, optFns ...func(*Options)) (*DescribeImageVersionOutput, error) {
	if params == nil {
		params = &DescribeImageVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImageVersion", params, optFns, c.addOperationDescribeImageVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImageVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImageVersionInput struct {

	// The name of the image.
	//
	// This member is required.
	ImageName *string

	// The alias of the image version.
	Alias *string

	// The version of the image. If not specified, the latest version is described.
	Version *int32

	noSmithyDocumentSerde
}

type DescribeImageVersionOutput struct {

	// The registry path of the container image on which this image version is based.
	BaseImage *string

	// The registry path of the container image that contains this image version.
	ContainerImage *string

	// When the version was created.
	CreationTime *time.Time

	// When a create or delete operation fails, the reason for the failure.
	FailureReason *string

	// Indicates Horovod compatibility.
	Horovod bool

	// The ARN of the image the version is based on.
	ImageArn *string

	// The ARN of the version.
	ImageVersionArn *string

	// The status of the version.
	ImageVersionStatus types.ImageVersionStatus

	// Indicates SageMaker job type compatibility.
	//
	// * TRAINING: The image version is
	// compatible with SageMaker training jobs.
	//
	// * INFERENCE: The image version is
	// compatible with SageMaker inference jobs.
	//
	// * NOTEBOOK_KERNEL: The image version
	// is compatible with SageMaker notebook kernels.
	JobType types.JobType

	// When the version was last modified.
	LastModifiedTime *time.Time

	// The machine learning framework vended in the image version.
	MLFramework *string

	// Indicates CPU or GPU compatibility.
	//
	// * CPU: The image version is compatible with
	// CPU.
	//
	// * GPU: The image version is compatible with GPU.
	Processor types.Processor

	// The supported programming language and its version.
	ProgrammingLang *string

	// The maintainer description of the image version.
	ReleaseNotes *string

	// The stability of the image version specified by the maintainer.
	//
	// * NOT_PROVIDED:
	// The maintainers did not provide a status for image version stability.
	//
	// * STABLE:
	// The image version is stable.
	//
	// * TO_BE_ARCHIVED: The image version is set to be
	// archived. Custom image versions that are set to be archived are automatically
	// archived after three months.
	//
	// * ARCHIVED: The image version is archived.
	// Archived image versions are not searchable and are no longer actively supported.
	VendorGuidance types.VendorGuidance

	// The version number.
	Version *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImageVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeImageVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeImageVersion{}, middleware.After)
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
	if err = addOpDescribeImageVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageVersion(options.Region), middleware.Before); err != nil {
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

// DescribeImageVersionAPIClient is a client that implements the
// DescribeImageVersion operation.
type DescribeImageVersionAPIClient interface {
	DescribeImageVersion(context.Context, *DescribeImageVersionInput, ...func(*Options)) (*DescribeImageVersionOutput, error)
}

var _ DescribeImageVersionAPIClient = (*Client)(nil)

// ImageVersionCreatedWaiterOptions are waiter options for
// ImageVersionCreatedWaiter
type ImageVersionCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ImageVersionCreatedWaiter will use default minimum delay of 60 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ImageVersionCreatedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeImageVersionInput, *DescribeImageVersionOutput, error) (bool, error)
}

// ImageVersionCreatedWaiter defines the waiters for ImageVersionCreated
type ImageVersionCreatedWaiter struct {
	client DescribeImageVersionAPIClient

	options ImageVersionCreatedWaiterOptions
}

// NewImageVersionCreatedWaiter constructs a ImageVersionCreatedWaiter.
func NewImageVersionCreatedWaiter(client DescribeImageVersionAPIClient, optFns ...func(*ImageVersionCreatedWaiterOptions)) *ImageVersionCreatedWaiter {
	options := ImageVersionCreatedWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = imageVersionCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ImageVersionCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ImageVersionCreated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *ImageVersionCreatedWaiter) Wait(ctx context.Context, params *DescribeImageVersionInput, maxWaitDur time.Duration, optFns ...func(*ImageVersionCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ImageVersionCreated waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ImageVersionCreatedWaiter) WaitForOutput(ctx context.Context, params *DescribeImageVersionInput, maxWaitDur time.Duration, optFns ...func(*ImageVersionCreatedWaiterOptions)) (*DescribeImageVersionOutput, error) {
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

		out, err := w.client.DescribeImageVersion(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ImageVersionCreated waiter")
}

func imageVersionCreatedStateRetryable(ctx context.Context, input *DescribeImageVersionInput, output *DescribeImageVersionOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("ImageVersionStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATED"
		value, ok := pathValue.(types.ImageVersionStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ImageVersionStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ImageVersionStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED"
		value, ok := pathValue.(types.ImageVersionStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ImageVersionStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "ValidationException" == apiErr.ErrorCode() {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// ImageVersionDeletedWaiterOptions are waiter options for
// ImageVersionDeletedWaiter
type ImageVersionDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ImageVersionDeletedWaiter will use default minimum delay of 60 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ImageVersionDeletedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeImageVersionInput, *DescribeImageVersionOutput, error) (bool, error)
}

// ImageVersionDeletedWaiter defines the waiters for ImageVersionDeleted
type ImageVersionDeletedWaiter struct {
	client DescribeImageVersionAPIClient

	options ImageVersionDeletedWaiterOptions
}

// NewImageVersionDeletedWaiter constructs a ImageVersionDeletedWaiter.
func NewImageVersionDeletedWaiter(client DescribeImageVersionAPIClient, optFns ...func(*ImageVersionDeletedWaiterOptions)) *ImageVersionDeletedWaiter {
	options := ImageVersionDeletedWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = imageVersionDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ImageVersionDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ImageVersionDeleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *ImageVersionDeletedWaiter) Wait(ctx context.Context, params *DescribeImageVersionInput, maxWaitDur time.Duration, optFns ...func(*ImageVersionDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for ImageVersionDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *ImageVersionDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeImageVersionInput, maxWaitDur time.Duration, optFns ...func(*ImageVersionDeletedWaiterOptions)) (*DescribeImageVersionOutput, error) {
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

		out, err := w.client.DescribeImageVersion(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for ImageVersionDeleted waiter")
}

func imageVersionDeletedStateRetryable(ctx context.Context, input *DescribeImageVersionInput, output *DescribeImageVersionOutput, err error) (bool, error) {

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "ResourceNotFoundException" == apiErr.ErrorCode() {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("ImageVersionStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETE_FAILED"
		value, ok := pathValue.(types.ImageVersionStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ImageVersionStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "ValidationException" == apiErr.ErrorCode() {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeImageVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeImageVersion",
	}
}
