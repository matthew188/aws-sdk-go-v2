// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an access configuration for your documents. This includes user and group
// access information for your documents. This is useful for user context
// filtering, where search results are filtered based on the user or their group
// access to documents. You can use this to re-configure your existing document
// level access control without indexing all of your documents again. For example,
// your index contains top-secret company documents that only certain employees or
// users should access. One of these users leaves the company or switches to a team
// that should be blocked from accessing top-secret documents. The user still has
// access to top-secret documents because the user had access when your documents
// were previously indexed. You can create a specific access control configuration
// for the user with deny access. You can later update the access control
// configuration to allow access if the user returns to the company and re-joins
// the 'top-secret' team. You can re-configure access control for your documents as
// circumstances change. To apply your access control configuration to certain
// documents, you call the BatchPutDocument
// (https://docs.aws.amazon.com/kendra/latest/dg/API_BatchPutDocument.html) API
// with the AccessControlConfigurationId included in the Document
// (https://docs.aws.amazon.com/kendra/latest/dg/API_Document.html) object. If you
// use an S3 bucket as a data source, you update the .metadata.json with the
// AccessControlConfigurationId and synchronize your data source. Amazon Kendra
// currently only supports access control configuration for S3 data sources and
// documents indexed using the BatchPutDocument API.
func (c *Client) CreateAccessControlConfiguration(ctx context.Context, params *CreateAccessControlConfigurationInput, optFns ...func(*Options)) (*CreateAccessControlConfigurationOutput, error) {
	if params == nil {
		params = &CreateAccessControlConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessControlConfiguration", params, optFns, c.addOperationCreateAccessControlConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessControlConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessControlConfigurationInput struct {

	// The identifier of the index to create an access control configuration for your
	// documents.
	//
	// This member is required.
	IndexId *string

	// A name for the access control configuration.
	//
	// This member is required.
	Name *string

	// Information on principals (users and/or groups) and which documents they should
	// have access to. This is useful for user context filtering, where search results
	// are filtered based on the user or their group access to documents.
	AccessControlList []types.Principal

	// A token that you provide to identify the request to create an access control
	// configuration. Multiple calls to the CreateAccessControlConfiguration API with
	// the same client token will create only one access control configuration.
	ClientToken *string

	// A description for the access control configuration.
	Description *string

	// The list of principal
	// (https://docs.aws.amazon.com/kendra/latest/dg/API_Principal.html) lists that
	// define the hierarchy for which documents users should have access to.
	HierarchicalAccessControlList []types.HierarchicalPrincipal

	noSmithyDocumentSerde
}

type CreateAccessControlConfigurationOutput struct {

	// The identifier of the access control configuration for your documents in an
	// index.
	//
	// This member is required.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessControlConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAccessControlConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAccessControlConfiguration{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAccessControlConfigurationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAccessControlConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessControlConfiguration(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAccessControlConfiguration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAccessControlConfiguration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAccessControlConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAccessControlConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAccessControlConfigurationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAccessControlConfigurationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAccessControlConfiguration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAccessControlConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateAccessControlConfiguration",
	}
}
