// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures an Auto Scaling group to send notifications when specified events
// take place. Subscribers to the specified topic can have messages delivered to an
// endpoint such as a web server or an email address. This configuration overwrites
// any existing configuration. For more information, see Getting Amazon SNS
// notifications when your Auto Scaling group scales
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html)
// in the Amazon EC2 Auto Scaling User Guide. If you exceed your maximum limit of
// SNS topics, which is 10 per Auto Scaling group, the call fails.
func (c *Client) PutNotificationConfiguration(ctx context.Context, params *PutNotificationConfigurationInput, optFns ...func(*Options)) (*PutNotificationConfigurationOutput, error) {
	if params == nil {
		params = &PutNotificationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutNotificationConfiguration", params, optFns, c.addOperationPutNotificationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutNotificationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutNotificationConfigurationInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The type of event that causes the notification to be sent. To query the
	// notification types supported by Amazon EC2 Auto Scaling, call the
	// DescribeAutoScalingNotificationTypes API.
	//
	// This member is required.
	NotificationTypes []string

	// The Amazon Resource Name (ARN) of the Amazon SNS topic.
	//
	// This member is required.
	TopicARN *string

	noSmithyDocumentSerde
}

type PutNotificationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutNotificationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpPutNotificationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpPutNotificationConfiguration{}, middleware.After)
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
	if err = addOpPutNotificationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutNotificationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutNotificationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutNotificationConfiguration",
	}
}
