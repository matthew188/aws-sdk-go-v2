// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon QuickSight account, or subscribes to Amazon QuickSight Q. The
// Amazon Web Services Region for the account is derived from what is configured in
// the CLI or SDK. This operation isn't supported in the US East (Ohio) Region,
// South America (Sao Paulo) Region, or Asia Pacific (Singapore) Region. Before you
// use this operation, make sure that you can connect to an existing Amazon Web
// Services account. If you don't have an Amazon Web Services account, see Sign up
// for Amazon Web Services
// (https://docs.aws.amazon.com/quicksight/latest/user/setting-up-aws-sign-up.html)
// in the Amazon QuickSight User Guide. The person who signs up for Amazon
// QuickSight needs to have the correct Identity and Access Management (IAM)
// permissions. For more information, see IAM Policy Examples for Amazon QuickSight
// (https://docs.aws.amazon.com/quicksight/latest/user/iam-policy-examples.html) in
// the Amazon QuickSight User Guide. If your IAM policy includes both the Subscribe
// and CreateAccountSubscription actions, make sure that both actions are set to
// Allow. If either action is set to Deny, the Deny action prevails and your API
// call fails. You can't pass an existing IAM role to access other Amazon Web
// Services services using this API operation. To pass your existing IAM role to
// Amazon QuickSight, see Passing IAM roles to Amazon QuickSight
// (https://docs.aws.amazon.com/quicksight/latest/user/security_iam_service-with-iam.html#security-create-iam-role)
// in the Amazon QuickSight User Guide. You can't set default resource access on
// the new account from the Amazon QuickSight API. Instead, add default resource
// access from the Amazon QuickSight console. For more information about setting
// default resource access to Amazon Web Services services, see Setting default
// resource access to Amazon Web Services services
// (https://docs.aws.amazon.com/quicksight/latest/user/scoping-policies-defaults.html)
// in the Amazon QuickSight User Guide.
func (c *Client) CreateAccountSubscription(ctx context.Context, params *CreateAccountSubscriptionInput, optFns ...func(*Options)) (*CreateAccountSubscriptionOutput, error) {
	if params == nil {
		params = &CreateAccountSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccountSubscription", params, optFns, c.addOperationCreateAccountSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccountSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccountSubscriptionInput struct {

	// The name of your Amazon QuickSight account. This name is unique over all of
	// Amazon Web Services, and it appears only when users sign in. You can't change
	// AccountName value after the Amazon QuickSight account is created.
	//
	// This member is required.
	AccountName *string

	// The method that you want to use to authenticate your Amazon QuickSight account.
	// Currently, the valid values for this parameter are IAM_AND_QUICKSIGHT, IAM_ONLY,
	// and ACTIVE_DIRECTORY. If you choose ACTIVE_DIRECTORY, provide an
	// ActiveDirectoryName and an AdminGroup associated with your Active Directory.
	//
	// This member is required.
	AuthenticationMethod types.AuthenticationMethodOption

	// The Amazon Web Services account ID of the account that you're using to create
	// your Amazon QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The edition of Amazon QuickSight that you want your account to have. Currently,
	// you can choose from ENTERPRISE or ENTERPRISE_AND_Q. If you choose
	// ENTERPRISE_AND_Q, the following parameters are required:
	//
	// * FirstName
	//
	// *
	// LastName
	//
	// * EmailAddress
	//
	// * ContactNumber
	//
	// This member is required.
	Edition types.Edition

	// The email address that you want Amazon QuickSight to send notifications to
	// regarding your Amazon QuickSight account or Amazon QuickSight subscription.
	//
	// This member is required.
	NotificationEmail *string

	// The name of your Active Directory. This field is required if ACTIVE_DIRECTORY is
	// the selected authentication method of the new Amazon QuickSight account.
	ActiveDirectoryName *string

	// The admin group associated with your Active Directory. This field is required if
	// ACTIVE_DIRECTORY is the selected authentication method of the new Amazon
	// QuickSight account. For more information about using Active Directory in Amazon
	// QuickSight, see Using Active Directory with Amazon QuickSight Enterprise Edition
	// (https://docs.aws.amazon.com/quicksight/latest/user/aws-directory-service.html)
	// in the Amazon QuickSight User Guide.
	AdminGroup []string

	// The author group associated with your Active Directory. For more information
	// about using Active Directory in Amazon QuickSight, see Using Active Directory
	// with Amazon QuickSight Enterprise Edition
	// (https://docs.aws.amazon.com/quicksight/latest/user/aws-directory-service.html)
	// in the Amazon QuickSight User Guide.
	AuthorGroup []string

	// A 10-digit phone number for the author of the Amazon QuickSight account to use
	// for future communications. This field is required if ENTERPPRISE_AND_Q is the
	// selected edition of the new Amazon QuickSight account.
	ContactNumber *string

	// The ID of the Active Directory that is associated with your Amazon QuickSight
	// account.
	DirectoryId *string

	// The email address of the author of the Amazon QuickSight account to use for
	// future communications. This field is required if ENTERPPRISE_AND_Q is the
	// selected edition of the new Amazon QuickSight account.
	EmailAddress *string

	// The first name of the author of the Amazon QuickSight account to use for future
	// communications. This field is required if ENTERPPRISE_AND_Q is the selected
	// edition of the new Amazon QuickSight account.
	FirstName *string

	// The last name of the author of the Amazon QuickSight account to use for future
	// communications. This field is required if ENTERPPRISE_AND_Q is the selected
	// edition of the new Amazon QuickSight account.
	LastName *string

	// The reader group associated with your Active Direcrtory. For more information
	// about using Active Directory in Amazon QuickSight, see Using Active Directory
	// with Amazon QuickSight Enterprise Edition
	// (https://docs.aws.amazon.com/quicksight/latest/user/aws-directory-service.html)
	// in the Amazon QuickSight User Guide.
	ReaderGroup []string

	// The realm of the Active Directory that is associated with your Amazon QuickSight
	// account. This field is required if ACTIVE_DIRECTORY is the selected
	// authentication method of the new Amazon QuickSight account.
	Realm *string

	noSmithyDocumentSerde
}

type CreateAccountSubscriptionOutput struct {

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// A SignupResponse object that returns information about a newly created Amazon
	// QuickSight account.
	SignupResponse *types.SignupResponse

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccountSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccountSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccountSubscription{}, middleware.After)
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
	if err = addOpCreateAccountSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccountSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAccountSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateAccountSubscription",
	}
}
