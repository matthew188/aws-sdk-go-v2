// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the data points of a specific metric for an Amazon Lightsail bucket.
// Metrics report the utilization of a bucket. View and collect metric data
// regularly to monitor the number of objects stored in a bucket (including object
// versions) and the storage space used by those objects.
func (c *Client) GetBucketMetricData(ctx context.Context, params *GetBucketMetricDataInput, optFns ...func(*Options)) (*GetBucketMetricDataOutput, error) {
	if params == nil {
		params = &GetBucketMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketMetricData", params, optFns, c.addOperationGetBucketMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketMetricDataInput struct {

	// The name of the bucket for which to get metric data.
	//
	// This member is required.
	BucketName *string

	// The timestamp indicating the latest data to be returned.
	//
	// This member is required.
	EndTime *time.Time

	// The metric for which you want to return information. Valid bucket metric names
	// are listed below, along with the most useful statistics to include in your
	// request, and the published unit value. These bucket metrics are reported once
	// per day.
	//
	// * BucketSizeBytes - The amount of data in bytes stored in a bucket.
	// This value is calculated by summing the size of all objects in the bucket
	// (including object versions), including the size of all parts for all incomplete
	// multipart uploads to the bucket. Statistics: The most useful statistic is
	// Maximum. Unit: The published unit is Bytes.
	//
	// * NumberOfObjects - The total
	// number of objects stored in a bucket. This value is calculated by counting all
	// objects in the bucket (including object versions) and the total number of parts
	// for all incomplete multipart uploads to the bucket. Statistics: The most useful
	// statistic is Average. Unit: The published unit is Count.
	//
	// This member is required.
	MetricName types.BucketMetricName

	// The granularity, in seconds, of the returned data points. Bucket storage metrics
	// are reported once per day. Therefore, you should specify a period of 86400
	// seconds, which is the number of seconds in a day.
	//
	// This member is required.
	Period int32

	// The timestamp indicating the earliest data to be returned.
	//
	// This member is required.
	StartTime *time.Time

	// The statistic for the metric. The following statistics are available:
	//
	// * Minimum
	// - The lowest value observed during the specified period. Use this value to
	// determine low volumes of activity for your application.
	//
	// * Maximum - The highest
	// value observed during the specified period. Use this value to determine high
	// volumes of activity for your application.
	//
	// * Sum - The sum of all values
	// submitted for the matching metric. You can use this statistic to determine the
	// total volume of a metric.
	//
	// * Average - The value of Sum / SampleCount during the
	// specified period. By comparing this statistic with the Minimum and Maximum
	// values, you can determine the full scope of a metric and how close the average
	// use is to the Minimum and Maximum values. This comparison helps you to know when
	// to increase or decrease your resources.
	//
	// * SampleCount - The count, or number,
	// of data points used for the statistical calculation.
	//
	// This member is required.
	Statistics []types.MetricStatistic

	// The unit for the metric data request. Valid units depend on the metric data
	// being requested. For the valid units with each available metric, see the
	// metricName parameter.
	//
	// This member is required.
	Unit types.MetricUnit

	noSmithyDocumentSerde
}

type GetBucketMetricDataOutput struct {

	// An array of objects that describe the metric data returned.
	MetricData []types.MetricDatapoint

	// The name of the metric returned.
	MetricName types.BucketMetricName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetBucketMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBucketMetricData{}, middleware.After)
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
	if err = addOpGetBucketMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketMetricData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBucketMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetBucketMetricData",
	}
}
