// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds the specified targets to the specified rule, or updates the targets if they
// are already associated with the rule. Targets are the resources that are invoked
// when a rule is triggered. Each rule can have up to five (5) targets associated
// with it at one time. You can configure the following as targets for Events:
//
// *
// API destination
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-api-destinations.html)
//
// *
// API Gateway
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-api-gateway-target.html)
//
// *
// Batch job queue
//
// * CloudWatch group
//
// * CodeBuild project
//
// * CodePipeline
//
// * EC2
// CreateSnapshot API call
//
// * EC2 Image Builder
//
// * EC2 RebootInstances API call
//
// *
// EC2 StopInstances API call
//
// * EC2 TerminateInstances API call
//
// * ECS task
//
// *
// Event bus in a different account or Region
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-cross-account.html)
//
// *
// Event bus in the same account and Region
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-bus-to-bus.html)
//
// *
// Firehose delivery stream
//
// * Glue workflow
//
// * Incident Manager response plan
// (https://docs.aws.amazon.com/incident-manager/latest/userguide/incident-creation.html#incident-tracking-auto-eventbridge)
//
// *
// Inspector assessment template
//
// * Kinesis stream
//
// * Lambda function
//
// * Redshift
// cluster
//
// * Redshift Serverless workgroup
//
// * SageMaker Pipeline
//
// * SNS topic
//
// *
// SQS queue
//
// * Step Functions state machine
//
// * Systems Manager Automation
//
// *
// Systems Manager OpsItem
//
// * Systems Manager Run Command
//
// Creating rules with
// built-in targets is supported only in the Amazon Web Services Management
// Console. The built-in targets are EC2 CreateSnapshot API call, EC2
// RebootInstances API call, EC2 StopInstances API call, and EC2 TerminateInstances
// API call. For some target types, PutTargets provides target-specific parameters.
// If the target is a Kinesis data stream, you can optionally specify which shard
// the event goes to by using the KinesisParameters argument. To invoke a command
// on multiple EC2 instances with one rule, you can use the RunCommandParameters
// field. To be able to make API calls against the resources that you own, Amazon
// EventBridge needs the appropriate permissions. For Lambda and Amazon SNS
// resources, EventBridge relies on resource-based policies. For EC2 instances,
// Kinesis Data Streams, Step Functions state machines and API Gateway APIs,
// EventBridge relies on IAM roles that you specify in the RoleARN argument in
// PutTargets. For more information, see Authentication and Access Control
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/auth-and-access-control-eventbridge.html)
// in the Amazon EventBridge User Guide. If another Amazon Web Services account is
// in the same region and has granted you permission (using PutPermission), you can
// send events to that account. Set that account's event bus as a target of the
// rules in your account. To send the matched events to the other account, specify
// that account's event bus as the Arn value when you run PutTargets. If your
// account sends events to another account, your account is charged for each sent
// event. Each event sent to another account is charged as a custom event. The
// account receiving the event is not charged. For more information, see Amazon
// EventBridge Pricing (http://aws.amazon.com/eventbridge/pricing/). Input,
// InputPath, and InputTransformer are not available with PutTarget if the target
// is an event bus of a different Amazon Web Services account. If you are setting
// the event bus of another account as the target, and that account granted
// permission to your account through an organization instead of directly by the
// account ID, then you must specify a RoleArn with proper permissions in the
// Target structure. For more information, see Sending and Receiving Events Between
// Amazon Web Services Accounts
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html)
// in the Amazon EventBridge User Guide. For more information about enabling
// cross-account events, see PutPermission
// (https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutPermission.html).
// Input, InputPath, and InputTransformer are mutually exclusive and optional
// parameters of a target. When a rule is triggered due to a matched event:
//
// * If
// none of the following arguments are specified for a target, then the entire
// event is passed to the target in JSON format (unless the target is Amazon EC2
// Run Command or Amazon ECS task, in which case nothing from the event is passed
// to the target).
//
// * If Input is specified in the form of valid JSON, then the
// matched event is overridden with this constant.
//
// * If InputPath is specified in
// the form of JSONPath (for example, $.detail), then only the part of the event
// specified in the path is passed to the target (for example, only the detail part
// of the event is passed).
//
// * If InputTransformer is specified, then one or more
// specified JSONPaths are extracted from the event and used as values in a
// template that you specify as the input to the target.
//
// When you specify
// InputPath or InputTransformer, you must use JSON dot notation, not bracket
// notation. When you add targets to a rule and the associated rule triggers soon
// after, new or updated targets might not be immediately invoked. Allow a short
// period of time for changes to take effect. This action can partially fail if too
// many requests are made at the same time. If that happens, FailedEntryCount is
// non-zero in the response and each entry in FailedEntries provides the ID of the
// failed target and the error code.
func (c *Client) PutTargets(ctx context.Context, params *PutTargetsInput, optFns ...func(*Options)) (*PutTargetsOutput, error) {
	if params == nil {
		params = &PutTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutTargets", params, optFns, c.addOperationPutTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutTargetsInput struct {

	// The name of the rule.
	//
	// This member is required.
	Rule *string

	// The targets to update or add to the rule.
	//
	// This member is required.
	Targets []types.Target

	// The name or ARN of the event bus associated with the rule. If you omit this, the
	// default event bus is used.
	EventBusName *string

	noSmithyDocumentSerde
}

type PutTargetsOutput struct {

	// The failed target entries.
	FailedEntries []types.PutTargetsResultEntry

	// The number of failed entries.
	FailedEntryCount int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutTargets{}, middleware.After)
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
	if err = addOpPutTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutTargets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "PutTargets",
	}
}
