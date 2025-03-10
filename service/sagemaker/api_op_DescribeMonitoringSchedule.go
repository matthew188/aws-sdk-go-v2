// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the schedule for a monitoring job.
func (c *Client) DescribeMonitoringSchedule(ctx context.Context, params *DescribeMonitoringScheduleInput, optFns ...func(*Options)) (*DescribeMonitoringScheduleOutput, error) {
	if params == nil {
		params = &DescribeMonitoringScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMonitoringSchedule", params, optFns, c.addOperationDescribeMonitoringScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMonitoringScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMonitoringScheduleInput struct {

	// Name of a previously created monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleName *string

	noSmithyDocumentSerde
}

type DescribeMonitoringScheduleOutput struct {

	// The time at which the monitoring job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The time at which the monitoring job was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The Amazon Resource Name (ARN) of the monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleArn *string

	// The configuration object that specifies the monitoring schedule and defines the
	// monitoring job.
	//
	// This member is required.
	MonitoringScheduleConfig *types.MonitoringScheduleConfig

	// Name of the monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleName *string

	// The status of an monitoring job.
	//
	// This member is required.
	MonitoringScheduleStatus types.ScheduleStatus

	// The name of the endpoint for the monitoring job.
	EndpointName *string

	// A string, up to one KB in size, that contains the reason a monitoring job
	// failed, if it failed.
	FailureReason *string

	// Describes metadata on the last execution to run, if there was one.
	LastMonitoringExecutionSummary *types.MonitoringExecutionSummary

	// The type of the monitoring job that this schedule runs. This is one of the
	// following values.
	//
	// * DATA_QUALITY - The schedule is for a data quality
	// monitoring job.
	//
	// * MODEL_QUALITY - The schedule is for a model quality
	// monitoring job.
	//
	// * MODEL_BIAS - The schedule is for a bias monitoring job.
	//
	// *
	// MODEL_EXPLAINABILITY - The schedule is for an explainability monitoring job.
	MonitoringType types.MonitoringType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMonitoringScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMonitoringSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMonitoringSchedule{}, middleware.After)
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
	if err = addOpDescribeMonitoringScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMonitoringSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMonitoringSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeMonitoringSchedule",
	}
}
