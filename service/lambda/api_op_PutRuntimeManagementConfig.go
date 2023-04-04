// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the runtime management configuration for a function's version. For more
// information, see Runtime updates
// (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html).
func (c *Client) PutRuntimeManagementConfig(ctx context.Context, params *PutRuntimeManagementConfigInput, optFns ...func(*Options)) (*PutRuntimeManagementConfigOutput, error) {
	if params == nil {
		params = &PutRuntimeManagementConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRuntimeManagementConfig", params, optFns, c.addOperationPutRuntimeManagementConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRuntimeManagementConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRuntimeManagementConfigInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name – my-function.
	//
	// *
	// Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// *
	// Partial ARN – 123456789012:function:my-function.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// Specify the runtime update mode.
	//
	// * Auto (default) - Automatically update to the
	// most recent and secure runtime version using a Two-phase runtime version rollout
	// (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-two-phase).
	// This is the best choice for most customers to ensure they always benefit from
	// runtime updates.
	//
	// * Function update - Lambda updates the runtime of your
	// function to the most recent and secure runtime version when you update your
	// function. This approach synchronizes runtime updates with function deployments,
	// giving you control over when runtime updates are applied and allowing you to
	// detect and mitigate rare runtime update incompatibilities early. When using this
	// setting, you need to regularly update your functions to keep their runtime
	// up-to-date.
	//
	// * Manual - You specify a runtime version in your function
	// configuration. The function will use this runtime version indefinitely. In the
	// rare case where a new runtime version is incompatible with an existing function,
	// this allows you to roll back your function to an earlier runtime version. For
	// more information, see Roll back a runtime version
	// (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-rollback).
	//
	// This member is required.
	UpdateRuntimeOn types.UpdateRuntimeOn

	// Specify a version of the function. This can be $LATEST or a published version
	// number. If no value is specified, the configuration for the $LATEST version is
	// returned.
	Qualifier *string

	// The ARN of the runtime version you want the function to use. This is only
	// required if you're using the Manual runtime update mode.
	RuntimeVersionArn *string

	noSmithyDocumentSerde
}

type PutRuntimeManagementConfigOutput struct {

	// The ARN of the function
	//
	// This member is required.
	FunctionArn *string

	// The runtime update mode.
	//
	// This member is required.
	UpdateRuntimeOn types.UpdateRuntimeOn

	// The ARN of the runtime the function is configured to use. If the runtime update
	// mode is manual, the ARN is returned, otherwise null is returned.
	RuntimeVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRuntimeManagementConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutRuntimeManagementConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRuntimeManagementConfig{}, middleware.After)
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
	if err = addOpPutRuntimeManagementConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRuntimeManagementConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRuntimeManagementConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "PutRuntimeManagementConfig",
	}
}
