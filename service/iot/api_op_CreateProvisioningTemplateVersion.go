// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of a provisioning template. Requires permission to access
// the CreateProvisioningTemplateVersion
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreateProvisioningTemplateVersion(ctx context.Context, params *CreateProvisioningTemplateVersionInput, optFns ...func(*Options)) (*CreateProvisioningTemplateVersionOutput, error) {
	if params == nil {
		params = &CreateProvisioningTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProvisioningTemplateVersion", params, optFns, c.addOperationCreateProvisioningTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProvisioningTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProvisioningTemplateVersionInput struct {

	// The JSON formatted contents of the provisioning template.
	//
	// This member is required.
	TemplateBody *string

	// The name of the provisioning template.
	//
	// This member is required.
	TemplateName *string

	// Sets a fleet provision template version as the default version.
	SetAsDefault bool

	noSmithyDocumentSerde
}

type CreateProvisioningTemplateVersionOutput struct {

	// True if the provisioning template version is the default version, otherwise
	// false.
	IsDefaultVersion bool

	// The ARN that identifies the provisioning template.
	TemplateArn *string

	// The name of the provisioning template.
	TemplateName *string

	// The version of the provisioning template.
	VersionId *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProvisioningTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProvisioningTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProvisioningTemplateVersion{}, middleware.After)
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
	if err = addOpCreateProvisioningTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProvisioningTemplateVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProvisioningTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateProvisioningTemplateVersion",
	}
}
