// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this API to register a user's entered time-based one-time password (TOTP)
// code and mark the user's software token MFA status as "verified" if successful.
// The request takes an access token or a session string, but not both.
func (c *Client) VerifySoftwareToken(ctx context.Context, params *VerifySoftwareTokenInput, optFns ...func(*Options)) (*VerifySoftwareTokenOutput, error) {
	if params == nil {
		params = &VerifySoftwareTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifySoftwareToken", params, optFns, c.addOperationVerifySoftwareTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifySoftwareTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type VerifySoftwareTokenInput struct {

	// The one- time password computed using the secret code returned by
	// AssociateSoftwareToken
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AssociateSoftwareToken.html).
	//
	// This member is required.
	UserCode *string

	// A valid access token that Amazon Cognito issued to the user whose software token
	// you want to verify.
	AccessToken *string

	// The friendly device name.
	FriendlyDeviceName *string

	// The session that should be passed both ways in challenge-response calls to the
	// service.
	Session *string

	noSmithyDocumentSerde
}

type VerifySoftwareTokenOutput struct {

	// The session that should be passed both ways in challenge-response calls to the
	// service.
	Session *string

	// The status of the verify software token.
	Status types.VerifySoftwareTokenResponseType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifySoftwareTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpVerifySoftwareToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerifySoftwareToken{}, middleware.After)
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
	if err = addOpVerifySoftwareTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifySoftwareToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifySoftwareToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "VerifySoftwareToken",
	}
}
