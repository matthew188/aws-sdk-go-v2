// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/m2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the details of a specific application.
func (c *Client) GetApplication(ctx context.Context, params *GetApplicationInput, optFns ...func(*Options)) (*GetApplicationOutput, error) {
	if params == nil {
		params = &GetApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetApplication", params, optFns, c.addOperationGetApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetApplicationInput struct {

	// The identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	noSmithyDocumentSerde
}

type GetApplicationOutput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationArn *string

	// The identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The timestamp when this application was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The type of the target platform for the application.
	//
	// This member is required.
	EngineType types.EngineType

	// The latest version of the application.
	//
	// This member is required.
	LatestVersion *types.ApplicationVersionSummary

	// The unique identifier of the application.
	//
	// This member is required.
	Name *string

	// The status of the application.
	//
	// This member is required.
	Status types.ApplicationLifecycle

	// The version of the application that is deployed.
	DeployedVersion *types.DeployedVersionSummary

	// The description of the application.
	Description *string

	// The identifier of the runtime environment where you want to deploy the
	// application.
	EnvironmentId *string

	// The identifier of a customer managed key.
	KmsKeyId *string

	// The timestamp when you last started the application. Null until the application
	// runs for the first time.
	LastStartTime *time.Time

	// The Amazon Resource Name (ARN) for the network load balancer listener created in
	// your Amazon Web Services account. Amazon Web Services Mainframe Modernization
	// creates this listener for you the first time you deploy an application.
	ListenerArns []string

	// The port associated with the network load balancer listener created in your
	// Amazon Web Services account.
	ListenerPorts []int32

	// The public DNS name of the load balancer created in your Amazon Web Services
	// account.
	LoadBalancerDnsName *string

	// The list of log summaries. Each log summary includes the log type as well as the
	// log group identifier. These are CloudWatch logs. Amazon Web Services Mainframe
	// Modernization pushes the application log to CloudWatch under the customer's
	// account.
	LogGroups []types.LogGroupSummary

	// The reason for the reported status.
	StatusReason *string

	// A list of tags associated with the application.
	Tags map[string]string

	// Returns the Amazon Resource Names (ARNs) of the target groups that are attached
	// to the network load balancer.
	TargetGroupArns []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetApplication{}, middleware.After)
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
	if err = addOpGetApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "GetApplication",
	}
}
