// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a domain name.
func (c *Client) UpdateDomainName(ctx context.Context, params *UpdateDomainNameInput, optFns ...func(*Options)) (*UpdateDomainNameOutput, error) {
	if params == nil {
		params = &UpdateDomainNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainName", params, optFns, c.addOperationUpdateDomainNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates a DomainName.
type UpdateDomainNameInput struct {

	// The domain name.
	//
	// This member is required.
	DomainName *string

	// The domain name configurations.
	DomainNameConfigurations []types.DomainNameConfiguration

	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication *types.MutualTlsAuthenticationInput

	noSmithyDocumentSerde
}

type UpdateDomainNameOutput struct {

	// The API mapping selection expression.
	ApiMappingSelectionExpression *string

	// The name of the DomainName resource.
	DomainName *string

	// The domain name configurations.
	DomainNameConfigurations []types.DomainNameConfiguration

	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication *types.MutualTlsAuthentication

	// The collection of tags associated with a domain name.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDomainNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDomainName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDomainName{}, middleware.After)
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
	if err = addOpUpdateDomainNameValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainName(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomainName(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateDomainName",
	}
}
