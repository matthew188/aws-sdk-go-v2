// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an inference scheduler.
func (c *Client) StartInferenceScheduler(ctx context.Context, params *StartInferenceSchedulerInput, optFns ...func(*Options)) (*StartInferenceSchedulerOutput, error) {
	if params == nil {
		params = &StartInferenceSchedulerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartInferenceScheduler", params, optFns, c.addOperationStartInferenceSchedulerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartInferenceSchedulerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartInferenceSchedulerInput struct {

	// The name of the inference scheduler to be started.
	//
	// This member is required.
	InferenceSchedulerName *string

	noSmithyDocumentSerde
}

type StartInferenceSchedulerOutput struct {

	// The Amazon Resource Name (ARN) of the inference scheduler being started.
	InferenceSchedulerArn *string

	// The name of the inference scheduler being started.
	InferenceSchedulerName *string

	// The Amazon Resource Name (ARN) of the ML model being used by the inference
	// scheduler.
	ModelArn *string

	// The name of the ML model being used by the inference scheduler.
	ModelName *string

	// Indicates the status of the inference scheduler.
	Status types.InferenceSchedulerStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartInferenceSchedulerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartInferenceScheduler{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartInferenceScheduler{}, middleware.After)
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
	if err = addOpStartInferenceSchedulerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartInferenceScheduler(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartInferenceScheduler(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "StartInferenceScheduler",
	}
}
