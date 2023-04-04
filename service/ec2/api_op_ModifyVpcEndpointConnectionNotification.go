// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies a connection notification for VPC endpoint or VPC endpoint service. You
// can change the SNS topic for the notification, or the events for which to be
// notified.
func (c *Client) ModifyVpcEndpointConnectionNotification(ctx context.Context, params *ModifyVpcEndpointConnectionNotificationInput, optFns ...func(*Options)) (*ModifyVpcEndpointConnectionNotificationOutput, error) {
	if params == nil {
		params = &ModifyVpcEndpointConnectionNotificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcEndpointConnectionNotification", params, optFns, c.addOperationModifyVpcEndpointConnectionNotificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcEndpointConnectionNotificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcEndpointConnectionNotificationInput struct {

	// The ID of the notification.
	//
	// This member is required.
	ConnectionNotificationId *string

	// The events for the endpoint. Valid values are Accept, Connect, Delete, and
	// Reject.
	ConnectionEvents []string

	// The ARN for the SNS topic for the notification.
	ConnectionNotificationArn *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyVpcEndpointConnectionNotificationOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVpcEndpointConnectionNotificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcEndpointConnectionNotification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcEndpointConnectionNotification{}, middleware.After)
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
	if err = addOpModifyVpcEndpointConnectionNotificationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcEndpointConnectionNotification(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpcEndpointConnectionNotification(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcEndpointConnectionNotification",
	}
}
