// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the details of the specified model version.
func (c *Client) GetModelVersion(ctx context.Context, params *GetModelVersionInput, optFns ...func(*Options)) (*GetModelVersionOutput, error) {
	if params == nil {
		params = &GetModelVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModelVersion", params, optFns, c.addOperationGetModelVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetModelVersionInput struct {

	// The model ID.
	//
	// This member is required.
	ModelId *string

	// The model type.
	//
	// This member is required.
	ModelType types.ModelTypeEnum

	// The model version number.
	//
	// This member is required.
	ModelVersionNumber *string

	noSmithyDocumentSerde
}

type GetModelVersionOutput struct {

	// The model version ARN.
	Arn *string

	// The details of the external events data used for training the model version.
	// This will be populated if the trainingDataSource is EXTERNAL_EVENTS
	ExternalEventsDetail *types.ExternalEventsDetail

	// The details of the ingested events data used for training the model version.
	// This will be populated if the trainingDataSource is INGESTED_EVENTS.
	IngestedEventsDetail *types.IngestedEventsDetail

	// The model ID.
	ModelId *string

	// The model type.
	ModelType types.ModelTypeEnum

	// The model version number.
	ModelVersionNumber *string

	// The model version status. Possible values are:
	//
	// * TRAINING_IN_PROGRESS
	//
	// *
	// TRAINING_COMPLETE
	//
	// * ACTIVATE_REQUESTED
	//
	// * ACTIVATE_IN_PROGRESS
	//
	// * ACTIVE
	//
	// *
	// INACTIVATE_REQUESTED
	//
	// * INACTIVATE_IN_PROGRESS
	//
	// * INACTIVE
	//
	// * ERROR
	Status *string

	// The training data schema.
	TrainingDataSchema *types.TrainingDataSchema

	// The training data source.
	TrainingDataSource types.TrainingDataSourceEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetModelVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetModelVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetModelVersion{}, middleware.After)
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
	if err = addOpGetModelVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetModelVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetModelVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetModelVersion",
	}
}
