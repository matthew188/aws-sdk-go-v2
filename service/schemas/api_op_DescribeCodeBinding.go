// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/schemas/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describe the code binding URI.
func (c *Client) DescribeCodeBinding(ctx context.Context, params *DescribeCodeBindingInput, optFns ...func(*Options)) (*DescribeCodeBindingOutput, error) {
	if params == nil {
		params = &DescribeCodeBindingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCodeBinding", params, optFns, c.addOperationDescribeCodeBindingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCodeBindingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCodeBindingInput struct {

	// The language of the code binding.
	//
	// This member is required.
	Language *string

	// The name of the registry.
	//
	// This member is required.
	RegistryName *string

	// The name of the schema.
	//
	// This member is required.
	SchemaName *string

	// Specifying this limits the results to only this schema version.
	SchemaVersion *string

	noSmithyDocumentSerde
}

type DescribeCodeBindingOutput struct {

	// The time and date that the code binding was created.
	CreationDate *time.Time

	// The date and time that code bindings were modified.
	LastModified *time.Time

	// The version number of the schema.
	SchemaVersion *string

	// The current status of code binding generation.
	Status types.CodeGenerationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCodeBindingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCodeBinding{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCodeBinding{}, middleware.After)
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
	if err = addOpDescribeCodeBindingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCodeBinding(options.Region), middleware.Before); err != nil {
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

// DescribeCodeBindingAPIClient is a client that implements the DescribeCodeBinding
// operation.
type DescribeCodeBindingAPIClient interface {
	DescribeCodeBinding(context.Context, *DescribeCodeBindingInput, ...func(*Options)) (*DescribeCodeBindingOutput, error)
}

var _ DescribeCodeBindingAPIClient = (*Client)(nil)

// CodeBindingExistsWaiterOptions are waiter options for CodeBindingExistsWaiter
type CodeBindingExistsWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// CodeBindingExistsWaiter will use default minimum delay of 2 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, CodeBindingExistsWaiter will use default max delay of 120 seconds. Note
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
	Retryable func(context.Context, *DescribeCodeBindingInput, *DescribeCodeBindingOutput, error) (bool, error)
}

// CodeBindingExistsWaiter defines the waiters for CodeBindingExists
type CodeBindingExistsWaiter struct {
	client DescribeCodeBindingAPIClient

	options CodeBindingExistsWaiterOptions
}

// NewCodeBindingExistsWaiter constructs a CodeBindingExistsWaiter.
func NewCodeBindingExistsWaiter(client DescribeCodeBindingAPIClient, optFns ...func(*CodeBindingExistsWaiterOptions)) *CodeBindingExistsWaiter {
	options := CodeBindingExistsWaiterOptions{}
	options.MinDelay = 2 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = codeBindingExistsStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &CodeBindingExistsWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for CodeBindingExists waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *CodeBindingExistsWaiter) Wait(ctx context.Context, params *DescribeCodeBindingInput, maxWaitDur time.Duration, optFns ...func(*CodeBindingExistsWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for CodeBindingExists waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *CodeBindingExistsWaiter) WaitForOutput(ctx context.Context, params *DescribeCodeBindingInput, maxWaitDur time.Duration, optFns ...func(*CodeBindingExistsWaiterOptions)) (*DescribeCodeBindingOutput, error) {
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

		out, err := w.client.DescribeCodeBinding(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for CodeBindingExists waiter")
}

func codeBindingExistsStateRetryable(ctx context.Context, input *DescribeCodeBindingInput, output *DescribeCodeBindingOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_COMPLETE"
		value, ok := pathValue.(types.CodeGenerationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CodeGenerationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_IN_PROGRESS"
		value, ok := pathValue.(types.CodeGenerationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CodeGenerationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED"
		value, ok := pathValue.(types.CodeGenerationStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.CodeGenerationStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var errorType *types.NotFoundException
		if errors.As(err, &errorType) {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeCodeBinding(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "DescribeCodeBinding",
	}
}
