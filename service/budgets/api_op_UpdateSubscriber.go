// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a subscriber.
func (c *Client) UpdateSubscriber(ctx context.Context, params *UpdateSubscriberInput, optFns ...func(*Options)) (*UpdateSubscriberOutput, error) {
	if params == nil {
		params = &UpdateSubscriberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSubscriber", params, optFns, c.addOperationUpdateSubscriberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSubscriberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of UpdateSubscriber
type UpdateSubscriberInput struct {

	// The accountId that is associated with the budget whose subscriber you want to
	// update.
	//
	// This member is required.
	AccountId *string

	// The name of the budget whose subscriber you want to update.
	//
	// This member is required.
	BudgetName *string

	// The updated subscriber that is associated with a budget notification.
	//
	// This member is required.
	NewSubscriber *types.Subscriber

	// The notification whose subscriber you want to update.
	//
	// This member is required.
	Notification *types.Notification

	// The previous subscriber that is associated with a budget notification.
	//
	// This member is required.
	OldSubscriber *types.Subscriber

	noSmithyDocumentSerde
}

// Response of UpdateSubscriber
type UpdateSubscriberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSubscriberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSubscriber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSubscriber{}, middleware.After)
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
	if err = addOpUpdateSubscriberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSubscriber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSubscriber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "UpdateSubscriber",
	}
}
