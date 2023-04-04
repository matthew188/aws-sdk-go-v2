// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Permanently deletes a RateBasedRule. You can't delete a rule if it's
// still used in any WebACL objects or if it still includes any predicates, such as
// ByteMatchSet objects. If you just want to remove a rule from a WebACL, use
// UpdateWebACL. To permanently delete a RateBasedRule from AWS WAF, perform the
// following steps:
//
// * Update the RateBasedRule to remove predicates, if any. For
// more information, see UpdateRateBasedRule.
//
// * Use GetChangeToken to get the
// change token that you provide in the ChangeToken parameter of a
// DeleteRateBasedRule request.
//
// * Submit a DeleteRateBasedRule request.
func (c *Client) DeleteRateBasedRule(ctx context.Context, params *DeleteRateBasedRuleInput, optFns ...func(*Options)) (*DeleteRateBasedRuleOutput, error) {
	if params == nil {
		params = &DeleteRateBasedRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRateBasedRule", params, optFns, c.addOperationDeleteRateBasedRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRateBasedRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRateBasedRuleInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// The RuleId of the RateBasedRule that you want to delete. RuleId is returned by
	// CreateRateBasedRule and by ListRateBasedRules.
	//
	// This member is required.
	RuleId *string

	noSmithyDocumentSerde
}

type DeleteRateBasedRuleOutput struct {

	// The ChangeToken that you used to submit the DeleteRateBasedRule request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRateBasedRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRateBasedRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRateBasedRule{}, middleware.After)
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
	if err = addOpDeleteRateBasedRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRateBasedRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRateBasedRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "DeleteRateBasedRule",
	}
}
