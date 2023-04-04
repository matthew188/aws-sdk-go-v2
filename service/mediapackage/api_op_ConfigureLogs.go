// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the Channel's properities to configure log subscription
func (c *Client) ConfigureLogs(ctx context.Context, params *ConfigureLogsInput, optFns ...func(*Options)) (*ConfigureLogsOutput, error) {
	if params == nil {
		params = &ConfigureLogsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfigureLogs", params, optFns, c.addOperationConfigureLogsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfigureLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// the option to configure log subscription.
type ConfigureLogsInput struct {

	// The ID of the channel to log subscription.
	//
	// This member is required.
	Id *string

	// Configure egress access logging.
	EgressAccessLogs *types.EgressAccessLogs

	// Configure ingress access logging.
	IngressAccessLogs *types.IngressAccessLogs

	noSmithyDocumentSerde
}

type ConfigureLogsOutput struct {

	// The Amazon Resource Name (ARN) assigned to the Channel.
	Arn *string

	// The date and time the Channel was created.
	CreatedAt *string

	// A short text description of the Channel.
	Description *string

	// Configure egress access logging.
	EgressAccessLogs *types.EgressAccessLogs

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *types.HlsIngest

	// The ID of the Channel.
	Id *string

	// Configure ingress access logging.
	IngressAccessLogs *types.IngressAccessLogs

	// A collection of tags associated with a resource
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConfigureLogsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpConfigureLogs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpConfigureLogs{}, middleware.After)
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
	if err = addOpConfigureLogsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConfigureLogs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConfigureLogs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "ConfigureLogs",
	}
}
