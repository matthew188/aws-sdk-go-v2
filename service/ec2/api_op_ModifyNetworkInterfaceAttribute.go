// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified network interface attribute. You can specify only one
// attribute at a time. You can use this action to attach and detach security
// groups from an existing EC2 instance.
func (c *Client) ModifyNetworkInterfaceAttribute(ctx context.Context, params *ModifyNetworkInterfaceAttributeInput, optFns ...func(*Options)) (*ModifyNetworkInterfaceAttributeOutput, error) {
	if params == nil {
		params = &ModifyNetworkInterfaceAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyNetworkInterfaceAttribute", params, optFns, c.addOperationModifyNetworkInterfaceAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyNetworkInterfaceAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyNetworkInterfaceAttribute.
type ModifyNetworkInterfaceAttributeInput struct {

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// Information about the interface attachment. If modifying the delete on
	// termination attribute, you must specify the ID of the interface attachment.
	Attachment *types.NetworkInterfaceAttachmentChanges

	// A description for the network interface.
	Description *types.AttributeValue

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Updates the ENA Express configuration for the network interface that’s attached
	// to the instance.
	EnaSrdSpecification *types.EnaSrdSpecification

	// Changes the security groups for the network interface. The new set of groups you
	// specify replaces the current set. You must specify at least one group, even if
	// it's just the default security group in the VPC. You must specify the ID of the
	// security group, not the name.
	Groups []string

	// Enable or disable source/destination checks, which ensure that the instance is
	// either the source or the destination of any traffic that it receives. If the
	// value is true, source/destination checks are enabled; otherwise, they are
	// disabled. The default value is true. You must disable source/destination checks
	// if the instance runs services such as network address translation, routing, or
	// firewalls.
	SourceDestCheck *types.AttributeBooleanValue

	noSmithyDocumentSerde
}

type ModifyNetworkInterfaceAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyNetworkInterfaceAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyNetworkInterfaceAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyNetworkInterfaceAttribute{}, middleware.After)
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
	if err = addOpModifyNetworkInterfaceAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyNetworkInterfaceAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyNetworkInterfaceAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyNetworkInterfaceAttribute",
	}
}
