// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/codestarnotifications/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a specified notification rule.
func (c *Client) DescribeNotificationRule(ctx context.Context, params *DescribeNotificationRuleInput, optFns ...func(*Options)) (*DescribeNotificationRuleOutput, error) {
	if params == nil {
		params = &DescribeNotificationRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNotificationRule", params, optFns, c.addOperationDescribeNotificationRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNotificationRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotificationRuleInput struct {

	// The Amazon Resource Name (ARN) of the notification rule.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

type DescribeNotificationRuleOutput struct {

	// The Amazon Resource Name (ARN) of the notification rule.
	//
	// This member is required.
	Arn *string

	// The name or email alias of the person who created the notification rule.
	CreatedBy *string

	// The date and time the notification rule was created, in timestamp format.
	CreatedTimestamp *time.Time

	// The level of detail included in the notifications for this resource. BASIC will
	// include only the contents of the event as it would appear in Amazon CloudWatch.
	// FULL will include any supplemental information provided by AWS CodeStar
	// Notifications and/or the service for the resource for which the notification is
	// created.
	DetailType types.DetailType

	// A list of the event types associated with the notification rule.
	EventTypes []types.EventTypeSummary

	// The date and time the notification rule was most recently updated, in timestamp
	// format.
	LastModifiedTimestamp *time.Time

	// The name of the notification rule.
	Name *string

	// The Amazon Resource Name (ARN) of the resource associated with the notification
	// rule.
	Resource *string

	// The status of the notification rule. Valid statuses are on (sending
	// notifications) or off (not sending notifications).
	Status types.NotificationRuleStatus

	// The tags associated with the notification rule.
	Tags map[string]string

	// A list of the Chatbot topics and Chatbot clients associated with the
	// notification rule.
	Targets []types.TargetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNotificationRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeNotificationRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeNotificationRule{}, middleware.After)
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
	if err = addOpDescribeNotificationRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotificationRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeNotificationRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "DescribeNotificationRule",
	}
}
