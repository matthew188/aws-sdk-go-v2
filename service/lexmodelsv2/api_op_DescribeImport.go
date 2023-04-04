// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets information about a specific import.
func (c *Client) DescribeImport(ctx context.Context, params *DescribeImportInput, optFns ...func(*Options)) (*DescribeImportOutput, error) {
	if params == nil {
		params = &DescribeImportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImport", params, optFns, c.addOperationDescribeImportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImportInput struct {

	// The unique identifier of the import to describe.
	//
	// This member is required.
	ImportId *string

	noSmithyDocumentSerde
}

type DescribeImportOutput struct {

	// The date and time that the import was created.
	CreationDateTime *time.Time

	// If the importStatus field is Failed, this provides one or more reasons for the
	// failure.
	FailureReasons []string

	// The unique identifier of the described import.
	ImportId *string

	// The status of the import process. When the status is Completed the resource is
	// imported and ready for use.
	ImportStatus types.ImportStatus

	// The unique identifier that Amazon Lex assigned to the resource created by the
	// import.
	ImportedResourceId *string

	// The name of the imported resource.
	ImportedResourceName *string

	// The date and time that the import was last updated.
	LastUpdatedDateTime *time.Time

	// The strategy used when there was a name conflict between the imported resource
	// and an existing resource. When the merge strategy is FailOnConflict existing
	// resources are not overwritten and the import fails.
	MergeStrategy types.MergeStrategy

	// The specifications of the imported bot, bot locale, or custom vocabulary.
	ResourceSpecification *types.ImportResourceSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeImport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeImport{}, middleware.After)
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
	if err = addOpDescribeImportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImport(options.Region), middleware.Before); err != nil {
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

// DescribeImportAPIClient is a client that implements the DescribeImport
// operation.
type DescribeImportAPIClient interface {
	DescribeImport(context.Context, *DescribeImportInput, ...func(*Options)) (*DescribeImportOutput, error)
}

var _ DescribeImportAPIClient = (*Client)(nil)

// BotImportCompletedWaiterOptions are waiter options for BotImportCompletedWaiter
type BotImportCompletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BotImportCompletedWaiter will use default minimum delay of 10 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, BotImportCompletedWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeImportInput, *DescribeImportOutput, error) (bool, error)
}

// BotImportCompletedWaiter defines the waiters for BotImportCompleted
type BotImportCompletedWaiter struct {
	client DescribeImportAPIClient

	options BotImportCompletedWaiterOptions
}

// NewBotImportCompletedWaiter constructs a BotImportCompletedWaiter.
func NewBotImportCompletedWaiter(client DescribeImportAPIClient, optFns ...func(*BotImportCompletedWaiterOptions)) *BotImportCompletedWaiter {
	options := BotImportCompletedWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = botImportCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BotImportCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BotImportCompleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *BotImportCompletedWaiter) Wait(ctx context.Context, params *DescribeImportInput, maxWaitDur time.Duration, optFns ...func(*BotImportCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for BotImportCompleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *BotImportCompletedWaiter) WaitForOutput(ctx context.Context, params *DescribeImportInput, maxWaitDur time.Duration, optFns ...func(*BotImportCompletedWaiterOptions)) (*DescribeImportOutput, error) {
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

		out, err := w.client.DescribeImport(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for BotImportCompleted waiter")
}

func botImportCompletedStateRetryable(ctx context.Context, input *DescribeImportInput, output *DescribeImportOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("importStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Completed"
		value, ok := pathValue.(types.ImportStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ImportStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("importStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Deleting"
		value, ok := pathValue.(types.ImportStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ImportStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("importStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.ImportStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ImportStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeImport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeImport",
	}
}
