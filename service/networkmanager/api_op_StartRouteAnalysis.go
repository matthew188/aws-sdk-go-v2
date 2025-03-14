// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts analyzing the routing path between the specified source and destination.
// For more information, see Route Analyzer
// (https://docs.aws.amazon.com/vpc/latest/tgw/route-analyzer.html).
func (c *Client) StartRouteAnalysis(ctx context.Context, params *StartRouteAnalysisInput, optFns ...func(*Options)) (*StartRouteAnalysisOutput, error) {
	if params == nil {
		params = &StartRouteAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartRouteAnalysis", params, optFns, c.addOperationStartRouteAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartRouteAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRouteAnalysisInput struct {

	// The destination.
	//
	// This member is required.
	Destination *types.RouteAnalysisEndpointOptionsSpecification

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The source from which traffic originates.
	//
	// This member is required.
	Source *types.RouteAnalysisEndpointOptionsSpecification

	// Indicates whether to analyze the return path. The default is false.
	IncludeReturnPath bool

	// Indicates whether to include the location of middlebox appliances in the route
	// analysis. The default is false.
	UseMiddleboxes bool

	noSmithyDocumentSerde
}

type StartRouteAnalysisOutput struct {

	// The route analysis.
	RouteAnalysis *types.RouteAnalysis

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartRouteAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartRouteAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartRouteAnalysis{}, middleware.After)
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
	if err = addOpStartRouteAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartRouteAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartRouteAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "StartRouteAnalysis",
	}
}
