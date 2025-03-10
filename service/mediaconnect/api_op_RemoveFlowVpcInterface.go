// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a VPC Interface from an existing flow. This request can be made only on
// a VPC interface that does not have a Source or Output associated with it. If the
// VPC interface is referenced by a Source or Output, you must first delete or
// update the Source or Output to no longer reference the VPC interface.
func (c *Client) RemoveFlowVpcInterface(ctx context.Context, params *RemoveFlowVpcInterfaceInput, optFns ...func(*Options)) (*RemoveFlowVpcInterfaceOutput, error) {
	if params == nil {
		params = &RemoveFlowVpcInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveFlowVpcInterface", params, optFns, c.addOperationRemoveFlowVpcInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveFlowVpcInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveFlowVpcInterfaceInput struct {

	// The flow that you want to remove a VPC interface from.
	//
	// This member is required.
	FlowArn *string

	// The name of the VPC interface that you want to remove.
	//
	// This member is required.
	VpcInterfaceName *string

	noSmithyDocumentSerde
}

type RemoveFlowVpcInterfaceOutput struct {

	// The ARN of the flow that is associated with the VPC interface you removed.
	FlowArn *string

	// IDs of network interfaces associated with the removed VPC interface that Media
	// Connect was unable to remove.
	NonDeletedNetworkInterfaceIds []string

	// The name of the VPC interface that was removed.
	VpcInterfaceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveFlowVpcInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveFlowVpcInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveFlowVpcInterface{}, middleware.After)
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
	if err = addOpRemoveFlowVpcInterfaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveFlowVpcInterface(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveFlowVpcInterface(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "RemoveFlowVpcInterface",
	}
}
