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

// Set up a template to create new template versions automatically by tracking a
// linked repository. A linked repository is a repository that has been registered
// with Proton. For more information, see CreateRepository. When a commit is pushed
// to your linked repository, Proton checks for changes to your repository template
// bundles. If it detects a template bundle change, a new major or minor version of
// its template is created, if the version doesn’t already exist. For more
// information, see Template sync configurations
// (https://docs.aws.amazon.com/proton/latest/userguide/ag-template-sync-configs.html)
// in the Proton User Guide.
func (c *Client) CreateTemplateSyncConfig(ctx context.Context, params *CreateTemplateSyncConfigInput, optFns ...func(*Options)) (*CreateTemplateSyncConfigOutput, error) {
	if params == nil {
		params = &CreateTemplateSyncConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTemplateSyncConfig", params, optFns, c.addOperationCreateTemplateSyncConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTemplateSyncConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTemplateSyncConfigInput struct {

	// The repository branch for your template.
	//
	// This member is required.
	Branch *string

	// The repository name (for example, myrepos/myrepo).
	//
	// This member is required.
	RepositoryName *string

	// The provider type for your repository.
	//
	// This member is required.
	RepositoryProvider types.RepositoryProvider

	// The name of your registered template.
	//
	// This member is required.
	TemplateName *string

	// The type of the registered template.
	//
	// This member is required.
	TemplateType types.TemplateType

	// A repository subdirectory path to your template bundle directory. When included,
	// Proton limits the template bundle search to this repository directory.
	Subdirectory *string

	noSmithyDocumentSerde
}

type CreateTemplateSyncConfigOutput struct {

	// The template sync configuration detail data that's returned by Proton.
	TemplateSyncConfig *types.TemplateSyncConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTemplateSyncConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateTemplateSyncConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateTemplateSyncConfig{}, middleware.After)
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
	if err = addOpCreateTemplateSyncConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTemplateSyncConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTemplateSyncConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "CreateTemplateSyncConfig",
	}
}
