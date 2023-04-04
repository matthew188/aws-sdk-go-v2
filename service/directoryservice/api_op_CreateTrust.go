// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Directory Service for Microsoft Active Directory allows you to configure trust
// relationships. For example, you can establish a trust between your Managed
// Microsoft AD directory, and your existing self-managed Microsoft Active
// Directory. This would allow you to provide users and groups access to resources
// in either domain, with a single set of credentials. This action initiates the
// creation of the Amazon Web Services side of a trust relationship between an
// Managed Microsoft AD directory and an external domain. You can create either a
// forest trust or an external trust.
func (c *Client) CreateTrust(ctx context.Context, params *CreateTrustInput, optFns ...func(*Options)) (*CreateTrustOutput, error) {
	if params == nil {
		params = &CreateTrustInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrust", params, optFns, c.addOperationCreateTrustMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrustOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Directory Service for Microsoft Active Directory allows you to configure trust
// relationships. For example, you can establish a trust between your Managed
// Microsoft AD directory, and your existing self-managed Microsoft Active
// Directory. This would allow you to provide users and groups access to resources
// in either domain, with a single set of credentials. This action initiates the
// creation of the Amazon Web Services side of a trust relationship between an
// Managed Microsoft AD directory and an external domain.
type CreateTrustInput struct {

	// The Directory ID of the Managed Microsoft AD directory for which to establish
	// the trust relationship.
	//
	// This member is required.
	DirectoryId *string

	// The Fully Qualified Domain Name (FQDN) of the external domain for which to
	// create the trust relationship.
	//
	// This member is required.
	RemoteDomainName *string

	// The direction of the trust relationship.
	//
	// This member is required.
	TrustDirection types.TrustDirection

	// The trust password. The must be the same password that was used when creating
	// the trust relationship on the external domain.
	//
	// This member is required.
	TrustPassword *string

	// The IP addresses of the remote DNS server associated with RemoteDomainName.
	ConditionalForwarderIpAddrs []string

	// Optional parameter to enable selective authentication for the trust.
	SelectiveAuth types.SelectiveAuth

	// The trust relationship type. Forest is the default.
	TrustType types.TrustType

	noSmithyDocumentSerde
}

// The result of a CreateTrust request.
type CreateTrustOutput struct {

	// A unique identifier for the trust relationship that was created.
	TrustId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrustMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrust{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrust{}, middleware.After)
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
	if err = addOpCreateTrustValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrust(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrust(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CreateTrust",
	}
}
