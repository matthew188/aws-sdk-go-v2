// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the standards specified by the provided StandardsArn. To obtain the ARN
// for a standard, use the DescribeStandards operation. For more information, see
// the Security Standards
// (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards.html)
// section of the Security Hub User Guide.
func (c *Client) BatchEnableStandards(ctx context.Context, params *BatchEnableStandardsInput, optFns ...func(*Options)) (*BatchEnableStandardsOutput, error) {
	if params == nil {
		params = &BatchEnableStandardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchEnableStandards", params, optFns, c.addOperationBatchEnableStandardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchEnableStandardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchEnableStandardsInput struct {

	// The list of standards checks to enable.
	//
	// This member is required.
	StandardsSubscriptionRequests []types.StandardsSubscriptionRequest

	noSmithyDocumentSerde
}

type BatchEnableStandardsOutput struct {

	// The details of the standards subscriptions that were enabled.
	StandardsSubscriptions []types.StandardsSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchEnableStandardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchEnableStandards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchEnableStandards{}, middleware.After)
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
	if err = addOpBatchEnableStandardsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchEnableStandards(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchEnableStandards(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "BatchEnableStandards",
	}
}
