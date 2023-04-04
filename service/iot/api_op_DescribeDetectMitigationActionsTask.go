// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a Device Defender ML Detect mitigation action. Requires
// permission to access the DescribeDetectMitigationActionsTask
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DescribeDetectMitigationActionsTask(ctx context.Context, params *DescribeDetectMitigationActionsTaskInput, optFns ...func(*Options)) (*DescribeDetectMitigationActionsTaskOutput, error) {
	if params == nil {
		params = &DescribeDetectMitigationActionsTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDetectMitigationActionsTask", params, optFns, c.addOperationDescribeDetectMitigationActionsTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDetectMitigationActionsTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDetectMitigationActionsTaskInput struct {

	// The unique identifier of the task.
	//
	// This member is required.
	TaskId *string

	noSmithyDocumentSerde
}

type DescribeDetectMitigationActionsTaskOutput struct {

	// The description of a task.
	TaskSummary *types.DetectMitigationActionsTaskSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDetectMitigationActionsTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDetectMitigationActionsTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDetectMitigationActionsTask{}, middleware.After)
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
	if err = addOpDescribeDetectMitigationActionsTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDetectMitigationActionsTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDetectMitigationActionsTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeDetectMitigationActionsTask",
	}
}
