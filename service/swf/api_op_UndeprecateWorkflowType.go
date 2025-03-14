// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/swf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Undeprecates a previously deprecated workflow type. After a workflow type has
// been undeprecated, you can create new executions of that type. This operation is
// eventually consistent. The results are best effort and may not exactly reflect
// recent updates and changes. Access Control You can use IAM policies to control
// this action's access to Amazon SWF resources as follows:
//
// * Use a Resource
// element with the domain name to limit the action to only specified domains.
//
// *
// Use an Action element to allow or deny permission to call this action.
//
// *
// Constrain the following parameters by using a Condition element with the
// appropriate keys.
//
// * workflowType.name: String constraint. The key is
// swf:workflowType.name.
//
// * workflowType.version: String constraint. The key is
// swf:workflowType.version.
//
// If the caller doesn't have sufficient permissions to
// invoke the action, or the parameter values fall outside the specified
// constraints, the action fails. The associated event attribute's cause parameter
// is set to OPERATION_NOT_PERMITTED. For details and example IAM policies, see
// Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) UndeprecateWorkflowType(ctx context.Context, params *UndeprecateWorkflowTypeInput, optFns ...func(*Options)) (*UndeprecateWorkflowTypeOutput, error) {
	if params == nil {
		params = &UndeprecateWorkflowTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UndeprecateWorkflowType", params, optFns, c.addOperationUndeprecateWorkflowTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UndeprecateWorkflowTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UndeprecateWorkflowTypeInput struct {

	// The name of the domain of the deprecated workflow type.
	//
	// This member is required.
	Domain *string

	// The name of the domain of the deprecated workflow type.
	//
	// This member is required.
	WorkflowType *types.WorkflowType

	noSmithyDocumentSerde
}

type UndeprecateWorkflowTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUndeprecateWorkflowTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUndeprecateWorkflowType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUndeprecateWorkflowType{}, middleware.After)
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
	if err = addOpUndeprecateWorkflowTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUndeprecateWorkflowType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUndeprecateWorkflowType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "UndeprecateWorkflowType",
	}
}
