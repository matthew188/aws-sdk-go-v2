// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Releases an existing origination phone number in your account. Once released, a
// phone number is no longer available for sending messages. If the origination
// phone number has deletion protection enabled or is associated with a pool, an
// Error is returned.
func (c *Client) ReleasePhoneNumber(ctx context.Context, params *ReleasePhoneNumberInput, optFns ...func(*Options)) (*ReleasePhoneNumberOutput, error) {
	if params == nil {
		params = &ReleasePhoneNumberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReleasePhoneNumber", params, optFns, c.addOperationReleasePhoneNumberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReleasePhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReleasePhoneNumberInput struct {

	// The PhoneNumberId or PhoneNumberArn of the phone number to release. You can use
	// DescribePhoneNumbers to get the values for PhoneNumberId and PhoneNumberArn.
	//
	// This member is required.
	PhoneNumberId *string

	noSmithyDocumentSerde
}

type ReleasePhoneNumberOutput struct {

	// The time when the phone number was created, in UNIX epoch time
	// (https://www.epochconverter.com/) format.
	CreatedTimestamp *time.Time

	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	IsoCountryCode *string

	// The message type that was associated with the phone number.
	MessageType types.MessageType

	// The monthly price of the phone number, in US dollars.
	MonthlyLeasingPrice *string

	// Specifies if the number could be used for text messages, voice, or both.
	NumberCapabilities []types.NumberCapability

	// The type of number that was released.
	NumberType types.NumberType

	// The name of the OptOutList that was associated with the phone number.
	OptOutListName *string

	// The phone number that was released.
	PhoneNumber *string

	// The PhoneNumberArn of the phone number that was released.
	PhoneNumberArn *string

	// The PhoneNumberId of the phone number that was released.
	PhoneNumberId *string

	// By default this is set to false. When an end recipient sends a message that
	// begins with HELP or STOP to one of your dedicated numbers, Amazon Pinpoint
	// automatically replies with a customizable message and adds the end recipient to
	// the OptOutList. When set to true you're responsible for responding to HELP and
	// STOP requests. You're also responsible for tracking and honoring opt-out
	// requests.
	SelfManagedOptOutsEnabled bool

	// The current status of the request.
	Status types.NumberStatus

	// The Amazon Resource Name (ARN) of the TwoWayChannel.
	TwoWayChannelArn *string

	// By default this is set to false. When set to true you can receive incoming text
	// messages from your end recipients.
	TwoWayEnabled bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReleasePhoneNumberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpReleasePhoneNumber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpReleasePhoneNumber{}, middleware.After)
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
	if err = addOpReleasePhoneNumberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReleasePhoneNumber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReleasePhoneNumber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "ReleasePhoneNumber",
	}
}
