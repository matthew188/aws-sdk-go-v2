// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to configure or change the DKIM authentication settings for an email domain
// identity. You can use this operation to do any of the following:
//
// * Update the
// signing attributes for an identity that uses Bring Your Own DKIM (BYODKIM).
//
// *
// Update the key length that should be used for Easy DKIM.
//
// * Change from using no
// DKIM authentication to using Easy DKIM.
//
// * Change from using no DKIM
// authentication to using BYODKIM.
//
// * Change from using Easy DKIM to using
// BYODKIM.
//
// * Change from using BYODKIM to using Easy DKIM.
func (c *Client) PutEmailIdentityDkimSigningAttributes(ctx context.Context, params *PutEmailIdentityDkimSigningAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityDkimSigningAttributesOutput, error) {
	if params == nil {
		params = &PutEmailIdentityDkimSigningAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEmailIdentityDkimSigningAttributes", params, optFns, c.addOperationPutEmailIdentityDkimSigningAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEmailIdentityDkimSigningAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change the DKIM attributes for an email identity.
type PutEmailIdentityDkimSigningAttributesInput struct {

	// The email identity.
	//
	// This member is required.
	EmailIdentity *string

	// The method to use to configure DKIM for the identity. There are the following
	// possible values:
	//
	// * AWS_SES – Configure DKIM for the identity by using Easy DKIM
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
	//
	// *
	// EXTERNAL – Configure DKIM for the identity by using Bring Your Own DKIM
	// (BYODKIM).
	//
	// This member is required.
	SigningAttributesOrigin types.DkimSigningAttributesOrigin

	// An object that contains information about the private key and selector that you
	// want to use to configure DKIM for the identity for Bring Your Own DKIM (BYODKIM)
	// for the identity, or, configures the key length to be used for Easy DKIM
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html).
	SigningAttributes *types.DkimSigningAttributes

	noSmithyDocumentSerde
}

// If the action is successful, the service sends back an HTTP 200 response. The
// following data is returned in JSON format by the service.
type PutEmailIdentityDkimSigningAttributesOutput struct {

	// The DKIM authentication status of the identity. Amazon SES determines the
	// authentication status by searching for specific records in the DNS configuration
	// for your domain. If you used Easy DKIM
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) to set up
	// DKIM authentication, Amazon SES tries to find three unique CNAME records in the
	// DNS configuration for your domain. If you provided a public key to perform DKIM
	// authentication, Amazon SES tries to find a TXT record that uses the selector
	// that you specified. The value of the TXT record must be a public key that's
	// paired with the private key that you specified in the process of creating the
	// identity. The status can be one of the following:
	//
	// * PENDING – The verification
	// process was initiated, but Amazon SES hasn't yet detected the DKIM records in
	// the DNS configuration for the domain.
	//
	// * SUCCESS – The verification process
	// completed successfully.
	//
	// * FAILED – The verification process failed. This
	// typically occurs when Amazon SES fails to find the DKIM records in the DNS
	// configuration of the domain.
	//
	// * TEMPORARY_FAILURE – A temporary issue is
	// preventing Amazon SES from determining the DKIM authentication status of the
	// domain.
	//
	// * NOT_STARTED – The DKIM verification process hasn't been initiated for
	// the domain.
	DkimStatus types.DkimStatus

	// If you used Easy DKIM
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html) to
	// configure DKIM authentication for the domain, then this object contains a set of
	// unique strings that you use to create a set of CNAME records that you add to the
	// DNS configuration for your domain. When Amazon SES detects these records in the
	// DNS configuration for your domain, the DKIM authentication process is complete.
	// If you configured DKIM authentication for the domain by providing your own
	// public-private key pair, then this object contains the selector that's
	// associated with your public key. Regardless of the DKIM authentication method
	// you use, Amazon SES searches for the appropriate records in the DNS
	// configuration of the domain for up to 72 hours.
	DkimTokens []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEmailIdentityDkimSigningAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEmailIdentityDkimSigningAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEmailIdentityDkimSigningAttributes{}, middleware.After)
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
	if err = addOpPutEmailIdentityDkimSigningAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEmailIdentityDkimSigningAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEmailIdentityDkimSigningAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutEmailIdentityDkimSigningAttributes",
	}
}
