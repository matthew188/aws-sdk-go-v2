// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new user within the specified identity store.
func (c *Client) CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) {
	if params == nil {
		params = &CreateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUser", params, optFns, c.addOperationCreateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserInput struct {

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// A list of Address objects containing addresses associated with the user.
	Addresses []types.Address

	// A string containing the user's name. This value is typically formatted for
	// display when the user is referenced. For example, "John Doe."
	DisplayName *string

	// A list of Email objects containing email addresses associated with the user.
	Emails []types.Email

	// A string containing the user's geographical region or location.
	Locale *string

	// An object containing the user's name.
	Name *types.Name

	// A string containing an alternate name for the user.
	NickName *string

	// A list of PhoneNumber objects containing phone numbers associated with the user.
	PhoneNumbers []types.PhoneNumber

	// A string containing the preferred language of the user. For example, "American
	// English" or "en-us."
	PreferredLanguage *string

	// A string containing a URL that may be associated with the user.
	ProfileUrl *string

	// A string containing the user's time zone.
	Timezone *string

	// A string containing the user's title. Possible values are left unspecified given
	// that they depend on each customer's specific needs.
	Title *string

	// A unique string used to identify the user. The length limit is 128 characters.
	// This value can consist of letters, accented characters, symbols, numbers, and
	// punctuation. This value is specified at the time the user is created and stored
	// as an attribute of the user object in the identity store.
	UserName *string

	// A string indicating the user's type. Possible values depend on each customer's
	// specific needs, so they are left unspecified.
	UserType *string

	noSmithyDocumentSerde
}

type CreateUserOutput struct {

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// The identifier of the newly created user in the identity store.
	//
	// This member is required.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUser{}, middleware.After)
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
	if err = addOpCreateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "identitystore",
		OperationName: "CreateUser",
	}
}
