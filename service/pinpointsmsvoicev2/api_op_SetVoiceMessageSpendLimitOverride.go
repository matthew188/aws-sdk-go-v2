// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets an account level monthly spend limit override for sending voice messages.
// The requested spend limit must be less than or equal to the MaxLimit, which is
// set by Amazon Web Services.
func (c *Client) SetVoiceMessageSpendLimitOverride(ctx context.Context, params *SetVoiceMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*SetVoiceMessageSpendLimitOverrideOutput, error) {
	if params == nil {
		params = &SetVoiceMessageSpendLimitOverrideInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetVoiceMessageSpendLimitOverride", params, optFns, c.addOperationSetVoiceMessageSpendLimitOverrideMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetVoiceMessageSpendLimitOverrideOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetVoiceMessageSpendLimitOverrideInput struct {

	// The new monthly limit to enforce on voice messages.
	//
	// This member is required.
	MonthlyLimit *int64

	noSmithyDocumentSerde
}

type SetVoiceMessageSpendLimitOverrideOutput struct {

	// The current monthly limit to enforce on sending voice messages.
	MonthlyLimit *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetVoiceMessageSpendLimitOverrideMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpSetVoiceMessageSpendLimitOverride{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpSetVoiceMessageSpendLimitOverride{}, middleware.After)
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
	if err = addOpSetVoiceMessageSpendLimitOverrideValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetVoiceMessageSpendLimitOverride(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetVoiceMessageSpendLimitOverride(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "SetVoiceMessageSpendLimitOverride",
	}
}
