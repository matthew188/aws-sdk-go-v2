// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a specified authorization key pair. This invalidates future viewer
// tokens generated using the key pair’s privateKey. For more information, see
// Setting Up Private Channels
// (https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html) in the
// Amazon IVS User Guide.
func (c *Client) DeletePlaybackKeyPair(ctx context.Context, params *DeletePlaybackKeyPairInput, optFns ...func(*Options)) (*DeletePlaybackKeyPairOutput, error) {
	if params == nil {
		params = &DeletePlaybackKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePlaybackKeyPair", params, optFns, c.addOperationDeletePlaybackKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePlaybackKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePlaybackKeyPairInput struct {

	// ARN of the key pair to be deleted.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type DeletePlaybackKeyPairOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePlaybackKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeletePlaybackKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeletePlaybackKeyPair{}, middleware.After)
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
	if err = addOpDeletePlaybackKeyPairValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePlaybackKeyPair(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePlaybackKeyPair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "DeletePlaybackKeyPair",
	}
}
