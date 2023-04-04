// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the working storage of a gateway. This operation is
// only supported in the stored volumes gateway type. This operation is deprecated
// in cached volumes API version (20120630). Use DescribeUploadBuffer instead.
// Working storage is also referred to as upload buffer. You can also use the
// DescribeUploadBuffer operation to add upload buffer to a stored volume gateway.
// The response includes disk IDs that are configured as working storage, and it
// includes the amount of working storage allocated and used.
func (c *Client) DescribeWorkingStorage(ctx context.Context, params *DescribeWorkingStorageInput, optFns ...func(*Options)) (*DescribeWorkingStorageOutput, error) {
	if params == nil {
		params = &DescribeWorkingStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkingStorage", params, optFns, c.addOperationDescribeWorkingStorageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkingStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway.
type DescribeWorkingStorageInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
type DescribeWorkingStorageOutput struct {

	// An array of the gateway's local disk IDs that are configured as working storage.
	// Each local disk ID is specified as a string (minimum length of 1 and maximum
	// length of 300). If no local disks are configured as working storage, then the
	// DiskIds array is empty.
	DiskIds []string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// The total working storage in bytes allocated for the gateway. If no working
	// storage is configured for the gateway, this field returns 0.
	WorkingStorageAllocatedInBytes int64

	// The total working storage in bytes in use by the gateway. If no working storage
	// is configured for the gateway, this field returns 0.
	WorkingStorageUsedInBytes int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWorkingStorageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkingStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkingStorage{}, middleware.After)
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
	if err = addOpDescribeWorkingStorageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkingStorage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWorkingStorage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeWorkingStorage",
	}
}
