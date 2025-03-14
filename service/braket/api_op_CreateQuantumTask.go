// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a quantum task.
func (c *Client) CreateQuantumTask(ctx context.Context, params *CreateQuantumTaskInput, optFns ...func(*Options)) (*CreateQuantumTaskOutput, error) {
	if params == nil {
		params = &CreateQuantumTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateQuantumTask", params, optFns, c.addOperationCreateQuantumTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateQuantumTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateQuantumTaskInput struct {

	// The action associated with the task.
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Action *string

	// The client token associated with the request.
	//
	// This member is required.
	ClientToken *string

	// The ARN of the device to run the task on.
	//
	// This member is required.
	DeviceArn *string

	// The S3 bucket to store task result files in.
	//
	// This member is required.
	OutputS3Bucket *string

	// The key prefix for the location in the S3 bucket to store task results in.
	//
	// This member is required.
	OutputS3KeyPrefix *string

	// The number of shots to use for the task.
	//
	// This member is required.
	Shots *int64

	// The parameters for the device to run the task on.
	//
	// This value conforms to the media type: application/json
	DeviceParameters *string

	// The token for an Amazon Braket job that associates it with the quantum task.
	JobToken *string

	// Tags to be added to the quantum task you're creating.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateQuantumTaskOutput struct {

	// The ARN of the task created by the request.
	//
	// This member is required.
	QuantumTaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateQuantumTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateQuantumTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateQuantumTask{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateQuantumTaskMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateQuantumTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateQuantumTask(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateQuantumTask struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateQuantumTask) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateQuantumTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateQuantumTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateQuantumTaskInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateQuantumTaskMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateQuantumTask{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateQuantumTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "braket",
		OperationName: "CreateQuantumTask",
	}
}
