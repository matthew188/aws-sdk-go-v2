// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/autoscalingplans/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more of your scaling plans.
func (c *Client) DescribeScalingPlans(ctx context.Context, params *DescribeScalingPlansInput, optFns ...func(*Options)) (*DescribeScalingPlansOutput, error) {
	if params == nil {
		params = &DescribeScalingPlansInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScalingPlans", params, optFns, c.addOperationDescribeScalingPlansMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScalingPlansOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScalingPlansInput struct {

	// The sources for the applications (up to 10). If you specify scaling plan names,
	// you cannot specify application sources.
	ApplicationSources []types.ApplicationSource

	// The maximum number of scalable resources to return. This value can be between 1
	// and 50. The default value is 50.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// The names of the scaling plans (up to 10). If you specify application sources,
	// you cannot specify scaling plan names.
	ScalingPlanNames []string

	// The version number of the scaling plan. Currently, the only valid value is 1. If
	// you specify a scaling plan version, you must also specify a scaling plan name.
	ScalingPlanVersion *int64

	noSmithyDocumentSerde
}

type DescribeScalingPlansOutput struct {

	// The token required to get the next set of results. This value is null if there
	// are no more results to return.
	NextToken *string

	// Information about the scaling plans.
	ScalingPlans []types.ScalingPlan

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScalingPlansMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeScalingPlans{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeScalingPlans{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingPlans(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeScalingPlans(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling-plans",
		OperationName: "DescribeScalingPlans",
	}
}
