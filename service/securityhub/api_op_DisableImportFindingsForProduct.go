// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables the integration of the specified product with Security Hub. After the
// integration is disabled, findings from that product are no longer sent to
// Security Hub.
func (c *Client) DisableImportFindingsForProduct(ctx context.Context, params *DisableImportFindingsForProductInput, optFns ...func(*Options)) (*DisableImportFindingsForProductOutput, error) {
	if params == nil {
		params = &DisableImportFindingsForProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableImportFindingsForProduct", params, optFns, c.addOperationDisableImportFindingsForProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableImportFindingsForProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableImportFindingsForProductInput struct {

	// The ARN of the integrated product to disable the integration for.
	//
	// This member is required.
	ProductSubscriptionArn *string

	noSmithyDocumentSerde
}

type DisableImportFindingsForProductOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableImportFindingsForProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisableImportFindingsForProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisableImportFindingsForProduct{}, middleware.After)
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
	if err = addOpDisableImportFindingsForProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableImportFindingsForProduct(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableImportFindingsForProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DisableImportFindingsForProduct",
	}
}
