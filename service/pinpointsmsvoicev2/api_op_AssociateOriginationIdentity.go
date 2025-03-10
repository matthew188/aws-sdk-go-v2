// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the specified origination identity with a pool. If the origination
// identity is a phone number and is already associated with another pool, an Error
// is returned. A sender ID can be associated with multiple pools. If the
// origination identity configuration doesn't match the pool's configuration, an
// Error is returned.
func (c *Client) AssociateOriginationIdentity(ctx context.Context, params *AssociateOriginationIdentityInput, optFns ...func(*Options)) (*AssociateOriginationIdentityOutput, error) {
	if params == nil {
		params = &AssociateOriginationIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateOriginationIdentity", params, optFns, c.addOperationAssociateOriginationIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateOriginationIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateOriginationIdentityInput struct {

	// The new two-character code, in ISO 3166-1 alpha-2 format, for the country or
	// region of the origination identity.
	//
	// This member is required.
	IsoCountryCode *string

	// The origination identity to use, such as PhoneNumberId, PhoneNumberArn,
	// SenderId, or SenderIdArn. You can use DescribePhoneNumbers to find the values
	// for PhoneNumberId and PhoneNumberArn, while DescribeSenderIds can be used to get
	// the values for SenderId and SenderIdArn.
	//
	// This member is required.
	OriginationIdentity *string

	// The pool to update with the new Identity. This value can be either the PoolId or
	// PoolArn, and you can find these values using DescribePools.
	//
	// This member is required.
	PoolId *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don't specify a client token, a randomly generated token is
	// used for the request to ensure idempotency.
	ClientToken *string

	noSmithyDocumentSerde
}

type AssociateOriginationIdentityOutput struct {

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	IsoCountryCode *string

	// The PhoneNumberId or SenderId of the origination identity.
	OriginationIdentity *string

	// The PhoneNumberArn or SenderIdArn of the origination identity.
	OriginationIdentityArn *string

	// The Amazon Resource Name (ARN) of the pool that is now associated with the
	// origination identity.
	PoolArn *string

	// The PoolId of the pool that is now associated with the origination identity.
	PoolId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateOriginationIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpAssociateOriginationIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpAssociateOriginationIdentity{}, middleware.After)
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
	if err = addIdempotencyToken_opAssociateOriginationIdentityMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAssociateOriginationIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateOriginationIdentity(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpAssociateOriginationIdentity struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAssociateOriginationIdentity) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAssociateOriginationIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AssociateOriginationIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AssociateOriginationIdentityInput ")
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
func addIdempotencyToken_opAssociateOriginationIdentityMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpAssociateOriginationIdentity{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAssociateOriginationIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "AssociateOriginationIdentity",
	}
}
