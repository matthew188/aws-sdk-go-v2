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

// Runs an on-demand remediation for the specified Config rules against the last
// known remediation configuration. It runs an execution against the current state
// of your resources. Remediation execution is asynchronous. You can specify up to
// 100 resource keys per request. An existing StartRemediationExecution call for
// the specified resource keys must complete before you can call the API again.
func (c *Client) StartRemediationExecution(ctx context.Context, params *StartRemediationExecutionInput, optFns ...func(*Options)) (*StartRemediationExecutionOutput, error) {
	if params == nil {
		params = &StartRemediationExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartRemediationExecution", params, optFns, c.addOperationStartRemediationExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartRemediationExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRemediationExecutionInput struct {

	// The list of names of Config rules that you want to run remediation execution
	// for.
	//
	// This member is required.
	ConfigRuleName *string

	// A list of resource keys to be processed with the current request. Each element
	// in the list consists of the resource type and resource ID.
	//
	// This member is required.
	ResourceKeys []types.ResourceKey

	noSmithyDocumentSerde
}

type StartRemediationExecutionOutput struct {

	// For resources that have failed to start execution, the API returns a resource
	// key object.
	FailedItems []types.ResourceKey

	// Returns a failure message. For example, the resource is already compliant.
	FailureMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartRemediationExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartRemediationExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartRemediationExecution{}, middleware.After)
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
	if err = addOpStartRemediationExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartRemediationExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartRemediationExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "StartRemediationExecution",
	}
}
