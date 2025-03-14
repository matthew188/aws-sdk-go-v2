// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use UpdateRole instead. Modifies only the description of a role. This operation
// performs the same function as the Description parameter in the UpdateRole
// operation.
func (c *Client) UpdateRoleDescription(ctx context.Context, params *UpdateRoleDescriptionInput, optFns ...func(*Options)) (*UpdateRoleDescriptionOutput, error) {
	if params == nil {
		params = &UpdateRoleDescriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoleDescription", params, optFns, c.addOperationUpdateRoleDescriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoleDescriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoleDescriptionInput struct {

	// The new description that you want to apply to the specified role.
	//
	// This member is required.
	Description *string

	// The name of the role that you want to modify.
	//
	// This member is required.
	RoleName *string

	noSmithyDocumentSerde
}

type UpdateRoleDescriptionOutput struct {

	// A structure that contains details about the modified role.
	Role *types.Role

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRoleDescriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateRoleDescription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateRoleDescription{}, middleware.After)
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
	if err = addOpUpdateRoleDescriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoleDescription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRoleDescription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateRoleDescription",
	}
}
