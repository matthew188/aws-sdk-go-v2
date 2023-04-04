// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a pipeline execution.
func (c *Client) StartPipelineExecution(ctx context.Context, params *StartPipelineExecutionInput, optFns ...func(*Options)) (*StartPipelineExecutionOutput, error) {
	if params == nil {
		params = &StartPipelineExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartPipelineExecution", params, optFns, c.addOperationStartPipelineExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartPipelineExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartPipelineExecutionInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than once.
	//
	// This member is required.
	ClientRequestToken *string

	// The name of the pipeline.
	//
	// This member is required.
	PipelineName *string

	// This configuration, if specified, overrides the parallelism configuration of the
	// parent pipeline for this specific run.
	ParallelismConfiguration *types.ParallelismConfiguration

	// The description of the pipeline execution.
	PipelineExecutionDescription *string

	// The display name of the pipeline execution.
	PipelineExecutionDisplayName *string

	// Contains a list of pipeline parameters. This list can be empty.
	PipelineParameters []types.Parameter

	noSmithyDocumentSerde
}

type StartPipelineExecutionOutput struct {

	// The Amazon Resource Name (ARN) of the pipeline execution.
	PipelineExecutionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartPipelineExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartPipelineExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartPipelineExecution{}, middleware.After)
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
	if err = addIdempotencyToken_opStartPipelineExecutionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartPipelineExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartPipelineExecution(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartPipelineExecution struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartPipelineExecution) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartPipelineExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartPipelineExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartPipelineExecutionInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartPipelineExecutionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartPipelineExecution{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartPipelineExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StartPipelineExecution",
	}
}
