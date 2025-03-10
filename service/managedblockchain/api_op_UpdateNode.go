// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a node configuration with new parameters. Applies only to Hyperledger
// Fabric.
func (c *Client) UpdateNode(ctx context.Context, params *UpdateNodeInput, optFns ...func(*Options)) (*UpdateNodeOutput, error) {
	if params == nil {
		params = &UpdateNodeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNode", params, optFns, c.addOperationUpdateNodeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNodeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNodeInput struct {

	// The unique identifier of the network that the node is on.
	//
	// This member is required.
	NetworkId *string

	// The unique identifier of the node.
	//
	// This member is required.
	NodeId *string

	// Configuration properties for publishing to Amazon CloudWatch Logs.
	LogPublishingConfiguration *types.NodeLogPublishingConfiguration

	// The unique identifier of the member that owns the node. Applies only to
	// Hyperledger Fabric.
	MemberId *string

	noSmithyDocumentSerde
}

type UpdateNodeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNodeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNode{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNode{}, middleware.After)
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
	if err = addOpUpdateNodeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNode(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNode(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "UpdateNode",
	}
}
