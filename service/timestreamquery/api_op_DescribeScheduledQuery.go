// Code generated by smithy-go-codegen DO NOT EDIT.

package timestreamquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	internalEndpointDiscovery "github.com/matthew188/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/matthew188/aws-sdk-go-v2/service/timestreamquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides detailed information about a scheduled query.
func (c *Client) DescribeScheduledQuery(ctx context.Context, params *DescribeScheduledQueryInput, optFns ...func(*Options)) (*DescribeScheduledQueryOutput, error) {
	if params == nil {
		params = &DescribeScheduledQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScheduledQuery", params, optFns, c.addOperationDescribeScheduledQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScheduledQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScheduledQueryInput struct {

	// The ARN of the scheduled query.
	//
	// This member is required.
	ScheduledQueryArn *string

	noSmithyDocumentSerde
}

type DescribeScheduledQueryOutput struct {

	// The scheduled query.
	//
	// This member is required.
	ScheduledQuery *types.ScheduledQueryDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeScheduledQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeScheduledQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeScheduledQuery{}, middleware.After)
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
	if err = addOpDescribeScheduledQueryDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addOpDescribeScheduledQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScheduledQuery(options.Region), middleware.Before); err != nil {
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

func addOpDescribeScheduledQueryDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
				opt.EndpointResolverUsedForDiscovery = o.EndpointDiscovery.EndpointResolverUsedForDiscovery
			},
		},
		DiscoverOperation:            c.fetchOpDescribeScheduledQueryDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    true,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpDescribeScheduledQueryDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*DescribeScheduledQueryInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)

	key := fmt.Sprintf("Timestream Query.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	endpoint, err := c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, key, opt)
	if err != nil {
		return internalEndpointDiscovery.WeightedAddress{}, err
	}

	weighted, ok := endpoint.GetValidAddress()
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("no valid endpoint address returned by the endpoint discovery api")
	}
	return weighted, nil
}

func newServiceMetadataMiddleware_opDescribeScheduledQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "timestream",
		OperationName: "DescribeScheduledQuery",
	}
}
