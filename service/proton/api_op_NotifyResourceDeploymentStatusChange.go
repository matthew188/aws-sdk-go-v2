// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Notify Proton of status changes to a provisioned resource when you use
// self-managed provisioning. For more information, see Self-managed provisioning
// (https://docs.aws.amazon.com/proton/latest/userguide/ag-works-prov-methods.html#ag-works-prov-methods-self)
// in the Proton User Guide.
func (c *Client) NotifyResourceDeploymentStatusChange(ctx context.Context, params *NotifyResourceDeploymentStatusChangeInput, optFns ...func(*Options)) (*NotifyResourceDeploymentStatusChangeOutput, error) {
	if params == nil {
		params = &NotifyResourceDeploymentStatusChangeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NotifyResourceDeploymentStatusChange", params, optFns, c.addOperationNotifyResourceDeploymentStatusChangeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*NotifyResourceDeploymentStatusChangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NotifyResourceDeploymentStatusChangeInput struct {

	// The provisioned resource Amazon Resource Name (ARN).
	//
	// This member is required.
	ResourceArn *string

	// The deployment ID for your provisioned resource.
	DeploymentId *string

	// The provisioned resource state change detail data that's returned by Proton.
	Outputs []types.Output

	// The status of your provisioned resource.
	Status types.ResourceDeploymentStatus

	// The deployment status message for your provisioned resource.
	StatusMessage *string

	noSmithyDocumentSerde
}

type NotifyResourceDeploymentStatusChangeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationNotifyResourceDeploymentStatusChangeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpNotifyResourceDeploymentStatusChange{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpNotifyResourceDeploymentStatusChange{}, middleware.After)
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
	if err = addOpNotifyResourceDeploymentStatusChangeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opNotifyResourceDeploymentStatusChange(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opNotifyResourceDeploymentStatusChange(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "NotifyResourceDeploymentStatusChange",
	}
}
