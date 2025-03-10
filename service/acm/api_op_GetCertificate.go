// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves an Amazon-issued certificate and its certificate chain. The chain
// consists of the certificate of the issuing CA and the intermediate certificates
// of any other subordinate CAs. All of the certificates are base64 encoded. You
// can use OpenSSL (https://wiki.openssl.org/index.php/Command_Line_Utilities) to
// decode the certificates and inspect individual fields.
func (c *Client) GetCertificate(ctx context.Context, params *GetCertificateInput, optFns ...func(*Options)) (*GetCertificateOutput, error) {
	if params == nil {
		params = &GetCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCertificate", params, optFns, c.addOperationGetCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCertificateInput struct {

	// String that contains a certificate ARN in the following format:
	// arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	// For more information about ARNs, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// This member is required.
	CertificateArn *string

	noSmithyDocumentSerde
}

type GetCertificateOutput struct {

	// The ACM-issued certificate corresponding to the ARN specified as input.
	Certificate *string

	// Certificates forming the requested certificate's chain of trust. The chain
	// consists of the certificate of the issuing CA and the intermediate certificates
	// of any other subordinate CAs.
	CertificateChain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCertificate{}, middleware.After)
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
	if err = addOpGetCertificateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCertificate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCertificate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm",
		OperationName: "GetCertificate",
	}
}
