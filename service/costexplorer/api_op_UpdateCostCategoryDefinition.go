// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Cost Category. Changes made to the Cost Category rules will
// be used to categorize the current month’s expenses and future expenses. This
// won’t change categorization for the previous months.
func (c *Client) UpdateCostCategoryDefinition(ctx context.Context, params *UpdateCostCategoryDefinitionInput, optFns ...func(*Options)) (*UpdateCostCategoryDefinitionOutput, error) {
	if params == nil {
		params = &UpdateCostCategoryDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCostCategoryDefinition", params, optFns, c.addOperationUpdateCostCategoryDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCostCategoryDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCostCategoryDefinitionInput struct {

	// The unique identifier for your Cost Category.
	//
	// This member is required.
	CostCategoryArn *string

	// The rule schema version in this particular Cost Category.
	//
	// This member is required.
	RuleVersion types.CostCategoryRuleVersion

	// The Expression object used to categorize costs. For more information, see
	// CostCategoryRule
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryRule.html).
	//
	// This member is required.
	Rules []types.CostCategoryRule

	// The default value for the cost category.
	DefaultValue *string

	// The Cost Category's effective start date. It can only be a billing start date
	// (first day of the month). If the date isn't provided, it's the first day of the
	// current month. Dates can't be before the previous twelve months, or in the
	// future.
	EffectiveStart *string

	// The split charge rules used to allocate your charges between your Cost Category
	// values.
	SplitChargeRules []types.CostCategorySplitChargeRule

	noSmithyDocumentSerde
}

type UpdateCostCategoryDefinitionOutput struct {

	// The unique identifier for your Cost Category.
	CostCategoryArn *string

	// The Cost Category's effective start date. It can only be a billing start date
	// (first day of the month).
	EffectiveStart *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCostCategoryDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCostCategoryDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCostCategoryDefinition{}, middleware.After)
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
	if err = addOpUpdateCostCategoryDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCostCategoryDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCostCategoryDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "UpdateCostCategoryDefinition",
	}
}
