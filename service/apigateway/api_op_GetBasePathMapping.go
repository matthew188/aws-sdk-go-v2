// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describe a BasePathMapping resource.
func (c *Client) GetBasePathMapping(ctx context.Context, params *GetBasePathMappingInput, optFns ...func(*Options)) (*GetBasePathMappingOutput, error) {
	if params == nil {
		params = &GetBasePathMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBasePathMapping", params, optFns, c.addOperationGetBasePathMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBasePathMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to describe a BasePathMapping resource.
type GetBasePathMappingInput struct {

	// The base path name that callers of the API must provide as part of the URL after
	// the domain name. This value must be unique for all of the mappings across a
	// single API. Specify '(none)' if you do not want callers to specify any base path
	// name after the domain name.
	//
	// This member is required.
	BasePath *string

	// The domain name of the BasePathMapping resource to be described.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

// Represents the base path that callers of the API must provide as part of the URL
// after the domain name.
type GetBasePathMappingOutput struct {

	// The base path name that callers of the API must provide as part of the URL after
	// the domain name.
	BasePath *string

	// The string identifier of the associated RestApi.
	RestApiId *string

	// The name of the associated stage.
	Stage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBasePathMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBasePathMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBasePathMapping{}, middleware.After)
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
	if err = addOpGetBasePathMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBasePathMapping(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBasePathMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetBasePathMapping",
	}
}
