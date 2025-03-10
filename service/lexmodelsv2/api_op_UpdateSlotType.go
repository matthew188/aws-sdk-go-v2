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

// Updates the configuration of an existing slot type.
func (c *Client) UpdateSlotType(ctx context.Context, params *UpdateSlotTypeInput, optFns ...func(*Options)) (*UpdateSlotTypeOutput, error) {
	if params == nil {
		params = &UpdateSlotTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSlotType", params, optFns, c.addOperationUpdateSlotTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSlotTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSlotTypeInput struct {

	// The identifier of the bot that contains the slot type.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the slot type. Must be DRAFT.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale that contains the slot type. The
	// string must match one of the supported locales. For more information, see
	// Supported languages
	// (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html).
	//
	// This member is required.
	LocaleId *string

	// The unique identifier of the slot type to update.
	//
	// This member is required.
	SlotTypeId *string

	// The new name of the slot type.
	//
	// This member is required.
	SlotTypeName *string

	// Specifications for a composite slot type.
	CompositeSlotTypeSetting *types.CompositeSlotTypeSetting

	// The new description of the slot type.
	Description *string

	// Provides information about the external source of the slot type's definition.
	ExternalSourceSetting *types.ExternalSourceSetting

	// The new built-in slot type that should be used as the parent of this slot type.
	ParentSlotTypeSignature *string

	// A new list of values and their optional synonyms that define the values that the
	// slot type can take.
	SlotTypeValues []types.SlotTypeValue

	// The strategy that Amazon Lex should use when deciding on a value from the list
	// of slot type values.
	ValueSelectionSetting *types.SlotValueSelectionSetting

	noSmithyDocumentSerde
}

type UpdateSlotTypeOutput struct {

	// The identifier of the bot that contains the slot type.
	BotId *string

	// The version of the bot that contains the slot type. This is always DRAFT.
	BotVersion *string

	// Specifications for a composite slot type.
	CompositeSlotTypeSetting *types.CompositeSlotTypeSetting

	// The timestamp of the date and time that the slot type was created.
	CreationDateTime *time.Time

	// The updated description of the slot type.
	Description *string

	// Provides information about the external source of the slot type's definition.
	ExternalSourceSetting *types.ExternalSourceSetting

	// A timestamp of the date and time that the slot type was last updated.
	LastUpdatedDateTime *time.Time

	// The language and locale of the updated slot type.
	LocaleId *string

	// The updated signature of the built-in slot type that is the parent of this slot
	// type.
	ParentSlotTypeSignature *string

	// The unique identifier of the updated slot type.
	SlotTypeId *string

	// The updated name of the slot type.
	SlotTypeName *string

	// The updated values that the slot type provides.
	SlotTypeValues []types.SlotTypeValue

	// The updated strategy that Amazon Lex uses to determine which value to select
	// from the slot type.
	ValueSelectionSetting *types.SlotValueSelectionSetting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSlotTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSlotType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSlotType{}, middleware.After)
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
	if err = addOpUpdateSlotTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSlotType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSlotType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "UpdateSlotType",
	}
}
