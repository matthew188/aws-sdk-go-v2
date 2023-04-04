// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a state machine's definition, its IAM role Amazon
// Resource Name (ARN), and configuration. If the state machine ARN is a qualified
// state machine ARN, the response returned includes the Map state's label. A
// qualified state machine ARN refers to a Distributed Map state defined within a
// state machine. For example, the qualified state machine ARN
// arn:partition:states:region:account-id:stateMachine:stateMachineName/mapStateLabel
// refers to a Distributed Map state with a label mapStateLabel in the state
// machine named stateMachineName. This operation is eventually consistent. The
// results are best effort and may not reflect very recent updates and changes.
func (c *Client) DescribeStateMachine(ctx context.Context, params *DescribeStateMachineInput, optFns ...func(*Options)) (*DescribeStateMachineOutput, error) {
	if params == nil {
		params = &DescribeStateMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStateMachine", params, optFns, c.addOperationDescribeStateMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStateMachineInput struct {

	// The Amazon Resource Name (ARN) of the state machine to describe.
	//
	// This member is required.
	StateMachineArn *string

	noSmithyDocumentSerde
}

type DescribeStateMachineOutput struct {

	// The date the state machine is created.
	//
	// This member is required.
	CreationDate *time.Time

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language
	// (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	//
	// This member is required.
	Definition *string

	// The name of the state machine. A name must not contain:
	//
	// * white space
	//
	// *
	// brackets < > { } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters " # % \ ^
	// | ~ ` $ & , ; : /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable
	// logging with CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and
	// _.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role used when creating this state
	// machine. (The IAM role maintains security by granting Step Functions access to
	// Amazon Web Services resources.)
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) that identifies the state machine.
	//
	// This member is required.
	StateMachineArn *string

	// The type of the state machine (STANDARD or EXPRESS).
	//
	// This member is required.
	Type types.StateMachineType

	// A user-defined or an auto-generated string that identifies a Map state. This
	// parameter is present only if the stateMachineArn specified in input is a
	// qualified state machine ARN.
	Label *string

	// The LoggingConfiguration data type is used to set CloudWatch Logs options.
	LoggingConfiguration *types.LoggingConfiguration

	// The current status of the state machine.
	Status types.StateMachineStatus

	// Selects whether X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStateMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeStateMachine{}, middleware.After)
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
	if err = addOpDescribeStateMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStateMachine(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStateMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DescribeStateMachine",
	}
}
