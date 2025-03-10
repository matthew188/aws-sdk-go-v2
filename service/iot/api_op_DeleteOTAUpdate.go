// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete an OTA update. Requires permission to access the DeleteOTAUpdate
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DeleteOTAUpdate(ctx context.Context, params *DeleteOTAUpdateInput, optFns ...func(*Options)) (*DeleteOTAUpdateOutput, error) {
	if params == nil {
		params = &DeleteOTAUpdateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOTAUpdate", params, optFns, c.addOperationDeleteOTAUpdateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOTAUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOTAUpdateInput struct {

	// The ID of the OTA update to delete.
	//
	// This member is required.
	OtaUpdateId *string

	// When true, the stream created by the OTAUpdate process is deleted when the OTA
	// update is deleted. Ignored if the stream specified in the OTAUpdate is supplied
	// by the user.
	DeleteStream bool

	// When true, deletes the IoT job created by the OTAUpdate process even if it is
	// "IN_PROGRESS". Otherwise, if the job is not in a terminal state ("COMPLETED" or
	// "CANCELED") an exception will occur. The default is false.
	ForceDeleteAWSJob bool

	noSmithyDocumentSerde
}

type DeleteOTAUpdateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOTAUpdateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteOTAUpdate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteOTAUpdate{}, middleware.After)
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
	if err = addOpDeleteOTAUpdateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOTAUpdate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOTAUpdate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteOTAUpdate",
	}
}
