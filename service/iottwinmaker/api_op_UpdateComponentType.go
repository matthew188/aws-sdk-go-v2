// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates information in a component type.
func (c *Client) UpdateComponentType(ctx context.Context, params *UpdateComponentTypeInput, optFns ...func(*Options)) (*UpdateComponentTypeOutput, error) {
	if params == nil {
		params = &UpdateComponentTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateComponentType", params, optFns, c.addOperationUpdateComponentTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateComponentTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateComponentTypeInput struct {

	// The ID of the component type.
	//
	// This member is required.
	ComponentTypeId *string

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	// The component type name.
	ComponentTypeName *string

	// The description of the component type.
	Description *string

	// Specifies the component type that this component type extends.
	ExtendsFrom []string

	// An object that maps strings to the functions in the component type. Each string
	// in the mapping must be unique to this object.
	Functions map[string]types.FunctionRequest

	// A Boolean value that specifies whether an entity can have more than one
	// component of this type.
	IsSingleton *bool

	// An object that maps strings to the property definitions in the component type.
	// Each string in the mapping must be unique to this object.
	PropertyDefinitions map[string]types.PropertyDefinitionRequest

	// The property groups.
	PropertyGroups map[string]types.PropertyGroupRequest

	noSmithyDocumentSerde
}

type UpdateComponentTypeOutput struct {

	// The ARN of the component type.
	//
	// This member is required.
	Arn *string

	// The ID of the component type.
	//
	// This member is required.
	ComponentTypeId *string

	// The current state of the component type.
	//
	// This member is required.
	State types.State

	// The ID of the workspace that contains the component type.
	//
	// This member is required.
	WorkspaceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateComponentTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateComponentType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateComponentType{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateComponentTypeMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateComponentTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateComponentType(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateComponentTypeMiddleware struct {
}

func (*endpointPrefix_opUpdateComponentTypeMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateComponentTypeMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUpdateComponentTypeMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateComponentTypeMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateComponentType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "UpdateComponentType",
	}
}
