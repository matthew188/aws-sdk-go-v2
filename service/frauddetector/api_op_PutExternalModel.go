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

// Creates or updates an Amazon SageMaker model endpoint. You can also use this
// action to update the configuration of the model endpoint, including the IAM role
// and/or the mapped variables.
func (c *Client) PutExternalModel(ctx context.Context, params *PutExternalModelInput, optFns ...func(*Options)) (*PutExternalModelOutput, error) {
	if params == nil {
		params = &PutExternalModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutExternalModel", params, optFns, c.addOperationPutExternalModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutExternalModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutExternalModelInput struct {

	// The model endpoint input configuration.
	//
	// This member is required.
	InputConfiguration *types.ModelInputConfiguration

	// The IAM role used to invoke the model endpoint.
	//
	// This member is required.
	InvokeModelEndpointRoleArn *string

	// The model endpoints name.
	//
	// This member is required.
	ModelEndpoint *string

	// The model endpoint’s status in Amazon Fraud Detector.
	//
	// This member is required.
	ModelEndpointStatus types.ModelEndpointStatus

	// The source of the model.
	//
	// This member is required.
	ModelSource types.ModelSource

	// The model endpoint output configuration.
	//
	// This member is required.
	OutputConfiguration *types.ModelOutputConfiguration

	// A collection of key and value pairs.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutExternalModelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutExternalModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutExternalModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutExternalModel{}, middleware.After)
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
	if err = addOpPutExternalModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutExternalModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutExternalModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "PutExternalModel",
	}
}
