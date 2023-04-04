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

// Modifies the placement attributes for a specified instance. You can do the
// following:
//
// * Modify the affinity between an instance and a Dedicated Host
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-overview.html).
// When affinity is set to host and the instance is not associated with a specific
// Dedicated Host, the next time the instance is launched, it is automatically
// associated with the host on which it lands. If the instance is restarted or
// rebooted, this relationship persists.
//
// * Change the Dedicated Host with which an
// instance is associated.
//
// * Change the instance tenancy of an instance.
//
// * Move
// an instance to or from a placement group
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).
//
// At
// least one attribute for affinity, host ID, tenancy, or placement group name must
// be specified in the request. Affinity and tenancy can be modified in the same
// request. To modify the host ID, tenancy, placement group, or partition for an
// instance, the instance must be in the stopped state.
func (c *Client) ModifyInstancePlacement(ctx context.Context, params *ModifyInstancePlacementInput, optFns ...func(*Options)) (*ModifyInstancePlacementOutput, error) {
	if params == nil {
		params = &ModifyInstancePlacementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstancePlacement", params, optFns, c.addOperationModifyInstancePlacementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstancePlacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstancePlacementInput struct {

	// The ID of the instance that you are modifying.
	//
	// This member is required.
	InstanceId *string

	// The affinity setting for the instance.
	Affinity types.Affinity

	// The Group Id of a placement group. You must specify the Placement Group Group Id
	// to launch an instance in a shared placement group.
	GroupId *string

	// The name of the placement group in which to place the instance. For spread
	// placement groups, the instance must have a tenancy of default. For cluster and
	// partition placement groups, the instance must have a tenancy of default or
	// dedicated. To remove an instance from a placement group, specify an empty string
	// ("").
	GroupName *string

	// The ID of the Dedicated Host with which to associate the instance.
	HostId *string

	// The ARN of the host resource group in which to place the instance.
	HostResourceGroupArn *string

	// The number of the partition in which to place the instance. Valid only if the
	// placement group strategy is set to partition.
	PartitionNumber *int32

	// The tenancy for the instance. For T3 instances, you can't change the tenancy
	// from dedicated to host, or from host to dedicated. Attempting to make one of
	// these unsupported tenancy changes results in the InvalidTenancy error code.
	Tenancy types.HostTenancy

	noSmithyDocumentSerde
}

type ModifyInstancePlacementOutput struct {

	// Is true if the request succeeds, and an error otherwise.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyInstancePlacementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstancePlacement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstancePlacement{}, middleware.After)
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
	if err = addOpModifyInstancePlacementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstancePlacement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstancePlacement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyInstancePlacement",
	}
}
