// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/applicationinsights/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the recommended monitoring configuration of the component.
func (c *Client) DescribeComponentConfigurationRecommendation(ctx context.Context, params *DescribeComponentConfigurationRecommendationInput, optFns ...func(*Options)) (*DescribeComponentConfigurationRecommendationOutput, error) {
	if params == nil {
		params = &DescribeComponentConfigurationRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeComponentConfigurationRecommendation", params, optFns, c.addOperationDescribeComponentConfigurationRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeComponentConfigurationRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeComponentConfigurationRecommendationInput struct {

	// The name of the component.
	//
	// This member is required.
	ComponentName *string

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string

	// The tier of the application component.
	//
	// This member is required.
	Tier types.Tier

	noSmithyDocumentSerde
}

type DescribeComponentConfigurationRecommendationOutput struct {

	// The recommended configuration settings of the component. The value is the
	// escaped JSON of the configuration.
	ComponentConfiguration *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeComponentConfigurationRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeComponentConfigurationRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeComponentConfigurationRecommendation{}, middleware.After)
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
	if err = addOpDescribeComponentConfigurationRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComponentConfigurationRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeComponentConfigurationRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "DescribeComponentConfigurationRecommendation",
	}
}
