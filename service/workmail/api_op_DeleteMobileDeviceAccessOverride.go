// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the mobile device access override for the given WorkMail organization,
// user, and device. Deleting already deleted and non-existing overrides does not
// produce an error. In those cases, the service sends back an HTTP 200 response
// with an empty HTTP body.
func (c *Client) DeleteMobileDeviceAccessOverride(ctx context.Context, params *DeleteMobileDeviceAccessOverrideInput, optFns ...func(*Options)) (*DeleteMobileDeviceAccessOverrideOutput, error) {
	if params == nil {
		params = &DeleteMobileDeviceAccessOverrideInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMobileDeviceAccessOverride", params, optFns, c.addOperationDeleteMobileDeviceAccessOverrideMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMobileDeviceAccessOverrideOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMobileDeviceAccessOverrideInput struct {

	// The mobile device for which you delete the override. DeviceId is case
	// insensitive.
	//
	// This member is required.
	DeviceId *string

	// The WorkMail organization for which the access override will be deleted.
	//
	// This member is required.
	OrganizationId *string

	// The WorkMail user for which you want to delete the override. Accepts the
	// following types of user identities:
	//
	// * User ID:
	// 12345678-1234-1234-1234-123456789012 or
	// S-1-1-12-1234567890-123456789-123456789-1234
	//
	// * Email address:
	// user@domain.tld
	//
	// * User name: user
	//
	// This member is required.
	UserId *string

	noSmithyDocumentSerde
}

type DeleteMobileDeviceAccessOverrideOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMobileDeviceAccessOverrideMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMobileDeviceAccessOverride{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMobileDeviceAccessOverride{}, middleware.After)
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
	if err = addOpDeleteMobileDeviceAccessOverrideValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMobileDeviceAccessOverride(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMobileDeviceAccessOverride(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DeleteMobileDeviceAccessOverride",
	}
}
