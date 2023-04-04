// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports a private certificate issued by a private certificate authority (CA) for
// use anywhere. The exported file contains the certificate, the certificate chain,
// and the encrypted private 2048-bit RSA key associated with the public key that
// is embedded in the certificate. For security, you must assign a passphrase for
// the private key when exporting it. For information about exporting and
// formatting a certificate using the ACM console or CLI, see Export a Private
// Certificate
// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-export-private.html).
func (c *Client) ExportCertificate(ctx context.Context, params *ExportCertificateInput, optFns ...func(*Options)) (*ExportCertificateOutput, error) {
	if params == nil {
		params = &ExportCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportCertificate", params, optFns, c.addOperationExportCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportCertificateInput struct {

	// An Amazon Resource Name (ARN) of the issued certificate. This must be of the
	// form:
	// arn:aws:acm:region:account:certificate/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateArn *string

	// Passphrase to associate with the encrypted exported private key. When creating
	// your passphrase, you can use any ASCII character except #, $, or %. If you want
	// to later decrypt the private key, you must have the passphrase. You can use the
	// following OpenSSL command to decrypt a private key. After entering the command,
	// you are prompted for the passphrase. openssl rsa -in encrypted_key.pem -out
	// decrypted_key.pem
	//
	// This member is required.
	Passphrase []byte

	noSmithyDocumentSerde
}

type ExportCertificateOutput struct {

	// The base64 PEM-encoded certificate.
	Certificate *string

	// The base64 PEM-encoded certificate chain. This does not include the certificate
	// that you are exporting.
	CertificateChain *string

	// The encrypted private key associated with the public key in the certificate. The
	// key is output in PKCS #8 format and is base64 PEM-encoded.
	PrivateKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExportCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExportCertificate{}, middleware.After)
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
	if err = addOpExportCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm",
		OperationName: "ExportCertificate",
	}
}
