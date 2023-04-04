// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates a scaling policy for an Auto Scaling group. Scaling policies
// are used to scale an Auto Scaling group based on configurable metrics. If no
// policies are defined, the dynamic scaling and predictive scaling features are
// not used. For more information about using dynamic scaling, see Target tracking
// scaling policies
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-target-tracking.html)
// and Step and simple scaling policies
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information about using
// predictive scaling, see Predictive scaling for Amazon EC2 Auto Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html)
// in the Amazon EC2 Auto Scaling User Guide. You can view the scaling policies for
// an Auto Scaling group using the DescribePolicies API call. If you are no longer
// using a scaling policy, you can delete it by calling the DeletePolicy API.
func (c *Client) PutScalingPolicy(ctx context.Context, params *PutScalingPolicyInput, optFns ...func(*Options)) (*PutScalingPolicyOutput, error) {
	if params == nil {
		params = &PutScalingPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutScalingPolicy", params, optFns, c.addOperationPutScalingPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutScalingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutScalingPolicyInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The name of the policy.
	//
	// This member is required.
	PolicyName *string

	// Specifies how the scaling adjustment is interpreted (for example, an absolute
	// number or a percentage). The valid values are ChangeInCapacity, ExactCapacity,
	// and PercentChangeInCapacity. Required if the policy type is StepScaling or
	// SimpleScaling. For more information, see Scaling adjustment types
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment)
	// in the Amazon EC2 Auto Scaling User Guide.
	AdjustmentType *string

	// A cooldown period, in seconds, that applies to a specific simple scaling policy.
	// When a cooldown period is specified here, it overrides the default cooldown.
	// Valid only if the policy type is SimpleScaling. For more information, see
	// Scaling cooldowns for Amazon EC2 Auto Scaling
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/Cooldown.html) in the
	// Amazon EC2 Auto Scaling User Guide. Default: None
	Cooldown *int32

	// Indicates whether the scaling policy is enabled or disabled. The default is
	// enabled. For more information, see Disabling a scaling policy for an Auto
	// Scaling group
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-enable-disable-scaling-policy.html)
	// in the Amazon EC2 Auto Scaling User Guide.
	Enabled *bool

	// Not needed if the default instance warmup is defined for the group. The
	// estimated time, in seconds, until a newly launched instance can contribute to
	// the CloudWatch metrics. This warm-up period applies to instances launched due to
	// a specific target tracking or step scaling policy. When a warm-up period is
	// specified here, it overrides the default instance warmup. Valid only if the
	// policy type is TargetTrackingScaling or StepScaling. The default is to use the
	// value for the default instance warmup defined for the group. If default instance
	// warmup is null, then EstimatedInstanceWarmup falls back to the value of default
	// cooldown.
	EstimatedInstanceWarmup *int32

	// The aggregation type for the CloudWatch metrics. The valid values are Minimum,
	// Maximum, and Average. If the aggregation type is null, the value is treated as
	// Average. Valid only if the policy type is StepScaling.
	MetricAggregationType *string

	// The minimum value to scale by when the adjustment type is
	// PercentChangeInCapacity. For example, suppose that you create a step scaling
	// policy to scale out an Auto Scaling group by 25 percent and you specify a
	// MinAdjustmentMagnitude of 2. If the group has 4 instances and the scaling policy
	// is performed, 25 percent of 4 is 1. However, because you specified a
	// MinAdjustmentMagnitude of 2, Amazon EC2 Auto Scaling scales out the group by 2
	// instances. Valid only if the policy type is StepScaling or SimpleScaling. For
	// more information, see Scaling adjustment types
	// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scaling-simple-step.html#as-scaling-adjustment)
	// in the Amazon EC2 Auto Scaling User Guide. Some Auto Scaling groups use instance
	// weights. In this case, set the MinAdjustmentMagnitude to a value that is at
	// least as large as your largest instance weight.
	MinAdjustmentMagnitude *int32

	// Available for backward compatibility. Use MinAdjustmentMagnitude instead.
	//
	// Deprecated: This member has been deprecated.
	MinAdjustmentStep *int32

	// One of the following policy types:
	//
	// * TargetTrackingScaling
	//
	// * StepScaling
	//
	// *
	// SimpleScaling (default)
	//
	// * PredictiveScaling
	PolicyType *string

	// A predictive scaling policy. Provides support for predefined and custom metrics.
	// Predefined metrics include CPU utilization, network in/out, and the Application
	// Load Balancer request count. For more information, see
	// PredictiveScalingConfiguration
	// (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PredictiveScalingConfiguration.html)
	// in the Amazon EC2 Auto Scaling API Reference. Required if the policy type is
	// PredictiveScaling.
	PredictiveScalingConfiguration *types.PredictiveScalingConfiguration

	// The amount by which to scale, based on the specified adjustment type. A positive
	// value adds to the current capacity while a negative number removes from the
	// current capacity. For exact capacity, you must specify a positive value.
	// Required if the policy type is SimpleScaling. (Not used with any other policy
	// type.)
	ScalingAdjustment *int32

	// A set of adjustments that enable you to scale based on the size of the alarm
	// breach. Required if the policy type is StepScaling. (Not used with any other
	// policy type.)
	StepAdjustments []types.StepAdjustment

	// A target tracking scaling policy. Provides support for predefined or custom
	// metrics. The following predefined metrics are available:
	//
	// *
	// ASGAverageCPUUtilization
	//
	// * ASGAverageNetworkIn
	//
	// * ASGAverageNetworkOut
	//
	// *
	// ALBRequestCountPerTarget
	//
	// If you specify ALBRequestCountPerTarget for the
	// metric, you must specify the ResourceLabel property with the
	// PredefinedMetricSpecification. For more information, see
	// TargetTrackingConfiguration
	// (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_TargetTrackingConfiguration.html)
	// in the Amazon EC2 Auto Scaling API Reference. Required if the policy type is
	// TargetTrackingScaling.
	TargetTrackingConfiguration *types.TargetTrackingConfiguration

	noSmithyDocumentSerde
}

// Contains the output of PutScalingPolicy.
type PutScalingPolicyOutput struct {

	// The CloudWatch alarms created for the target tracking scaling policy.
	Alarms []types.Alarm

	// The Amazon Resource Name (ARN) of the policy.
	PolicyARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutScalingPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutScalingPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutScalingPolicy{}, middleware.After)
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
	if err = addOpPutScalingPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutScalingPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutScalingPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutScalingPolicy",
	}
}
