// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Inserts metadata into the active stream of the specified channel. At most 5
// requests per second per channel are allowed, each with a maximum 1 KB payload.
// (If 5 TPS is not sufficient for your needs, we recommend batching your data into
// a single PutMetadata call.) At most 155 requests per second per account are
// allowed. Also see Embedding Metadata within a Video Stream
// (https://docs.aws.amazon.com/ivs/latest/userguide/metadata.html) in the Amazon
// IVS User Guide.
func (c *Client) PutMetadata(ctx context.Context, params *PutMetadataInput, optFns ...func(*Options)) (*PutMetadataOutput, error) {
	if params == nil {
		params = &PutMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutMetadata", params, optFns, c.addOperationPutMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutMetadataInput struct {

	// ARN of the channel into which metadata is inserted. This channel must have an
	// active stream.
	//
	// This member is required.
	ChannelArn *string

	// Metadata to insert into the stream. Maximum: 1 KB per request.
	//
	// This member is required.
	Metadata *string

	noSmithyDocumentSerde
}

type PutMetadataOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutMetadata{}, middleware.After)
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
	if err = addOpPutMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "PutMetadata",
	}
}
