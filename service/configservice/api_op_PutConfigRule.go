// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates an Config rule to evaluate if your Amazon Web Services resources
// comply with your desired configurations. For information on how many Config
// rules you can have per account, see  Service Limits
// (https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in
// the Config Developer Guide. There are two types of rules: Config Managed Rules
// and Config Custom Rules. You can use PutConfigRule to create both Config Managed
// Rules and Config Custom Rules. Config Managed Rules are predefined, customizable
// rules created by Config. For a list of managed rules, see List of Config Managed
// Rules
// (https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html).
// If you are adding an Config managed rule, you must specify the rule's identifier
// for the SourceIdentifier key. Config Custom Rules are rules that you create from
// scratch. There are two ways to create Config custom rules: with Lambda functions
// ( Lambda Developer Guide
// (https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function))
// and with Guard (Guard GitHub Repository
// (https://github.com/aws-cloudformation/cloudformation-guard)), a policy-as-code
// language. Config custom rules created with Lambda are called Config Custom
// Lambda Rules and Config custom rules created with Guard are called Config Custom
// Policy Rules. If you are adding a new Config Custom Lambda rule, you first need
// to create an Lambda function that the rule invokes to evaluate your resources.
// When you use PutConfigRule to add a Custom Lambda rule to Config, you must
// specify the Amazon Resource Name (ARN) that Lambda assigns to the function. You
// specify the ARN in the SourceIdentifier key. This key is part of the Source
// object, which is part of the ConfigRule object. For any new Config rule that you
// add, specify the ConfigRuleName in the ConfigRule object. Do not specify the
// ConfigRuleArn or the ConfigRuleId. These values are generated by Config for new
// rules. If you are updating a rule that you added previously, you can specify the
// rule by ConfigRuleName, ConfigRuleId, or ConfigRuleArn in the ConfigRule data
// type that you use in this request. For more information about developing and
// using Config rules, see Evaluating Resources with Config Rules
// (https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config.html)
// in the Config Developer Guide. PutConfigRule is an idempotent API. Subsequent
// requests won’t create a duplicate resource if one was already created. If a
// following request has different tags values, Config will ignore these
// differences and treat it as an idempotent request of the previous. In this case,
// tags will not be updated, even if they are different.
func (c *Client) PutConfigRule(ctx context.Context, params *PutConfigRuleInput, optFns ...func(*Options)) (*PutConfigRuleOutput, error) {
	if params == nil {
		params = &PutConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigRule", params, optFns, c.addOperationPutConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutConfigRuleInput struct {

	// The rule that you want to add to your account.
	//
	// This member is required.
	ConfigRule *types.ConfigRule

	// An array of tag object.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type PutConfigRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutConfigRule{}, middleware.After)
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
	if err = addOpPutConfigRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConfigRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutConfigRule",
	}
}
