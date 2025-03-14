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

// Gets the content of the network service descriptor. A network service descriptor
// is a .yaml file in a network package that uses the TOSCA standard to describe
// the network functions you want to deploy and the Amazon Web Services
// infrastructure you want to deploy the network functions on.
func (c *Client) GetSolNetworkPackageDescriptor(ctx context.Context, params *GetSolNetworkPackageDescriptorInput, optFns ...func(*Options)) (*GetSolNetworkPackageDescriptorOutput, error) {
	if params == nil {
		params = &GetSolNetworkPackageDescriptorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSolNetworkPackageDescriptor", params, optFns, c.addOperationGetSolNetworkPackageDescriptorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSolNetworkPackageDescriptorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSolNetworkPackageDescriptorInput struct {

	// ID of the network service descriptor in the network package.
	//
	// This member is required.
	NsdInfoId *string

	noSmithyDocumentSerde
}

type GetSolNetworkPackageDescriptorOutput struct {

	// Indicates the media type of the resource.
	ContentType types.DescriptorContentType

	// Contents of the network service descriptor in the network package.
	Nsd []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSolNetworkPackageDescriptorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSolNetworkPackageDescriptor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSolNetworkPackageDescriptor{}, middleware.After)
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
	if err = addOpGetSolNetworkPackageDescriptorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSolNetworkPackageDescriptor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSolNetworkPackageDescriptor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "GetSolNetworkPackageDescriptor",
	}
}
