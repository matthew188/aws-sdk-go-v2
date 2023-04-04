// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a Stage resource.
func (c *Client) GetStage(ctx context.Context, params *GetStageInput, optFns ...func(*Options)) (*GetStageOutput, error) {
	if params == nil {
		params = &GetStageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStage", params, optFns, c.addOperationGetStageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to get information about a Stage resource.
type GetStageInput struct {

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The name of the Stage resource to get information about.
	//
	// This member is required.
	StageName *string

	noSmithyDocumentSerde
}

// Represents a unique identifier for a version of a deployed RestApi that is
// callable by users.
type GetStageOutput struct {

	// Settings for logging access in this stage.
	AccessLogSettings *types.AccessLogSettings

	// Specifies whether a cache cluster is enabled for the stage.
	CacheClusterEnabled bool

	// The stage's cache capacity in GB. For more information about choosing a cache
	// size, see Enabling API caching to enhance responsiveness
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).
	CacheClusterSize types.CacheClusterSize

	// The status of the cache cluster for the stage, if enabled.
	CacheClusterStatus types.CacheClusterStatus

	// Settings for the canary deployment in this stage.
	CanarySettings *types.CanarySettings

	// The identifier of a client certificate for an API stage.
	ClientCertificateId *string

	// The timestamp when the stage was created.
	CreatedDate *time.Time

	// The identifier of the Deployment that the stage points to.
	DeploymentId *string

	// The stage's description.
	Description *string

	// The version of the associated API documentation.
	DocumentationVersion *string

	// The timestamp when the stage last updated.
	LastUpdatedDate *time.Time

	// A map that defines the method settings for a Stage resource. Keys (designated as
	// /{method_setting_key below) are method paths defined as
	// {resource_path}/{http_method} for an individual method override, or /\*/\* for
	// overriding all methods in the stage.
	MethodSettings map[string]types.MethodSetting

	// The name of the stage is the first path segment in the Uniform Resource
	// Identifier (URI) of a call to API Gateway. Stage names can only contain
	// alphanumeric characters, hyphens, and underscores. Maximum length is 128
	// characters.
	StageName *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string

	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled bool

	// A map that defines the stage variables for a Stage resource. Variable names can
	// have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	Variables map[string]string

	// The ARN of the WebAcl associated with the Stage.
	WebAclArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetStage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStage{}, middleware.After)
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
	if err = addOpGetStageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetStage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetStage",
	}
}
