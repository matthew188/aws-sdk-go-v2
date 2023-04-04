// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new domain association for an Amplify app. This action associates a
// custom domain with the Amplify app
func (c *Client) CreateDomainAssociation(ctx context.Context, params *CreateDomainAssociationInput, optFns ...func(*Options)) (*CreateDomainAssociationOutput, error) {
	if params == nil {
		params = &CreateDomainAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainAssociation", params, optFns, c.addOperationCreateDomainAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the create domain association request.
type CreateDomainAssociationInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The domain name for the domain association.
	//
	// This member is required.
	DomainName *string

	// The setting for the subdomain.
	//
	// This member is required.
	SubDomainSettings []types.SubDomainSetting

	// Sets the branch patterns for automatic subdomain creation.
	AutoSubDomainCreationPatterns []string

	// The required AWS Identity and Access Management (IAM) service role for the
	// Amazon Resource Name (ARN) for automatically creating subdomains.
	AutoSubDomainIAMRole *string

	// Enables the automated creation of subdomains for branches.
	EnableAutoSubDomain *bool

	noSmithyDocumentSerde
}

// The result structure for the create domain association request.
type CreateDomainAssociationOutput struct {

	// Describes the structure of a domain association, which associates a custom
	// domain with an Amplify app.
	//
	// This member is required.
	DomainAssociation *types.DomainAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomainAssociation{}, middleware.After)
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
	if err = addOpCreateDomainAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDomainAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "CreateDomainAssociation",
	}
}
