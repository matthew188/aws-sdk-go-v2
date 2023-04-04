// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an Api resource.
func (c *Client) UpdateApi(ctx context.Context, params *UpdateApiInput, optFns ...func(*Options)) (*UpdateApiOutput, error) {
	if params == nil {
		params = &UpdateApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApi", params, optFns, c.addOperationUpdateApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates an Api.
type UpdateApiInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// An API key selection expression. Supported only for WebSocket APIs. See API Key
	// Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression *string

	// A CORS configuration. Supported only for HTTP APIs.
	CorsConfiguration *types.Cors

	// This property is part of quick create. It specifies the credentials required for
	// the integration, if any. For a Lambda integration, three options are available.
	// To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource
	// Name (ARN). To require that the caller's identity be passed through from the
	// request, specify arn:aws:iam::*:user/*. To use resource-based permissions on
	// supported AWS services, don't specify this parameter. Currently, this property
	// is not used for HTTP integrations. If provided, this value replaces the
	// credentials associated with the quick create integration. Supported only for
	// HTTP APIs.
	CredentialsArn *string

	// The description of the API.
	Description *string

	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint bool

	// Avoid validating models when creating a deployment. Supported only for WebSocket
	// APIs.
	DisableSchemaValidation bool

	// The name of the API.
	Name *string

	// This property is part of quick create. If not specified, the route created using
	// quick create is kept. Otherwise, this value replaces the route key of the quick
	// create route. Additional routes may still be added after the API is updated.
	// Supported only for HTTP APIs.
	RouteKey *string

	// The route selection expression for the API. For HTTP APIs, the
	// routeSelectionExpression must be ${request.method} ${request.path}. If not
	// provided, this will be the default for HTTP APIs. This property is required for
	// WebSocket APIs.
	RouteSelectionExpression *string

	// This property is part of quick create. For HTTP integrations, specify a fully
	// qualified URL. For Lambda integrations, specify a function ARN. The type of the
	// integration will be HTTP_PROXY or AWS_PROXY, respectively. The value provided
	// updates the integration URI and integration type. You can update a quick-created
	// target, but you can't remove it from an API. Supported only for HTTP APIs.
	Target *string

	// A version identifier for the API.
	Version *string

	noSmithyDocumentSerde
}

type UpdateApiOutput struct {

	// The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com. The
	// stage name is typically appended to this URI to form a complete path to a
	// deployed API stage.
	ApiEndpoint *string

	// Specifies whether an API is managed by API Gateway. You can't update or delete a
	// managed API by using API Gateway. A managed API can be deleted only through the
	// tooling or service that created it.
	ApiGatewayManaged bool

	// The API ID.
	ApiId *string

	// An API key selection expression. Supported only for WebSocket APIs. See API Key
	// Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).
	ApiKeySelectionExpression *string

	// A CORS configuration. Supported only for HTTP APIs.
	CorsConfiguration *types.Cors

	// The timestamp when the API was created.
	CreatedDate *time.Time

	// The description of the API.
	Description *string

	// Specifies whether clients can invoke your API by using the default execute-api
	// endpoint. By default, clients can invoke your API with the default
	// https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default
	// endpoint.
	DisableExecuteApiEndpoint bool

	// Avoid validating models when creating a deployment. Supported only for WebSocket
	// APIs.
	DisableSchemaValidation bool

	// The validation information during API import. This may include particular
	// properties of your OpenAPI definition which are ignored during import. Supported
	// only for HTTP APIs.
	ImportInfo []string

	// The name of the API.
	Name *string

	// The API protocol.
	ProtocolType types.ProtocolType

	// The route selection expression for the API. For HTTP APIs, the
	// routeSelectionExpression must be ${request.method} ${request.path}. If not
	// provided, this will be the default for HTTP APIs. This property is required for
	// WebSocket APIs.
	RouteSelectionExpression *string

	// A collection of tags associated with the API.
	Tags map[string]string

	// A version identifier for the API.
	Version *string

	// The warning messages reported when failonwarnings is turned on during API
	// import.
	Warnings []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateApi{}, middleware.After)
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
	if err = addOpUpdateApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateApi",
	}
}
