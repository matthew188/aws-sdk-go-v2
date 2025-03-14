// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Use this to provide your transcript data, and to start the bot recommendation
// process.
func (c *Client) StartBotRecommendation(ctx context.Context, params *StartBotRecommendationInput, optFns ...func(*Options)) (*StartBotRecommendationOutput, error) {
	if params == nil {
		params = &StartBotRecommendationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBotRecommendation", params, optFns, c.addOperationStartBotRecommendationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBotRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBotRecommendationInput struct {

	// The unique identifier of the bot containing the bot recommendation.
	//
	// This member is required.
	BotId *string

	// The version of the bot containing the bot recommendation.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale of the bot recommendation to start.
	// The string must match one of the supported locales. For more information, see
	// Supported languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	//
	// This member is required.
	LocaleId *string

	// The object representing the Amazon S3 bucket containing the transcript, as well
	// as the associated metadata.
	//
	// This member is required.
	TranscriptSourceSetting *types.TranscriptSourceSetting

	// The object representing the passwords that will be used to encrypt the data
	// related to the bot recommendation results, as well as the KMS key ARN used to
	// encrypt the associated metadata.
	EncryptionSetting *types.EncryptionSetting

	noSmithyDocumentSerde
}

type StartBotRecommendationOutput struct {

	// The unique identifier of the bot containing the bot recommendation.
	BotId *string

	// The identifier of the bot recommendation that you have created.
	BotRecommendationId *string

	// The status of the bot recommendation. If the status is Failed, then the reasons
	// for the failure are listed in the failureReasons field.
	BotRecommendationStatus types.BotRecommendationStatus

	// The version of the bot containing the bot recommendation.
	BotVersion *string

	// A timestamp of the date and time that the bot recommendation was created.
	CreationDateTime *time.Time

	// The object representing the passwords that were used to encrypt the data related
	// to the bot recommendation results, as well as the KMS key ARN used to encrypt
	// the associated metadata.
	EncryptionSetting *types.EncryptionSetting

	// The identifier of the language and locale of the bot recommendation to start.
	// The string must match one of the supported locales. For more information, see
	// Supported languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	LocaleId *string

	// The object representing the Amazon S3 bucket containing the transcript, as well
	// as the associated metadata.
	TranscriptSourceSetting *types.TranscriptSourceSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartBotRecommendationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartBotRecommendation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBotRecommendation{}, middleware.After)
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
	if err = addOpStartBotRecommendationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBotRecommendation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBotRecommendation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "StartBotRecommendation",
	}
}
