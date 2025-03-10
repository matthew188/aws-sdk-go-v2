// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and defines the settings for an allow list.
func (c *Client) CreateAllowList(ctx context.Context, params *CreateAllowListInput, optFns ...func(*Options)) (*CreateAllowListOutput, error) {
	if params == nil {
		params = &CreateAllowListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAllowList", params, optFns, c.addOperationCreateAllowListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAllowListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAllowListInput struct {

	// A unique, case-sensitive token that you provide to ensure the idempotency of the
	// request.
	//
	// This member is required.
	ClientToken *string

	// The criteria that specify the text or text pattern to ignore. The criteria can
	// be the location and name of an S3 object that lists specific text to ignore
	// (s3WordsList), or a regular expression (regex) that defines a text pattern to
	// ignore.
	//
	// This member is required.
	Criteria *types.AllowListCriteria

	// A custom name for the allow list. The name can contain as many as 128
	// characters.
	//
	// This member is required.
	Name *string

	// A custom description of the allow list. The description can contain as many as
	// 512 characters.
	Description *string

	// A map of key-value pairs that specifies the tags to associate with the allow
	// list. An allow list can have a maximum of 50 tags. Each tag consists of a tag
	// key and an associated tag value. The maximum length of a tag key is 128
	// characters. The maximum length of a tag value is 256 characters.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAllowListOutput struct {

	// The Amazon Resource Name (ARN) of the allow list.
	Arn *string

	// The unique identifier for the allow list.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAllowListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAllowList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAllowList{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAllowListMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAllowListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAllowList(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAllowList struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAllowList) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAllowList) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAllowListInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAllowListInput ")
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
func addIdempotencyToken_opCreateAllowListMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAllowList{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAllowList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateAllowList",
	}
}
