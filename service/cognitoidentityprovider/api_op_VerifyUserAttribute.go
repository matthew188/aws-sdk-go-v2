// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Verifies the specified user attributes in the user pool. If your user pool
// requires verification before Amazon Cognito updates the attribute value,
// VerifyUserAttribute updates the affected attribute to its pending value. For
// more information, see  UserAttributeUpdateSettingsType
// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserAttributeUpdateSettingsType.html).
func (c *Client) VerifyUserAttribute(ctx context.Context, params *VerifyUserAttributeInput, optFns ...func(*Options)) (*VerifyUserAttributeOutput, error) {
	if params == nil {
		params = &VerifyUserAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyUserAttribute", params, optFns, c.addOperationVerifyUserAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyUserAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to verify user attributes.
type VerifyUserAttributeInput struct {

	// A valid access token that Amazon Cognito issued to the user whose user
	// attributes you want to verify.
	//
	// This member is required.
	AccessToken *string

	// The attribute name in the request to verify user attributes.
	//
	// This member is required.
	AttributeName *string

	// The verification code in the request to verify user attributes.
	//
	// This member is required.
	Code *string

	noSmithyDocumentSerde
}

// A container representing the response from the server from the request to verify
// user attributes.
type VerifyUserAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyUserAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpVerifyUserAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerifyUserAttribute{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpVerifyUserAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyUserAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifyUserAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifyUserAttribute",
	}
}
