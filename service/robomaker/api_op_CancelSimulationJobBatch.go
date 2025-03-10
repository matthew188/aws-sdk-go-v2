// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels a simulation job batch. When you cancel a simulation job batch, you are
// also cancelling all of the active simulation jobs created as part of the batch.
func (c *Client) CancelSimulationJobBatch(ctx context.Context, params *CancelSimulationJobBatchInput, optFns ...func(*Options)) (*CancelSimulationJobBatchOutput, error) {
	if params == nil {
		params = &CancelSimulationJobBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelSimulationJobBatch", params, optFns, c.addOperationCancelSimulationJobBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelSimulationJobBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelSimulationJobBatchInput struct {

	// The id of the batch to cancel.
	//
	// This member is required.
	Batch *string

	noSmithyDocumentSerde
}

type CancelSimulationJobBatchOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelSimulationJobBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelSimulationJobBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelSimulationJobBatch{}, middleware.After)
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
	if err = addOpCancelSimulationJobBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelSimulationJobBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelSimulationJobBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CancelSimulationJobBatch",
	}
}
