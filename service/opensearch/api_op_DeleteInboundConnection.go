// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows the destination Amazon OpenSearch Service domain owner to delete an
// existing inbound cross-cluster search connection. For more information, see
// Cross-cluster search for Amazon OpenSearch Service
// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html).
func (c *Client) DeleteInboundConnection(ctx context.Context, params *DeleteInboundConnectionInput, optFns ...func(*Options)) (*DeleteInboundConnectionOutput, error) {
	if params == nil {
		params = &DeleteInboundConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteInboundConnection", params, optFns, c.addOperationDeleteInboundConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteInboundConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteInboundConnection operation.
type DeleteInboundConnectionInput struct {

	// The ID of the inbound connection to permanently delete.
	//
	// This member is required.
	ConnectionId *string

	noSmithyDocumentSerde
}

// The results of a DeleteInboundConnection operation. Contains details about the
// deleted inbound connection.
type DeleteInboundConnectionOutput struct {

	// The deleted inbound connection.
	Connection *types.InboundConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteInboundConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteInboundConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteInboundConnection{}, middleware.After)
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
	if err = addOpDeleteInboundConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInboundConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteInboundConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DeleteInboundConnection",
	}
}
