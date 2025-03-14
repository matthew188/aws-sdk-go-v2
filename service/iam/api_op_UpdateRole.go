// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the description or maximum session duration setting of a role.
func (c *Client) UpdateRole(ctx context.Context, params *UpdateRoleInput, optFns ...func(*Options)) (*UpdateRoleOutput, error) {
	if params == nil {
		params = &UpdateRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRole", params, optFns, c.addOperationUpdateRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoleInput struct {

	// The name of the role that you want to modify.
	//
	// This member is required.
	RoleName *string

	// The new description that you want to apply to the specified role.
	Description *string

	// The maximum session duration (in seconds) that you want to set for the specified
	// role. If you do not specify a value for this setting, the default value of one
	// hour is applied. This setting can have a value from 1 hour to 12 hours. Anyone
	// who assumes the role from the CLI or API can use the DurationSeconds API
	// parameter or the duration-seconds CLI parameter to request a longer session. The
	// MaxSessionDuration setting determines the maximum duration that can be requested
	// using the DurationSeconds parameter. If users don't specify a value for the
	// DurationSeconds parameter, their security credentials are valid for one hour by
	// default. This applies when you use the AssumeRole* API operations or the
	// assume-role* CLI operations but does not apply when you use those operations to
	// create a console URL. For more information, see Using IAM roles
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html) in the IAM
	// User Guide.
	MaxSessionDuration *int32

	noSmithyDocumentSerde
}

type UpdateRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateRole{}, middleware.After)
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
	if err = addOpUpdateRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateRole",
	}
}
