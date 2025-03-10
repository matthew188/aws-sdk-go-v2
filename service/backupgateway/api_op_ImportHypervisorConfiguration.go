// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/backupgateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Connect to a hypervisor by importing its configuration.
func (c *Client) ImportHypervisorConfiguration(ctx context.Context, params *ImportHypervisorConfigurationInput, optFns ...func(*Options)) (*ImportHypervisorConfigurationOutput, error) {
	if params == nil {
		params = &ImportHypervisorConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportHypervisorConfiguration", params, optFns, c.addOperationImportHypervisorConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportHypervisorConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportHypervisorConfigurationInput struct {

	// The server host of the hypervisor. This can be either an IP address or a
	// fully-qualified domain name (FQDN).
	//
	// This member is required.
	Host *string

	// The name of the hypervisor.
	//
	// This member is required.
	Name *string

	// The Key Management Service for the hypervisor.
	KmsKeyArn *string

	// The password for the hypervisor.
	Password *string

	// The tags of the hypervisor configuration to import.
	Tags []types.Tag

	// The username for the hypervisor.
	Username *string

	noSmithyDocumentSerde
}

type ImportHypervisorConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the hypervisor you disassociated.
	HypervisorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportHypervisorConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpImportHypervisorConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpImportHypervisorConfiguration{}, middleware.After)
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
	if err = addOpImportHypervisorConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportHypervisorConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportHypervisorConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup-gateway",
		OperationName: "ImportHypervisorConfiguration",
	}
}
