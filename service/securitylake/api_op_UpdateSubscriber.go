// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing subscription for the given Amazon Security Lake account ID.
// You can update a subscriber by changing the sources that the subscriber consumes
// data from.
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

type UpdateSubscriberInput struct {

	// A value created by Security Lake that uniquely identifies your subscription.
	//
	// This member is required.
	Id *string

	// The supported Amazon Web Services from which logs and events are collected. For
	// the list of supported Amazon Web Services, see the Amazon Security Lake User
	// Guide
	// (https://docs.aws.amazon.com/security-lake/latest/userguide/internal-sources.html).
	//
	// This member is required.
	SourceTypes []types.SourceType

	// The external ID of the Security Lake account.
	ExternalId *string

	// The description of the Security Lake account subscriber.
	SubscriberDescription *string

	// The name of the Security Lake account subscriber.
	SubscriberName *string

	noSmithyDocumentSerde
}

type UpdateSubscriberOutput struct {

	// The account of the subscriber.
	Subscriber *types.SubscriberResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSubscriberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSubscriber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSubscriber{}, middleware.After)
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
		SigningName:   "securitylake",
		OperationName: "UpdateSubscriber",
	}
}
