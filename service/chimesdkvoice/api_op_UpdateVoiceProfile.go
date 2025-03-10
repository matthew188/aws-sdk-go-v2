// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified voice profile’s voice print and refreshes its expiration
// timestamp. As a condition of using this feature, you acknowledge that the
// collection, use, storage, and retention of your caller’s biometric identifiers
// and biometric information (“biometric data”) in the form of a digital voiceprint
// requires the caller’s informed consent via a written release. Such consent is
// required under various state laws, including biometrics laws in Illinois, Texas,
// Washington and other state privacy laws. You must provide a written release to
// each caller through a process that clearly reflects each caller’s informed
// consent before using Amazon Chime SDK Voice Insights service, as required under
// the terms of your agreement with AWS governing your use of the service.
func (c *Client) UpdateVoiceProfile(ctx context.Context, params *UpdateVoiceProfileInput, optFns ...func(*Options)) (*UpdateVoiceProfileOutput, error) {
	if params == nil {
		params = &UpdateVoiceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVoiceProfile", params, optFns, c.addOperationUpdateVoiceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVoiceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVoiceProfileInput struct {

	// The ID of the speaker search task.
	//
	// This member is required.
	SpeakerSearchTaskId *string

	// The profile ID.
	//
	// This member is required.
	VoiceProfileId *string

	noSmithyDocumentSerde
}

type UpdateVoiceProfileOutput struct {

	// The updated voice profile settings.
	VoiceProfile *types.VoiceProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVoiceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVoiceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVoiceProfile{}, middleware.After)
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
	if err = addOpUpdateVoiceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVoiceProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVoiceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateVoiceProfile",
	}
}
