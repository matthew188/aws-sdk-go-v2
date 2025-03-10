// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the allocations from the specified customer-owned address pool.
func (c *Client) GetCoipPoolUsage(ctx context.Context, params *GetCoipPoolUsageInput, optFns ...func(*Options)) (*GetCoipPoolUsageOutput, error) {
	if params == nil {
		params = &GetCoipPoolUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCoipPoolUsage", params, optFns, c.addOperationGetCoipPoolUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCoipPoolUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCoipPoolUsageInput struct {

	// The ID of the address pool.
	//
	// This member is required.
	PoolId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * coip-address-usage.allocation-id - The allocation ID of
	// the address.
	//
	// * coip-address-usage.aws-account-id - The ID of the Amazon Web
	// Services account that is using the customer-owned IP address.
	//
	// *
	// coip-address-usage.aws-service - The Amazon Web Services service that is using
	// the customer-owned IP address.
	//
	// * coip-address-usage.co-ip - The customer-owned
	// IP address.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCoipPoolUsageOutput struct {

	// Information about the address usage.
	CoipAddressUsages []types.CoipAddressUsage

	// The ID of the customer-owned address pool.
	CoipPoolId *string

	// The ID of the local gateway route table.
	LocalGatewayRouteTableId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCoipPoolUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpGetCoipPoolUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpGetCoipPoolUsage{}, middleware.After)
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
	if err = addOpGetCoipPoolUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCoipPoolUsage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCoipPoolUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "GetCoipPoolUsage",
	}
}
