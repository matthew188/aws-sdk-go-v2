// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a user under an Amazon Chime AppInstance. The request consists of a
// unique appInstanceUserId and Name for that user.
func (c *Client) CreateAppInstanceUser(ctx context.Context, params *CreateAppInstanceUserInput, optFns ...func(*Options)) (*CreateAppInstanceUserOutput, error) {
	if params == nil {
		params = &CreateAppInstanceUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppInstanceUser", params, optFns, c.addOperationCreateAppInstanceUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppInstanceUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppInstanceUserInput struct {

	// The ARN of the AppInstance request.
	//
	// This member is required.
	AppInstanceArn *string

	// The user ID of the AppInstance.
	//
	// This member is required.
	AppInstanceUserId *string

	// The unique ID of the request. Use different tokens to request additional
	// AppInstances.
	//
	// This member is required.
	ClientRequestToken *string

	// The user's name.
	//
	// This member is required.
	Name *string

	// Settings that control the interval after which the AppInstanceUser is
	// automatically deleted.
	ExpirationSettings *types.ExpirationSettings

	// The request's metadata. Limited to a 1KB string in UTF-8.
	Metadata *string

	// Tags assigned to the AppInstanceUser.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAppInstanceUserOutput struct {

	// The user's ARN.
	AppInstanceUserArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppInstanceUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAppInstanceUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAppInstanceUser{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAppInstanceUserMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAppInstanceUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppInstanceUser(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAppInstanceUser struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAppInstanceUser) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAppInstanceUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAppInstanceUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAppInstanceUserInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAppInstanceUserMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAppInstanceUser{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAppInstanceUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateAppInstanceUser",
	}
}
