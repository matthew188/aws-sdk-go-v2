// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// If the DNS server for your self-managed domain uses a publicly addressable IP
// address, you must add a CIDR address block to correctly route traffic to and
// from your Microsoft AD on Amazon Web Services. AddIpRoutes adds this address
// block. You can also use AddIpRoutes to facilitate routing traffic that uses
// public IP ranges from your Microsoft AD on Amazon Web Services to a peer VPC.
// Before you call AddIpRoutes, ensure that all of the required permissions have
// been explicitly granted through a policy. For details about what permissions are
// required to run the AddIpRoutes operation, see Directory Service API
// Permissions: Actions, Resources, and Conditions Reference
// (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
func (c *Client) AddIpRoutes(ctx context.Context, params *AddIpRoutesInput, optFns ...func(*Options)) (*AddIpRoutesOutput, error) {
	if params == nil {
		params = &AddIpRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddIpRoutes", params, optFns, c.addOperationAddIpRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddIpRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddIpRoutesInput struct {

	// Identifier (ID) of the directory to which to add the address block.
	//
	// This member is required.
	DirectoryId *string

	// IP address blocks, using CIDR format, of the traffic to route. This is often the
	// IP address block of the DNS server used for your self-managed domain.
	//
	// This member is required.
	IpRoutes []types.IpRoute

	// If set to true, updates the inbound and outbound rules of the security group
	// that has the description: "Amazon Web Services created security group for
	// directory ID directory controllers." Following are the new rules: Inbound:
	//
	// *
	// Type: Custom UDP Rule, Protocol: UDP, Range: 88, Source: 0.0.0.0/0
	//
	// * Type:
	// Custom UDP Rule, Protocol: UDP, Range: 123, Source: 0.0.0.0/0
	//
	// * Type: Custom
	// UDP Rule, Protocol: UDP, Range: 138, Source: 0.0.0.0/0
	//
	// * Type: Custom UDP Rule,
	// Protocol: UDP, Range: 389, Source: 0.0.0.0/0
	//
	// * Type: Custom UDP Rule, Protocol:
	// UDP, Range: 464, Source: 0.0.0.0/0
	//
	// * Type: Custom UDP Rule, Protocol: UDP,
	// Range: 445, Source: 0.0.0.0/0
	//
	// * Type: Custom TCP Rule, Protocol: TCP, Range:
	// 88, Source: 0.0.0.0/0
	//
	// * Type: Custom TCP Rule, Protocol: TCP, Range: 135,
	// Source: 0.0.0.0/0
	//
	// * Type: Custom TCP Rule, Protocol: TCP, Range: 445, Source:
	// 0.0.0.0/0
	//
	// * Type: Custom TCP Rule, Protocol: TCP, Range: 464, Source:
	// 0.0.0.0/0
	//
	// * Type: Custom TCP Rule, Protocol: TCP, Range: 636, Source:
	// 0.0.0.0/0
	//
	// * Type: Custom TCP Rule, Protocol: TCP, Range: 1024-65535, Source:
	// 0.0.0.0/0
	//
	// * Type: Custom TCP Rule, Protocol: TCP, Range: 3268-33269, Source:
	// 0.0.0.0/0
	//
	// * Type: DNS (UDP), Protocol: UDP, Range: 53, Source: 0.0.0.0/0
	//
	// *
	// Type: DNS (TCP), Protocol: TCP, Range: 53, Source: 0.0.0.0/0
	//
	// * Type: LDAP,
	// Protocol: TCP, Range: 389, Source: 0.0.0.0/0
	//
	// * Type: All ICMP, Protocol: All,
	// Range: N/A, Source: 0.0.0.0/0
	//
	// Outbound:
	//
	// * Type: All traffic, Protocol: All,
	// Range: All, Destination: 0.0.0.0/0
	//
	// These security rules impact an internal
	// network interface that is not exposed publicly.
	UpdateSecurityGroupForDirectoryControllers bool

	noSmithyDocumentSerde
}

type AddIpRoutesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddIpRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddIpRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddIpRoutes{}, middleware.After)
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
	if err = addOpAddIpRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddIpRoutes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddIpRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "AddIpRoutes",
	}
}
