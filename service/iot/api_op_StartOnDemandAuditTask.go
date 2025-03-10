// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an on-demand Device Defender audit. Requires permission to access the
// StartOnDemandAuditTask
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) StartOnDemandAuditTask(ctx context.Context, params *StartOnDemandAuditTaskInput, optFns ...func(*Options)) (*StartOnDemandAuditTaskOutput, error) {
	if params == nil {
		params = &StartOnDemandAuditTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartOnDemandAuditTask", params, optFns, c.addOperationStartOnDemandAuditTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartOnDemandAuditTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOnDemandAuditTaskInput struct {

	// Which checks are performed during the audit. The checks you specify must be
	// enabled for your account or an exception occurs. Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or UpdateAccountAuditConfiguration to select which checks are
	// enabled.
	//
	// This member is required.
	TargetCheckNames []string

	noSmithyDocumentSerde
}

type StartOnDemandAuditTaskOutput struct {

	// The ID of the on-demand audit you started.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartOnDemandAuditTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartOnDemandAuditTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartOnDemandAuditTask{}, middleware.After)
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
	if err = addOpStartOnDemandAuditTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartOnDemandAuditTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartOnDemandAuditTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "StartOnDemandAuditTask",
	}
}
