// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Resets the specified user's password in a user pool as an administrator. Works
// on any user. When a developer calls this API, the current password is
// invalidated, so it must be changed. If a user tries to sign in after the API is
// called, the app will get a PasswordResetRequiredException exception back and
// should direct the user down the flow to reset the password, which is the same as
// the forgot password flow. In addition, if the user pool has phone verification
// selected and a verified phone number exists for the user, or if email
// verification is selected and a verified email exists for the user, calling this
// API will also result in sending a message to the end user with the code to
// change their password. This action might generate an SMS text message. Starting
// June 1, 2021, US telecom carriers require you to register an origination phone
// number before you can send SMS messages to US phone numbers. If you use SMS text
// messages in Amazon Cognito, you must register a phone number with Amazon
// Pinpoint (https://console.aws.amazon.com/pinpoint/home/). Amazon Cognito uses
// the registered number automatically. Otherwise, Amazon Cognito users who must
// receive SMS messages might not be able to sign up, activate their accounts, or
// sign in. If you have never used SMS text messages with Amazon Cognito or any
// other Amazon Web Service, Amazon Simple Notification Service might place your
// account in the SMS sandbox. In sandbox mode
// (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html) , you can send
// messages only to verified phone numbers. After you test your app while in the
// sandbox environment, you can move out of the sandbox and into production. For
// more information, see  SMS message settings for Amazon Cognito user pools
// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-sms-userpool-settings.html)
// in the Amazon Cognito Developer Guide. Calling this action requires developer
// credentials.
func (c *Client) AdminResetUserPassword(ctx context.Context, params *AdminResetUserPasswordInput, optFns ...func(*Options)) (*AdminResetUserPasswordOutput, error) {
	if params == nil {
		params = &AdminResetUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminResetUserPassword", params, optFns, c.addOperationAdminResetUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminResetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to reset a user's password as an administrator.
type AdminResetUserPasswordInput struct {

	// The user pool ID for the user pool where you want to reset the user's password.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user whose password you want to reset.
	//
	// This member is required.
	Username *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// Lambda functions to user pool triggers. When you use the AdminResetUserPassword
	// API action, Amazon Cognito invokes the function that is assigned to the custom
	// message trigger. When Amazon Cognito invokes this function, it passes a JSON
	// payload, which the function receives as input. This payload contains a
	// clientMetadata attribute, which provides the data that you assigned to the
	// ClientMetadata parameter in your AdminResetUserPassword request. In your
	// function code in Lambda, you can process the clientMetadata value to enhance
	// your workflow for your specific needs. For more information, see  Customizing
	// user pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. When you use the ClientMetadata
	// parameter, remember that Amazon Cognito won't do the following:
	//
	// * Store the
	// ClientMetadata value. This data is available only to Lambda triggers that are
	// assigned to a user pool to support custom workflows. If your user pool
	// configuration doesn't include triggers, the ClientMetadata parameter serves no
	// purpose.
	//
	// * Validate the ClientMetadata value.
	//
	// * Encrypt the ClientMetadata
	// value. Don't use Amazon Cognito to provide sensitive information.
	ClientMetadata map[string]string

	noSmithyDocumentSerde
}

// Represents the response from the server to reset a user password as an
// administrator.
type AdminResetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminResetUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminResetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminResetUserPassword{}, middleware.After)
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
	if err = addOpAdminResetUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminResetUserPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminResetUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminResetUserPassword",
	}
}
