// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an ML model for data inference. A machine-learning (ML) model is a
// mathematical model that finds patterns in your data. In Amazon Lookout for
// Equipment, the model learns the patterns of normal behavior and detects abnormal
// behavior that could be potential equipment failure (or maintenance events). The
// models are made by analyzing normal data and abnormalities in machine behavior
// that have already occurred. Your model is trained using a portion of the data
// from your dataset and uses that data to learn patterns of normal behavior and
// abnormal patterns that lead to equipment failure. Another portion of the data is
// used to evaluate the model's accuracy.
func (c *Client) CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) {
	if params == nil {
		params = &CreateModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModel", params, optFns, c.addOperationCreateModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelInput struct {

	// A unique identifier for the request. If you do not set the client request token,
	// Amazon Lookout for Equipment generates one.
	//
	// This member is required.
	ClientToken *string

	// The name of the dataset for the ML model being created.
	//
	// This member is required.
	DatasetName *string

	// The name for the ML model to be created.
	//
	// This member is required.
	ModelName *string

	// The configuration is the TargetSamplingRate, which is the sampling rate of the
	// data after post processing by Amazon Lookout for Equipment. For example, if you
	// provide data that has been collected at a 1 second level and you want the system
	// to resample the data at a 1 minute rate before training, the TargetSamplingRate
	// is 1 minute. When providing a value for the TargetSamplingRate, you must attach
	// the prefix "PT" to the rate you want. The value for a 1 second rate is therefore
	// PT1S, the value for a 15 minute rate is PT15M, and the value for a 1 hour rate
	// is PT1H
	DataPreProcessingConfiguration *types.DataPreProcessingConfiguration

	// The data schema for the ML model being created.
	DatasetSchema *types.DatasetSchema

	// Indicates the time reference in the dataset that should be used to end the
	// subset of evaluation data for the ML model.
	EvaluationDataEndTime *time.Time

	// Indicates the time reference in the dataset that should be used to begin the
	// subset of evaluation data for the ML model.
	EvaluationDataStartTime *time.Time

	// The input configuration for the labels being used for the ML model that's being
	// created.
	LabelsInputConfiguration *types.LabelsInputConfiguration

	// Indicates that the asset associated with this sensor has been shut off. As long
	// as this condition is met, Lookout for Equipment will not use data from this
	// asset for training, evaluation, or inference.
	OffCondition *string

	// The Amazon Resource Name (ARN) of a role with permission to access the data
	// source being used to create the ML model.
	RoleArn *string

	// Provides the identifier of the KMS key used to encrypt model data by Amazon
	// Lookout for Equipment.
	ServerSideKmsKeyId *string

	// Any tags associated with the ML model being created.
	Tags []types.Tag

	// Indicates the time reference in the dataset that should be used to end the
	// subset of training data for the ML model.
	TrainingDataEndTime *time.Time

	// Indicates the time reference in the dataset that should be used to begin the
	// subset of training data for the ML model.
	TrainingDataStartTime *time.Time

	noSmithyDocumentSerde
}

type CreateModelOutput struct {

	// The Amazon Resource Name (ARN) of the model being created.
	ModelArn *string

	// Indicates the status of the CreateModel operation.
	Status types.ModelStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateModel{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateModelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModel(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateModelInput ")
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
func addIdempotencyToken_opCreateModelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "CreateModel",
	}
}
