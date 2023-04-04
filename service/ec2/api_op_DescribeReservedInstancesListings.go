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

// Describes your account's Reserved Instance listings in the Reserved Instance
// Marketplace. The Reserved Instance Marketplace matches sellers who want to
// resell Reserved Instance capacity that they no longer need with buyers who want
// to purchase additional capacity. Reserved Instances bought and sold through the
// Reserved Instance Marketplace work like any other Reserved Instances. As a
// seller, you choose to list some or all of your Reserved Instances, and you
// specify the upfront price to receive for them. Your Reserved Instances are then
// listed in the Reserved Instance Marketplace and are available for purchase. As a
// buyer, you specify the configuration of the Reserved Instance to purchase, and
// the Marketplace matches what you're searching for with what's available. The
// Marketplace first sells the lowest priced Reserved Instances to you, and
// continues to sell available Reserved Instance listings to you until your demand
// is met. You are charged based on the total price of all of the listings that you
// purchase. For more information, see Reserved Instance Marketplace
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html) in
// the Amazon EC2 User Guide.
func (c *Client) DescribeReservedInstancesListings(ctx context.Context, params *DescribeReservedInstancesListingsInput, optFns ...func(*Options)) (*DescribeReservedInstancesListingsOutput, error) {
	if params == nil {
		params = &DescribeReservedInstancesListingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedInstancesListings", params, optFns, c.addOperationDescribeReservedInstancesListingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedInstancesListingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeReservedInstancesListings.
type DescribeReservedInstancesListingsInput struct {

	// One or more filters.
	//
	// * reserved-instances-id - The ID of the Reserved
	// Instances.
	//
	// * reserved-instances-listing-id - The ID of the Reserved Instances
	// listing.
	//
	// * status - The status of the Reserved Instance listing (pending |
	// active | cancelled | closed).
	//
	// * status-message - The reason for the status.
	Filters []types.Filter

	// One or more Reserved Instance IDs.
	ReservedInstancesId *string

	// One or more Reserved Instance listing IDs.
	ReservedInstancesListingId *string

	noSmithyDocumentSerde
}

// Contains the output of DescribeReservedInstancesListings.
type DescribeReservedInstancesListingsOutput struct {

	// Information about the Reserved Instance listing.
	ReservedInstancesListings []types.ReservedInstancesListing

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedInstancesListingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeReservedInstancesListings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeReservedInstancesListings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedInstancesListings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReservedInstancesListings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeReservedInstancesListings",
	}
}
