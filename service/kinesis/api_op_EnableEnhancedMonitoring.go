// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables enhanced Kinesis data stream monitoring for shard-level metrics. When
// invoking this API, it is recommended you use the StreamARN input parameter
// rather than the StreamName input parameter.
func (c *Client) EnableEnhancedMonitoring(ctx context.Context, params *EnableEnhancedMonitoringInput, optFns ...func(*Options)) (*EnableEnhancedMonitoringOutput, error) {
	if params == nil {
		params = &EnableEnhancedMonitoringInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableEnhancedMonitoring", params, optFns, c.addOperationEnableEnhancedMonitoringMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableEnhancedMonitoringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for EnableEnhancedMonitoring.
type EnableEnhancedMonitoringInput struct {

	// List of shard-level metrics to enable. The following are the valid shard-level
	// metrics. The value "ALL" enables every metric.
	//
	// * IncomingBytes
	//
	// *
	// IncomingRecords
	//
	// * OutgoingBytes
	//
	// * OutgoingRecords
	//
	// *
	// WriteProvisionedThroughputExceeded
	//
	// * ReadProvisionedThroughputExceeded
	//
	// *
	// IteratorAgeMilliseconds
	//
	// * ALL
	//
	// For more information, see Monitoring the Amazon
	// Kinesis Data Streams Service with Amazon CloudWatch
	// (https://docs.aws.amazon.com/kinesis/latest/dev/monitoring-with-cloudwatch.html)
	// in the Amazon Kinesis Data Streams Developer Guide.
	//
	// This member is required.
	ShardLevelMetrics []types.MetricsName

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream for which to enable enhanced monitoring.
	StreamName *string

	noSmithyDocumentSerde
}

// Represents the output for EnableEnhancedMonitoring and
// DisableEnhancedMonitoring.
type EnableEnhancedMonitoringOutput struct {

	// Represents the current state of the metrics that are in the enhanced state
	// before the operation.
	CurrentShardLevelMetrics []types.MetricsName

	// Represents the list of all the metrics that would be in the enhanced state after
	// the operation.
	DesiredShardLevelMetrics []types.MetricsName

	// The ARN of the stream.
	StreamARN *string

	// The name of the Kinesis data stream.
	StreamName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableEnhancedMonitoringMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableEnhancedMonitoring{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableEnhancedMonitoring{}, middleware.After)
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
	if err = addOpEnableEnhancedMonitoringValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableEnhancedMonitoring(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableEnhancedMonitoring(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "EnableEnhancedMonitoring",
	}
}
