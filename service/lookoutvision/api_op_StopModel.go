// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutvision

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lookoutvision/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops the hosting of a running model. The operation might take a while to
// complete. To check the current status, call DescribeModel. After the model
// hosting stops, the Status of the model is TRAINED. This operation requires
// permissions to perform the lookoutvision:StopModel operation.
func (c *Client) StopModel(ctx context.Context, params *StopModelInput, optFns ...func(*Options)) (*StopModelOutput, error) {
	if params == nil {
		params = &StopModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopModel", params, optFns, c.addOperationStopModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopModelInput struct {

	// The version of the model that you want to stop.
	//
	// This member is required.
	ModelVersion *string

	// The name of the project that contains the model that you want to stop.
	//
	// This member is required.
	ProjectName *string

	// ClientToken is an idempotency token that ensures a call to StopModel completes
	// only once. You choose the value to pass. For example, An issue might prevent you
	// from getting a response from StopModel. In this case, safely retry your call to
	// StopModel by using the same ClientToken parameter value. If you don't supply a
	// value for ClientToken, the AWS SDK you are using inserts a value for you. This
	// prevents retries after a network error from making multiple stop requests.
	// You'll need to provide your own value for other use cases. An error occurs if
	// the other input parameters are not the same as in the first request. Using a
	// different value for ClientToken is considered a new call to StopModel. An
	// idempotency token is active for 8 hours.
	ClientToken *string

	noSmithyDocumentSerde
}

type StopModelOutput struct {

	// The status of the model.
	Status types.ModelHostingStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopModel{}, middleware.After)
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
	if err = addIdempotencyToken_opStopModelMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStopModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopModel(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStopModel struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStopModel) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStopModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StopModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StopModelInput ")
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
func addIdempotencyToken_opStopModelMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStopModel{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStopModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutvision",
		OperationName: "StopModel",
	}
}
