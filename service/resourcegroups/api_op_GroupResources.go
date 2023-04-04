// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds the specified resources to the specified group. You can use this operation
// with only resource groups that are configured with the following types:
//
// *
// AWS::EC2::HostManagement
//
// * AWS::EC2::CapacityReservationPool
//
// Other resource
// group type and resource types aren't currently supported by this operation.
// Minimum permissions To run this command, you must have the following
// permissions:
//
// * resource-groups:GroupResources
func (c *Client) GroupResources(ctx context.Context, params *GroupResourcesInput, optFns ...func(*Options)) (*GroupResourcesOutput, error) {
	if params == nil {
		params = &GroupResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GroupResources", params, optFns, c.addOperationGroupResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GroupResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GroupResourcesInput struct {

	// The name or the ARN of the resource group to add resources to.
	//
	// This member is required.
	Group *string

	// The list of ARNs of the resources to be added to the group.
	//
	// This member is required.
	ResourceArns []string

	noSmithyDocumentSerde
}

type GroupResourcesOutput struct {

	// A list of ARNs of any resources that this operation failed to add to the group.
	Failed []types.FailedResource

	// A list of ARNs of any resources that this operation is still in the process
	// adding to the group. These pending additions continue asynchronously. You can
	// check the status of pending additions by using the ListGroupResources operation,
	// and checking the Resources array in the response and the Status field of each
	// object in that array.
	Pending []types.PendingResource

	// A list of ARNs of the resources that this operation successfully added to the
	// group.
	Succeeded []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGroupResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGroupResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGroupResources{}, middleware.After)
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
	if err = addOpGroupResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGroupResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGroupResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "GroupResources",
	}
}
