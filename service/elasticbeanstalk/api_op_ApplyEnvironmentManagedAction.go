// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies a scheduled managed action immediately. A managed action can be applied
// only if its status is Scheduled. Get the status and action ID of a managed
// action with DescribeEnvironmentManagedActions.
func (c *Client) ApplyEnvironmentManagedAction(ctx context.Context, params *ApplyEnvironmentManagedActionInput, optFns ...func(*Options)) (*ApplyEnvironmentManagedActionOutput, error) {
	if params == nil {
		params = &ApplyEnvironmentManagedActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApplyEnvironmentManagedAction", params, optFns, c.addOperationApplyEnvironmentManagedActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ApplyEnvironmentManagedActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to execute a scheduled managed action immediately.
type ApplyEnvironmentManagedActionInput struct {

	// The action ID of the scheduled managed action to execute.
	//
	// This member is required.
	ActionId *string

	// The environment ID of the target environment.
	EnvironmentId *string

	// The name of the target environment.
	EnvironmentName *string

	noSmithyDocumentSerde
}

// The result message containing information about the managed action.
type ApplyEnvironmentManagedActionOutput struct {

	// A description of the managed action.
	ActionDescription *string

	// The action ID of the managed action.
	ActionId *string

	// The type of managed action.
	ActionType types.ActionType

	// The status of the managed action.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationApplyEnvironmentManagedActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpApplyEnvironmentManagedAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpApplyEnvironmentManagedAction{}, middleware.After)
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
	if err = addOpApplyEnvironmentManagedActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opApplyEnvironmentManagedAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opApplyEnvironmentManagedAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ApplyEnvironmentManagedAction",
	}
}
