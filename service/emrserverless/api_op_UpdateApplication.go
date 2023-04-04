// Code generated by smithy-go-codegen DO NOT EDIT.

package emrserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/emrserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a specified application. An application has to be in a stopped or
// created state in order to be updated.
func (c *Client) UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) {
	if params == nil {
		params = &UpdateApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApplication", params, optFns, c.addOperationUpdateApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationInput struct {

	// The ID of the application to update.
	//
	// This member is required.
	ApplicationId *string

	// The client idempotency token of the application to update. Its value must be
	// unique for each request.
	//
	// This member is required.
	ClientToken *string

	// The CPU architecture of an application.
	Architecture types.Architecture

	// The configuration for an application to automatically start on job submission.
	AutoStartConfiguration *types.AutoStartConfig

	// The configuration for an application to automatically stop after a certain
	// amount of time being idle.
	AutoStopConfiguration *types.AutoStopConfig

	// The image configuration to be used for all worker types. You can either set this
	// parameter or imageConfiguration for each worker type in
	// WorkerTypeSpecificationInput.
	ImageConfiguration *types.ImageConfigurationInput

	// The capacity to initialize when the application is updated.
	InitialCapacity map[string]types.InitialCapacityConfig

	// The maximum capacity to allocate when the application is updated. This is
	// cumulative across all workers at any given point in time during the lifespan of
	// the application. No new resources will be created once any one of the defined
	// limits is hit.
	MaximumCapacity *types.MaximumAllowedResources

	// The network configuration for customer VPC connectivity.
	NetworkConfiguration *types.NetworkConfiguration

	// The key-value pairs that specify worker type to WorkerTypeSpecificationInput.
	// This parameter must contain all valid worker types for a Spark or Hive
	// application. Valid worker types include Driver and Executor for Spark
	// applications and HiveDriver and TezTask for Hive applications. You can either
	// set image details in this parameter for each worker type, or in
	// imageConfiguration for all worker types.
	WorkerTypeSpecifications map[string]types.WorkerTypeSpecificationInput

	noSmithyDocumentSerde
}

type UpdateApplicationOutput struct {

	// Information about the updated application.
	//
	// This member is required.
	Application *types.Application

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApplication{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateApplicationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplication(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateApplication struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateApplication) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateApplicationInput ")
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
func addIdempotencyToken_opUpdateApplicationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateApplication{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "emr-serverless",
		OperationName: "UpdateApplication",
	}
}
