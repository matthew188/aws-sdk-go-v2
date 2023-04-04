// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Accepts ownership of a public virtual interface created by another Amazon Web
// Services account. After the virtual interface owner makes this call, the
// specified virtual interface is created and made available to handle traffic.
func (c *Client) ConfirmPublicVirtualInterface(ctx context.Context, params *ConfirmPublicVirtualInterfaceInput, optFns ...func(*Options)) (*ConfirmPublicVirtualInterfaceOutput, error) {
	if params == nil {
		params = &ConfirmPublicVirtualInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfirmPublicVirtualInterface", params, optFns, c.addOperationConfirmPublicVirtualInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfirmPublicVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ConfirmPublicVirtualInterfaceInput struct {

	// The ID of the virtual interface.
	//
	// This member is required.
	VirtualInterfaceId *string

	noSmithyDocumentSerde
}

type ConfirmPublicVirtualInterfaceOutput struct {

	// The state of the virtual interface. The following are the possible values:
	//
	// *
	// confirming: The creation of the virtual interface is pending confirmation from
	// the virtual interface owner. If the owner of the virtual interface is different
	// from the owner of the connection on which it is provisioned, then the virtual
	// interface will remain in this state until it is confirmed by the virtual
	// interface owner.
	//
	// * verifying: This state only applies to public virtual
	// interfaces. Each public virtual interface needs validation before the virtual
	// interface can be created.
	//
	// * pending: A virtual interface is in this state from
	// the time that it is created until the virtual interface is ready to forward
	// traffic.
	//
	// * available: A virtual interface that is able to forward traffic.
	//
	// *
	// down: A virtual interface that is BGP down.
	//
	// * deleting: A virtual interface is
	// in this state immediately after calling DeleteVirtualInterface until it can no
	// longer forward traffic.
	//
	// * deleted: A virtual interface that cannot forward
	// traffic.
	//
	// * rejected: The virtual interface owner has declined creation of the
	// virtual interface. If a virtual interface in the Confirming state is deleted by
	// the virtual interface owner, the virtual interface enters the Rejected state.
	//
	// *
	// unknown: The state of the virtual interface is not available.
	VirtualInterfaceState types.VirtualInterfaceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConfirmPublicVirtualInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpConfirmPublicVirtualInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpConfirmPublicVirtualInterface{}, middleware.After)
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
	if err = addOpConfirmPublicVirtualInterfaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConfirmPublicVirtualInterface(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConfirmPublicVirtualInterface(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "ConfirmPublicVirtualInterface",
	}
}
