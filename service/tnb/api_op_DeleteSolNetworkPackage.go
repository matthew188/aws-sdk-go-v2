// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes network package. A network package is a .zip file in CSAR (Cloud Service
// Archive) format defines the function packages you want to deploy and the Amazon
// Web Services infrastructure you want to deploy them on. To delete a network
// package, the package must be in a disable state. To disable a network package,
// see UpdateSolNetworkPackage
// (https://docs.aws.amazon.com/tnb/latest/APIReference/API_UpdateSolNetworkPackage.html).
func (c *Client) DeleteSolNetworkPackage(ctx context.Context, params *DeleteSolNetworkPackageInput, optFns ...func(*Options)) (*DeleteSolNetworkPackageOutput, error) {
	if params == nil {
		params = &DeleteSolNetworkPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSolNetworkPackage", params, optFns, c.addOperationDeleteSolNetworkPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSolNetworkPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSolNetworkPackageInput struct {

	// ID of the network service descriptor in the network package.
	//
	// This member is required.
	NsdInfoId *string

	noSmithyDocumentSerde
}

type DeleteSolNetworkPackageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSolNetworkPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSolNetworkPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSolNetworkPackage{}, middleware.After)
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
	if err = addOpDeleteSolNetworkPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSolNetworkPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSolNetworkPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "DeleteSolNetworkPackage",
	}
}
