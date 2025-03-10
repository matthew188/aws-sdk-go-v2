// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get the information about an existing global endpoint. For more information
// about global endpoints, see Making applications Regional-fault tolerant with
// global endpoints and event replication
// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-global-endpoints.html)
// in the Amazon EventBridge User Guide..
func (c *Client) DescribeEndpoint(ctx context.Context, params *DescribeEndpointInput, optFns ...func(*Options)) (*DescribeEndpointOutput, error) {
	if params == nil {
		params = &DescribeEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpoint", params, optFns, c.addOperationDescribeEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointInput struct {

	// The name of the endpoint you want to get information about. For example,
	// "Name":"us-east-2-custom_bus_A-endpoint".
	//
	// This member is required.
	Name *string

	// The primary Region of the endpoint you want to get information about. For
	// example "HomeRegion": "us-east-1".
	HomeRegion *string

	noSmithyDocumentSerde
}

type DescribeEndpointOutput struct {

	// The ARN of the endpoint you asked for information about.
	Arn *string

	// The time the endpoint you asked for information about was created.
	CreationTime *time.Time

	// The description of the endpoint you asked for information about.
	Description *string

	// The ID of the endpoint you asked for information about.
	EndpointId *string

	// The URL of the endpoint you asked for information about.
	EndpointUrl *string

	// The event buses being used by the endpoint you asked for information about.
	EventBuses []types.EndpointEventBus

	// The last time the endpoint you asked for information about was modified.
	LastModifiedTime *time.Time

	// The name of the endpoint you asked for information about.
	Name *string

	// Whether replication is enabled or disabled for the endpoint you asked for
	// information about.
	ReplicationConfig *types.ReplicationConfig

	// The ARN of the role used by the endpoint you asked for information about.
	RoleArn *string

	// The routing configuration of the endpoint you asked for information about.
	RoutingConfig *types.RoutingConfig

	// The current state of the endpoint you asked for information about.
	State types.EndpointState

	// The reason the endpoint you asked for information about is in its current state.
	StateReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEndpoint{}, middleware.After)
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
	if err = addOpDescribeEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DescribeEndpoint",
	}
}
