// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"time"
)

// Retrieves the certificate signing request (CSR) for your private certificate
// authority (CA). The CSR is created when you call the CreateCertificateAuthority
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
// action. Sign the CSR with your Amazon Web Services Private CA-hosted or
// on-premises root or subordinate CA. Then import the signed certificate back into
// Amazon Web Services Private CA by calling the
// ImportCertificateAuthorityCertificate
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html)
// action. The CSR is returned as a base64 PEM-encoded string.
func (c *Client) GetCertificateAuthorityCsr(ctx context.Context, params *GetCertificateAuthorityCsrInput, optFns ...func(*Options)) (*GetCertificateAuthorityCsrOutput, error) {
	if params == nil {
		params = &GetCertificateAuthorityCsrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCertificateAuthorityCsr", params, optFns, c.addOperationGetCertificateAuthorityCsrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCertificateAuthorityCsrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificateAuthorityCsrInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called the
	// CreateCertificateAuthority
	// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html)
	// action. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	noSmithyDocumentSerde
}

type GetCertificateAuthorityCsrOutput struct {

	// The base64 PEM-encoded certificate signing request (CSR) for your private CA
	// certificate.
	Csr *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCertificateAuthorityCsrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificateAuthorityCsr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificateAuthorityCsr{}, middleware.After)
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
	if err = addOpGetCertificateAuthorityCsrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificateAuthorityCsr(options.Region), middleware.Before); err != nil {
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

// GetCertificateAuthorityCsrAPIClient is a client that implements the
// GetCertificateAuthorityCsr operation.
type GetCertificateAuthorityCsrAPIClient interface {
	GetCertificateAuthorityCsr(context.Context, *GetCertificateAuthorityCsrInput, ...func(*Options)) (*GetCertificateAuthorityCsrOutput, error)
}

var _ GetCertificateAuthorityCsrAPIClient = (*Client)(nil)

// CertificateAuthorityCSRCreatedWaiterOptions are waiter options for
// CertificateAuthorityCSRCreatedWaiter
type CertificateAuthorityCSRCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// CertificateAuthorityCSRCreatedWaiter will use default minimum delay of 3
	// seconds. Note that MinDelay must resolve to a value lesser than or equal to the
	// MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, CertificateAuthorityCSRCreatedWaiter will use default max delay of 120
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
	Retryable func(context.Context, *GetCertificateAuthorityCsrInput, *GetCertificateAuthorityCsrOutput, error) (bool, error)
}

// CertificateAuthorityCSRCreatedWaiter defines the waiters for
// CertificateAuthorityCSRCreated
type CertificateAuthorityCSRCreatedWaiter struct {
	client GetCertificateAuthorityCsrAPIClient

	options CertificateAuthorityCSRCreatedWaiterOptions
}

// NewCertificateAuthorityCSRCreatedWaiter constructs a
// CertificateAuthorityCSRCreatedWaiter.
func NewCertificateAuthorityCSRCreatedWaiter(client GetCertificateAuthorityCsrAPIClient, optFns ...func(*CertificateAuthorityCSRCreatedWaiterOptions)) *CertificateAuthorityCSRCreatedWaiter {
	options := CertificateAuthorityCSRCreatedWaiterOptions{}
	options.MinDelay = 3 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = certificateAuthorityCSRCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &CertificateAuthorityCSRCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for CertificateAuthorityCSRCreated waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *CertificateAuthorityCSRCreatedWaiter) Wait(ctx context.Context, params *GetCertificateAuthorityCsrInput, maxWaitDur time.Duration, optFns ...func(*CertificateAuthorityCSRCreatedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for CertificateAuthorityCSRCreated
// waiter and returns the output of the successful operation. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *CertificateAuthorityCSRCreatedWaiter) WaitForOutput(ctx context.Context, params *GetCertificateAuthorityCsrInput, maxWaitDur time.Duration, optFns ...func(*CertificateAuthorityCSRCreatedWaiterOptions)) (*GetCertificateAuthorityCsrOutput, error) {
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

		out, err := w.client.GetCertificateAuthorityCsr(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for CertificateAuthorityCSRCreated waiter")
}

func certificateAuthorityCSRCreatedStateRetryable(ctx context.Context, input *GetCertificateAuthorityCsrInput, output *GetCertificateAuthorityCsrOutput, err error) (bool, error) {

	if err == nil {
		return false, nil
	}

	if err != nil {
		var errorType *types.RequestInProgressException
		if errors.As(err, &errorType) {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetCertificateAuthorityCsr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "GetCertificateAuthorityCsr",
	}
}
