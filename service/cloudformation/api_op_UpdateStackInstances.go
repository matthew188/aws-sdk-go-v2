// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the parameter values for stack instances for the specified accounts,
// within the specified Amazon Web Services Regions. A stack instance refers to a
// stack in a specific account and Region. You can only update stack instances in
// Amazon Web Services Regions and accounts where they already exist; to create
// additional stack instances, use CreateStackInstances
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_CreateStackInstances.html).
// During stack set updates, any parameters overridden for a stack instance aren't
// updated, but retain their overridden value. You can only update the parameter
// values that are specified in the stack set; to add or delete a parameter itself,
// use UpdateStackSet
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
// to update the stack set template. If you add a parameter to a template, before
// you can override the parameter value specified in the stack set you must first
// use UpdateStackSet
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
// to update all stack instances with the updated template and parameter value
// specified in the stack set. Once a stack instance has been updated with the new
// parameter, you can then override the parameter value using UpdateStackInstances.
func (c *Client) UpdateStackInstances(ctx context.Context, params *UpdateStackInstancesInput, optFns ...func(*Options)) (*UpdateStackInstancesOutput, error) {
	if params == nil {
		params = &UpdateStackInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStackInstances", params, optFns, c.addOperationUpdateStackInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStackInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStackInstancesInput struct {

	// The names of one or more Amazon Web Services Regions in which you want to update
	// parameter values for stack instances. The overridden parameter values will be
	// applied to all stack instances in the specified accounts and Amazon Web Services
	// Regions.
	//
	// This member is required.
	Regions []string

	// The name or unique ID of the stack set associated with the stack instances.
	//
	// This member is required.
	StackSetName *string

	// [Self-managed permissions] The names of one or more Amazon Web Services accounts
	// for which you want to update parameter values for stack instances. The
	// overridden parameter values will be applied to all stack instances in the
	// specified accounts and Amazon Web Services Regions. You can specify Accounts or
	// DeploymentTargets, but not both.
	Accounts []string

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account. By default, SELF is specified. Use SELF for
	// stack sets with self-managed permissions.
	//
	// * If you are signed in to the
	// management account, specify SELF.
	//
	// * If you are signed in to a delegated
	// administrator account, specify DELEGATED_ADMIN. Your Amazon Web Services account
	// must be registered as a delegated administrator in the management account. For
	// more information, see Register a delegated administrator
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	// in the CloudFormation User Guide.
	CallAs types.CallAs

	// [Service-managed permissions] The Organizations accounts for which you want to
	// update parameter values for stack instances. If your update targets OUs, the
	// overridden parameter values only apply to the accounts that are currently in the
	// target OUs and their child OUs. Accounts added to the target OUs and their child
	// OUs in the future won't use the overridden values. You can specify Accounts or
	// DeploymentTargets, but not both.
	DeploymentTargets *types.DeploymentTargets

	// The unique identifier for this stack set operation. The operation ID also
	// functions as an idempotency token, to ensure that CloudFormation performs the
	// stack set operation only once, even if you retry the request multiple times. You
	// might retry stack set operation requests to ensure that CloudFormation
	// successfully received them. If you don't specify an operation ID, the SDK
	// generates one automatically.
	OperationId *string

	// Preferences for how CloudFormation performs this stack set operation.
	OperationPreferences *types.StackSetOperationPreferences

	// A list of input parameters whose values you want to update for the specified
	// stack instances. Any overridden parameter values will be applied to all stack
	// instances in the specified accounts and Amazon Web Services Regions. When
	// specifying parameters and their values, be aware of how CloudFormation sets
	// parameter values during stack instance update operations:
	//
	// * To override the
	// current value for a parameter, include the parameter and specify its value.
	//
	// *
	// To leave an overridden parameter set to its present value, include the parameter
	// and specify UsePreviousValue as true. (You can't specify both a value and set
	// UsePreviousValue to true.)
	//
	// * To set an overridden parameter back to the value
	// specified in the stack set, specify a parameter list but don't include the
	// parameter in the list.
	//
	// * To leave all parameters set to their present values,
	// don't specify this property at all.
	//
	// During stack set updates, any parameter
	// values overridden for a stack instance aren't updated, but retain their
	// overridden value. You can only override the parameter values that are specified
	// in the stack set; to add or delete a parameter itself, use UpdateStackSet to
	// update the stack set template. If you add a parameter to a template, before you
	// can override the parameter value specified in the stack set you must first use
	// UpdateStackSet
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStackSet.html)
	// to update all stack instances with the updated template and parameter value
	// specified in the stack set. Once a stack instance has been updated with the new
	// parameter, you can then override the parameter value using UpdateStackInstances.
	ParameterOverrides []types.Parameter

	noSmithyDocumentSerde
}

type UpdateStackInstancesOutput struct {

	// The unique identifier for this stack set operation.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStackInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateStackInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateStackInstances{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateStackInstancesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateStackInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStackInstances(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateStackInstances struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateStackInstances) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateStackInstances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateStackInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateStackInstancesInput ")
	}

	if input.OperationId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.OperationId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateStackInstancesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateStackInstances{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateStackInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "UpdateStackInstances",
	}
}
