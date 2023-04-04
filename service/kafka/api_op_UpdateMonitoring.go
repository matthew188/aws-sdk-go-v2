// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the monitoring settings for the cluster. You can use this operation to
// specify which Apache Kafka metrics you want Amazon MSK to send to Amazon
// CloudWatch. You can also specify settings for open monitoring with Prometheus.
func (c *Client) UpdateMonitoring(ctx context.Context, params *UpdateMonitoringInput, optFns ...func(*Options)) (*UpdateMonitoringOutput, error) {
	if params == nil {
		params = &UpdateMonitoringInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMonitoring", params, optFns, c.addOperationUpdateMonitoringMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMonitoringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request body for UpdateMonitoring.
type UpdateMonitoringInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	//
	// This member is required.
	ClusterArn *string

	// The version of the MSK cluster to update. Cluster versions aren't simple
	// numbers. You can describe an MSK cluster to find its version. When this update
	// operation is successful, it generates a new cluster version.
	//
	// This member is required.
	CurrentVersion *string

	// Specifies which Apache Kafka metrics Amazon MSK gathers and sends to Amazon
	// CloudWatch for this cluster.
	EnhancedMonitoring types.EnhancedMonitoring

	LoggingInfo *types.LoggingInfo

	// The settings for open monitoring.
	OpenMonitoring *types.OpenMonitoringInfo

	noSmithyDocumentSerde
}

type UpdateMonitoringOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMonitoringMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMonitoring{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMonitoring{}, middleware.After)
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
	if err = addOpUpdateMonitoringValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMonitoring(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMonitoring(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "UpdateMonitoring",
	}
}
