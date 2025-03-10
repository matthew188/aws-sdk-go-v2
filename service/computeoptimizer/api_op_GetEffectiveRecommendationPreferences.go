// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the recommendation preferences that are in effect for a given resource,
// such as enhanced infrastructure metrics. Considers all applicable preferences
// that you might have set at the resource, account, and organization level. When
// you create a recommendation preference, you can set its status to Active or
// Inactive. Use this action to view the recommendation preferences that are in
// effect, or Active.
func (c *Client) GetEffectiveRecommendationPreferences(ctx context.Context, params *GetEffectiveRecommendationPreferencesInput, optFns ...func(*Options)) (*GetEffectiveRecommendationPreferencesOutput, error) {
	if params == nil {
		params = &GetEffectiveRecommendationPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEffectiveRecommendationPreferences", params, optFns, c.addOperationGetEffectiveRecommendationPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEffectiveRecommendationPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEffectiveRecommendationPreferencesInput struct {

	// The Amazon Resource Name (ARN) of the resource for which to confirm effective
	// recommendation preferences. Only EC2 instance and Auto Scaling group ARNs are
	// currently supported.
	//
	// This member is required.
	ResourceArn *string

	noSmithyDocumentSerde
}

type GetEffectiveRecommendationPreferencesOutput struct {

	// The status of the enhanced infrastructure metrics recommendation preference.
	// Considers all applicable preferences that you might have set at the resource,
	// account, and organization level. A status of Active confirms that the preference
	// is applied in the latest recommendation refresh, and a status of Inactive
	// confirms that it's not yet applied to recommendations. To validate whether the
	// preference is applied to your last generated set of recommendations, review the
	// effectiveRecommendationPreferences value in the response of the
	// GetAutoScalingGroupRecommendations and GetEC2InstanceRecommendations actions.
	// For more information, see Enhanced infrastructure metrics
	// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html)
	// in the Compute Optimizer User Guide.
	EnhancedInfrastructureMetrics types.EnhancedInfrastructureMetrics

	// The provider of the external metrics recommendation preference. Considers all
	// applicable preferences that you might have set at the account and organization
	// level. If the preference is applied in the latest recommendation refresh, an
	// object with a valid source value appears in the response. If the preference
	// isn't applied to the recommendations already, then this object doesn't appear in
	// the response. To validate whether the preference is applied to your last
	// generated set of recommendations, review the effectiveRecommendationPreferences
	// value in the response of the GetEC2InstanceRecommendations actions. For more
	// information, see Enhanced infrastructure metrics
	// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/external-metrics-ingestion.html)
	// in the Compute Optimizer User Guide.
	ExternalMetricsPreference *types.ExternalMetricsPreference

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEffectiveRecommendationPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEffectiveRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEffectiveRecommendationPreferences{}, middleware.After)
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
	if err = addOpGetEffectiveRecommendationPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEffectiveRecommendationPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEffectiveRecommendationPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetEffectiveRecommendationPreferences",
	}
}
