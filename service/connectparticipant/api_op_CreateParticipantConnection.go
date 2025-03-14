// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connectparticipant/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the participant's connection. ParticipantToken is used for invoking this
// API instead of ConnectionToken. The participant token is valid for the lifetime
// of the participant – until they are part of a contact. The response URL for
// WEBSOCKET Type has a connect expiry timeout of 100s. Clients must manually
// connect to the returned websocket URL and subscribe to the desired topic. For
// chat, you need to publish the following on the established websocket connection:
// {"topic":"aws/subscribe","content":{"topics":["aws/chat"]}} Upon websocket URL
// expiry, as specified in the response ConnectionExpiry parameter, clients need to
// call this API again to obtain a new websocket URL and perform the same steps as
// before. Message streaming support: This API can also be used together with the
// StartContactStreaming
// (https://docs.aws.amazon.com/connect/latest/APIReference/API_StartContactStreaming.html)
// API to create a participant connection for chat contacts that are not using a
// websocket. For more information about message streaming, Enable real-time chat
// message streaming
// (https://docs.aws.amazon.com/connect/latest/adminguide/chat-message-streaming.html)
// in the Amazon Connect Administrator Guide. Feature specifications: For
// information about feature specifications, such as the allowed number of open
// websocket connections per participant, see Feature specifications
// (https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#feature-limits)
// in the Amazon Connect Administrator Guide. The Amazon Connect Participant
// Service APIs do not use Signature Version 4 authentication
// (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
func (c *Client) CreateParticipantConnection(ctx context.Context, params *CreateParticipantConnectionInput, optFns ...func(*Options)) (*CreateParticipantConnectionOutput, error) {
	if params == nil {
		params = &CreateParticipantConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateParticipantConnection", params, optFns, c.addOperationCreateParticipantConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateParticipantConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateParticipantConnectionInput struct {

	// This is a header parameter. The ParticipantToken as obtained from
	// StartChatContact
	// (https://docs.aws.amazon.com/connect/latest/APIReference/API_StartChatContact.html)
	// API response.
	//
	// This member is required.
	ParticipantToken *string

	// Amazon Connect Participant is used to mark the participant as connected for
	// customer participant in message streaming, as well as for agent or manager
	// participant in non-streaming chats.
	ConnectParticipant *bool

	// Type of connection information required. This can be omitted if
	// ConnectParticipant is true.
	Type []types.ConnectionType

	noSmithyDocumentSerde
}

type CreateParticipantConnectionOutput struct {

	// Creates the participant's connection credentials. The authentication token
	// associated with the participant's connection.
	ConnectionCredentials *types.ConnectionCredentials

	// Creates the participant's websocket connection.
	Websocket *types.Websocket

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateParticipantConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateParticipantConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateParticipantConnection{}, middleware.After)
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
	if err = addOpCreateParticipantConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateParticipantConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateParticipantConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateParticipantConnection",
	}
}
