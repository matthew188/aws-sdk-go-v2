// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a self-service action.
func (c *Client) CreateServiceAction(ctx context.Context, params *CreateServiceActionInput, optFns ...func(*Options)) (*CreateServiceActionOutput, error) {
	if params == nil {
		params = &CreateServiceActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateServiceAction", params, optFns, c.addOperationCreateServiceActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceActionInput struct {

	// The self-service action definition. Can be one of the following: Name The name
	// of the Amazon Web Services Systems Manager document (SSM document). For example,
	// AWS-RestartEC2Instance. If you are using a shared SSM document, you must provide
	// the ARN instead of the name. Version The Amazon Web Services Systems Manager
	// automation document version. For example, "Version": "1" AssumeRole The Amazon
	// Resource Name (ARN) of the role that performs the self-service actions on your
	// behalf. For example, "AssumeRole": "arn:aws:iam::12345678910:role/ActionRole".
	// To reuse the provisioned product launch role, set to "AssumeRole":
	// "LAUNCH_ROLE". Parameters The list of parameters in JSON format. For example:
	// [{\"Name\":\"InstanceId\",\"Type\":\"TARGET\"}] or
	// [{\"Name\":\"InstanceId\",\"Type\":\"TEXT_VALUE\"}].
	//
	// This member is required.
	Definition map[string]string

	// The service action definition type. For example, SSM_AUTOMATION.
	//
	// This member is required.
	DefinitionType types.ServiceActionDefinitionType

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// This member is required.
	IdempotencyToken *string

	// The self-service action name.
	//
	// This member is required.
	Name *string

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The self-service action description.
	Description *string

	noSmithyDocumentSerde
}

type CreateServiceActionOutput struct {

	// An object containing information about the self-service action.
	ServiceActionDetail *types.ServiceActionDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateServiceAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateServiceAction{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateServiceActionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateServiceActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateServiceAction(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateServiceAction struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateServiceAction) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateServiceAction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateServiceActionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateServiceActionInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateServiceActionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateServiceAction{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateServiceAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "CreateServiceAction",
	}
}
