// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this API to define a Custom Metric published by your devices to Device
// Defender. Requires permission to access the CreateCustomMetric
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreateCustomMetric(ctx context.Context, params *CreateCustomMetricInput, optFns ...func(*Options)) (*CreateCustomMetricOutput, error) {
	if params == nil {
		params = &CreateCustomMetricInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomMetric", params, optFns, c.addOperationCreateCustomMetricMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomMetricOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCustomMetricInput struct {

	// Each custom metric must have a unique client request token. If you try to create
	// a new custom metric that already exists with a different token, an exception
	// occurs. If you omit this value, Amazon Web Services SDKs will automatically
	// generate a unique client request.
	//
	// This member is required.
	ClientRequestToken *string

	// The name of the custom metric. This will be used in the metric report submitted
	// from the device/thing. The name can't begin with aws:. You can't change the name
	// after you define it.
	//
	// This member is required.
	MetricName *string

	// The type of the custom metric. The type number only takes a single metric value
	// as an input, but when you submit the metrics value in the DeviceMetrics report,
	// you must pass it as an array with a single value.
	//
	// This member is required.
	MetricType types.CustomMetricType

	// The friendly name in the console for the custom metric. This name doesn't have
	// to be unique. Don't use this name as the metric identifier in the device metric
	// report. You can update the friendly name after you define it.
	DisplayName *string

	// Metadata that can be used to manage the custom metric.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateCustomMetricOutput struct {

	// The Amazon Resource Number (ARN) of the custom metric. For example,
	// arn:aws-partition:iot:region:accountId:custommetric/metricName
	MetricArn *string

	// The name of the custom metric to be used in the metric report.
	MetricName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomMetricMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCustomMetric{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCustomMetric{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateCustomMetricMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCustomMetricValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomMetric(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCustomMetric struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCustomMetric) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCustomMetric) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCustomMetricInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCustomMetricInput ")
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
func addIdempotencyToken_opCreateCustomMetricMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCustomMetric{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCustomMetric(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateCustomMetric",
	}
}
