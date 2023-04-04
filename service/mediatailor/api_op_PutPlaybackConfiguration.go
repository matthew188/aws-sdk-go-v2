// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a playback configuration. For information about MediaTailor
// configurations, see Working with configurations in AWS Elemental MediaTailor
// (https://docs.aws.amazon.com/mediatailor/latest/ug/configurations.html).
func (c *Client) PutPlaybackConfiguration(ctx context.Context, params *PutPlaybackConfigurationInput, optFns ...func(*Options)) (*PutPlaybackConfigurationOutput, error) {
	if params == nil {
		params = &PutPlaybackConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPlaybackConfiguration", params, optFns, c.addOperationPutPlaybackConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPlaybackConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPlaybackConfigurationInput struct {

	// The identifier for the playback configuration.
	//
	// This member is required.
	Name *string

	// The URL for the ad decision server (ADS). This includes the specification of
	// static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing you can provide a static
	// VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The configuration for avail suppression, also known as ad suppression. For more
	// information about ad suppression, see Ad Suppression
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	AvailSuppression *types.AvailSuppression

	// The configuration for bumpers. Bumpers are short audio or video clips that play
	// at the start or before the end of an ad break. To learn more about bumpers, see
	// Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
	Bumper *types.Bumper

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *types.CdnConfiguration

	// The player parameters and aliases used as dynamic variables during session
	// initialization. For more information, see Domain Variables
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html).
	ConfigurationAliases map[string]map[string]string

	// The configuration for DASH content.
	DashConfiguration *types.DashConfigurationForPut

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *types.LivePreRollConfiguration

	// The configuration for manifest processing rules. Manifest processing rules
	// enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *types.ManifestProcessingRules

	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an
	// ad break. If the duration of underfilled ad time exceeds the personalization
	// threshold, then the personalization of the ad break is abandoned and the
	// underlying content is shown. This feature applies to ad replacement in live and
	// VOD streams, rather than ad insertion, because it relies on an underlying
	// content stream. For more information about ad break behavior, including ad
	// replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	PersonalizationThresholdSeconds int32

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in
	// gaps in media content. Configuring the slate is optional for non-VPAID
	// configurations. For VPAID, the slate is required because MediaTailor provides it
	// in the slots that are designated for dynamic ad content. The slate must be a
	// high-quality asset that contains both audio and video.
	SlateAdUrl *string

	// The tags to assign to the playback configuration. Tags are key-value pairs that
	// you can associate with Amazon resources to help with organization, access
	// control, and cost tracking. For more information, see Tagging AWS Elemental
	// MediaTailor Resources
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html).
	Tags map[string]string

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of
	// MediaTailor. Use this only if you have already set up custom profiles with the
	// help of AWS Support.
	TranscodeProfileName *string

	// The URL prefix for the parent manifest for the stream, minus the asset ID. The
	// maximum length is 512 characters.
	VideoContentSourceUrl *string

	noSmithyDocumentSerde
}

type PutPlaybackConfigurationOutput struct {

	// The URL for the ad decision server (ADS). This includes the specification of
	// static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing you can provide a static
	// VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The configuration for avail suppression, also known as ad suppression. For more
	// information about ad suppression, see Ad Suppression
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	AvailSuppression *types.AvailSuppression

	// The configuration for bumpers. Bumpers are short audio or video clips that play
	// at the start or before the end of an ad break. To learn more about bumpers, see
	// Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
	Bumper *types.Bumper

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *types.CdnConfiguration

	// The player parameters and aliases used as dynamic variables during session
	// initialization. For more information, see Domain Variables
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/variables-domain.html).
	ConfigurationAliases map[string]map[string]string

	// The configuration for DASH content.
	DashConfiguration *types.DashConfiguration

	// The configuration for HLS content.
	HlsConfiguration *types.HlsConfiguration

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *types.LivePreRollConfiguration

	// The Amazon CloudWatch log settings for a playback configuration.
	LogConfiguration *types.LogConfiguration

	// The configuration for manifest processing rules. Manifest processing rules
	// enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *types.ManifestProcessingRules

	// The identifier for the playback configuration.
	Name *string

	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an
	// ad break. If the duration of underfilled ad time exceeds the personalization
	// threshold, then the personalization of the ad break is abandoned and the
	// underlying content is shown. This feature applies to ad replacement in live and
	// VOD streams, rather than ad insertion, because it relies on an underlying
	// content stream. For more information about ad break behavior, including ad
	// replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	PersonalizationThresholdSeconds int32

	// The Amazon Resource Name (ARN) associated with the playback configuration.
	PlaybackConfigurationArn *string

	// The playback endpoint prefix associated with the playback configuration.
	PlaybackEndpointPrefix *string

	// The session initialization endpoint prefix associated with the playback
	// configuration.
	SessionInitializationEndpointPrefix *string

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in
	// gaps in media content. Configuring the slate is optional for non-VPAID
	// configurations. For VPAID, the slate is required because MediaTailor provides it
	// in the slots that are designated for dynamic ad content. The slate must be a
	// high-quality asset that contains both audio and video.
	SlateAdUrl *string

	// The tags to assign to the playback configuration. Tags are key-value pairs that
	// you can associate with Amazon resources to help with organization, access
	// control, and cost tracking. For more information, see Tagging AWS Elemental
	// MediaTailor Resources
	// (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html).
	Tags map[string]string

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of
	// MediaTailor. Use this only if you have already set up custom profiles with the
	// help of AWS Support.
	TranscodeProfileName *string

	// The URL prefix for the parent manifest for the stream, minus the asset ID. The
	// maximum length is 512 characters.
	VideoContentSourceUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutPlaybackConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutPlaybackConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutPlaybackConfiguration{}, middleware.After)
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
	if err = addOpPutPlaybackConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutPlaybackConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutPlaybackConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "PutPlaybackConfiguration",
	}
}
