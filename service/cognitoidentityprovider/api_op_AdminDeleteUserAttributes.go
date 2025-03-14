// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the user attributes in a user pool as an administrator. Works on any
// user. Calling this action requires developer credentials.
func (c *Client) AdminDeleteUserAttributes(ctx context.Context, params *AdminDeleteUserAttributesInput, optFns ...func(*Options)) (*AdminDeleteUserAttributesOutput, error) {
	if params == nil {
		params = &AdminDeleteUserAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminDeleteUserAttributes", params, optFns, c.addOperationAdminDeleteUserAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminDeleteUserAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to delete user attributes as an administrator.
type AdminDeleteUserAttributesInput struct {

	// An array of strings representing the user attribute names you want to delete.
	// For custom attributes, you must prepend the custom: prefix to the attribute
	// name.
	//
	// This member is required.
	UserAttributeNames []string

	// The user pool ID for the user pool where you want to delete user attributes.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user from which you would like to delete attributes.
	//
	// This member is required.
	Username *string

	noSmithyDocumentSerde
}

// Represents the response received from the server for a request to delete user
// attributes.
type AdminDeleteUserAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminDeleteUserAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminDeleteUserAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminDeleteUserAttributes{}, middleware.After)
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
	if err = addOpAdminDeleteUserAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminDeleteUserAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminDeleteUserAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminDeleteUserAttributes",
	}
}
