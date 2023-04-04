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

// Delete an App Runner VPC Ingress Connection resource that's associated with an
// App Runner service. The VPC Ingress Connection must be in one of the following
// states to be deleted:
//
// * AVAILABLE
//
// * FAILED_CREATION
//
// * FAILED_UPDATE
//
// *
// FAILED_DELETION
func (c *Client) DeleteVpcIngressConnection(ctx context.Context, params *DeleteVpcIngressConnectionInput, optFns ...func(*Options)) (*DeleteVpcIngressConnectionOutput, error) {
	if params == nil {
		params = &DeleteVpcIngressConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVpcIngressConnection", params, optFns, c.addOperationDeleteVpcIngressConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVpcIngressConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVpcIngressConnectionInput struct {

	// The Amazon Resource Name (ARN) of the App Runner VPC Ingress Connection that you
	// want to delete.
	//
	// This member is required.
	VpcIngressConnectionArn *string

	noSmithyDocumentSerde
}

type DeleteVpcIngressConnectionOutput struct {

	// A description of the App Runner VPC Ingress Connection that this request just
	// deleted.
	//
	// This member is required.
	VpcIngressConnection *types.VpcIngressConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVpcIngressConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteVpcIngressConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteVpcIngressConnection{}, middleware.After)
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
	if err = addOpDeleteVpcIngressConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVpcIngressConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVpcIngressConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "DeleteVpcIngressConnection",
	}
}
