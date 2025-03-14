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

// This API call has been replaced with a new "traffic sources" API call
// (DescribeTrafficSources) that can describe multiple traffic sources types. While
// we continue to support DescribeLoadBalancers, and you can use both the original
// DescribeLoadBalancers API call and the new DescribeTrafficSources API call on
// the same Auto Scaling group, we recommend using the new "traffic sources" API
// call to simplify how you manage traffic sources. Gets information about the load
// balancers for the specified Auto Scaling group. This operation describes only
// Classic Load Balancers. If you have Application Load Balancers, Network Load
// Balancers, or Gateway Load Balancers, use the DescribeLoadBalancerTargetGroups
// API instead. To determine the attachment status of the load balancer, use the
// State element in the response. When you attach a load balancer to an Auto
// Scaling group, the initial State value is Adding. The state transitions to Added
// after all Auto Scaling instances are registered with the load balancer. If
// Elastic Load Balancing health checks are enabled for the Auto Scaling group, the
// state transitions to InService after at least one Auto Scaling instance passes
// the health check. When the load balancer is in the InService state, Amazon EC2
// Auto Scaling can terminate and replace any instances that are reported as
// unhealthy. If no registered instances pass the health checks, the load balancer
// doesn't enter the InService state. Load balancers also have an InService state
// if you attach them in the CreateAutoScalingGroup API call. If your load balancer
// state is InService, but it is not working properly, check the scaling activities
// by calling DescribeScalingActivities and take any corrective actions necessary.
// For help with failed health checks, see Troubleshooting Amazon EC2 Auto Scaling:
// Health checks
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ts-as-healthchecks.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Use Elastic
// Load Balancing to distribute traffic across the instances in your Auto Scaling
// group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeLoadBalancers(ctx context.Context, params *DescribeLoadBalancersInput, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) {
	if params == nil {
		params = &DescribeLoadBalancersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoadBalancers", params, optFns, c.addOperationDescribeLoadBalancersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoadBalancersInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100.
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeLoadBalancersOutput struct {

	// The load balancers.
	LoadBalancers []types.LoadBalancerState

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoadBalancersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancers{}, middleware.After)
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
	if err = addOpDescribeLoadBalancersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLoadBalancers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeLoadBalancers",
	}
}
