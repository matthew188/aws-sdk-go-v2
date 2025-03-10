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
	"time"
)

// Retrieves information about an entity.
func (c *Client) GetEntity(ctx context.Context, params *GetEntityInput, optFns ...func(*Options)) (*GetEntityOutput, error) {
	if params == nil {
		params = &GetEntityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEntity", params, optFns, c.addOperationGetEntityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEntityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEntityInput struct {

	// The ID of the entity.
	//
	// This member is required.
	EntityId *string

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

type GetEntityOutput struct {

	// The ARN of the entity.
	//
	// This member is required.
	Arn *string

	// The date and time when the entity was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The ID of the entity.
	//
	// This member is required.
	EntityId *string

	// The name of the entity.
	//
	// This member is required.
	EntityName *string

	// A Boolean value that specifies whether the entity has associated child entities.
	//
	// This member is required.
	HasChildEntities *bool

	// The ID of the parent entity for this entity.
	//
	// This member is required.
	ParentEntityId *string

	// The current status of the entity.
	//
	// This member is required.
	Status *types.Status

	// The date and time when the entity was last updated.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// The ID of the workspace.
	//
	// This member is required.
	WorkspaceId *string

	// An object that maps strings to the components in the entity. Each string in the
	// mapping must be unique to this object.
	Components map[string]types.ComponentResponse

	// The description of the entity.
	Description *string

	// The syncSource of the sync job, if this entity was created by a sync job.
	SyncSource *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEntityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEntity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEntity{}, middleware.After)
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
	if err = addEndpointPrefix_opGetEntityMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetEntityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEntity(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetEntityMiddleware struct {
}

func (*endpointPrefix_opGetEntityMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetEntityMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addEndpointPrefix_opGetEntityMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetEntityMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetEntity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iottwinmaker",
		OperationName: "GetEntity",
	}
}
