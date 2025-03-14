// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the specified properties for a given route calculator resource.
func (c *Client) UpdateRouteCalculator(ctx context.Context, params *UpdateRouteCalculatorInput, optFns ...func(*Options)) (*UpdateRouteCalculatorOutput, error) {
	if params == nil {
		params = &UpdateRouteCalculatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRouteCalculator", params, optFns, c.addOperationUpdateRouteCalculatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRouteCalculatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRouteCalculatorInput struct {

	// The name of the route calculator resource to update.
	//
	// This member is required.
	CalculatorName *string

	// Updates the description for the route calculator resource.
	Description *string

	// No longer used. If included, the only allowed value is RequestBasedUsage.
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	noSmithyDocumentSerde
}

type UpdateRouteCalculatorOutput struct {

	// The Amazon Resource Name (ARN) of the updated route calculator resource. Used to
	// specify a resource across AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:route- calculator/ExampleCalculator
	//
	// This member is required.
	CalculatorArn *string

	// The name of the updated route calculator resource.
	//
	// This member is required.
	CalculatorName *string

	// The timestamp for when the route calculator was last updated in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRouteCalculatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRouteCalculator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRouteCalculator{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateRouteCalculatorMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateRouteCalculatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRouteCalculator(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateRouteCalculatorMiddleware struct {
}

func (*endpointPrefix_opUpdateRouteCalculatorMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateRouteCalculatorMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "routes." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUpdateRouteCalculatorMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateRouteCalculatorMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRouteCalculator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "UpdateRouteCalculator",
	}
}
