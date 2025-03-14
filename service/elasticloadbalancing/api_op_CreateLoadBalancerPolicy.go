// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a policy with the specified attributes for the specified load balancer.
// Policies are settings that are saved for your load balancer and that can be
// applied to the listener or the application server, depending on the policy type.
func (c *Client) CreateLoadBalancerPolicy(ctx context.Context, params *CreateLoadBalancerPolicyInput, optFns ...func(*Options)) (*CreateLoadBalancerPolicyOutput, error) {
	if params == nil {
		params = &CreateLoadBalancerPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoadBalancerPolicy", params, optFns, c.addOperationCreateLoadBalancerPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoadBalancerPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateLoadBalancerPolicy.
type CreateLoadBalancerPolicyInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The name of the load balancer policy to be created. This name must be unique
	// within the set of policies for this load balancer.
	//
	// This member is required.
	PolicyName *string

	// The name of the base policy type. To get the list of policy types, use
	// DescribeLoadBalancerPolicyTypes.
	//
	// This member is required.
	PolicyTypeName *string

	// The policy attributes.
	PolicyAttributes []types.PolicyAttribute

	noSmithyDocumentSerde
}

// Contains the output of CreateLoadBalancerPolicy.
type CreateLoadBalancerPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLoadBalancerPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoadBalancerPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoadBalancerPolicy{}, middleware.After)
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
	if err = addOpCreateLoadBalancerPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancerPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoadBalancerPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateLoadBalancerPolicy",
	}
}
