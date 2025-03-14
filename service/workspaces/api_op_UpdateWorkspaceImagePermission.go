// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Shares or unshares an image with one account in the same Amazon Web Services
// Region by specifying whether that account has permission to copy the image. If
// the copy image permission is granted, the image is shared with that account. If
// the copy image permission is revoked, the image is unshared with the account.
// After an image has been shared, the recipient account can copy the image to
// other Regions as needed. In the China (Ningxia) Region, you can copy images only
// within the same Region. In Amazon Web Services GovCloud (US), to copy images to
// and from other Regions, contact Amazon Web Services Support. For more
// information about sharing images, see  Share or Unshare a Custom WorkSpaces
// Image
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/share-custom-image.html).
//
// *
// To delete an image that has been shared, you must unshare the image before you
// delete it.
//
// * Sharing Bring Your Own License (BYOL) images across Amazon Web
// Services accounts isn't supported at this time in Amazon Web Services GovCloud
// (US). To share BYOL images across accounts in Amazon Web Services GovCloud (US),
// contact Amazon Web Services Support.
func (c *Client) UpdateWorkspaceImagePermission(ctx context.Context, params *UpdateWorkspaceImagePermissionInput, optFns ...func(*Options)) (*UpdateWorkspaceImagePermissionOutput, error) {
	if params == nil {
		params = &UpdateWorkspaceImagePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkspaceImagePermission", params, optFns, c.addOperationUpdateWorkspaceImagePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkspaceImagePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkspaceImagePermissionInput struct {

	// The permission to copy the image. This permission can be revoked only after an
	// image has been shared.
	//
	// This member is required.
	AllowCopyImage *bool

	// The identifier of the image.
	//
	// This member is required.
	ImageId *string

	// The identifier of the Amazon Web Services account to share or unshare the image
	// with. Before sharing the image, confirm that you are sharing to the correct
	// Amazon Web Services account ID.
	//
	// This member is required.
	SharedAccountId *string

	noSmithyDocumentSerde
}

type UpdateWorkspaceImagePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkspaceImagePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWorkspaceImagePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWorkspaceImagePermission{}, middleware.After)
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
	if err = addOpUpdateWorkspaceImagePermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkspaceImagePermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkspaceImagePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "UpdateWorkspaceImagePermission",
	}
}
