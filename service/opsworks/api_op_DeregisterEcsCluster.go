// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deregisters a specified Amazon ECS cluster from a stack. For more information,
// see  Resource Management
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-ecscluster.html#workinglayers-ecscluster-delete).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack or an attached policy that explicitly grants
// permissions. For more information on user permissions, see
// https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DeregisterEcsCluster(ctx context.Context, params *DeregisterEcsClusterInput, optFns ...func(*Options)) (*DeregisterEcsClusterOutput, error) {
	if params == nil {
		params = &DeregisterEcsClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterEcsCluster", params, optFns, c.addOperationDeregisterEcsClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterEcsClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterEcsClusterInput struct {

	// The cluster's Amazon Resource Number (ARN).
	//
	// This member is required.
	EcsClusterArn *string

	noSmithyDocumentSerde
}

type DeregisterEcsClusterOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterEcsClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterEcsCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterEcsCluster{}, middleware.After)
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
	if err = addOpDeregisterEcsClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterEcsCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterEcsCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DeregisterEcsCluster",
	}
}
