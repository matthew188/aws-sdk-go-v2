// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a rule for the listener. You can't modify a default listener rule. To
// modify a default listener rule, use UpdateListener.
func (c *Client) UpdateRule(ctx context.Context, params *UpdateRuleInput, optFns ...func(*Options)) (*UpdateRuleOutput, error) {
	if params == nil {
		params = &UpdateRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRule", params, optFns, c.addOperationUpdateRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRuleInput struct {

	// The ID or Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerIdentifier *string

	// The ID or Amazon Resource Name (ARN) of the rule.
	//
	// This member is required.
	RuleIdentifier *string

	// The ID or Amazon Resource Name (ARN) of the service.
	//
	// This member is required.
	ServiceIdentifier *string

	// Information about the action for the specified listener rule.
	Action types.RuleAction

	// The rule match.
	Match types.RuleMatch

	// The rule priority. A listener can't have multiple rules with the same priority.
	Priority *int32

	noSmithyDocumentSerde
}

type UpdateRuleOutput struct {

	// Information about the action for the specified listener rule.
	Action types.RuleAction

	// The Amazon Resource Name (ARN) of the listener.
	Arn *string

	// The ID of the listener.
	Id *string

	// Indicates whether this is the default rule.
	IsDefault *bool

	// The rule match.
	Match types.RuleMatch

	// The name of the listener.
	Name *string

	// The rule priority.
	Priority *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRule{}, middleware.After)
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
	if err = addOpUpdateRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "UpdateRule",
	}
}
