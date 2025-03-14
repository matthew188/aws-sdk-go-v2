// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an InputProcessingConfiguration from an input.
func (c *Client) DeleteApplicationInputProcessingConfiguration(ctx context.Context, params *DeleteApplicationInputProcessingConfigurationInput, optFns ...func(*Options)) (*DeleteApplicationInputProcessingConfigurationOutput, error) {
	if params == nil {
		params = &DeleteApplicationInputProcessingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApplicationInputProcessingConfiguration", params, optFns, c.addOperationDeleteApplicationInputProcessingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApplicationInputProcessingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationInputProcessingConfigurationInput struct {

	// The name of the application.
	//
	// This member is required.
	ApplicationName *string

	// The application version. You can use the DescribeApplication operation to get
	// the current application version. If the version specified is not the current
	// version, the ConcurrentModificationException is returned.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The ID of the input configuration from which to delete the input processing
	// configuration. You can get a list of the input IDs for an application by using
	// the DescribeApplication operation.
	//
	// This member is required.
	InputId *string

	noSmithyDocumentSerde
}

type DeleteApplicationInputProcessingConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the application.
	ApplicationARN *string

	// The current application version ID.
	ApplicationVersionId *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteApplicationInputProcessingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationInputProcessingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationInputProcessingConfiguration{}, middleware.After)
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
	if err = addOpDeleteApplicationInputProcessingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationInputProcessingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteApplicationInputProcessingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationInputProcessingConfiguration",
	}
}
