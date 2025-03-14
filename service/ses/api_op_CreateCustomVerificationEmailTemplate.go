// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new custom verification email template. For more information about
// custom verification email templates, see Using Custom Verification Email
// Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) CreateCustomVerificationEmailTemplate(ctx context.Context, params *CreateCustomVerificationEmailTemplateInput, optFns ...func(*Options)) (*CreateCustomVerificationEmailTemplateOutput, error) {
	if params == nil {
		params = &CreateCustomVerificationEmailTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomVerificationEmailTemplate", params, optFns, c.addOperationCreateCustomVerificationEmailTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomVerificationEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create a custom verification email template.
type CreateCustomVerificationEmailTemplateInput struct {

	// The URL that the recipient of the verification email is sent to if his or her
	// address is not successfully verified.
	//
	// This member is required.
	FailureRedirectionURL *string

	// The email address that the custom verification email is sent from.
	//
	// This member is required.
	FromEmailAddress *string

	// The URL that the recipient of the verification email is sent to if his or her
	// address is successfully verified.
	//
	// This member is required.
	SuccessRedirectionURL *string

	// The content of the custom verification email. The total size of the email must
	// be less than 10 MB. The message body may contain HTML, with some limitations.
	// For more information, see Custom Verification Email Frequently Asked Questions
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html#custom-verification-emails-faq)
	// in the Amazon SES Developer Guide.
	//
	// This member is required.
	TemplateContent *string

	// The name of the custom verification email template.
	//
	// This member is required.
	TemplateName *string

	// The subject line of the custom verification email.
	//
	// This member is required.
	TemplateSubject *string

	noSmithyDocumentSerde
}

type CreateCustomVerificationEmailTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomVerificationEmailTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateCustomVerificationEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateCustomVerificationEmailTemplate{}, middleware.After)
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
	if err = addOpCreateCustomVerificationEmailTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomVerificationEmailTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCustomVerificationEmailTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateCustomVerificationEmailTemplate",
	}
}
