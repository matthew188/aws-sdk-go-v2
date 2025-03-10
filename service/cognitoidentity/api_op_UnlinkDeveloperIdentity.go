// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Unlinks a DeveloperUserIdentifier from an existing identity. Unlinked developer
// users will be considered new identities next time they are seen. If, for a given
// Cognito identity, you remove all federated identities as well as the developer
// user identifier, the Cognito identity becomes inaccessible. You must use AWS
// Developer credentials to call this API.
func (c *Client) UnlinkDeveloperIdentity(ctx context.Context, params *UnlinkDeveloperIdentityInput, optFns ...func(*Options)) (*UnlinkDeveloperIdentityOutput, error) {
	if params == nil {
		params = &UnlinkDeveloperIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnlinkDeveloperIdentity", params, optFns, c.addOperationUnlinkDeveloperIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UnlinkDeveloperIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the UnlinkDeveloperIdentity action.
type UnlinkDeveloperIdentityInput struct {

	// The "domain" by which Cognito will refer to your users.
	//
	// This member is required.
	DeveloperProviderName *string

	// A unique ID used by your backend authentication process to identify a user.
	//
	// This member is required.
	DeveloperUserIdentifier *string

	// A unique identifier in the format REGION:GUID.
	//
	// This member is required.
	IdentityId *string

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	noSmithyDocumentSerde
}

type UnlinkDeveloperIdentityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUnlinkDeveloperIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUnlinkDeveloperIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUnlinkDeveloperIdentity{}, middleware.After)
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
	if err = addOpUnlinkDeveloperIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUnlinkDeveloperIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUnlinkDeveloperIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "UnlinkDeveloperIdentity",
	}
}
