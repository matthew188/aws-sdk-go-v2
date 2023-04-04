// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a receipt rule. For information about setting up receipt rules, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rules.html).
// You can execute this operation no more than once per second.
func (c *Client) CreateReceiptRule(ctx context.Context, params *CreateReceiptRuleInput, optFns ...func(*Options)) (*CreateReceiptRuleOutput, error) {
	if params == nil {
		params = &CreateReceiptRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReceiptRule", params, optFns, c.addOperationCreateReceiptRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReceiptRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to create a receipt rule. You use receipt rules to receive
// email with Amazon SES. For more information, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
type CreateReceiptRuleInput struct {

	// A data structure that contains the specified rule's name, actions, recipients,
	// domains, enabled status, scan status, and TLS policy.
	//
	// This member is required.
	Rule *types.ReceiptRule

	// The name of the rule set that the receipt rule will be added to.
	//
	// This member is required.
	RuleSetName *string

	// The name of an existing rule after which the new rule will be placed. If this
	// parameter is null, the new rule will be inserted at the beginning of the rule
	// list.
	After *string

	noSmithyDocumentSerde
}

// An empty element returned on a successful request.
type CreateReceiptRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReceiptRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateReceiptRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateReceiptRule{}, middleware.After)
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
	if err = addOpCreateReceiptRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReceiptRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReceiptRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateReceiptRule",
	}
}
