// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation assigns feature variation to user sessions. For each user
// session, you pass in an entityID that represents the user. Evidently then checks
// the evaluation rules and assigns the variation. The first rules that are
// evaluated are the override rules. If the user's entityID matches an override
// rule, the user is served the variation specified by that rule. Next, if there is
// a launch of the feature, the user might be assigned to a variation in the
// launch. The chance of this depends on the percentage of users that are allocated
// to that launch. If the user is enrolled in the launch, the variation they are
// served depends on the allocation of the various feature variations used for the
// launch. If the user is not assigned to a launch, and there is an ongoing
// experiment for this feature, the user might be assigned to a variation in the
// experiment. The chance of this depends on the percentage of users that are
// allocated to that experiment. If the user is enrolled in the experiment, the
// variation they are served depends on the allocation of the various feature
// variations used for the experiment. If the user is not assigned to a launch or
// experiment, they are served the default variation.
func (c *Client) BatchEvaluateFeature(ctx context.Context, params *BatchEvaluateFeatureInput, optFns ...func(*Options)) (*BatchEvaluateFeatureOutput, error) {
	if params == nil {
		params = &BatchEvaluateFeatureInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchEvaluateFeature", params, optFns, c.addOperationBatchEvaluateFeatureMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchEvaluateFeatureOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchEvaluateFeatureInput struct {

	// The name or ARN of the project that contains the feature being evaluated.
	//
	// This member is required.
	Project *string

	// An array of structures, where each structure assigns a feature variation to one
	// user session.
	//
	// This member is required.
	Requests []types.EvaluationRequest

	noSmithyDocumentSerde
}

type BatchEvaluateFeatureOutput struct {

	// An array of structures, where each structure displays the results of one feature
	// evaluation assignment to one user session.
	Results []types.EvaluationResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchEvaluateFeatureMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchEvaluateFeature{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchEvaluateFeature{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchEvaluateFeatureMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchEvaluateFeatureValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchEvaluateFeature(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchEvaluateFeatureMiddleware struct {
}

func (*endpointPrefix_opBatchEvaluateFeatureMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchEvaluateFeatureMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "dataplane." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchEvaluateFeatureMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchEvaluateFeatureMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opBatchEvaluateFeature(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "BatchEvaluateFeature",
	}
}
