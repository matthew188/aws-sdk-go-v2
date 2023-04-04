// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/pinpointsmsvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new event destination in a configuration set.
func (c *Client) CreateConfigurationSetEventDestination(ctx context.Context, params *CreateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*CreateConfigurationSetEventDestinationOutput, error) {
	if params == nil {
		params = &CreateConfigurationSetEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfigurationSetEventDestination", params, optFns, c.addOperationCreateConfigurationSetEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationSetEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Create a new event destination in a configuration set.
type CreateConfigurationSetEventDestinationInput struct {

	// ConfigurationSetName
	//
	// This member is required.
	ConfigurationSetName *string

	// An object that defines a single event destination.
	EventDestination *types.EventDestinationDefinition

	// A name that identifies the event destination.
	EventDestinationName *string

	noSmithyDocumentSerde
}

// An empty object that indicates that the event destination was created
// successfully.
type CreateConfigurationSetEventDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfigurationSetEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfigurationSetEventDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfigurationSetEventDestination{}, middleware.After)
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
	if err = addOpCreateConfigurationSetEventDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfigurationSetEventDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfigurationSetEventDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "CreateConfigurationSetEventDestination",
	}
}
