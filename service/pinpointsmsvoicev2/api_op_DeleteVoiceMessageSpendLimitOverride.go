// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an account level monthly spend limit override for sending voice
// messages. Deleting a spend limit override sets the EnforcedLimit equal to the
// MaxLimit, which is controlled by Amazon Web Services. For more information on
// spending limits (quotas) see Amazon Pinpoint quotas
// (https://docs.aws.amazon.com/pinpoint/latest/developerguide/quotas.html) in the
// Amazon Pinpoint Developer Guide.
func (c *Client) DeleteVoiceMessageSpendLimitOverride(ctx context.Context, params *DeleteVoiceMessageSpendLimitOverrideInput, optFns ...func(*Options)) (*DeleteVoiceMessageSpendLimitOverrideOutput, error) {
	if params == nil {
		params = &DeleteVoiceMessageSpendLimitOverrideInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteVoiceMessageSpendLimitOverride", params, optFns, c.addOperationDeleteVoiceMessageSpendLimitOverrideMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteVoiceMessageSpendLimitOverrideOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVoiceMessageSpendLimitOverrideInput struct {
	noSmithyDocumentSerde
}

type DeleteVoiceMessageSpendLimitOverrideOutput struct {

	// The current monthly limit, in US dollars.
	MonthlyLimit *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteVoiceMessageSpendLimitOverrideMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteVoiceMessageSpendLimitOverride{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteVoiceMessageSpendLimitOverride{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVoiceMessageSpendLimitOverride(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteVoiceMessageSpendLimitOverride(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DeleteVoiceMessageSpendLimitOverride",
	}
}
