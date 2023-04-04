// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches one or more Lightsail instances to a load balancer. After some time,
// the instances are attached to the load balancer and the health check status is
// available. The attach instances to load balancer operation supports tag-based
// access control via resource tags applied to the resource identified by load
// balancer name. For more information, see the Lightsail Developer Guide
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) AttachInstancesToLoadBalancer(ctx context.Context, params *AttachInstancesToLoadBalancerInput, optFns ...func(*Options)) (*AttachInstancesToLoadBalancerOutput, error) {
	if params == nil {
		params = &AttachInstancesToLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachInstancesToLoadBalancer", params, optFns, c.addOperationAttachInstancesToLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachInstancesToLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachInstancesToLoadBalancerInput struct {

	// An array of strings representing the instance name(s) you want to attach to your
	// load balancer. An instance must be running before you can attach it to your load
	// balancer. There are no additional limits on the number of instances you can
	// attach to your load balancer, aside from the limit of Lightsail instances you
	// can create in your account (20).
	//
	// This member is required.
	InstanceNames []string

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	noSmithyDocumentSerde
}

type AttachInstancesToLoadBalancerOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachInstancesToLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachInstancesToLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachInstancesToLoadBalancer{}, middleware.After)
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
	if err = addOpAttachInstancesToLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachInstancesToLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachInstancesToLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "AttachInstancesToLoadBalancer",
	}
}
