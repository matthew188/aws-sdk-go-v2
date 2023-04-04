// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Confirms user registration as an admin without using a confirmation code. Works
// on any user. Calling this action requires developer credentials.
func (c *Client) AdminConfirmSignUp(ctx context.Context, params *AdminConfirmSignUpInput, optFns ...func(*Options)) (*AdminConfirmSignUpOutput, error) {
	if params == nil {
		params = &AdminConfirmSignUpInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminConfirmSignUp", params, optFns, c.addOperationAdminConfirmSignUpMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminConfirmSignUpOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to confirm user registration.
type AdminConfirmSignUpInput struct {

	// The user pool ID for which you want to confirm user registration.
	//
	// This member is required.
	UserPoolId *string

	// The user name for which you want to confirm user registration.
	//
	// This member is required.
	Username *string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. If your user pool configuration includes
	// triggers, the AdminConfirmSignUp API action invokes the Lambda function that is
	// specified for the post confirmation trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. In
	// this payload, the clientMetadata attribute provides the data that you assigned
	// to the ClientMetadata parameter in your AdminConfirmSignUp request. In your
	// function code in Lambda, you can process the ClientMetadata value to enhance
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

// Represents the response from the server for the request to confirm registration.
type AdminConfirmSignUpOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminConfirmSignUpMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminConfirmSignUp{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminConfirmSignUp{}, middleware.After)
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
	if err = addOpAdminConfirmSignUpValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminConfirmSignUp(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminConfirmSignUp(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminConfirmSignUp",
	}
}
