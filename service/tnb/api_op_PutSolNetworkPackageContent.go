// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Uploads the contents of a network package. A network package is a .zip file in
// CSAR (Cloud Service Archive) format defines the function packages you want to
// deploy and the Amazon Web Services infrastructure you want to deploy them on.
func (c *Client) PutSolNetworkPackageContent(ctx context.Context, params *PutSolNetworkPackageContentInput, optFns ...func(*Options)) (*PutSolNetworkPackageContentOutput, error) {
	if params == nil {
		params = &PutSolNetworkPackageContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSolNetworkPackageContent", params, optFns, c.addOperationPutSolNetworkPackageContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSolNetworkPackageContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSolNetworkPackageContentInput struct {

	// Network package file.
	//
	// This member is required.
	File []byte

	// Network service descriptor info ID.
	//
	// This member is required.
	NsdInfoId *string

	// Network package content type.
	ContentType types.PackageContentType

	noSmithyDocumentSerde
}

type PutSolNetworkPackageContentOutput struct {

	// Network package ARN.
	//
	// This member is required.
	Arn *string

	// Network package ID.
	//
	// This member is required.
	Id *string

	// Network package metadata.
	//
	// This member is required.
	Metadata *types.PutSolNetworkPackageContentMetadata

	// Network service descriptor ID.
	//
	// This member is required.
	NsdId *string

	// Network service descriptor name.
	//
	// This member is required.
	NsdName *string

	// Network service descriptor version.
	//
	// This member is required.
	NsdVersion *string

	// Function package IDs.
	//
	// This member is required.
	VnfPkgIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSolNetworkPackageContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSolNetworkPackageContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSolNetworkPackageContent{}, middleware.After)
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
	if err = addOpPutSolNetworkPackageContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSolNetworkPackageContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSolNetworkPackageContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "PutSolNetworkPackageContent",
	}
}
