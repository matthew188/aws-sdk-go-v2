// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new room profile with the specified details.
func (c *Client) CreateProfile(ctx context.Context, params *CreateProfileInput, optFns ...func(*Options)) (*CreateProfileOutput, error) {
	if params == nil {
		params = &CreateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProfile", params, optFns, c.addOperationCreateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProfileInput struct {

	// The valid address for the room.
	//
	// This member is required.
	Address *string

	// The distance unit to be used by devices in the profile.
	//
	// This member is required.
	DistanceUnit types.DistanceUnit

	// The name of a room profile.
	//
	// This member is required.
	ProfileName *string

	// The temperature unit to be used by devices in the profile.
	//
	// This member is required.
	TemperatureUnit types.TemperatureUnit

	// The time zone used by a room profile.
	//
	// This member is required.
	Timezone *string

	// A wake word for Alexa, Echo, Amazon, or a computer.
	//
	// This member is required.
	WakeWord types.WakeWord

	// The user-specified token that is used during the creation of a profile.
	ClientRequestToken *string

	// Whether data retention of the profile is enabled.
	DataRetentionOptIn *bool

	// The locale of the room profile. (This is currently only available to a limited
	// preview audience.)
	Locale *string

	// The maximum volume limit for a room profile.
	MaxVolumeLimit *int32

	// The meeting room settings of a room profile.
	MeetingRoomConfiguration *types.CreateMeetingRoomConfiguration

	// Whether PSTN calling is enabled.
	PSTNEnabled *bool

	// Whether room profile setup is enabled.
	SetupModeDisabled *bool

	// The tags for the profile.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateProfileOutput struct {

	// The ARN of the newly created room profile in the response.
	ProfileArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProfile{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateProfileMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProfile(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateProfile struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProfile) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProfileInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateProfileMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateProfile{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "CreateProfile",
	}
}
