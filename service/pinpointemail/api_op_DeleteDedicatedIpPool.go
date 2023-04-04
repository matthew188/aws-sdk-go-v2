// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a dedicated IP pool.
func (c *Client) DeleteDedicatedIpPool(ctx context.Context, params *DeleteDedicatedIpPoolInput, optFns ...func(*Options)) (*DeleteDedicatedIpPoolOutput, error) {
	if params == nil {
		params = &DeleteDedicatedIpPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDedicatedIpPool", params, optFns, c.addOperationDeleteDedicatedIpPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDedicatedIpPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a dedicated IP pool.
type DeleteDedicatedIpPoolInput struct {

	// The name of the dedicated IP pool that you want to delete.
	//
	// This member is required.
	PoolName *string

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type DeleteDedicatedIpPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDedicatedIpPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDedicatedIpPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDedicatedIpPool{}, middleware.After)
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
	if err = addOpDeleteDedicatedIpPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDedicatedIpPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDedicatedIpPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteDedicatedIpPool",
	}
}
