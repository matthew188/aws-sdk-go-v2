// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Return a full description of an App Runner observability configuration resource.
func (c *Client) DescribeObservabilityConfiguration(ctx context.Context, params *DescribeObservabilityConfigurationInput, optFns ...func(*Options)) (*DescribeObservabilityConfigurationOutput, error) {
	if params == nil {
		params = &DescribeObservabilityConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeObservabilityConfiguration", params, optFns, c.addOperationDescribeObservabilityConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeObservabilityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeObservabilityConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the App Runner observability configuration
	// that you want a description for. The ARN can be a full observability
	// configuration ARN, or a partial ARN ending with either .../name  or
	// .../name/revision . If a revision isn't specified, the latest active revision is
	// described.
	//
	// This member is required.
	ObservabilityConfigurationArn *string

	noSmithyDocumentSerde
}

type DescribeObservabilityConfigurationOutput struct {

	// A full description of the App Runner observability configuration that you
	// specified in this request.
	//
	// This member is required.
	ObservabilityConfiguration *types.ObservabilityConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeObservabilityConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeObservabilityConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeObservabilityConfiguration{}, middleware.After)
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
	if err = addOpDescribeObservabilityConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeObservabilityConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeObservabilityConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "DescribeObservabilityConfiguration",
	}
}
