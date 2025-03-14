// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an impersonation role for the given WorkMail organization.
func (c *Client) UpdateImpersonationRole(ctx context.Context, params *UpdateImpersonationRoleInput, optFns ...func(*Options)) (*UpdateImpersonationRoleOutput, error) {
	if params == nil {
		params = &UpdateImpersonationRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateImpersonationRole", params, optFns, c.addOperationUpdateImpersonationRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateImpersonationRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateImpersonationRoleInput struct {

	// The ID of the impersonation role to update.
	//
	// This member is required.
	ImpersonationRoleId *string

	// The updated impersonation role name.
	//
	// This member is required.
	Name *string

	// The WorkMail organization that contains the impersonation role to update.
	//
	// This member is required.
	OrganizationId *string

	// The updated list of rules.
	//
	// This member is required.
	Rules []types.ImpersonationRule

	// The updated impersonation role type.
	//
	// This member is required.
	Type types.ImpersonationRoleType

	// The updated impersonation role description.
	Description *string

	noSmithyDocumentSerde
}

type UpdateImpersonationRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateImpersonationRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateImpersonationRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateImpersonationRole{}, middleware.After)
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
	if err = addOpUpdateImpersonationRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateImpersonationRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateImpersonationRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "UpdateImpersonationRole",
	}
}
