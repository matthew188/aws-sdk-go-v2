// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the License Manager settings for the current Region.
func (c *Client) GetServiceSettings(ctx context.Context, params *GetServiceSettingsInput, optFns ...func(*Options)) (*GetServiceSettingsOutput, error) {
	if params == nil {
		params = &GetServiceSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceSettings", params, optFns, c.addOperationGetServiceSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceSettingsInput struct {
	noSmithyDocumentSerde
}

type GetServiceSettingsOutput struct {

	// Indicates whether cross-account discovery is enabled.
	EnableCrossAccountsDiscovery *bool

	// Amazon Resource Name (ARN) of the resource share. The License Manager management
	// account provides member accounts with access to this share.
	LicenseManagerResourceShareArn *string

	// Indicates whether Organizations is integrated with License Manager for
	// cross-account discovery.
	OrganizationConfiguration *types.OrganizationConfiguration

	// Regional S3 bucket path for storing reports, license trail event data, discovery
	// data, and so on.
	S3BucketArn *string

	// SNS topic configured to receive notifications from License Manager.
	SnsTopicArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetServiceSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServiceSettings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "GetServiceSettings",
	}
}
