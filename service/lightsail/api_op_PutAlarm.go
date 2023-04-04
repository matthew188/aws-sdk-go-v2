// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates an alarm, and associates it with the specified metric. An
// alarm is used to monitor a single metric for one of your resources. When a
// metric condition is met, the alarm can notify you by email, SMS text message,
// and a banner displayed on the Amazon Lightsail console. For more information,
// see Alarms in Amazon Lightsail
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-alarms).
// When this action creates an alarm, the alarm state is immediately set to
// INSUFFICIENT_DATA. The alarm is then evaluated and its state is set
// appropriately. Any actions associated with the new state are then executed. When
// you update an existing alarm, its state is left unchanged, but the update
// completely overwrites the previous configuration of the alarm. The alarm is then
// evaluated with the updated configuration.
func (c *Client) PutAlarm(ctx context.Context, params *PutAlarmInput, optFns ...func(*Options)) (*PutAlarmOutput, error) {
	if params == nil {
		params = &PutAlarmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAlarm", params, optFns, c.addOperationPutAlarmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAlarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAlarmInput struct {

	// The name for the alarm. Specify the name of an existing alarm to update, and
	// overwrite the previous configuration of the alarm.
	//
	// This member is required.
	AlarmName *string

	// The arithmetic operation to use when comparing the specified statistic to the
	// threshold. The specified statistic value is used as the first operand.
	//
	// This member is required.
	ComparisonOperator types.ComparisonOperator

	// The number of most recent periods over which data is compared to the specified
	// threshold. If you are setting an "M out of N" alarm, this value
	// (evaluationPeriods) is the N. If you are setting an alarm that requires that a
	// number of consecutive data points be breaching to trigger the alarm, this value
	// specifies the rolling period of time in which data points are evaluated. Each
	// evaluation period is five minutes long. For example, specify an evaluation
	// period of 24 to evaluate a metric over a rolling period of two hours. You can
	// specify a minimum valuation period of 1 (5 minutes), and a maximum evaluation
	// period of 288 (24 hours).
	//
	// This member is required.
	EvaluationPeriods *int32

	// The name of the metric to associate with the alarm. You can configure up to two
	// alarms per metric. The following metrics are available for each resource
	// type:
	//
	// * Instances: BurstCapacityPercentage, BurstCapacityTime, CPUUtilization,
	// NetworkIn, NetworkOut, StatusCheckFailed, StatusCheckFailed_Instance, and
	// StatusCheckFailed_System.
	//
	// * Load balancers: ClientTLSNegotiationErrorCount,
	// HealthyHostCount, UnhealthyHostCount, HTTPCode_LB_4XX_Count,
	// HTTPCode_LB_5XX_Count, HTTPCode_Instance_2XX_Count, HTTPCode_Instance_3XX_Count,
	// HTTPCode_Instance_4XX_Count, HTTPCode_Instance_5XX_Count, InstanceResponseTime,
	// RejectedConnectionCount, and RequestCount.
	//
	// * Relational databases:
	// CPUUtilization, DatabaseConnections, DiskQueueDepth, FreeStorageSpace,
	// NetworkReceiveThroughput, and NetworkTransmitThroughput.
	//
	// For more information
	// about these metrics, see Metrics available in Lightsail
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-resource-health-metrics#available-metrics).
	//
	// This member is required.
	MetricName types.MetricName

	// The name of the Lightsail resource that will be monitored. Instances, load
	// balancers, and relational databases are the only Lightsail resources that can
	// currently be monitored by alarms.
	//
	// This member is required.
	MonitoredResourceName *string

	// The value against which the specified statistic is compared.
	//
	// This member is required.
	Threshold *float64

	// The contact protocols to use for the alarm, such as Email, SMS (text messaging),
	// or both. A notification is sent via the specified contact protocol if
	// notifications are enabled for the alarm, and when the alarm is triggered. A
	// notification is not sent if a contact protocol is not specified, if the
	// specified contact protocol is not configured in the Amazon Web Services Region,
	// or if notifications are not enabled for the alarm using the notificationEnabled
	// paramater. Use the CreateContactMethod action to configure a contact protocol in
	// an Amazon Web Services Region.
	ContactProtocols []types.ContactProtocol

	// The number of data points that must be not within the specified threshold to
	// trigger the alarm. If you are setting an "M out of N" alarm, this value
	// (datapointsToAlarm) is the M.
	DatapointsToAlarm *int32

	// Indicates whether the alarm is enabled. Notifications are enabled by default if
	// you don't specify this parameter.
	NotificationEnabled *bool

	// The alarm states that trigger a notification. An alarm has the following
	// possible states:
	//
	// * ALARM - The metric is outside of the defined threshold.
	//
	// *
	// INSUFFICIENT_DATA - The alarm has just started, the metric is not available, or
	// not enough data is available for the metric to determine the alarm state.
	//
	// * OK
	// - The metric is within the defined threshold.
	//
	// When you specify a notification
	// trigger, the ALARM state must be specified. The INSUFFICIENT_DATA and OK states
	// can be specified in addition to the ALARM state.
	//
	// * If you specify OK as an
	// alarm trigger, a notification is sent when the alarm switches from an ALARM or
	// INSUFFICIENT_DATA alarm state to an OK state. This can be thought of as an all
	// clear alarm notification.
	//
	// * If you specify INSUFFICIENT_DATA as the alarm
	// trigger, a notification is sent when the alarm switches from an OK or ALARM
	// alarm state to an INSUFFICIENT_DATA state.
	//
	// The notification trigger defaults to
	// ALARM if you don't specify this parameter.
	NotificationTriggers []types.AlarmState

	// Sets how this alarm will handle missing data points. An alarm can treat missing
	// data in the following ways:
	//
	// * breaching - Assume the missing data is not within
	// the threshold. Missing data counts towards the number of times the metric is not
	// within the threshold.
	//
	// * notBreaching - Assume the missing data is within the
	// threshold. Missing data does not count towards the number of times the metric is
	// not within the threshold.
	//
	// * ignore - Ignore the missing data. Maintains the
	// current alarm state.
	//
	// * missing - Missing data is treated as missing.
	//
	// If
	// treatMissingData is not specified, the default behavior of missing is used.
	TreatMissingData types.TreatMissingData

	noSmithyDocumentSerde
}

type PutAlarmOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAlarmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAlarm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAlarm{}, middleware.After)
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
	if err = addOpPutAlarmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAlarm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAlarm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "PutAlarm",
	}
}
