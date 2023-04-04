// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to create a workflow with specified steps and step details the
// workflow invokes after file transfer completes. After creating a workflow, you
// can associate the workflow created with any transfer servers by specifying the
// workflow-details field in CreateServer and UpdateServer operations.
func (c *Client) CreateWorkflow(ctx context.Context, params *CreateWorkflowInput, optFns ...func(*Options)) (*CreateWorkflowOutput, error) {
	if params == nil {
		params = &CreateWorkflowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkflow", params, optFns, c.addOperationCreateWorkflowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkflowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkflowInput struct {

	// Specifies the details for the steps that are in the specified workflow. The TYPE
	// specifies which of the following actions is being taken for this step.
	//
	// * COPY -
	// Copy the file to another location.
	//
	// * CUSTOM - Perform a custom step with an
	// Lambda function target.
	//
	// * DECRYPT - Decrypt a file that was encrypted before it
	// was uploaded.
	//
	// * DELETE - Delete the file.
	//
	// * TAG - Add a tag to the
	// file.
	//
	// Currently, copying and tagging are supported only on S3. For file
	// location, you specify either the Amazon S3 bucket and key, or the Amazon EFS
	// file system ID and path.
	//
	// This member is required.
	Steps []types.WorkflowStep

	// A textual description for the workflow.
	Description *string

	// Specifies the steps (actions) to take if errors are encountered during execution
	// of the workflow. For custom steps, the lambda function needs to send FAILURE to
	// the call back API to kick off the exception steps. Additionally, if the lambda
	// does not send SUCCESS before it times out, the exception steps are executed.
	OnExceptionSteps []types.WorkflowStep

	// Key-value pairs that can be used to group and search for workflows. Tags are
	// metadata attached to workflows for any purpose.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateWorkflowOutput struct {

	// A unique identifier for the workflow.
	//
	// This member is required.
	WorkflowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkflowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWorkflow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWorkflow{}, middleware.After)
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
	if err = addOpCreateWorkflowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkflow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkflow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateWorkflow",
	}
}
