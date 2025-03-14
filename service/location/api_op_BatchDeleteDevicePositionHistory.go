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
)

// Deletes the position history of one or more devices from a tracker resource.
func (c *Client) BatchDeleteDevicePositionHistory(ctx context.Context, params *BatchDeleteDevicePositionHistoryInput, optFns ...func(*Options)) (*BatchDeleteDevicePositionHistoryOutput, error) {
	if params == nil {
		params = &BatchDeleteDevicePositionHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteDevicePositionHistory", params, optFns, c.addOperationBatchDeleteDevicePositionHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteDevicePositionHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteDevicePositionHistoryInput struct {

	// Devices whose position history you want to delete.
	//
	// * For example, for two
	// devices: “DeviceIds” : [DeviceId1,DeviceId2]
	//
	// This member is required.
	DeviceIds []string

	// The name of the tracker resource to delete the device position history from.
	//
	// This member is required.
	TrackerName *string

	noSmithyDocumentSerde
}

type BatchDeleteDevicePositionHistoryOutput struct {

	// Contains error details for each device history that failed to delete.
	//
	// This member is required.
	Errors []types.BatchDeleteDevicePositionHistoryError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteDevicePositionHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDeleteDevicePositionHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDeleteDevicePositionHistory{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchDeleteDevicePositionHistoryMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchDeleteDevicePositionHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteDevicePositionHistory(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchDeleteDevicePositionHistoryMiddleware struct {
}

func (*endpointPrefix_opBatchDeleteDevicePositionHistoryMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchDeleteDevicePositionHistoryMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchDeleteDevicePositionHistoryMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchDeleteDevicePositionHistoryMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opBatchDeleteDevicePositionHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "BatchDeleteDevicePositionHistory",
	}
}
