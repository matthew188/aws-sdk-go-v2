// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Runs an on-demand evaluation for the specified Config rules against the last
// known configuration state of the resources. Use StartConfigRulesEvaluation when
// you want to test that a rule you updated is working as expected.
// StartConfigRulesEvaluation does not re-record the latest configuration state for
// your resources. It re-runs an evaluation against the last known state of your
// resources. You can specify up to 25 Config rules per request. An existing
// StartConfigRulesEvaluation call for the specified rules must complete before you
// can call the API again. If you chose to have Config stream to an Amazon SNS
// topic, you will receive a ConfigRuleEvaluationStarted notification when the
// evaluation starts. You don't need to call the StartConfigRulesEvaluation API to
// run an evaluation for a new rule. When you create a rule, Config evaluates your
// resources against the rule automatically. The StartConfigRulesEvaluation API is
// useful if you want to run on-demand evaluations, such as the following
// example:
//
// * You have a custom rule that evaluates your IAM resources every 24
// hours.
//
// * You update your Lambda function to add additional conditions to your
// rule.
//
// * Instead of waiting for the next periodic evaluation, you call the
// StartConfigRulesEvaluation API.
//
// * Config invokes your Lambda function and
// evaluates your IAM resources.
//
// * Your custom rule will still run periodic
// evaluations every 24 hours.
func (c *Client) StartConfigRulesEvaluation(ctx context.Context, params *StartConfigRulesEvaluationInput, optFns ...func(*Options)) (*StartConfigRulesEvaluationOutput, error) {
	if params == nil {
		params = &StartConfigRulesEvaluationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartConfigRulesEvaluation", params, optFns, c.addOperationStartConfigRulesEvaluationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartConfigRulesEvaluationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartConfigRulesEvaluationInput struct {

	// The list of names of Config rules that you want to run evaluations for.
	ConfigRuleNames []string

	noSmithyDocumentSerde
}

// The output when you start the evaluation for the specified Config rule.
type StartConfigRulesEvaluationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartConfigRulesEvaluationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartConfigRulesEvaluation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartConfigRulesEvaluation{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartConfigRulesEvaluation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartConfigRulesEvaluation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "StartConfigRulesEvaluation",
	}
}
