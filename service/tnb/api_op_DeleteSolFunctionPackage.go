// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a function package. A function package is a .zip file in CSAR (Cloud
// Service Archive) format that contains a network function (an ETSI standard
// telecommunication application) and function package descriptor that uses the
// TOSCA standard to describe how the network functions should run on your network.
// To delete a function package, the package must be in a disabled state. To
// disable a function package, see UpdateSolFunctionPackage
// (https://docs.aws.amazon.com/tnb/latest/APIReference/API_UpdateSolFunctionPackage.html).
func (c *Client) DeleteSolFunctionPackage(ctx context.Context, params *DeleteSolFunctionPackageInput, optFns ...func(*Options)) (*DeleteSolFunctionPackageOutput, error) {
	if params == nil {
		params = &DeleteSolFunctionPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSolFunctionPackage", params, optFns, c.addOperationDeleteSolFunctionPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSolFunctionPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSolFunctionPackageInput struct {

	// ID of the function package.
	//
	// This member is required.
	VnfPkgId *string

	noSmithyDocumentSerde
}

type DeleteSolFunctionPackageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSolFunctionPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSolFunctionPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSolFunctionPackage{}, middleware.After)
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
	if err = addOpDeleteSolFunctionPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSolFunctionPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSolFunctionPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "DeleteSolFunctionPackage",
	}
}
