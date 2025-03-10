// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Given a list of verified identities (email addresses and/or domains), returns a
// structure describing identity notification attributes. This operation is
// throttled at one request per second and can only get notification attributes for
// up to 100 identities at a time. For more information about using notifications
// with Amazon SES, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
func (c *Client) GetIdentityNotificationAttributes(ctx context.Context, params *GetIdentityNotificationAttributesInput, optFns ...func(*Options)) (*GetIdentityNotificationAttributesOutput, error) {
	if params == nil {
		params = &GetIdentityNotificationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentityNotificationAttributes", params, optFns, c.addOperationGetIdentityNotificationAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentityNotificationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the notification attributes for a list of
// identities you verified with Amazon SES. For information about Amazon SES
// notifications, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/notifications.html).
type GetIdentityNotificationAttributesInput struct {

	// A list of one or more identities. You can specify an identity by using its name
	// or by using its Amazon Resource Name (ARN). Examples: user@example.com,
	// example.com, arn:aws:ses:us-east-1:123456789012:identity/example.com.
	//
	// This member is required.
	Identities []string

	noSmithyDocumentSerde
}

// Represents the notification attributes for a list of identities.
type GetIdentityNotificationAttributesOutput struct {

	// A map of Identity to IdentityNotificationAttributes.
	//
	// This member is required.
	NotificationAttributes map[string]types.IdentityNotificationAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdentityNotificationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetIdentityNotificationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetIdentityNotificationAttributes{}, middleware.After)
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
	if err = addOpGetIdentityNotificationAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityNotificationAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdentityNotificationAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetIdentityNotificationAttributes",
	}
}
