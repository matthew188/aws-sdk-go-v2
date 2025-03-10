// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a mobile device access rule for the specified WorkMail organization.
func (c *Client) UpdateMobileDeviceAccessRule(ctx context.Context, params *UpdateMobileDeviceAccessRuleInput, optFns ...func(*Options)) (*UpdateMobileDeviceAccessRuleOutput, error) {
	if params == nil {
		params = &UpdateMobileDeviceAccessRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMobileDeviceAccessRule", params, optFns, c.addOperationUpdateMobileDeviceAccessRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMobileDeviceAccessRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMobileDeviceAccessRuleInput struct {

	// The effect of the rule when it matches. Allowed values are ALLOW or DENY.
	//
	// This member is required.
	Effect types.MobileDeviceAccessRuleEffect

	// The identifier of the rule to be updated.
	//
	// This member is required.
	MobileDeviceAccessRuleId *string

	// The updated rule name.
	//
	// This member is required.
	Name *string

	// The WorkMail organization under which the rule will be updated.
	//
	// This member is required.
	OrganizationId *string

	// The updated rule description.
	Description *string

	// Device models that the updated rule will match.
	DeviceModels []string

	// Device operating systems that the updated rule will match.
	DeviceOperatingSystems []string

	// Device types that the updated rule will match.
	DeviceTypes []string

	// User agents that the updated rule will match.
	DeviceUserAgents []string

	// Device models that the updated rule will not match. All other device models will
	// match.
	NotDeviceModels []string

	// Device operating systems that the updated rule will not match. All other device
	// operating systems will match.
	NotDeviceOperatingSystems []string

	// Device types that the updated rule will not match. All other device types will
	// match.
	NotDeviceTypes []string

	// User agents that the updated rule will not match. All other user agents will
	// match.
	NotDeviceUserAgents []string

	noSmithyDocumentSerde
}

type UpdateMobileDeviceAccessRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMobileDeviceAccessRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMobileDeviceAccessRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMobileDeviceAccessRule{}, middleware.After)
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
	if err = addOpUpdateMobileDeviceAccessRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMobileDeviceAccessRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMobileDeviceAccessRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "UpdateMobileDeviceAccessRule",
	}
}
