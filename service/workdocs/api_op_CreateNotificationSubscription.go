// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configure Amazon WorkDocs to use Amazon SNS notifications. The endpoint receives
// a confirmation message, and must confirm the subscription. For more information,
// see Setting up notifications for an IAM user or role
// (https://docs.aws.amazon.com/workdocs/latest/developerguide/manage-notifications.html)
// in the Amazon WorkDocs Developer Guide.
func (c *Client) CreateNotificationSubscription(ctx context.Context, params *CreateNotificationSubscriptionInput, optFns ...func(*Options)) (*CreateNotificationSubscriptionOutput, error) {
	if params == nil {
		params = &CreateNotificationSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNotificationSubscription", params, optFns, c.addOperationCreateNotificationSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNotificationSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNotificationSubscriptionInput struct {

	// The endpoint to receive the notifications. If the protocol is HTTPS, the
	// endpoint is a URL that begins with https.
	//
	// This member is required.
	Endpoint *string

	// The ID of the organization.
	//
	// This member is required.
	OrganizationId *string

	// The protocol to use. The supported value is https, which delivers JSON-encoded
	// messages using HTTPS POST.
	//
	// This member is required.
	Protocol types.SubscriptionProtocolType

	// The notification type.
	//
	// This member is required.
	SubscriptionType types.SubscriptionType

	noSmithyDocumentSerde
}

type CreateNotificationSubscriptionOutput struct {

	// The subscription.
	Subscription *types.Subscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNotificationSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateNotificationSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNotificationSubscription{}, middleware.After)
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
	if err = addOpCreateNotificationSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNotificationSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNotificationSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "CreateNotificationSubscription",
	}
}
