// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Device Defender audit suppression. Requires permission to access the
// DeleteAuditSuppression
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DeleteAuditSuppression(ctx context.Context, params *DeleteAuditSuppressionInput, optFns ...func(*Options)) (*DeleteAuditSuppressionOutput, error) {
	if params == nil {
		params = &DeleteAuditSuppressionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAuditSuppression", params, optFns, c.addOperationDeleteAuditSuppressionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAuditSuppressionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAuditSuppressionInput struct {

	// An audit check name. Checks must be enabled for your account. (Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or use UpdateAccountAuditConfiguration to select which checks
	// are enabled.)
	//
	// This member is required.
	CheckName *string

	// Information that identifies the noncompliant resource.
	//
	// This member is required.
	ResourceIdentifier *types.ResourceIdentifier

	noSmithyDocumentSerde
}

type DeleteAuditSuppressionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAuditSuppressionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteAuditSuppression{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteAuditSuppression{}, middleware.After)
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
	if err = addOpDeleteAuditSuppressionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAuditSuppression(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAuditSuppression(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteAuditSuppression",
	}
}
