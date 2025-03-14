// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Revokes a certificate that was issued inside Amazon Web Services Private CA. If
// you enable a certificate revocation list (CRL) when you create or update your
// private CA, information about the revoked certificates will be included in the
// CRL. Amazon Web Services Private CA writes the CRL to an S3 bucket that you
// specify. A CRL is typically updated approximately 30 minutes after a certificate
// is revoked. If for any reason the CRL update fails, Amazon Web Services Private
// CA attempts makes further attempts every 15 minutes. With Amazon CloudWatch, you
// can create alarms for the metrics CRLGenerated and MisconfiguredCRLBucket. For
// more information, see Supported CloudWatch Metrics
// (https://docs.aws.amazon.com/privateca/latest/userguide/PcaCloudWatch.html).
// Both Amazon Web Services Private CA and the IAM principal must have permission
// to write to the S3 bucket that you specify. If the IAM principal making the call
// does not have permission to write to the bucket, then an exception is thrown.
// For more information, see Access policies for CRLs in Amazon S3
// (https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies).
// Amazon Web Services Private CA also writes revocation information to the audit
// report. For more information, see CreateCertificateAuthorityAuditReport
// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthorityAuditReport.html).
// You cannot revoke a root CA self-signed certificate.
func (c *Client) RevokeCertificate(ctx context.Context, params *RevokeCertificateInput, optFns ...func(*Options)) (*RevokeCertificateOutput, error) {
	if params == nil {
		params = &RevokeCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeCertificate", params, optFns, c.addOperationRevokeCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeCertificateInput struct {

	// Amazon Resource Name (ARN) of the private CA that issued the certificate to be
	// revoked. This must be of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// This member is required.
	CertificateAuthorityArn *string

	// Serial number of the certificate to be revoked. This must be in hexadecimal
	// format. You can retrieve the serial number by calling GetCertificate
	// (https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificate.html)
	// with the Amazon Resource Name (ARN) of the certificate you want and the ARN of
	// your private CA. The GetCertificate action retrieves the certificate in the PEM
	// format. You can use the following OpenSSL command to list the certificate in
	// text format and copy the hexadecimal serial number. openssl x509 -in file_path
	// -text -noout You can also copy the serial number from the console or use the
	// DescribeCertificate
	// (https://docs.aws.amazon.com/acm/latest/APIReference/API_DescribeCertificate.html)
	// action in the Certificate Manager API Reference.
	//
	// This member is required.
	CertificateSerial *string

	// Specifies why you revoked the certificate.
	//
	// This member is required.
	RevocationReason types.RevocationReason

	noSmithyDocumentSerde
}

type RevokeCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRevokeCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRevokeCertificate{}, middleware.After)
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
	if err = addOpRevokeCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "RevokeCertificate",
	}
}
