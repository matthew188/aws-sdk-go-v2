// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Puts the specified proxy configuration to the specified Amazon Chime SDK Voice
// Connector.
func (c *Client) PutVoiceConnectorProxy(ctx context.Context, params *PutVoiceConnectorProxyInput, optFns ...func(*Options)) (*PutVoiceConnectorProxyOutput, error) {
	if params == nil {
		params = &PutVoiceConnectorProxyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutVoiceConnectorProxy", params, optFns, c.addOperationPutVoiceConnectorProxyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutVoiceConnectorProxyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorProxyInput struct {

	// The default number of minutes allowed for proxy session.
	//
	// This member is required.
	DefaultSessionExpiryMinutes *int32

	// The countries for proxy phone numbers to be selected from.
	//
	// This member is required.
	PhoneNumberPoolCountries []string

	// The Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// When true, stops proxy sessions from being created on the specified Amazon Chime
	// SDK Voice Connector.
	Disabled *bool

	// The phone number to route calls to after a proxy session expires.
	FallBackPhoneNumber *string

	noSmithyDocumentSerde
}

type PutVoiceConnectorProxyOutput struct {

	// The proxy configuration details.
	Proxy *types.Proxy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutVoiceConnectorProxyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutVoiceConnectorProxy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutVoiceConnectorProxy{}, middleware.After)
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
	if err = addOpPutVoiceConnectorProxyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutVoiceConnectorProxy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutVoiceConnectorProxy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorProxy",
	}
}
