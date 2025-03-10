// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a channel schedule
func (c *Client) BatchUpdateSchedule(ctx context.Context, params *BatchUpdateScheduleInput, optFns ...func(*Options)) (*BatchUpdateScheduleOutput, error) {
	if params == nil {
		params = &BatchUpdateScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateSchedule", params, optFns, c.addOperationBatchUpdateScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// List of actions to create and list of actions to delete.
type BatchUpdateScheduleInput struct {

	// Id of the channel whose schedule is being updated.
	//
	// This member is required.
	ChannelId *string

	// Schedule actions to create in the schedule.
	Creates *types.BatchScheduleActionCreateRequest

	// Schedule actions to delete from the schedule.
	Deletes *types.BatchScheduleActionDeleteRequest

	noSmithyDocumentSerde
}

// Placeholder documentation for BatchUpdateScheduleResponse
type BatchUpdateScheduleOutput struct {

	// Schedule actions created in the schedule.
	Creates *types.BatchScheduleActionCreateResult

	// Schedule actions deleted from the schedule.
	Deletes *types.BatchScheduleActionDeleteResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateSchedule{}, middleware.After)
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
	if err = addOpBatchUpdateScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "BatchUpdateSchedule",
	}
}
