// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an Amazon EKS managed node group configuration. Your node group
// continues to function during the update. The response output includes an update
// ID that you can use to track the status of your node group update with the
// DescribeUpdate API operation. Currently you can update the Kubernetes labels for
// a node group or the scaling configuration.
func (c *Client) UpdateNodegroupConfig(ctx context.Context, params *UpdateNodegroupConfigInput, optFns ...func(*Options)) (*UpdateNodegroupConfigOutput, error) {
	if params == nil {
		params = &UpdateNodegroupConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNodegroupConfig", params, optFns, c.addOperationUpdateNodegroupConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNodegroupConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNodegroupConfigInput struct {

	// The name of the Amazon EKS cluster that the managed node group resides in.
	//
	// This member is required.
	ClusterName *string

	// The name of the managed node group to update.
	//
	// This member is required.
	NodegroupName *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The Kubernetes labels to be applied to the nodes in the node group after the
	// update.
	Labels *types.UpdateLabelsPayload

	// The scaling configuration details for the Auto Scaling group after the update.
	ScalingConfig *types.NodegroupScalingConfig

	// The Kubernetes taints to be applied to the nodes in the node group after the
	// update. For more information, see Node taints on managed node groups
	// (https://docs.aws.amazon.com/eks/latest/userguide/node-taints-managed-node-groups.html).
	Taints *types.UpdateTaintsPayload

	// The node group update configuration.
	UpdateConfig *types.NodegroupUpdateConfig

	noSmithyDocumentSerde
}

type UpdateNodegroupConfigOutput struct {

	// An object representing an asynchronous update.
	Update *types.Update

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNodegroupConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNodegroupConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNodegroupConfig{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateNodegroupConfigMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateNodegroupConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNodegroupConfig(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateNodegroupConfig struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateNodegroupConfig) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateNodegroupConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateNodegroupConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateNodegroupConfigInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateNodegroupConfigMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateNodegroupConfig{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateNodegroupConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "UpdateNodegroupConfig",
	}
}
