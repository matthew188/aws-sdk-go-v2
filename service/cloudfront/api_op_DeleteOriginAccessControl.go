// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a CloudFront origin access control. You cannot delete an origin access
// control if it's in use. First, update all distributions to remove the origin
// access control from all origins, then delete the origin access control.
func (c *Client) DeleteOriginAccessControl(ctx context.Context, params *DeleteOriginAccessControlInput, optFns ...func(*Options)) (*DeleteOriginAccessControlOutput, error) {
	if params == nil {
		params = &DeleteOriginAccessControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOriginAccessControl", params, optFns, c.addOperationDeleteOriginAccessControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOriginAccessControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOriginAccessControlInput struct {

	// The unique identifier of the origin access control that you are deleting.
	//
	// This member is required.
	Id *string

	// The current version (ETag value) of the origin access control that you are
	// deleting.
	IfMatch *string

	noSmithyDocumentSerde
}

type DeleteOriginAccessControlOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOriginAccessControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteOriginAccessControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteOriginAccessControl{}, middleware.After)
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
	if err = addOpDeleteOriginAccessControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOriginAccessControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOriginAccessControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteOriginAccessControl",
	}
}
