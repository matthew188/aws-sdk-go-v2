// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified storage volume that you previously created using the
// CreateCachediSCSIVolume or CreateStorediSCSIVolume API. This operation is only
// supported in the cached volume and stored volume types. For stored volume
// gateways, the local disk that was configured as the storage volume is not
// deleted. You can reuse the local disk to create another storage volume. Before
// you delete a volume, make sure there are no iSCSI connections to the volume you
// are deleting. You should also make sure there is no snapshot in progress. You
// can use the Amazon Elastic Compute Cloud (Amazon EC2) API to query snapshots on
// the volume you are deleting and check the snapshot status. For more information,
// go to DescribeSnapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html)
// in the Amazon Elastic Compute Cloud API Reference. In the request, you must
// provide the Amazon Resource Name (ARN) of the storage volume you want to delete.
func (c *Client) DeleteVolume(ctx context.Context, params *DeleteVolumeInput, optFns ...func(*Options)) (*DeleteVolumeOutput, error) {
	if params == nil {
		params = &DeleteVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVolume", params, optFns, c.addOperationDeleteVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the DeleteVolumeInput$VolumeARN to delete.
type DeleteVolumeInput struct {

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation to
	// return a list of gateway volumes.
	//
	// This member is required.
	VolumeARN *string

	noSmithyDocumentSerde
}

// A JSON object containing the Amazon Resource Name (ARN) of the storage volume
// that was deleted.
type DeleteVolumeOutput struct {

	// The Amazon Resource Name (ARN) of the storage volume that was deleted. It is the
	// same ARN you provided in the request.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteVolume{}, middleware.After)
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
	if err = addOpDeleteVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVolume(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteVolume",
	}
}
