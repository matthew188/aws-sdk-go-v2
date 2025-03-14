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

// Creates an Amazon Kendra index. Index creation is an asynchronous API. To
// determine if index creation has completed, check the Status field returned from
// a call to DescribeIndex. The Status field is set to ACTIVE when the index is
// ready to use. Once the index is active you can index your documents using the
// BatchPutDocument API or using one of the supported data sources. For an example
// of creating an index and data source using the Python SDK, see Getting started
// with Python SDK (https://docs.aws.amazon.com/kendra/latest/dg/gs-python.html).
// For an example of creating an index and data source using the Java SDK, see
// Getting started with Java SDK
// (https://docs.aws.amazon.com/kendra/latest/dg/gs-java.html).
func (c *Client) CreateIndex(ctx context.Context, params *CreateIndexInput, optFns ...func(*Options)) (*CreateIndexOutput, error) {
	if params == nil {
		params = &CreateIndexInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIndex", params, optFns, c.addOperationCreateIndexMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIndexOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIndexInput struct {

	// A name for the index.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of an IAM role with permission to access your
	// Amazon CloudWatch logs and metrics. For more information, see IAM access roles
	// for Amazon Kendra (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	//
	// This member is required.
	RoleArn *string

	// A token that you provide to identify the request to create an index. Multiple
	// calls to the CreateIndex API with the same client token will create only one
	// index.
	ClientToken *string

	// A description for the index.
	Description *string

	// The Amazon Kendra edition to use for the index. Choose DEVELOPER_EDITION for
	// indexes intended for development, testing, or proof of concept. Use
	// ENTERPRISE_EDITION for production. Once you set the edition for an index, it
	// can't be changed. The Edition parameter is optional. If you don't supply a
	// value, the default is ENTERPRISE_EDITION. For more information on quota limits
	// for Enterprise and Developer editions, see Quotas
	// (https://docs.aws.amazon.com/kendra/latest/dg/quotas.html).
	Edition types.IndexEdition

	// The identifier of the KMS customer managed key (CMK) that's used to encrypt data
	// indexed by Amazon Kendra. Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// A list of key-value pairs that identify or categorize the index. You can also
	// use tags to help control access to the index. Tag keys and values can consist of
	// Unicode letters, digits, white space, and any of the following symbols: _ . : /
	// = + - @.
	Tags []types.Tag

	// The user context policy. ATTRIBUTE_FILTER All indexed content is searchable and
	// displayable for all users. If you want to filter search results on user context,
	// you can use the attribute filters of _user_id and _group_ids or you can provide
	// user and group information in UserContext. USER_TOKEN Enables token-based user
	// access control to filter search results on user context. All documents with no
	// access control and all documents accessible to the user will be searchable and
	// displayable.
	UserContextPolicy types.UserContextPolicy

	// Gets users and groups from IAM Identity Center (successor to Single Sign-On)
	// identity source. To configure this, see UserGroupResolutionConfiguration
	// (https://docs.aws.amazon.com/kendra/latest/dg/API_UserGroupResolutionConfiguration.html).
	UserGroupResolutionConfiguration *types.UserGroupResolutionConfiguration

	// The user token configuration.
	UserTokenConfigurations []types.UserTokenConfiguration

	noSmithyDocumentSerde
}

type CreateIndexOutput struct {

	// The identifier of the index. Use this identifier when you query an index, set up
	// a data source, or index a document.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIndexMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateIndex{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateIndex{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateIndexMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateIndexValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIndex(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateIndex struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIndex) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIndex) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIndexInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIndexInput ")
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
func addIdempotencyToken_opCreateIndexMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateIndex{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIndex(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateIndex",
	}
}
