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

// Describes the scalable resources in the specified scaling plan.
func (c *Client) DescribeScalingPlanResources(ctx context.Context, params *DescribeScalingPlanResourcesInput, optFns ...func(*Options)) (*DescribeScalingPlanResourcesOutput, error) {
	if params == nil {
		params = &DescribeScalingPlanResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScalingPlanResources", params, optFns, c.addOperationDescribeScalingPlanResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScalingPlanResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScalingPlanResourcesInput struct {

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan. Currently, the only valid value is 1.
	//
	// This member is required.
	ScalingPlanVersion *int64

	// The maximum number of scalable resources to return. The value must be between 1
	// and 50. The default value is 50.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeScalingPlanResourcesOutput struct {

	// The token required to get the next set of results. This value is null if there
	// are no more results to return.
	NextToken *string

	// Information about the scalable resources.
	ScalingPlanResources []types.ScalingPlanResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScalingPlanResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeScalingPlanResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeScalingPlanResources{}, middleware.After)
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
	if err = addOpDescribeScalingPlanResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingPlanResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeScalingPlanResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling-plans",
		OperationName: "DescribeScalingPlanResources",
	}
}
