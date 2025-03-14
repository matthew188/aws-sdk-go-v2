// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerlinuxsubscriptions

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/licensemanagerlinuxsubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the service settings for Linux subscriptions.
func (c *Client) UpdateServiceSettings(ctx context.Context, params *UpdateServiceSettingsInput, optFns ...func(*Options)) (*UpdateServiceSettingsOutput, error) {
	if params == nil {
		params = &UpdateServiceSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServiceSettings", params, optFns, c.addOperationUpdateServiceSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceSettingsInput struct {

	// Describes if the discovery of Linux subscriptions is enabled.
	//
	// This member is required.
	LinuxSubscriptionsDiscovery types.LinuxSubscriptionsDiscovery

	// The settings defined for Linux subscriptions discovery. The settings include if
	// Organizations integration has been enabled, and which Regions data will be
	// aggregated from.
	//
	// This member is required.
	LinuxSubscriptionsDiscoverySettings *types.LinuxSubscriptionsDiscoverySettings

	// Describes if updates are allowed to the service settings for Linux
	// subscriptions. If you allow updates, you can aggregate Linux subscription data
	// in more than one home Region.
	AllowUpdate *bool

	noSmithyDocumentSerde
}

type UpdateServiceSettingsOutput struct {

	// The Region in which License Manager displays the aggregated data for Linux
	// subscriptions.
	HomeRegions []string

	// Lists if discovery has been enabled for Linux subscriptions.
	LinuxSubscriptionsDiscovery types.LinuxSubscriptionsDiscovery

	// The settings defined for Linux subscriptions discovery. The settings include if
	// Organizations integration has been enabled, and which Regions data will be
	// aggregated from.
	LinuxSubscriptionsDiscoverySettings *types.LinuxSubscriptionsDiscoverySettings

	// Indicates the status of Linux subscriptions settings being applied.
	Status types.Status

	// A message which details the Linux subscriptions service settings current status.
	StatusMessage map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateServiceSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateServiceSettings{}, middleware.After)
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
	if err = addOpUpdateServiceSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServiceSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager-linux-subscriptions",
		OperationName: "UpdateServiceSettings",
	}
}
