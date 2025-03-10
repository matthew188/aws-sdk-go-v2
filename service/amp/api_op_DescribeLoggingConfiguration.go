// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes logging configuration.
func (c *Client) DescribeLoggingConfiguration(ctx context.Context, params *DescribeLoggingConfigurationInput, optFns ...func(*Options)) (*DescribeLoggingConfigurationOutput, error) {
	if params == nil {
		params = &DescribeLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoggingConfiguration", params, optFns, c.addOperationDescribeLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeLoggingConfiguration operation.
type DescribeLoggingConfigurationInput struct {

	// The ID of the workspace to vend logs to.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

// Represents the output of a DescribeLoggingConfiguration operation.
type DescribeLoggingConfigurationOutput struct {

	// Metadata object containing information about the logging configuration of a
	// workspace.
	//
	// This member is required.
	LoggingConfiguration *types.LoggingConfigurationMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeLoggingConfiguration{}, middleware.After)
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
	if err = addOpDescribeLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aps",
		OperationName: "DescribeLoggingConfiguration",
	}
}
