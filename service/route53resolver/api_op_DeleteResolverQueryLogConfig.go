// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a query logging configuration. When you delete a configuration, Resolver
// stops logging DNS queries for all of the Amazon VPCs that are associated with
// the configuration. This also applies if the query logging configuration is
// shared with other Amazon Web Services accounts, and the other accounts have
// associated VPCs with the shared configuration. Before you can delete a query
// logging configuration, you must first disassociate all VPCs from the
// configuration. See DisassociateResolverQueryLogConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverQueryLogConfig.html).
// If you used Resource Access Manager (RAM) to share a query logging configuration
// with other accounts, you must stop sharing the configuration before you can
// delete a configuration. The accounts that you shared the configuration with can
// first disassociate VPCs that they associated with the configuration, but that's
// not necessary. If you stop sharing the configuration, those VPCs are
// automatically disassociated from the configuration.
func (c *Client) DeleteResolverQueryLogConfig(ctx context.Context, params *DeleteResolverQueryLogConfigInput, optFns ...func(*Options)) (*DeleteResolverQueryLogConfigOutput, error) {
	if params == nil {
		params = &DeleteResolverQueryLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResolverQueryLogConfig", params, optFns, c.addOperationDeleteResolverQueryLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResolverQueryLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResolverQueryLogConfigInput struct {

	// The ID of the query logging configuration that you want to delete.
	//
	// This member is required.
	ResolverQueryLogConfigId *string

	noSmithyDocumentSerde
}

type DeleteResolverQueryLogConfigOutput struct {

	// Information about the query logging configuration that you deleted, including
	// the status of the request.
	ResolverQueryLogConfig *types.ResolverQueryLogConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteResolverQueryLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteResolverQueryLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteResolverQueryLogConfig{}, middleware.After)
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
	if err = addOpDeleteResolverQueryLogConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResolverQueryLogConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteResolverQueryLogConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "DeleteResolverQueryLogConfig",
	}
}
