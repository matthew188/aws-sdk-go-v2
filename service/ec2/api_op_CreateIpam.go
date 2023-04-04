// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an IPAM. Amazon VPC IP Address Manager (IPAM) is a VPC feature that you
// can use to automate your IP address management workflows including assigning,
// tracking, troubleshooting, and auditing IP addresses across Amazon Web Services
// Regions and accounts throughout your Amazon Web Services Organization. For more
// information, see Create an IPAM
// (https://docs.aws.amazon.com/vpc/latest/ipam/create-ipam.html) in the Amazon VPC
// IPAM User Guide.
func (c *Client) CreateIpam(ctx context.Context, params *CreateIpamInput, optFns ...func(*Options)) (*CreateIpamOutput, error) {
	if params == nil {
		params = &CreateIpamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIpam", params, optFns, c.addOperationCreateIpamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIpamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIpamInput struct {

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see Ensuring Idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string

	// A description for the IPAM.
	Description *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The operating Regions for the IPAM. Operating Regions are Amazon Web Services
	// Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only
	// discovers and monitors resources in the Amazon Web Services Regions you select
	// as operating Regions. For more information about operating Regions, see Create
	// an IPAM (https://docs.aws.amazon.com/vpc/latest/ipam/create-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	OperatingRegions []types.AddIpamOperatingRegion

	// The key/value combination of a tag assigned to the resource. Use the tag key in
	// the filter name and the tag value as the filter value. For example, to find all
	// resources that have a tag with the key Owner and the value TeamA, specify
	// tag:Owner for the filter name and TeamA for the filter value.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateIpamOutput struct {

	// Information about the IPAM created.
	Ipam *types.Ipam

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIpamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateIpam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateIpam{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateIpamMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIpam(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateIpam struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIpam) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIpam) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIpamInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIpamInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateIpamMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateIpam{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIpam(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateIpam",
	}
}
