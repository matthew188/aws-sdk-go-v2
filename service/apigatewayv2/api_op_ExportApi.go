// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func (c *Client) ExportApi(ctx context.Context, params *ExportApiInput, optFns ...func(*Options)) (*ExportApiOutput, error) {
	if params == nil {
		params = &ExportApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportApi", params, optFns, c.addOperationExportApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportApiInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The output type of the exported definition file. Valid values are JSON and YAML.
	//
	// This member is required.
	OutputType *string

	// The version of the API specification to use. OAS30, for OpenAPI 3.0, is the only
	// supported value.
	//
	// This member is required.
	Specification *string

	// The version of the API Gateway export algorithm. API Gateway uses the latest
	// version by default. Currently, the only supported version is 1.0.
	ExportVersion *string

	// Specifies whether to include API Gateway extensions
	// (https://docs.aws.amazon.com//apigateway/latest/developerguide/api-gateway-swagger-extensions.html)
	// in the exported API definition. API Gateway extensions are included by default.
	IncludeExtensions bool

	// The name of the API stage to export. If you don't specify this property, a
	// representation of the latest API configuration is exported.
	StageName *string

	noSmithyDocumentSerde
}

type ExportApiOutput struct {

	// Represents an exported definition of an API in a particular output format, for
	// example, YAML. The API is serialized to the requested specification, for
	// example, OpenAPI 3.0.
	Body []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportApi{}, middleware.After)
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
	if err = addOpExportApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportApi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "ExportApi",
	}
}
