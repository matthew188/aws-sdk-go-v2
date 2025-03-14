// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a listener for the specified Application Load Balancer, Network Load
// Balancer, or Gateway Load Balancer. For more information, see the following:
//
// *
// Listeners for your Application Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-listeners.html)
//
// *
// Listeners for your Network Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-listeners.html)
//
// *
// Listeners for your Gateway Load Balancers
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/gateway/gateway-listeners.html)
//
// This
// operation is idempotent, which means that it completes at most one time. If you
// attempt to create multiple listeners with the same settings, each call succeeds.
func (c *Client) CreateListener(ctx context.Context, params *CreateListenerInput, optFns ...func(*Options)) (*CreateListenerOutput, error) {
	if params == nil {
		params = &CreateListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateListener", params, optFns, c.addOperationCreateListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateListenerInput struct {

	// The actions for the default rule.
	//
	// This member is required.
	DefaultActions []types.Action

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// This member is required.
	LoadBalancerArn *string

	// [TLS listeners] The name of the Application-Layer Protocol Negotiation (ALPN)
	// policy. You can specify one policy name. The following are the possible
	// values:
	//
	// * HTTP1Only
	//
	// * HTTP2Only
	//
	// * HTTP2Optional
	//
	// * HTTP2Preferred
	//
	// *
	// None
	//
	// For more information, see ALPN policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#alpn-policies)
	// in the Network Load Balancers Guide.
	AlpnPolicy []string

	// [HTTPS and TLS listeners] The default certificate for the listener. You must
	// provide exactly one certificate. Set CertificateArn to the certificate ARN but
	// do not set IsDefault.
	Certificates []types.Certificate

	// The port on which the load balancer is listening. You cannot specify a port for
	// a Gateway Load Balancer.
	Port *int32

	// The protocol for connections from clients to the load balancer. For Application
	// Load Balancers, the supported protocols are HTTP and HTTPS. For Network Load
	// Balancers, the supported protocols are TCP, TLS, UDP, and TCP_UDP. You can’t
	// specify the UDP or TCP_UDP protocol if dual-stack mode is enabled. You cannot
	// specify a protocol for a Gateway Load Balancer.
	Protocol types.ProtocolEnum

	// [HTTPS and TLS listeners] The security policy that defines which protocols and
	// ciphers are supported. For more information, see Security policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html#describe-ssl-policies)
	// in the Application Load Balancers Guide and Security policies
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html#describe-ssl-policies)
	// in the Network Load Balancers Guide.
	SslPolicy *string

	// The tags to assign to the listener.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateListenerOutput struct {

	// Information about the listener.
	Listeners []types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateListener{}, middleware.After)
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
	if err = addOpCreateListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateListener(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateListener",
	}
}
