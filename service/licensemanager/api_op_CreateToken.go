// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a long-lived token. A refresh token is a JWT token used to get an access
// token. With an access token, you can call AssumeRoleWithWebIdentity to get role
// credentials that you can use to call License Manager to manage the specified
// license.
func (c *Client) CreateToken(ctx context.Context, params *CreateTokenInput, optFns ...func(*Options)) (*CreateTokenOutput, error) {
	if params == nil {
		params = &CreateTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateToken", params, optFns, c.addOperationCreateTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTokenInput struct {

	// Idempotency token, valid for 10 minutes.
	//
	// This member is required.
	ClientToken *string

	// Amazon Resource Name (ARN) of the license. The ARN is mapped to the aud claim of
	// the JWT token.
	//
	// This member is required.
	LicenseArn *string

	// Token expiration, in days, counted from token creation. The default is 365 days.
	ExpirationInDays *int32

	// Amazon Resource Name (ARN) of the IAM roles to embed in the token. License
	// Manager does not check whether the roles are in use.
	RoleArns []string

	// Data specified by the caller to be included in the JWT token. The data is mapped
	// to the amr claim of the JWT token.
	TokenProperties []string

	noSmithyDocumentSerde
}

type CreateTokenOutput struct {

	// Refresh token, encoded as a JWT token.
	Token *string

	// Token ID.
	TokenId *string

	// Token type.
	TokenType types.TokenType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateToken{}, middleware.After)
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
	if err = addOpCreateTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "CreateToken",
	}
}
