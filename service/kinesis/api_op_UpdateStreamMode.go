// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the capacity mode of the data stream. Currently, in Kinesis Data
// Streams, you can choose between an on-demand capacity mode and a provisioned
// capacity mode for your data stream.
func (c *Client) UpdateStreamMode(ctx context.Context, params *UpdateStreamModeInput, optFns ...func(*Options)) (*UpdateStreamModeOutput, error) {
	if params == nil {
		params = &UpdateStreamModeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStreamMode", params, optFns, c.addOperationUpdateStreamModeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStreamModeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStreamModeInput struct {

	// Specifies the ARN of the data stream whose capacity mode you want to update.
	//
	// This member is required.
	StreamARN *string

	// Specifies the capacity mode to which you want to set your data stream.
	// Currently, in Kinesis Data Streams, you can choose between an on-demand capacity
	// mode and a provisioned capacity mode for your data streams.
	//
	// This member is required.
	StreamModeDetails *types.StreamModeDetails

	noSmithyDocumentSerde
}

type UpdateStreamModeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStreamModeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateStreamMode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateStreamMode{}, middleware.After)
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
	if err = addOpUpdateStreamModeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStreamMode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStreamMode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "UpdateStreamMode",
	}
}
