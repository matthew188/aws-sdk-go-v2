// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the position information for a given resource. This action is no longer
// supported. Calls to retrieve the position information should use the
// GetResourcePosition
// (https://docs.aws.amazon.com/iot-wireless/2020-11-22/apireference/API_GetResourcePosition.html)
// API operation instead.
//
// Deprecated: This operation is no longer supported.
func (c *Client) GetPosition(ctx context.Context, params *GetPositionInput, optFns ...func(*Options)) (*GetPositionOutput, error) {
	if params == nil {
		params = &GetPositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPosition", params, optFns, c.addOperationGetPositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPositionInput struct {

	// Resource identifier used to retrieve the position information.
	//
	// This member is required.
	ResourceIdentifier *string

	// Resource type of the resource for which position information is retrieved.
	//
	// This member is required.
	ResourceType types.PositionResourceType

	noSmithyDocumentSerde
}

type GetPositionOutput struct {

	// The accuracy of the estimated position in meters. An empty value indicates that
	// no position data is available. A value of ‘0.0’ value indicates that position
	// data is available. This data corresponds to the position information that you
	// specified instead of the position computed by solver.
	Accuracy *types.Accuracy

	// The position information of the resource.
	Position []float32

	// The vendor of the positioning solver.
	SolverProvider types.PositionSolverProvider

	// The type of solver used to identify the position of the resource.
	SolverType types.PositionSolverType

	// The version of the positioning solver.
	SolverVersion *string

	// The timestamp at which the device's position was determined.
	Timestamp *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPosition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPosition{}, middleware.After)
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
	if err = addOpGetPositionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPosition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPosition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetPosition",
	}
}
