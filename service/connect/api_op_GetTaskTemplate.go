// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets details about a specific task template in the specified Amazon Connect
// instance.
func (c *Client) GetTaskTemplate(ctx context.Context, params *GetTaskTemplateInput, optFns ...func(*Options)) (*GetTaskTemplateOutput, error) {
	if params == nil {
		params = &GetTaskTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTaskTemplate", params, optFns, c.addOperationGetTaskTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTaskTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTaskTemplateInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// A unique identifier for the task template.
	//
	// This member is required.
	TaskTemplateId *string

	// The system generated version of a task template that is associated with a task,
	// when the task is created.
	SnapshotVersion *string

	noSmithyDocumentSerde
}

type GetTaskTemplateOutput struct {

	// The Amazon Resource Name (ARN).
	//
	// This member is required.
	Arn *string

	// A unique identifier for the task template.
	//
	// This member is required.
	Id *string

	// The name of the task template.
	//
	// This member is required.
	Name *string

	// Constraints that are applicable to the fields listed.
	Constraints *types.TaskTemplateConstraints

	// The identifier of the flow that runs by default when a task is created by
	// referencing this template.
	ContactFlowId *string

	// The timestamp when the task template was created.
	CreatedTime *time.Time

	// The default values for fields when a task is created by referencing this
	// template.
	Defaults *types.TaskTemplateDefaults

	// The description of the task template.
	Description *string

	// Fields that are part of the template.
	Fields []types.TaskTemplateField

	// The identifier of the Amazon Connect instance. You can find the instance ID
	// (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	InstanceId *string

	// The timestamp when the task template was last modified.
	LastModifiedTime *time.Time

	// Marks a template as ACTIVE or INACTIVE for a task to refer to it. Tasks can only
	// be created from ACTIVE templates. If a template is marked as INACTIVE, then a
	// task that refers to this template cannot be created.
	Status types.TaskTemplateStatus

	// The tags used to organize, track, or control access for this resource. For
	// example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTaskTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTaskTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTaskTemplate{}, middleware.After)
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
	if err = addOpGetTaskTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTaskTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTaskTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "GetTaskTemplate",
	}
}
