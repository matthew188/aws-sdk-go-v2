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

// Associate a virtual private cloud (VPC) subnet endpoint with your custom routing
// accelerator. The listener port range must be large enough to support the number
// of IP addresses that can be specified in your subnet. The number of ports
// required is: subnet size times the number of ports per destination EC2
// instances. For example, a subnet defined as /24 requires a listener port range
// of at least 255 ports. Note: You must have enough remaining listener ports
// available to map to the subnet ports, or the call will fail with a
// LimitExceededException. By default, all destinations in a subnet in a custom
// routing accelerator cannot receive traffic. To enable all destinations to
// receive traffic, or to specify individual port mappings that can receive
// traffic, see the  AllowCustomRoutingTraffic
// (https://docs.aws.amazon.com/global-accelerator/latest/api/API_AllowCustomRoutingTraffic.html)
// operation.
func (c *Client) AddCustomRoutingEndpoints(ctx context.Context, params *AddCustomRoutingEndpointsInput, optFns ...func(*Options)) (*AddCustomRoutingEndpointsOutput, error) {
	if params == nil {
		params = &AddCustomRoutingEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddCustomRoutingEndpoints", params, optFns, c.addOperationAddCustomRoutingEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddCustomRoutingEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddCustomRoutingEndpointsInput struct {

	// The list of endpoint objects to add to a custom routing accelerator.
	//
	// This member is required.
	EndpointConfigurations []types.CustomRoutingEndpointConfiguration

	// The Amazon Resource Name (ARN) of the endpoint group for the custom routing
	// endpoint.
	//
	// This member is required.
	EndpointGroupArn *string

	noSmithyDocumentSerde
}

type AddCustomRoutingEndpointsOutput struct {

	// The endpoint objects added to the custom routing accelerator.
	EndpointDescriptions []types.CustomRoutingEndpointDescription

	// The Amazon Resource Name (ARN) of the endpoint group for the custom routing
	// endpoint.
	EndpointGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddCustomRoutingEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddCustomRoutingEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddCustomRoutingEndpoints{}, middleware.After)
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
	if err = addOpAddCustomRoutingEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddCustomRoutingEndpoints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddCustomRoutingEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "AddCustomRoutingEndpoints",
	}
}
