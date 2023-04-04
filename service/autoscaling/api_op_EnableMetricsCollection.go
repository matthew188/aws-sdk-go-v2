// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables group metrics collection for the specified Auto Scaling group. You can
// use these metrics to track changes in an Auto Scaling group and to set alarms on
// threshold values. You can view group metrics using the Amazon EC2 Auto Scaling
// console or the CloudWatch console. For more information, see Monitor CloudWatch
// metrics for your Auto Scaling groups and instances
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-cloudwatch-monitoring.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) EnableMetricsCollection(ctx context.Context, params *EnableMetricsCollectionInput, optFns ...func(*Options)) (*EnableMetricsCollectionOutput, error) {
	if params == nil {
		params = &EnableMetricsCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableMetricsCollection", params, optFns, c.addOperationEnableMetricsCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableMetricsCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableMetricsCollectionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The frequency at which Amazon EC2 Auto Scaling sends aggregated data to
	// CloudWatch. The only valid value is 1Minute.
	//
	// This member is required.
	Granularity *string

	// Identifies the metrics to enable. You can specify one or more of the following
	// metrics:
	//
	// * GroupMinSize
	//
	// * GroupMaxSize
	//
	// * GroupDesiredCapacity
	//
	// *
	// GroupInServiceInstances
	//
	// * GroupPendingInstances
	//
	// * GroupStandbyInstances
	//
	// *
	// GroupTerminatingInstances
	//
	// * GroupTotalInstances
	//
	// * GroupInServiceCapacity
	//
	// *
	// GroupPendingCapacity
	//
	// * GroupStandbyCapacity
	//
	// * GroupTerminatingCapacity
	//
	// *
	// GroupTotalCapacity
	//
	// * WarmPoolDesiredCapacity
	//
	// * WarmPoolWarmedCapacity
	//
	// *
	// WarmPoolPendingCapacity
	//
	// * WarmPoolTerminatingCapacity
	//
	// *
	// WarmPoolTotalCapacity
	//
	// * GroupAndWarmPoolDesiredCapacity
	//
	// *
	// GroupAndWarmPoolTotalCapacity
	//
	// If you specify Granularity and don't specify any
	// metrics, all metrics are enabled. For more information, see Auto Scaling group
	// metrics
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-cloudwatch-monitoring.html#as-group-metrics)
	// in the Amazon EC2 Auto Scaling User Guide.
	Metrics []string

	noSmithyDocumentSerde
}

type EnableMetricsCollectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableMetricsCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnableMetricsCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableMetricsCollection{}, middleware.After)
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
	if err = addOpEnableMetricsCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableMetricsCollection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableMetricsCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "EnableMetricsCollection",
	}
}
