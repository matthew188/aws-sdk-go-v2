// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a version of a subscription definition which has already been defined.
func (c *Client) CreateSubscriptionDefinitionVersion(ctx context.Context, params *CreateSubscriptionDefinitionVersionInput, optFns ...func(*Options)) (*CreateSubscriptionDefinitionVersionOutput, error) {
	if params == nil {
		params = &CreateSubscriptionDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubscriptionDefinitionVersion", params, optFns, c.addOperationCreateSubscriptionDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubscriptionDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubscriptionDefinitionVersionInput struct {

	// The ID of the subscription definition.
	//
	// This member is required.
	SubscriptionDefinitionId *string

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// A list of subscriptions.
	Subscriptions []types.Subscription

	noSmithyDocumentSerde
}

type CreateSubscriptionDefinitionVersionOutput struct {

	// The ARN of the version.
	Arn *string

	// The time, in milliseconds since the epoch, when the version was created.
	CreationTimestamp *string

	// The ID of the parent definition that the version is associated with.
	Id *string

	// The ID of the version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubscriptionDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSubscriptionDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSubscriptionDefinitionVersion{}, middleware.After)
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
	if err = addOpCreateSubscriptionDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubscriptionDefinitionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSubscriptionDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateSubscriptionDefinitionVersion",
	}
}
