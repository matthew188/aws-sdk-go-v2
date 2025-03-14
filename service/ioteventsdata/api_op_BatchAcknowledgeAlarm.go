// Code generated by smithy-go-codegen DO NOT EDIT.

package ioteventsdata

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ioteventsdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Acknowledges one or more alarms. The alarms change to the ACKNOWLEDGED state
// after you acknowledge them.
func (c *Client) BatchAcknowledgeAlarm(ctx context.Context, params *BatchAcknowledgeAlarmInput, optFns ...func(*Options)) (*BatchAcknowledgeAlarmOutput, error) {
	if params == nil {
		params = &BatchAcknowledgeAlarmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchAcknowledgeAlarm", params, optFns, c.addOperationBatchAcknowledgeAlarmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchAcknowledgeAlarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchAcknowledgeAlarmInput struct {

	// The list of acknowledge action requests. You can specify up to 10 requests per
	// operation.
	//
	// This member is required.
	AcknowledgeActionRequests []types.AcknowledgeAlarmActionRequest

	noSmithyDocumentSerde
}

type BatchAcknowledgeAlarmOutput struct {

	// A list of errors associated with the request, or null if there are no errors.
	// Each error entry contains an entry ID that helps you identify the entry that
	// failed.
	ErrorEntries []types.BatchAlarmActionErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchAcknowledgeAlarmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchAcknowledgeAlarm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchAcknowledgeAlarm{}, middleware.After)
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
	if err = addOpBatchAcknowledgeAlarmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchAcknowledgeAlarm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchAcknowledgeAlarm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ioteventsdata",
		OperationName: "BatchAcknowledgeAlarm",
	}
}
