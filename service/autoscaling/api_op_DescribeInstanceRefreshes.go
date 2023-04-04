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

// Gets information about the instance refreshes for the specified Auto Scaling
// group. This operation is part of the instance refresh feature
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html)
// in Amazon EC2 Auto Scaling, which helps you update instances in your Auto
// Scaling group after you make configuration changes. To help you determine the
// status of an instance refresh, Amazon EC2 Auto Scaling returns information about
// the instance refreshes you previously initiated, including their status, start
// time, end time, the percentage of the instance refresh that is complete, and the
// number of instances remaining to update before the instance refresh is complete.
// If a rollback is initiated while an instance refresh is in progress, Amazon EC2
// Auto Scaling also returns information about the rollback of the instance
// refresh.
func (c *Client) DescribeInstanceRefreshes(ctx context.Context, params *DescribeInstanceRefreshesInput, optFns ...func(*Options)) (*DescribeInstanceRefreshesOutput, error) {
	if params == nil {
		params = &DescribeInstanceRefreshesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceRefreshes", params, optFns, c.addOperationDescribeInstanceRefreshesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceRefreshesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceRefreshesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// One or more instance refresh IDs.
	InstanceRefreshIds []string

	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100.
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstanceRefreshesOutput struct {

	// The instance refreshes for the specified group, sorted by creation timestamp in
	// descending order.
	InstanceRefreshes []types.InstanceRefresh

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceRefreshesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInstanceRefreshes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInstanceRefreshes{}, middleware.After)
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
	if err = addOpDescribeInstanceRefreshesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceRefreshes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstanceRefreshes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeInstanceRefreshes",
	}
}
