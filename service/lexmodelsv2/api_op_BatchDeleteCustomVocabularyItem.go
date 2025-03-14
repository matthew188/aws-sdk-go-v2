// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a batch of custom vocabulary items for a given bot locale's custom
// vocabulary.
func (c *Client) BatchDeleteCustomVocabularyItem(ctx context.Context, params *BatchDeleteCustomVocabularyItemInput, optFns ...func(*Options)) (*BatchDeleteCustomVocabularyItemOutput, error) {
	if params == nil {
		params = &BatchDeleteCustomVocabularyItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteCustomVocabularyItem", params, optFns, c.addOperationBatchDeleteCustomVocabularyItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteCustomVocabularyItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteCustomVocabularyItemInput struct {

	// The identifier of the bot associated with this custom vocabulary.
	//
	// This member is required.
	BotId *string

	// The identifier of the version of the bot associated with this custom vocabulary.
	//
	// This member is required.
	BotVersion *string

	// A list of custom vocabulary items requested to be deleted. Each entry must
	// contain the unique custom vocabulary entry identifier.
	//
	// This member is required.
	CustomVocabularyItemList []types.CustomVocabularyEntryId

	// The identifier of the language and locale where this custom vocabulary is used.
	// The string must match one of the supported locales. For more information, see
	// Supported Languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type BatchDeleteCustomVocabularyItemOutput struct {

	// The identifier of the bot associated with this custom vocabulary.
	BotId *string

	// The identifier of the version of the bot associated with this custom vocabulary.
	BotVersion *string

	// A list of custom vocabulary items that failed to delete during the operation.
	// The reason for the error is contained within each error object.
	Errors []types.FailedCustomVocabularyItem

	// The identifier of the language and locale where this custom vocabulary is used.
	// The string must match one of the supported locales. For more information, see
	// Supported languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	LocaleId *string

	// A list of custom vocabulary items that were successfully deleted during the
	// operation.
	Resources []types.CustomVocabularyItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteCustomVocabularyItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDeleteCustomVocabularyItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDeleteCustomVocabularyItem{}, middleware.After)
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
	if err = addOpBatchDeleteCustomVocabularyItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteCustomVocabularyItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteCustomVocabularyItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "BatchDeleteCustomVocabularyItem",
	}
}
