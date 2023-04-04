// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociate a custom domain name from an App Runner service. Certificates
// tracking domain validity are associated with a custom domain and are stored in
// AWS Certificate Manager (ACM)
// (https://docs.aws.amazon.com/acm/latest/userguide). These certificates aren't
// deleted as part of this action. App Runner delays certificate deletion for 30
// days after a domain is disassociated from your service.
func (c *Client) DisassociateCustomDomain(ctx context.Context, params *DisassociateCustomDomainInput, optFns ...func(*Options)) (*DisassociateCustomDomainOutput, error) {
	if params == nil {
		params = &DisassociateCustomDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateCustomDomain", params, optFns, c.addOperationDisassociateCustomDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateCustomDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateCustomDomainInput struct {

	// The domain name that you want to disassociate from the App Runner service.
	//
	// This member is required.
	DomainName *string

	// The Amazon Resource Name (ARN) of the App Runner service that you want to
	// disassociate a custom domain name from.
	//
	// This member is required.
	ServiceArn *string

	noSmithyDocumentSerde
}

type DisassociateCustomDomainOutput struct {

	// A description of the domain name that's being disassociated.
	//
	// This member is required.
	CustomDomain *types.CustomDomain

	// The App Runner subdomain of the App Runner service. The disassociated custom
	// domain name was mapped to this target name.
	//
	// This member is required.
	DNSTarget *string

	// The Amazon Resource Name (ARN) of the App Runner service that a custom domain
	// name is disassociated from.
	//
	// This member is required.
	ServiceArn *string

	// DNS Target records for the custom domains of this Amazon VPC.
	//
	// This member is required.
	VpcDNSTargets []types.VpcDNSTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateCustomDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDisassociateCustomDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDisassociateCustomDomain{}, middleware.After)
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
	if err = addOpDisassociateCustomDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateCustomDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateCustomDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "DisassociateCustomDomain",
	}
}
