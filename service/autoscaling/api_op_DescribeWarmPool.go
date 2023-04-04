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

// Gets information about a warm pool and its instances. For more information, see
// Warm pools for Amazon EC2 Auto Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-warm-pools.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeWarmPool(ctx context.Context, params *DescribeWarmPoolInput, optFns ...func(*Options)) (*DescribeWarmPoolOutput, error) {
	if params == nil {
		params = &DescribeWarmPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWarmPool", params, optFns, c.addOperationDescribeWarmPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWarmPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWarmPoolInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The maximum number of instances to return with this call. The maximum value is
	// 50.
	MaxRecords *int32

	// The token for the next set of instances to return. (You received this token from
	// a previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeWarmPoolOutput struct {

	// The instances that are currently in the warm pool.
	Instances []types.Instance

	// This string indicates that the response contains more items than can be returned
	// in a single response. To receive additional items, specify this string for the
	// NextToken value when requesting the next set of items. This value is null when
	// there are no more items to return.
	NextToken *string

	// The warm pool configuration details.
	WarmPoolConfiguration *types.WarmPoolConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWarmPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeWarmPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeWarmPool{}, middleware.After)
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
	if err = addOpDescribeWarmPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWarmPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWarmPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeWarmPool",
	}
}
