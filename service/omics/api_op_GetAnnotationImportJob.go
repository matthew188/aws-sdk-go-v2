// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets information about an annotation import job.
func (c *Client) GetAnnotationImportJob(ctx context.Context, params *GetAnnotationImportJobInput, optFns ...func(*Options)) (*GetAnnotationImportJobOutput, error) {
	if params == nil {
		params = &GetAnnotationImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAnnotationImportJob", params, optFns, c.addOperationGetAnnotationImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAnnotationImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAnnotationImportJobInput struct {

	// The job's ID.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type GetAnnotationImportJobOutput struct {

	// When the job completed.
	//
	// This member is required.
	CompletionTime *time.Time

	// When the job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The job's destination annotation store.
	//
	// This member is required.
	DestinationName *string

	// Formatting options for a file.
	//
	// This member is required.
	FormatOptions types.FormatOptions

	// The job's ID.
	//
	// This member is required.
	Id *string

	// The job's imported items.
	//
	// This member is required.
	Items []types.AnnotationImportItemDetail

	// The job's service role ARN.
	//
	// This member is required.
	RoleArn *string

	// The job's left normalization setting.
	//
	// This member is required.
	RunLeftNormalization bool

	// The job's status.
	//
	// This member is required.
	Status types.JobStatus

	// The job's status message.
	//
	// This member is required.
	StatusMessage *string

	// When the job was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAnnotationImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAnnotationImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAnnotationImportJob{}, middleware.After)
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
	if err = addEndpointPrefix_opGetAnnotationImportJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetAnnotationImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAnnotationImportJob(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetAnnotationImportJobMiddleware struct {
}

func (*endpointPrefix_opGetAnnotationImportJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetAnnotationImportJobMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetAnnotationImportJobMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetAnnotationImportJobMiddleware{}, `OperationSerializer`, middleware.After)
}

// GetAnnotationImportJobAPIClient is a client that implements the
// GetAnnotationImportJob operation.
type GetAnnotationImportJobAPIClient interface {
	GetAnnotationImportJob(context.Context, *GetAnnotationImportJobInput, ...func(*Options)) (*GetAnnotationImportJobOutput, error)
}

var _ GetAnnotationImportJobAPIClient = (*Client)(nil)

// AnnotationImportJobCreatedWaiterOptions are waiter options for
// AnnotationImportJobCreatedWaiter
type AnnotationImportJobCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AnnotationImportJobCreatedWaiter will use default minimum delay of 30 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, AnnotationImportJobCreatedWaiter will use default max delay of 600
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
	Retryable func(context.Context, *GetAnnotationImportJobInput, *GetAnnotationImportJobOutput, error) (bool, error)
}

// AnnotationImportJobCreatedWaiter defines the waiters for
// AnnotationImportJobCreated
type AnnotationImportJobCreatedWaiter struct {
	client GetAnnotationImportJobAPIClient

	options AnnotationImportJobCreatedWaiterOptions
}

// NewAnnotationImportJobCreatedWaiter constructs a
// AnnotationImportJobCreatedWaiter.
func NewAnnotationImportJobCreatedWaiter(client GetAnnotationImportJobAPIClient, optFns ...func(*AnnotationImportJobCreatedWaiterOptions)) *AnnotationImportJobCreatedWaiter {
	options := AnnotationImportJobCreatedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = annotationImportJobCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AnnotationImportJobCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AnnotationImportJobCreated waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *AnnotationImportJobCreatedWaiter) Wait(ctx context.Context, params *GetAnnotationImportJobInput, maxWaitDur time.Duration, optFns ...func(*AnnotationImportJobCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for AnnotationImportJobCreated waiter
// and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *AnnotationImportJobCreatedWaiter) WaitForOutput(ctx context.Context, params *GetAnnotationImportJobInput, maxWaitDur time.Duration, optFns ...func(*AnnotationImportJobCreatedWaiterOptions)) (*GetAnnotationImportJobOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
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

		out, err := w.client.GetAnnotationImportJob(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for AnnotationImportJobCreated waiter")
}

func annotationImportJobCreatedStateRetryable(ctx context.Context, input *GetAnnotationImportJobInput, output *GetAnnotationImportJobOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUBMITTED"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "IN_PROGRESS"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CANCELLED"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "COMPLETED"
		value, ok := pathValue.(types.JobStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetAnnotationImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "GetAnnotationImportJob",
	}
}
