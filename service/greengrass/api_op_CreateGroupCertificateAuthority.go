// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a CA for the group. If a CA already exists, it will rotate the existing
// CA.
func (c *Client) CreateGroupCertificateAuthority(ctx context.Context, params *CreateGroupCertificateAuthorityInput, optFns ...func(*Options)) (*CreateGroupCertificateAuthorityOutput, error) {
	if params == nil {
		params = &CreateGroupCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGroupCertificateAuthority", params, optFns, c.addOperationCreateGroupCertificateAuthorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGroupCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGroupCertificateAuthorityInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	noSmithyDocumentSerde
}

type CreateGroupCertificateAuthorityOutput struct {

	// The ARN of the group certificate authority.
	GroupCertificateAuthorityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGroupCertificateAuthorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGroupCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGroupCertificateAuthority{}, middleware.After)
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
	if err = addOpCreateGroupCertificateAuthorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGroupCertificateAuthority(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGroupCertificateAuthority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateGroupCertificateAuthority",
	}
}
