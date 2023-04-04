// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a listener.
func (c *Client) UpdateListener(ctx context.Context, params *UpdateListenerInput, optFns ...func(*Options)) (*UpdateListenerOutput, error) {
	if params == nil {
		params = &UpdateListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateListener", params, optFns, c.addOperationUpdateListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateListenerInput struct {

	// The Amazon Resource Name (ARN) of the listener to update.
	//
	// This member is required.
	ListenerArn *string

	// Client affinity lets you direct all requests from a user to the same endpoint,
	// if you have stateful applications, regardless of the port and protocol of the
	// client request. Client affinity gives you control over whether to always route
	// each client to the same specific endpoint. Global Accelerator uses a
	// consistent-flow hashing algorithm to choose the optimal endpoint for a
	// connection. If client affinity is NONE, Global Accelerator uses the "five-tuple"
	// (5-tuple) properties—source IP address, source port, destination IP address,
	// destination port, and protocol—to select the hash value, and then chooses the
	// best endpoint. However, with this setting, if someone uses different ports to
	// connect to Global Accelerator, their connections might not be always routed to
	// the same endpoint because the hash value changes. If you want a given client to
	// always be routed to the same endpoint, set client affinity to SOURCE_IP instead.
	// When you use the SOURCE_IP setting, Global Accelerator uses the "two-tuple"
	// (2-tuple) properties— source (client) IP address and destination IP address—to
	// select the hash value. The default value is NONE.
	ClientAffinity types.ClientAffinity

	// The updated list of port ranges for the connections from clients to the
	// accelerator.
	PortRanges []types.PortRange

	// The updated protocol for the connections from clients to the accelerator.
	Protocol types.Protocol

	noSmithyDocumentSerde
}

type UpdateListenerOutput struct {

	// Information for the updated listener.
	Listener *types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateListener{}, middleware.After)
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
	if err = addOpUpdateListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateListener(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "UpdateListener",
	}
}
