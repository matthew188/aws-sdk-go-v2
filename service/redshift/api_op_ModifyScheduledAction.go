// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies a scheduled action.
func (c *Client) ModifyScheduledAction(ctx context.Context, params *ModifyScheduledActionInput, optFns ...func(*Options)) (*ModifyScheduledActionOutput, error) {
	if params == nil {
		params = &ModifyScheduledActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyScheduledAction", params, optFns, c.addOperationModifyScheduledActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyScheduledActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyScheduledActionInput struct {

	// The name of the scheduled action to modify.
	//
	// This member is required.
	ScheduledActionName *string

	// A modified enable flag of the scheduled action. If true, the scheduled action is
	// active. If false, the scheduled action is disabled.
	Enable *bool

	// A modified end time of the scheduled action. For more information about this
	// parameter, see ScheduledAction.
	EndTime *time.Time

	// A different IAM role to assume to run the target action. For more information
	// about this parameter, see ScheduledAction.
	IamRole *string

	// A modified schedule in either at( ) or cron( ) format. For more information
	// about this parameter, see ScheduledAction.
	Schedule *string

	// A modified description of the scheduled action.
	ScheduledActionDescription *string

	// A modified start time of the scheduled action. For more information about this
	// parameter, see ScheduledAction.
	StartTime *time.Time

	// A modified JSON format of the scheduled action. For more information about this
	// parameter, see ScheduledAction.
	TargetAction *types.ScheduledActionType

	noSmithyDocumentSerde
}

// Describes a scheduled action. You can use a scheduled action to trigger some
// Amazon Redshift API operations on a schedule. For information about which API
// operations can be scheduled, see ScheduledActionType.
type ModifyScheduledActionOutput struct {

	// The end time in UTC when the schedule is no longer active. After this time, the
	// scheduled action does not trigger.
	EndTime *time.Time

	// The IAM role to assume to run the scheduled action. This IAM role must have
	// permission to run the Amazon Redshift API operation in the scheduled action.
	// This IAM role must allow the Amazon Redshift scheduler (Principal
	// scheduler.redshift.amazonaws.com) to assume permissions on your behalf. For more
	// information about the IAM role to use with the Amazon Redshift scheduler, see
	// Using Identity-Based Policies for Amazon Redshift
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html)
	// in the Amazon Redshift Cluster Management Guide.
	IamRole *string

	// List of times when the scheduled action will run.
	NextInvocations []time.Time

	// The schedule for a one-time (at format) or recurring (cron format) scheduled
	// action. Schedule invocations must be separated by at least one hour. Format of
	// at expressions is "at(yyyy-mm-ddThh:mm:ss)". For example,
	// "at(2016-03-04T17:27:00)". Format of cron expressions is "cron(Minutes Hours
	// Day-of-month Month Day-of-week Year)". For example, "cron(0 10 ? * MON *)". For
	// more information, see Cron Expressions
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions)
	// in the Amazon CloudWatch Events User Guide.
	Schedule *string

	// The description of the scheduled action.
	ScheduledActionDescription *string

	// The name of the scheduled action.
	ScheduledActionName *string

	// The start time in UTC when the schedule is active. Before this time, the
	// scheduled action does not trigger.
	StartTime *time.Time

	// The state of the scheduled action. For example, DISABLED.
	State types.ScheduledActionState

	// A JSON format string of the Amazon Redshift API operation with input parameters.
	// "{\"ResizeCluster\":{\"NodeType\":\"ds2.8xlarge\",\"ClusterIdentifier\":\"my-test-cluster\",\"NumberOfNodes\":3}}".
	TargetAction *types.ScheduledActionType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyScheduledActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyScheduledAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyScheduledAction{}, middleware.After)
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
	if err = addOpModifyScheduledActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyScheduledAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyScheduledAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyScheduledAction",
	}
}
