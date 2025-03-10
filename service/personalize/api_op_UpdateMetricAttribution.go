// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a metric attribution.
func (c *Client) UpdateMetricAttribution(ctx context.Context, params *UpdateMetricAttributionInput, optFns ...func(*Options)) (*UpdateMetricAttributionOutput, error) {
	if params == nil {
		params = &UpdateMetricAttributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMetricAttribution", params, optFns, c.addOperationUpdateMetricAttributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMetricAttributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMetricAttributionInput struct {

	// Add new metric attributes to the metric attribution.
	AddMetrics []types.MetricAttribute

	// The Amazon Resource Name (ARN) for the metric attribution to update.
	MetricAttributionArn *string

	// An output config for the metric attribution.
	MetricsOutputConfig *types.MetricAttributionOutput

	// Remove metric attributes from the metric attribution.
	RemoveMetrics []string

	noSmithyDocumentSerde
}

type UpdateMetricAttributionOutput struct {

	// The Amazon Resource Name (ARN) for the metric attribution that you updated.
	MetricAttributionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMetricAttributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMetricAttribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMetricAttribution{}, middleware.After)
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
	if err = addOpUpdateMetricAttributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMetricAttribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMetricAttribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "UpdateMetricAttribution",
	}
}
