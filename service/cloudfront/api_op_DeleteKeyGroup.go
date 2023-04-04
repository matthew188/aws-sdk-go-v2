// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a key group. You cannot delete a key group that is referenced in a cache
// behavior. First update your distributions to remove the key group from all cache
// behaviors, then delete the key group. To delete a key group, you must provide
// the key group's identifier and version. To get these values, use ListKeyGroups
// followed by GetKeyGroup or GetKeyGroupConfig.
func (c *Client) DeleteKeyGroup(ctx context.Context, params *DeleteKeyGroupInput, optFns ...func(*Options)) (*DeleteKeyGroupOutput, error) {
	if params == nil {
		params = &DeleteKeyGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteKeyGroup", params, optFns, c.addOperationDeleteKeyGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteKeyGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKeyGroupInput struct {

	// The identifier of the key group that you are deleting. To get the identifier,
	// use ListKeyGroups.
	//
	// This member is required.
	Id *string

	// The version of the key group that you are deleting. The version is the key
	// group's ETag value. To get the ETag, use GetKeyGroup or GetKeyGroupConfig.
	IfMatch *string

	noSmithyDocumentSerde
}

type DeleteKeyGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteKeyGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteKeyGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteKeyGroup{}, middleware.After)
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
	if err = addOpDeleteKeyGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKeyGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteKeyGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteKeyGroup",
	}
}
