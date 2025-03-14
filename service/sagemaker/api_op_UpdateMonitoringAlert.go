// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the parameters of a model monitor alert.
func (c *Client) UpdateMonitoringAlert(ctx context.Context, params *UpdateMonitoringAlertInput, optFns ...func(*Options)) (*UpdateMonitoringAlertOutput, error) {
	if params == nil {
		params = &UpdateMonitoringAlertInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMonitoringAlert", params, optFns, c.addOperationUpdateMonitoringAlertMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMonitoringAlertOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMonitoringAlertInput struct {

	// Within EvaluationPeriod, how many execution failures will raise an alert.
	//
	// This member is required.
	DatapointsToAlert *int32

	// The number of most recent monitoring executions to consider when evaluating
	// alert status.
	//
	// This member is required.
	EvaluationPeriod *int32

	// The name of a monitoring alert.
	//
	// This member is required.
	MonitoringAlertName *string

	// The name of a monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleName *string

	noSmithyDocumentSerde
}

type UpdateMonitoringAlertOutput struct {

	// The Amazon Resource Name (ARN) of the monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleArn *string

	// The name of a monitoring alert.
	MonitoringAlertName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMonitoringAlertMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMonitoringAlert{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMonitoringAlert{}, middleware.After)
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
	if err = addOpUpdateMonitoringAlertValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMonitoringAlert(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMonitoringAlert(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateMonitoringAlert",
	}
}
