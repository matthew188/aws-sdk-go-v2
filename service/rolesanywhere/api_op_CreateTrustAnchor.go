// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rolesanywhere/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a trust anchor. You establish trust between IAM Roles Anywhere and your
// certificate authority (CA) by configuring a trust anchor. A Trust Anchor is
// defined either as a reference to a AWS Certificate Manager Private Certificate
// Authority (ACM PCA), or by uploading a Certificate Authority (CA) certificate.
// Your AWS workloads can authenticate with the trust anchor using certificates
// issued by the trusted Certificate Authority (CA) in exchange for temporary AWS
// credentials. Required permissions: rolesanywhere:CreateTrustAnchor.
func (c *Client) CreateTrustAnchor(ctx context.Context, params *CreateTrustAnchorInput, optFns ...func(*Options)) (*CreateTrustAnchorOutput, error) {
	if params == nil {
		params = &CreateTrustAnchorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrustAnchor", params, optFns, c.addOperationCreateTrustAnchorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTrustAnchorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrustAnchorInput struct {

	// The name of the trust anchor.
	//
	// This member is required.
	Name *string

	// The trust anchor type and its related certificate data.
	//
	// This member is required.
	Source *types.Source

	// Specifies whether the trust anchor is enabled.
	Enabled *bool

	// The tags to attach to the trust anchor.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateTrustAnchorOutput struct {

	// The state of the trust anchor after a read or write operation.
	//
	// This member is required.
	TrustAnchor *types.TrustAnchorDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTrustAnchorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTrustAnchor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTrustAnchor{}, middleware.After)
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
	if err = addOpCreateTrustAnchorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrustAnchor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrustAnchor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rolesanywhere",
		OperationName: "CreateTrustAnchor",
	}
}
