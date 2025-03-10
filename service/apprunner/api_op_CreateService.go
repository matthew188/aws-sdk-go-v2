// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an App Runner service. After the service is created, the action also
// automatically starts a deployment. This is an asynchronous operation. On a
// successful call, you can use the returned OperationId and the ListOperations
// (https://docs.aws.amazon.com/apprunner/latest/api/API_ListOperations.html) call
// to track the operation's progress.
func (c *Client) CreateService(ctx context.Context, params *CreateServiceInput, optFns ...func(*Options)) (*CreateServiceOutput, error) {
	if params == nil {
		params = &CreateServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateService", params, optFns, c.addOperationCreateServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateServiceInput struct {

	// A name for the App Runner service. It must be unique across all the running App
	// Runner services in your Amazon Web Services account in the Amazon Web Services
	// Region.
	//
	// This member is required.
	ServiceName *string

	// The source to deploy to the App Runner service. It can be a code or an image
	// repository.
	//
	// This member is required.
	SourceConfiguration *types.SourceConfiguration

	// The Amazon Resource Name (ARN) of an App Runner automatic scaling configuration
	// resource that you want to associate with your service. If not provided, App
	// Runner associates the latest revision of a default auto scaling configuration.
	// Specify an ARN with a name and a revision number to associate that revision. For
	// example:
	// arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability/3
	// Specify just the name to associate the latest revision. For example:
	// arn:aws:apprunner:us-east-1:123456789012:autoscalingconfiguration/high-availability
	AutoScalingConfigurationArn *string

	// An optional custom encryption key that App Runner uses to encrypt the copy of
	// your source repository that it maintains and your service logs. By default, App
	// Runner uses an Amazon Web Services managed key.
	EncryptionConfiguration *types.EncryptionConfiguration

	// The settings for the health check that App Runner performs to monitor the health
	// of the App Runner service.
	HealthCheckConfiguration *types.HealthCheckConfiguration

	// The runtime configuration of instances (scaling units) of your service.
	InstanceConfiguration *types.InstanceConfiguration

	// Configuration settings related to network traffic of the web application that
	// the App Runner service runs.
	NetworkConfiguration *types.NetworkConfiguration

	// The observability configuration of your service.
	ObservabilityConfiguration *types.ServiceObservabilityConfiguration

	// An optional list of metadata items that you can associate with the App Runner
	// service resource. A tag is a key-value pair.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateServiceOutput struct {

	// The unique ID of the asynchronous operation that this request started. You can
	// use it combined with the ListOperations
	// (https://docs.aws.amazon.com/apprunner/latest/api/API_ListOperations.html) call
	// to track the operation's progress.
	//
	// This member is required.
	OperationId *string

	// A description of the App Runner service that's created by this request.
	//
	// This member is required.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateService{}, middleware.After)
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
	if err = addOpCreateServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "CreateService",
	}
}
