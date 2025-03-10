// Code generated by smithy-go-codegen DO NOT EDIT.

package rbin

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rbin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Unlocks a retention rule. After a retention rule is unlocked, it can be modified
// or deleted only after the unlock delay period expires.
func (c *Client) UnlockRule(ctx context.Context, params *UnlockRuleInput, optFns ...func(*Options)) (*UnlockRuleOutput, error) {
	if params == nil {
		params = &UnlockRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnlockRule", params, optFns, c.addOperationUnlockRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnlockRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnlockRuleInput struct {

	// The unique ID of the retention rule.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type UnlockRuleOutput struct {

	// The retention rule description.
	Description *string

	// The unique ID of the retention rule.
	Identifier *string

	// Information about the retention rule lock configuration.
	LockConfiguration *types.LockConfiguration

	// The date and time at which the unlock delay is set to expire. Only returned for
	// retention rules that have been unlocked and that are still within the unlock
	// delay period.
	LockEndTime *time.Time

	// The lock state for the retention rule.
	//
	// * locked - The retention rule is locked
	// and can't be modified or deleted.
	//
	// * pending_unlock - The retention rule has
	// been unlocked but it is still within the unlock delay period. The retention rule
	// can be modified or deleted only after the unlock delay period has expired.
	//
	// *
	// unlocked - The retention rule is unlocked and it can be modified or deleted by
	// any user with the required permissions.
	//
	// * null - The retention rule has never
	// been locked. Once a retention rule has been locked, it can transition between
	// the locked and unlocked states only; it can never transition back to null.
	LockState types.LockState

	// Information about the resource tags used to identify resources that are retained
	// by the retention rule.
	ResourceTags []types.ResourceTag

	// The resource type retained by the retention rule.
	ResourceType types.ResourceType

	// Information about the retention period for which the retention rule is to retain
	// resources.
	RetentionPeriod *types.RetentionPeriod

	// The state of the retention rule. Only retention rules that are in the available
	// state retain resources.
	Status types.RuleStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUnlockRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUnlockRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUnlockRule{}, middleware.After)
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
	if err = addOpUnlockRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnlockRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUnlockRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rbin",
		OperationName: "UnlockRule",
	}
}
