// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a knowledge base. When using this API, you cannot reuse Amazon
// AppIntegrations
// (https://docs.aws.amazon.com/appintegrations/latest/APIReference/Welcome.html)
// DataIntegrations with external knowledge bases such as Salesforce and
// ServiceNow. If you do, you'll get an InvalidRequestException error. For example,
// you're programmatically managing your external knowledge base, and you want to
// add or remove one of the fields that is being ingested from Salesforce. Do the
// following:
//
// * Call DeleteKnowledgeBase
// (https://docs.aws.amazon.com/wisdom/latest/APIReference/API_DeleteKnowledgeBase.html).
//
// *
// Call DeleteDataIntegration
// (https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html).
//
// *
// Call CreateDataIntegration
// (https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html)
// to recreate the DataIntegration or a create different one.
//
// * Call
// CreateKnowledgeBase.
func (c *Client) CreateKnowledgeBase(ctx context.Context, params *CreateKnowledgeBaseInput, optFns ...func(*Options)) (*CreateKnowledgeBaseOutput, error) {
	if params == nil {
		params = &CreateKnowledgeBaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKnowledgeBase", params, optFns, c.addOperationCreateKnowledgeBaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKnowledgeBaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKnowledgeBaseInput struct {

	// The type of knowledge base. Only CUSTOM knowledge bases allow you to upload your
	// own content. EXTERNAL knowledge bases support integrations with third-party
	// systems whose content is synchronized automatically.
	//
	// This member is required.
	KnowledgeBaseType types.KnowledgeBaseType

	// The name of the knowledge base.
	//
	// This member is required.
	Name *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs
	// (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/).
	ClientToken *string

	// The description.
	Description *string

	// Information about how to render the content.
	RenderingConfiguration *types.RenderingConfiguration

	// The KMS key used for encryption.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The source of the knowledge base content. Only set this argument for EXTERNAL
	// knowledge bases.
	SourceConfiguration types.SourceConfiguration

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateKnowledgeBaseOutput struct {

	// The knowledge base.
	KnowledgeBase *types.KnowledgeBaseData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKnowledgeBaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKnowledgeBase{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateKnowledgeBaseMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKnowledgeBaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKnowledgeBase(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateKnowledgeBase struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKnowledgeBase) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKnowledgeBase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKnowledgeBaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKnowledgeBaseInput ")
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
func addIdempotencyToken_opCreateKnowledgeBaseMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKnowledgeBase{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKnowledgeBase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wisdom",
		OperationName: "CreateKnowledgeBase",
	}
}
