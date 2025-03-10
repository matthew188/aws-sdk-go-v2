// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the mobile device access override for the given WorkMail organization,
// user, and device.
func (c *Client) GetMobileDeviceAccessOverride(ctx context.Context, params *GetMobileDeviceAccessOverrideInput, optFns ...func(*Options)) (*GetMobileDeviceAccessOverrideOutput, error) {
	if params == nil {
		params = &GetMobileDeviceAccessOverrideInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMobileDeviceAccessOverride", params, optFns, c.addOperationGetMobileDeviceAccessOverrideMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMobileDeviceAccessOverrideOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMobileDeviceAccessOverrideInput struct {

	// The mobile device to which the override applies. DeviceId is case insensitive.
	//
	// This member is required.
	DeviceId *string

	// The WorkMail organization to which you want to apply the override.
	//
	// This member is required.
	OrganizationId *string

	// Identifies the WorkMail user for the override. Accepts the following types of
	// user identities:
	//
	// * User ID: 12345678-1234-1234-1234-123456789012 or
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

type GetMobileDeviceAccessOverrideOutput struct {

	// The date the override was first created.
	DateCreated *time.Time

	// The date the description was last modified.
	DateModified *time.Time

	// A description of the override.
	Description *string

	// The device to which the access override applies.
	DeviceId *string

	// The effect of the override, ALLOW or DENY.
	Effect types.MobileDeviceAccessRuleEffect

	// The WorkMail user to which the access override applies.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMobileDeviceAccessOverrideMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMobileDeviceAccessOverride{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMobileDeviceAccessOverride{}, middleware.After)
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
	if err = addOpGetMobileDeviceAccessOverrideValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMobileDeviceAccessOverride(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMobileDeviceAccessOverride(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "GetMobileDeviceAccessOverride",
	}
}
