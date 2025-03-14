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

// Creates a subscriber. You must create the associated budget and notification
// before you create the subscriber.
func (c *Client) CreateSubscriber(ctx context.Context, params *CreateSubscriberInput, optFns ...func(*Options)) (*CreateSubscriberOutput, error) {
	if params == nil {
		params = &CreateSubscriberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubscriber", params, optFns, c.addOperationCreateSubscriberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubscriberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request of CreateSubscriber
type CreateSubscriberInput struct {

	// The accountId that is associated with the budget that you want to create a
	// subscriber for.
	//
	// This member is required.
	AccountId *string

	// The name of the budget that you want to subscribe to. Budget names must be
	// unique within an account.
	//
	// This member is required.
	BudgetName *string

	// The notification that you want to create a subscriber for.
	//
	// This member is required.
	Notification *types.Notification

	// The subscriber that you want to associate with a budget notification.
	//
	// This member is required.
	Subscriber *types.Subscriber

	noSmithyDocumentSerde
}

// Response of CreateSubscriber
type CreateSubscriberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubscriberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSubscriber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSubscriber{}, middleware.After)
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
	if err = addOpCreateSubscriberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubscriber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSubscriber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "CreateSubscriber",
	}
}
