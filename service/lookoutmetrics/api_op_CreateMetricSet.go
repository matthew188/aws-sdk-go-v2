// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a dataset.
func (c *Client) CreateMetricSet(ctx context.Context, params *CreateMetricSetInput, optFns ...func(*Options)) (*CreateMetricSetOutput, error) {
	if params == nil {
		params = &CreateMetricSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMetricSet", params, optFns, c.addOperationCreateMetricSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMetricSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMetricSetInput struct {

	// The ARN of the anomaly detector that will use the dataset.
	//
	// This member is required.
	AnomalyDetectorArn *string

	// A list of metrics that the dataset will contain.
	//
	// This member is required.
	MetricList []types.Metric

	// The name of the dataset.
	//
	// This member is required.
	MetricSetName *string

	// Contains information about how the source data should be interpreted.
	//
	// This member is required.
	MetricSource *types.MetricSource

	// A list of filters that specify which data is kept for anomaly detection.
	DimensionFilterList []types.MetricSetDimensionFilter

	// A list of the fields you want to treat as dimensions.
	DimensionList []string

	// A description of the dataset you are creating.
	MetricSetDescription *string

	// The frequency with which the source data will be analyzed for anomalies.
	MetricSetFrequency types.Frequency

	// After an interval ends, the amount of seconds that the detector waits before
	// importing data. Offset is only supported for S3, Redshift, Athena and
	// datasources.
	Offset *int32

	// A list of tags
	// (https://docs.aws.amazon.com/lookoutmetrics/latest/dev/detectors-tags.html) to
	// apply to the dataset.
	Tags map[string]string

	// Contains information about the column used for tracking time in your source
	// data.
	TimestampColumn *types.TimestampColumn

	// The time zone in which your source data was recorded.
	Timezone *string

	noSmithyDocumentSerde
}

type CreateMetricSetOutput struct {

	// The ARN of the dataset.
	MetricSetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMetricSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMetricSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMetricSet{}, middleware.After)
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
	if err = addOpCreateMetricSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMetricSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMetricSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutmetrics",
		OperationName: "CreateMetricSet",
	}
}
