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

// Deletes a subscriber. Deleting the last subscriber to a notification also
// deletes the notification.
func (c *Client) DeleteSubscriber(ctx context.Context, params *DeleteSubscriberInput, optFns ...func(*Options)) (*DeleteSubscriberOutput, error) {
	if params == nil {
		params = &DeleteSubscriberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSubscriber", params, optFns, c.addOperationDeleteSubscriberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSubscriberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of DeleteSubscriber
type DeleteSubscriberInput struct {

	// The accountId that is associated with the budget whose subscriber you want to
	// delete.
	//
	// This member is required.
	AccountId *string

	// The name of the budget whose subscriber you want to delete.
	//
	// This member is required.
	BudgetName *string

	// The notification whose subscriber you want to delete.
	//
	// This member is required.
	Notification *types.Notification

	// The subscriber that you want to delete.
	//
	// This member is required.
	Subscriber *types.Subscriber

	noSmithyDocumentSerde
}

// Response of DeleteSubscriber
type DeleteSubscriberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSubscriberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSubscriber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSubscriber{}, middleware.After)
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
	if err = addOpDeleteSubscriberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSubscriber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSubscriber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DeleteSubscriber",
	}
}
