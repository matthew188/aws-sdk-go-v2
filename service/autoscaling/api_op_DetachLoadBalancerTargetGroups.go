// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API call has been replaced with a new "traffic sources" API call
// (DetachTrafficSources) that can detach multiple traffic sources types. While we
// continue to support DetachLoadBalancerTargetGroups, and you can use both the
// original DetachLoadBalancerTargetGroups API call and the new
// DetachTrafficSources API call on the same Auto Scaling group, we recommend using
// the new "traffic sources" API call to simplify how you manage traffic sources.
// Detaches one or more target groups from the specified Auto Scaling group. When
// you detach a target group, it enters the Removing state while deregistering the
// instances in the group. When all instances are deregistered, then you can no
// longer describe the target group using the DescribeLoadBalancerTargetGroups API
// call. The instances remain running. You can use this operation to detach target
// groups that were attached by using AttachLoadBalancerTargetGroups, but not for
// target groups that were attached by using AttachTrafficSources.
func (c *Client) DetachLoadBalancerTargetGroups(ctx context.Context, params *DetachLoadBalancerTargetGroupsInput, optFns ...func(*Options)) (*DetachLoadBalancerTargetGroupsOutput, error) {
	if params == nil {
		params = &DetachLoadBalancerTargetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachLoadBalancerTargetGroups", params, optFns, c.addOperationDetachLoadBalancerTargetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachLoadBalancerTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachLoadBalancerTargetGroupsInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The Amazon Resource Names (ARN) of the target groups. You can specify up to 10
	// target groups.
	//
	// This member is required.
	TargetGroupARNs []string

	noSmithyDocumentSerde
}

type DetachLoadBalancerTargetGroupsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachLoadBalancerTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDetachLoadBalancerTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDetachLoadBalancerTargetGroups{}, middleware.After)
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
	if err = addOpDetachLoadBalancerTargetGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachLoadBalancerTargetGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachLoadBalancerTargetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DetachLoadBalancerTargetGroups",
	}
}
