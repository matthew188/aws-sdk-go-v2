// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the networking configuration to access the internal websites
// associated with the specified fleet.
//
// Deprecated: Amazon WorkLink is no longer supported. This will be removed in a
// future version of the SDK.
func (c *Client) DescribeCompanyNetworkConfiguration(ctx context.Context, params *DescribeCompanyNetworkConfigurationInput, optFns ...func(*Options)) (*DescribeCompanyNetworkConfigurationOutput, error) {
	if params == nil {
		params = &DescribeCompanyNetworkConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCompanyNetworkConfiguration", params, optFns, c.addOperationDescribeCompanyNetworkConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCompanyNetworkConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCompanyNetworkConfigurationInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	noSmithyDocumentSerde
}

type DescribeCompanyNetworkConfigurationOutput struct {

	// The security groups associated with access to the provided subnets.
	SecurityGroupIds []string

	// The subnets used for X-ENI connections from Amazon WorkLink rendering
	// containers.
	SubnetIds []string

	// The VPC with connectivity to associated websites.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCompanyNetworkConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCompanyNetworkConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCompanyNetworkConfiguration{}, middleware.After)
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
	if err = addOpDescribeCompanyNetworkConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCompanyNetworkConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCompanyNetworkConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DescribeCompanyNetworkConfiguration",
	}
}
