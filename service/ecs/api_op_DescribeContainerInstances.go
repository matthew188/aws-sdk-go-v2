// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ecs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more container instances. Returns metadata about each container
// instance requested.
func (c *Client) DescribeContainerInstances(ctx context.Context, params *DescribeContainerInstancesInput, optFns ...func(*Options)) (*DescribeContainerInstancesOutput, error) {
	if params == nil {
		params = &DescribeContainerInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeContainerInstances", params, optFns, c.addOperationDescribeContainerInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeContainerInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeContainerInstancesInput struct {

	// A list of up to 100 container instance IDs or full Amazon Resource Name (ARN)
	// entries.
	//
	// This member is required.
	ContainerInstances []string

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// container instances to describe. If you do not specify a cluster, the default
	// cluster is assumed. This parameter is required if the container instance or
	// container instances you are describing were launched in any cluster other than
	// the default cluster.
	Cluster *string

	// Specifies whether you want to see the resource tags for the container instance.
	// If TAGS is specified, the tags are included in the response. If
	// CONTAINER_INSTANCE_HEALTH is specified, the container instance health is
	// included in the response. If this field is omitted, tags and container instance
	// health status aren't included in the response.
	Include []types.ContainerInstanceField

	noSmithyDocumentSerde
}

type DescribeContainerInstancesOutput struct {

	// The list of container instances.
	ContainerInstances []types.ContainerInstance

	// Any failures associated with the call.
	Failures []types.Failure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeContainerInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeContainerInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeContainerInstances{}, middleware.After)
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
	if err = addOpDescribeContainerInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeContainerInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeContainerInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DescribeContainerInstances",
	}
}
