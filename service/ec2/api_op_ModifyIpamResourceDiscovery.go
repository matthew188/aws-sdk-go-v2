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

// Modifies a resource discovery. A resource discovery is an IPAM component that
// enables IPAM to manage and monitor resources that belong to the owning account.
func (c *Client) ModifyIpamResourceDiscovery(ctx context.Context, params *ModifyIpamResourceDiscoveryInput, optFns ...func(*Options)) (*ModifyIpamResourceDiscoveryOutput, error) {
	if params == nil {
		params = &ModifyIpamResourceDiscoveryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyIpamResourceDiscovery", params, optFns, c.addOperationModifyIpamResourceDiscoveryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyIpamResourceDiscoveryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyIpamResourceDiscoveryInput struct {

	// A resource discovery ID.
	//
	// This member is required.
	IpamResourceDiscoveryId *string

	// Add operating Regions to the resource discovery. Operating Regions are Amazon
	// Web Services Regions where the IPAM is allowed to manage IP address CIDRs. IPAM
	// only discovers and monitors resources in the Amazon Web Services Regions you
	// select as operating Regions.
	AddOperatingRegions []types.AddIpamOperatingRegion

	// A resource discovery description.
	Description *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Remove operating Regions.
	RemoveOperatingRegions []types.RemoveIpamOperatingRegion

	noSmithyDocumentSerde
}

type ModifyIpamResourceDiscoveryOutput struct {

	// A resource discovery.
	IpamResourceDiscovery *types.IpamResourceDiscovery

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyIpamResourceDiscoveryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyIpamResourceDiscovery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyIpamResourceDiscovery{}, middleware.After)
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
	if err = addOpModifyIpamResourceDiscoveryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyIpamResourceDiscovery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyIpamResourceDiscovery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyIpamResourceDiscovery",
	}
}
