// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the specified MFA device and associates it with the specified IAM user.
// When enabled, the MFA device is required for every subsequent login by the IAM
// user associated with the device.
func (c *Client) EnableMFADevice(ctx context.Context, params *EnableMFADeviceInput, optFns ...func(*Options)) (*EnableMFADeviceOutput, error) {
	if params == nil {
		params = &EnableMFADeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableMFADevice", params, optFns, c.addOperationEnableMFADeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableMFADeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableMFADeviceInput struct {

	// An authentication code emitted by the device. The format for this parameter is a
	// string of six digits. Submit your request immediately after generating the
	// authentication codes. If you generate the codes and then wait too long to submit
	// the request, the MFA device successfully associates with the user but the MFA
	// device becomes out of sync. This happens because time-based one-time passwords
	// (TOTP) expire after a short period of time. If this happens, you can resync the
	// device
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_sync.html).
	//
	// This member is required.
	AuthenticationCode1 *string

	// A subsequent authentication code emitted by the device. The format for this
	// parameter is a string of six digits. Submit your request immediately after
	// generating the authentication codes. If you generate the codes and then wait too
	// long to submit the request, the MFA device successfully associates with the user
	// but the MFA device becomes out of sync. This happens because time-based one-time
	// passwords (TOTP) expire after a short period of time. If this happens, you can
	// resync the device
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_sync.html).
	//
	// This member is required.
	AuthenticationCode2 *string

	// The serial number that uniquely identifies the MFA device. For virtual MFA
	// devices, the serial number is the device ARN. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: =,.@:/-
	//
	// This member is required.
	SerialNumber *string

	// The name of the IAM user for whom you want to enable the MFA device. This
	// parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a
	// string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

type EnableMFADeviceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableMFADeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnableMFADevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableMFADevice{}, middleware.After)
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
	if err = addOpEnableMFADeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableMFADevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableMFADevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "EnableMFADevice",
	}
}
