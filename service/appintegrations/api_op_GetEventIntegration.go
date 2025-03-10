// Code generated by smithy-go-codegen DO NOT EDIT.

package appintegrations

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/appintegrations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the event integration.
func (c *Client) GetEventIntegration(ctx context.Context, params *GetEventIntegrationInput, optFns ...func(*Options)) (*GetEventIntegrationOutput, error) {
	if params == nil {
		params = &GetEventIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventIntegration", params, optFns, c.addOperationGetEventIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventIntegrationInput struct {

	// The name of the event integration.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetEventIntegrationOutput struct {

	// The description of the event integration.
	Description *string

	// The EventBridge bus.
	EventBridgeBus *string

	// The event filter.
	EventFilter *types.EventFilter

	// The Amazon Resource Name (ARN) for the event integration.
	EventIntegrationArn *string

	// The name of the event integration.
	Name *string

	// The tags used to organize, track, or control access for this resource. For
	// example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEventIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEventIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEventIntegration{}, middleware.After)
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
	if err = addOpGetEventIntegrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventIntegration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "app-integrations",
		OperationName: "GetEventIntegration",
	}
}
